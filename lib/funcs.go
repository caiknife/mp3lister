package lib

import (
	"bufio"
	"encoding/csv"
	"os"
	"path/filepath"
	"strings"

	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/duke-git/lancet/v2/slice"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"github.com/vicanso/go-charts/v2"

	"github.com/caiknife/mp3lister/lib/types"
)

func GetOutputFileNoExt(originFile string) string {
	base := filepath.Base(originFile)
	s := filepath.Ext(originFile)
	return base[0 : len(base)-len(s)]
}

func GetOutputFileWithExt(originFile, ext string) string {
	return GetOutputFileNoExt(originFile) + "." + ext
}

func ReadM3U(file string) (s types.Slice[*types.MP3], err error) {
	abs, err := filepath.Abs(file)
	if err != nil {
		return nil, errors.WithMessage(err, "file abs path error")
	}
	if !fileutil.IsExist(abs) {
		return nil, errors.WithMessage(err, "file does not exist")
	}
	if strings.ToLower(filepath.Ext(abs)) != ".m3u" {
		return nil, errors.WithMessage(err, "file is not m3u")
	}

	u, err := readM3U(abs)
	if err != nil {
		return nil, err
	}

	s = slice.Map[string, *types.MP3](u, func(_ int, item string) *types.MP3 {
		mp3, err := types.NewMP3(item)
		if err != nil {
			return nil
		}
		return mp3
	})
	s = slice.Filter(s, func(_ int, item *types.MP3) bool {
		return item != nil
	})
	return s, nil
}

func readM3U(file string) (s types.Slice[string], err error) {
	open, err := os.Open(file)
	if err != nil {
		return nil, errors.WithMessage(err, "open file error")
	}
	defer open.Close()

	scanner := bufio.NewScanner(open)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") || strings.TrimSpace(line) == "" {
			continue
		}
		s = append(s, line)
	}
	return s, nil
}

func WriteChart(s types.Slice[*types.MP3], outputFile string) error {
	values := [][]float64{
		slice.Map[*types.MP3, float64](s, func(_ int, item *types.MP3) float64 {
			return cast.ToFloat64(item.BPM)
		}),
	}
	xValues := slice.Map[*types.MP3, string](s, func(index int, _ *types.MP3) string {
		return cast.ToString(index + 1)
	})
	render, err := charts.LineRender(
		values,
		charts.TitleTextOptionFunc(GetOutputFileNoExt(outputFile)),
		charts.XAxisDataOptionFunc(xValues),
		func(opt *charts.ChartOption) {
			opt.Width = 600 * 2
			opt.Height = 400 * 2
			opt.FillArea = true
		},
	)
	if err != nil {
		return errors.WithMessage(err, "charts render error")
	}
	bytes, err := render.Bytes()
	if err != nil {
		return errors.WithMessage(err, "render bytes error")
	}
	err = fileutil.WriteBytesToFile(outputFile, bytes)
	if err != nil {
		return errors.WithMessage(err, "write output file error")
	}
	return nil
}

func WriteCSV(s types.Slice[*types.MP3], outputFile string) error {
	create, err := os.Create(outputFile)
	if err != nil {
		return errors.WithMessage(err, "create file error")
	}
	defer create.Close()

	writer := csv.NewWriter(create)
	defer writer.Flush()

	err = writer.Write([]string{"No.", "BPM", "Title", "Artist", "Album", "Memo", "OriginFile"})
	if err != nil {
		return errors.WithMessage(err, "write csv header error")
	}
	for i, file := range s {
		err := writer.Write([]string{
			cast.ToString(i + 1),
			file.BPM,
			file.Title,
			file.Artist,
			file.Album,
			file.Memo,
			file.OriginFile,
		})
		if err != nil {
			return errors.WithMessage(err, "write csv content error")
		}
	}

	return nil
}
