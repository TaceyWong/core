package cmd

import (
	"fmt"
         "os"

	"github.com/urfave/cli/v2"
)

// FalseCMDVersion false version
const FalseCMDVersion = "0.0.1"

// FalseCMD define `false` command
var FalseCMD = cli.Command{
	Name:    "false",
	Aliases: []string{"FALSE"},
	Usage:   "do nothing, successfully",
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
		os.Exit(1)
		return nil
	},
}
