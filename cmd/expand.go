package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// ExpandCMDVersion expand version
const ExpandCMDVersion = "0.0.1"

// ExpandCMD define `expand` cmd
var ExpandCMD = cli.Command{
	Name:    "expand",
	Aliases: []string{"EXPAND"},
	Usage:   "convert spaces to tabs",
	Description: `Convert tabs in each FILE to spaces, writing to standard output.
	如果没有指定文件，或者文件为"-"，则从标准输入读取。`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "output version information and exit",
		},
	},
	Action: func(c *cli.Context) error {
		if c.Bool("version") {
			fmt.Println(ExpandCMDVersion)
			return nil
		}
		return nil
	},
}
