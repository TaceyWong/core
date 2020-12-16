package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// TSortCMDVersion tsort version
const TSortCMDVersion = "0.0.1"

// TSortCMD define `tsort` cmd
var TSortCMD = cli.Command{
	Name:    "tsort",
	Aliases: []string{"TSORT"},
	Usage:   "perform topological sort",
	Description: `Write totally ordered list consistent with the partial ordering in FILE.

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
			fmt.Println(TSortCMDVersion)
			return nil
		}
		return nil
	},
}
