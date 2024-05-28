package main

import (
	"time"

	"github.com/duke-git/lancet/v2/random"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"

	"github.com/caiknife/mp3lister/cmd/fcm/fcm"
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
	return fcm.DefaultFangChenMi.Auth(code, AuthSuccess.Random())
}

func t2(code string) error {
	return fcm.DefaultFangChenMi.Auth(code, AuthNotYet.Random())
}

func t3(code string) error {
	return fcm.DefaultFangChenMi.Auth(code, AuthFailed.Random())
}

func t4(code string) error {
	return fcm.DefaultFangChenMi.Query(code, QuerySuccess.Random())
}

func t5(code string) error {
	return fcm.DefaultFangChenMi.Query(code, QueryNotYet.Random())
}

func t6(code string) error {
	return fcm.DefaultFangChenMi.Query(code, QueryFail.Random())
}

func t7(code string) error {
	s := &fcm.Behavior{
		No: 1,
		Si: random.RandString(32),
		Bt: 0,
		Ot: time.Now().Unix(),
		Ct: 2,
		Di: random.RandString(32),
		Pi: "",
	}
	return fcm.DefaultFangChenMi.LoginOrOut(code, &fcm.Collections{Collections: []*fcm.Behavior{
		s,
	}})
}

func t8(code string) error {
	r := Report.Random()
	s := &fcm.Behavior{
		No: 100,
		Si: random.RandString(32),
		Bt: 1,
		Ot: time.Now().Unix(),
		Ct: 0,
		Di: "",
		Pi: r.Pi,
	}
	return fcm.DefaultFangChenMi.LoginOrOut(code, &fcm.Collections{Collections: []*fcm.Behavior{
		s,
	}})
}
