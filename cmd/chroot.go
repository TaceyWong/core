package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// ChRootCMDVersion chroor version
const ChRootCMDVersion = "0.0.1"

// ChRootCMD define `chroot` cmd
var ChRootCMD = cli.Command{
	Name:        "chroot",
	Aliases:     []string{"CHROOT"},
	Usage:       "以指定的新根为运行指定命令时的的根目录。",
	Description: `If no command is given, run '"$SHELL" -i' (default: '/bin/sh -i').`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "output version information and exit",
		},
	},
	Action: func(c *cli.Context) error {
		if c.Bool("version") {
			fmt.Println(ChRootCMDVersion)
			return nil
		}
		return nil
	},
}
