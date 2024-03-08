package app

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/bogem/id3v2/v2"
	"github.com/fatih/color"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"

	"github.com/caiknife/mp3lister/lib"
	"github.com/caiknife/mp3lister/lib/types"
)

func action(ctx *cli.Context) error {
	if ctx.Bool("debug") {
		cli.VersionPrinter(ctx)
	}

	path, err := lib.GetInputPath(ctx.String("input"))
	if err != nil {
		return err
	}

	lib.ColorPrintf("working in %s\n", color.CyanString("%s", path))

	err = filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
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

		fmt.Println(mp3)

		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
