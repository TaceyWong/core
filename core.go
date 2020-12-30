package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"time"

	"github.com/TaceyWong/core/cmd"

	"github.com/urfave/cli/v2"
)

var logo = `
   ______   ____     ____     ______
  / ____/  / __ \   / __ \   / ____/
 / /      / / / /  / /_/ /  / __/
/ /___   / /_/ /  / _, _/  / /___
\____/   \____/  /_/ |_|  /_____/

Please type "ctools -h/--help" for the help of usage

`

// Setup the main-in of Core
func Setup() {
	app := cli.NewApp()
	app.Name = "core"
	app.Version = "0.0.1"
	app.Compiled = time.Now()
	app.Authors = []*cli.Author{
		{
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

	app.Commands = []*cli.Command{
		&cmd.ArchCMD,
		&cmd.Base32CMD,
		&cmd.Base64CMD,
		&cmd.BaseNameCMD,
		&cmd.CatCMD,
		&cmd.DfCMD,
		&cmd.DuCMD,
		&cmd.EnvCMD,
		&cmd.GroupsCMD,
		&cmd.NlCMD,
		&cmd.UNameCMD,
		&cmd.SeqCMD,
		&cmd.StatCMD,
		&cmd.TouchCMD,
		&cmd.TrueCMD,
		&cmd.TtyCMD,
		&cmd.FalseCMD,
		&cmd.YesCMD,
		&cmd.HostnameCMD,
		&cmd.HostidCMD,
		&cmd.IDCMD,
		&cmd.UsersCMD,
		&cmd.UptimeCMD,
		&cmd.MkTempCMD,
		&cmd.PrintEnvCMD,
		&cmd.LogNameCMD,
		&cmd.LsCPUCMD,
		&cmd.PWDCMD,
		&cmd.RMCMD,
		&cmd.CPCMD,
		&cmd.TailCMD,
		&cmd.MVCMD,
		&cmd.CommCMD,
		&cmd.ExpandCMD,
		&cmd.UnExpandCMD,
		&cmd.ChGrpCMD,
		&cmd.ChOwnCMD,
		&cmd.ChRootCMD,
		&cmd.CkSumCMD,
		&cmd.CSplitCMD,
		&cmd.CutCMD,
		&cmd.DateCMD,
		&cmd.DDCMD,
		&cmd.DirNameCMD,
		&cmd.ExprCMD,
		&cmd.FactorCMD,
		&cmd.HeadCMD,
		&cmd.JoinCMD,
		&cmd.NiceCMD,
		&cmd.NProcCMD,
		&cmd.NumFmtCMD,
		&cmd.PasteCMD,
		&cmd.PathChkCMD,
		&cmd.PtxCMD,
		&cmd.RealPathCMD,
		&cmd.ShufCMD,
		&cmd.SortCMD,
		&cmd.SplitCMD,
		&cmd.SyncCMD,
		&cmd.TeeCMD,
		&cmd.TestCMD,
		&cmd.TimeoutCMD,
		&cmd.TrCMD,
		&cmd.TruncateCMD,
		&cmd.TSortCMD,
		&cmd.ShredCMD,
	}
	app.Action = func(c *cli.Context) error {
		fmt.Print(logo)
		return nil
	}
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
