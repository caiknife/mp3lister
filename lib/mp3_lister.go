package lib

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/bogem/id3v2/v2"
	"github.com/duke-git/lancet/v2/slice"
	"github.com/fatih/color"
	"github.com/spf13/cast"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/caiknife/mp3lister/lib/types"
	"github.com/caiknife/mp3lister/orm/music"
	"github.com/caiknife/mp3lister/orm/music/model"
)

type MP3Lister struct {
	InputPath  string
	OutputName string
	OutputExt  string

	finalOutput string
	all         types.MP3Collection
}

type OptionApply interface {
	Apply(l *MP3Lister)
}

var (
	_ OptionApply = (*Option)(nil)
)

type Option func(lister *MP3Lister)

func (o Option) Apply(l *MP3Lister) {
	o(l)
}

func WithInputPath(input string) Option {
	return func(lister *MP3Lister) {
		lister.InputPath = input
	}
}

func WithOutputName(output string) Option {
	return func(lister *MP3Lister) {
		lister.OutputName = output
	}
}

func WithOutputExt(ext string) Option {
	return func(lister *MP3Lister) {
		lister.OutputExt = ext
	}
}

func NewMP3Lister(ops ...Option) *MP3Lister {
	lister := &MP3Lister{
		InputPath:  "",
		OutputName: "",
		OutputExt:  "csv",
	}

	for _, op := range ops {
		op.Apply(lister)
	}
	lister.all = make(types.MP3Collection, 0)
	lister.finalOutput = lister.OutputName + "." + lister.OutputExt
	return lister
}

func (m *MP3Lister) Do() error {
	if m.InputPath == "" {
		return ErrInputPathIsEmpty
	}
	if m.OutputName == "" {
		return ErrOutputNameIsEmpty
	}
	if m.OutputExt == "" {
		return ErrOutputExtIsEmpty
	}

	err := filepath.WalkDir(m.InputPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			if errors.Is(err, os.ErrPermission) {
				return nil
			}
			return err
		}
		// 忽略目录
		if d.IsDir() {
			return nil
		}
		// 忽略非mp3文件
		if !strings.HasSuffix(d.Name(), ".mp3") {
			return nil
		}
		mp3, err := types.NewMP3(path)
		if err != nil {
			if errors.Is(err, id3v2.ErrUnsupportedVersion) {
				return nil
			}
			return err
		}
		m.all = append(m.all, mp3)

		return nil
	})
	if err != nil {
		return err
	}

	sort.Sort(m.all)

	return m.WriteToFile()
}

func (m *MP3Lister) SaveToDB(dsn string) error {
	if m.all.Len() < 1 {
		return ErrDataIsEmpty
	}

	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return nil
	}

	if !db.Migrator().HasTable(&model.Song{}) {
		err := db.AutoMigrate(&model.Song{})
		if err != nil {
			return err
		}
	}

	music.SetDefault(db)

	// 清空表
	_, err = music.Song.Unscoped().Where(music.Song.ID).Delete()
	if err != nil {
		return err
	}
	// 插入数据
	songs := slice.Map[*types.MP3, *model.Song](m.all, func(index int, item *types.MP3) *model.Song {
		song := &model.Song{
			Title:      item.Title,
			Artist:     item.Artist,
			Album:      item.Album,
			Bpm:        cast.ToInt32(item.BPM),
			OriginFile: item.OriginFile,
		}
		return song
	})

	err = music.Song.CreateInBatches(songs, 100)
	if err != nil {
		return err
	}

	return nil
}

func (m *MP3Lister) WriteToFile() error {
	switch m.OutputExt {
	case "csv":
		return m.writeToCSV()
	default:
	}
	return nil
}

func (m *MP3Lister) writeToCSV() error {
	file, err := os.Create(m.finalOutput)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	headline := []string{"No.", "Artist", "Album", "Title", "BPM", "OriginFile"}
	err = writer.Write(headline)
	if err != nil {
		return err
	}

	for i, mp3 := range m.all {
		err := writer.Write([]string{
			cast.ToString(i + 1),
			mp3.Artist,
			mp3.Album,
			mp3.Title,
			mp3.BPM,
			mp3.OriginFile,
		})
		if err != nil {
			return err
		}
	}
	writer.Flush()
	return nil
}

func (m *MP3Lister) Print() {
	colorPrintf("%s\t%s\t%s\t%s\t%s\t%s\n", "No.", "Artist", "Album", "Title", "BPM", "OriginFile")
	for i, mp3 := range m.all {
		colorPrintf("%s\t%s\t%s\t%s\t%s\t%s\n",
			cast.ToString(i+1),
			mp3.Artist,
			mp3.Album,
			mp3.Title,
			mp3.BPM,
			mp3.OriginFile,
		)
	}
}

func colorPrintf(format string, args ...interface{}) {
	results := make([]any, len(args))
	for i := 0; i < len(args); i++ {
		results[i] = colorString[i%len(colorString)]("%+v", args[i])
	}
	fmt.Printf(format, results...)
}

var (
	colorString = []func(format string, a ...interface{}) string{
		color.HiRedString, color.HiGreenString, color.HiYellowString, color.HiBlueString, color.HiMagentaString, color.HiCyanString,
		color.RedString, color.GreenString, color.YellowString, color.BlueString, color.MagentaString, color.CyanString,
	}
)
