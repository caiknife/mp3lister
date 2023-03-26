package app

import (
	"encoding/csv"
	"errors"
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

	"gitea.caiknife.live/caiknife/mp3lister/orm/dal"
	"gitea.caiknife.live/caiknife/mp3lister/orm/model"
)

type MP3Lister struct {
	InputPath  string
	OutputName string
	OutputExt  string

	finalOutput string
	all         MP3Collection
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
	lister.all = make(MP3Collection, 0)
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
		mp3, err := NewMP3(path)
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

	dal.SetDefault(db)

	// 清空表
	_, err = dal.Song.Unscoped().Where(dal.Song.ID).Delete()
	if err != nil {
		return nil
	}
	// 插入数据
	songs := slice.Map[*MP3, *model.Song](m.all, func(index int, item *MP3) *model.Song {
		song := &model.Song{
			Title:      item.Title,
			Artist:     item.Artist,
			Album:      item.Album,
			Bpm:        item.BPM,
			OriginFile: item.OriginFile,
		}
		return song
	})

	err = dal.Song.CreateInBatches(songs, 100)
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
	color.Yellow("%s\t%s\t%s\t%s\t%s\t%s", "No.", "Artist", "Album", "Title", "BPM", "OriginFile")
	for i, mp3 := range m.all {
		color.Cyan("%s\t%s\t%s\t%s\t%s\t%s",
			cast.ToString(i+1),
			mp3.Artist,
			mp3.Album,
			mp3.Title,
			mp3.BPM,
			mp3.OriginFile,
		)
	}
}
