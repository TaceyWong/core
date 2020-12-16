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
	Usage:   "Remove (unlink) the FILE(s).",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "output version information and exit",
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
