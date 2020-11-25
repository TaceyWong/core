package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// LogNameCMDVersion logname version no
var LogNameCMDVersion = "v0.0.1"

// LogNameCMD define logname cli command
var LogNameCMD = cli.Command{
	Name:    "logname",
	Aliases: []string{"LOGNAME"},
	Usage:   `显示当前用户的名称`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "output version information and exit",
		},
	},
	Action: lognameAction,
}

func lognameAction(c *cli.Context) error {
	if c.Bool("version") {
		fmt.Println(c.Command.Name, LogNameCMDVersion)
		return nil
	}
	return nil

}
