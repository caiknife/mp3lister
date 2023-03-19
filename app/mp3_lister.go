package app

import (
	"io/fs"
	"path/filepath"
	"strings"
)

type MP3Lister struct {
	InputPath  string
	OutputName string
	OutputExt  string

	finalOutput string
	all         []*MP3
}

type Option func(lister *MP3Lister)

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
		op(lister)
	}

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
		if !strings.HasPrefix(d.Name(), ".mp3") {
			return nil
		}
		mp3, err := NewMP3(path)
		if err != nil {
			return err
		}

		m.all = append(m.all, mp3)

		return nil
	})
	if err != nil {
		return err
	}

	return m.writeToFile()
}

func (m *MP3Lister) writeToFile() error {
	switch m.OutputExt {
	case "csv":
		return m.writeToCSV()
	default:
	}
	return nil
}

func (m *MP3Lister) writeToCSV() error {
	return nil
}
