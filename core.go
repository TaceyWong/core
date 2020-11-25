package main

import (
	"log"
	"os"
	"sort"
	"time"

	"github.com/TaceyWong/core/cmd"

	"github.com/urfave/cli/v2"
)

// Setup the main-in of Core
func Setup() {
	app := cli.NewApp()
	app.Name = "core"
	app.Version = "0.0.1"
	app.Compiled = time.Now()
	app.Authors = []*cli.Author{
		&cli.Author{
			Name:  "Tacey Wong",
			Email: "xinyong.wang@qq.com",
		},
	}
	app.Copyright = "(c) 2020 - Forever Tacey Wong"
	app.Usage = "Clone of GNU Coreutils"
	app.EnableBashCompletion = true
	app.Action = func(c *cli.Context) error {
		return nil
	}

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "lang, l",
			Value: "english",
			Usage: "language for the app",
		},
		&cli.StringFlag{
			Name:  "config, c",
			Usage: "load configuration from `FILE`",
		},
	}

	app.Commands = []*cli.Command{}
	app.Commands = append(app.Commands, &cmd.ArchCMD)
	app.Commands = append(app.Commands, &cmd.Base32CMD)
	app.Commands = append(app.Commands, &cmd.Base64CMD)
	app.Commands = append(app.Commands, &cmd.BaseNameCMD)
	app.Commands = append(app.Commands, &cmd.CatCMD)
	app.Commands = append(app.Commands, &cmd.DfCMD)
	app.Commands = append(app.Commands, &cmd.DuCMD)
	app.Commands = append(app.Commands, &cmd.NlCMD)
	app.Commands = append(app.Commands, &cmd.UNameCMD)
	app.Commands = append(app.Commands, &cmd.SeqCMD)
	app.Commands = append(app.Commands, &cmd.StatCMD)
	app.Commands = append(app.Commands, &cmd.TrueCMD)
	app.Commands = append(app.Commands, &cmd.FalseCMD)
	app.Commands = append(app.Commands, &cmd.YesCMD)
	app.Commands = append(app.Commands, &cmd.HostnameCMD)
	app.Commands = append(app.Commands, &cmd.HostidCMD)
	app.Commands = append(app.Commands, &cmd.UptimeCMD)
	app.Commands = append(app.Commands, &cmd.MkTempCMD)
	sort.Sort(cli.CommandsByName(app.Commands))
	sort.Sort(cli.FlagsByName(app.Flags))
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	Setup()
}
