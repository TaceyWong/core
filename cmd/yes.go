package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// YesCMDVersion yes version
const YesCMDVersion = "0.0.1"

// YesCMD define `yes` cmd
var YesCMD = cli.Command{
	Name:    "yes",
	Aliases: []string{"YES"},
	Usage:   "output a string repeatedly until killed",
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
		strContent := "y"
		if c.Args().Get(0) != ""  {
			strContent = c.Args().Get(0)
		}
		for {
			fmt.Println(strContent)
		}
	},
}
