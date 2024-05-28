package main

import (
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"

	"github.com/caiknife/mp3lister/cmd/fcm/fcm"
	"github.com/caiknife/mp3lister/config"
	"github.com/caiknife/mp3lister/lib/logger"
)

var commands = cli.Commands{
	&cli.Command{
		Name:    "testcase01",
		Aliases: []string{"t1"},
		Action:  testcase(),
	},
	&cli.Command{
		Name:    "testcase02",
		Aliases: []string{"t2"},
		Action:  testcase(),
	},
	&cli.Command{
		Name:    "testcase03",
		Aliases: []string{"t3"},
		Action:  testcase(),
	},
	&cli.Command{
		Name:    "testcase04",
		Aliases: []string{"t4"},
		Action:  testcase(),
	},
	&cli.Command{
		Name:    "testcase05",
		Aliases: []string{"t5"},
		Action:  testcase(),
	},
	&cli.Command{
		Name:    "testcase06",
		Aliases: []string{"t6"},
		Action:  testcase(),
	},
	&cli.Command{
		Name:    "testcase07",
		Aliases: []string{"t7"},
		Action:  testcase(),
	},
	&cli.Command{
		Name:    "testcase08",
		Aliases: []string{"t8"},
		Action:  testcase(),
	},
}

func testcase() cli.ActionFunc {
	return func(ctx *cli.Context) error {
		args := ctx.Args()
		if args.Len() == 0 {
			return errors.New("请输入测试码")
		}
		code := args.First()
		switch ctx.Command.Name {
		case "testcase01":
			logger.ConsoleLogger.Infoln("testcase01")
			return t1(code)
		case "testcase02":
			logger.ConsoleLogger.Infoln("testcase02")
			return t2(code)
		case "testcase03":
			logger.ConsoleLogger.Infoln("testcase03")
			return t3(code)
		case "testcase04":
			logger.ConsoleLogger.Infoln("testcase04")
			return t4(code)
		case "testcase05":
			logger.ConsoleLogger.Infoln("testcase05")
			return t5(code)
		case "testcase06":
			logger.ConsoleLogger.Infoln("testcase06")
			return t6(code)
		case "testcase07":
			logger.ConsoleLogger.Infoln("testcase07")
			return t7(code)
		case "testcase08":
			logger.ConsoleLogger.Infoln("testcase08")
			return t8(code)
		default:
			return errors.New("command not found")
		}
		return nil
	}
}

func t1(code string) error {
	return fcm.DefaultFangChenMi.Auth(code, config.AuthSuccess.Random())
}

func t2(code string) error {
	return fcm.DefaultFangChenMi.Auth(code, config.AuthNotYet.Random())
}

func t3(code string) error {
	return fcm.DefaultFangChenMi.Auth(code, config.AuthFailed.Random())
}

func t4(code string) error {
	return nil
}

func t5(code string) error {
	return nil
}

func t6(code string) error {
	return nil
}

func t7(code string) error {
	return nil
}

func t8(code string) error {
	return nil
}
