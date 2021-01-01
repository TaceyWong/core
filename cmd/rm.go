package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// RMCMDVersion rm version no
var RMCMDVersion = "v0.0.1"

// RMCMD define rm cli command
var RMCMD = cli.Command{
	Name:    "rm",
	Aliases: []string{"RM"},
	Usage:   "删除文件",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "输出版本信息并推出",
		},
	},
	Action: rmAction,
}

func rmAction(c *cli.Context) error {
	if c.Bool("version") {
		fmt.Println(c.Command.Name, RMCMDVersion)
		return nil
	}
	return nil

}
