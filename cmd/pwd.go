package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// PWDCMDVersion pwd version no
var PWDCMDVersion = "v0.0.1"

// PWDCMD define pwd cli command
var PWDCMD = cli.Command{
	Name:      "pwd",
	Aliases:   []string{"PWD"},
	Usage:     `print name of current/working directory`,
	UsageText: `core pwd [command options]`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "output version information and exit",
		}, &cli.BoolFlag{
			Name:    "logical",
			Aliases: []string{"L"},
			Usage:   `use PWD from environment, even if it contains symlinks`,
		}, &cli.BoolFlag{
			Name:    `physical`,
			Aliases: []string{"P"},
			Value:   true,
			Usage:   `avoid all symlinks`,
		},
	},
	Action: pwdAction,
}

func pwdAction(c *cli.Context) error {
	if c.Bool("version") {
		fmt.Println(c.Command.Name, PWDCMDVersion)
		return nil
	}
	return nil

}
