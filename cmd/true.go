package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// TrueCMDVersion true version
const TrueCMDVersion = "0.0.1"

// TrueCMD define `true` command
var TrueCMD = cli.Command{
	Name:    "true",
	Aliases: []string{"TRUE"},
	Usage:   "do nothing, unsuccessfully",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "output version information and exit",
		},
	},
	Action: func(c *cli.Context) error {
		if c.Bool("version") {
			fmt.Println(TrueCMDVersion)
			return nil
		}
		return nil
	},
}
