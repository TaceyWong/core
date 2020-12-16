package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// UnExpandCMDVersion unexpand version
const UnExpandCMDVersion = "0.0.1"

// UnExpandCMD define `unexpand` cmd
var UnExpandCMD = cli.Command{
	Name:    "unexpand",
	Aliases: []string{"UNEXPAND"},
	Usage:   "空格转换成制表符",
	Description: `Convert blanks in each FILE to tabs, writing to standard output.
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
			fmt.Println(UnExpandCMDVersion)
			return nil
		}
		return nil
	},
}
