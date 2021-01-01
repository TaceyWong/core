package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// PasteCMDVersion paste version
const PasteCMDVersion = "0.0.1"

// PasteCMD define `paste` cmd
var PasteCMD = cli.Command{
	Name:    "paste",
	Aliases: []string{"PASTE"},
	Usage:   "合并文件文本行",
	Description: `Write lines consisting of the sequentially corresponding lines from
	each FILE, separated by TABs, to standard output.
	
	如果没有指定文件，或者文件为"-"，则从标准输入读取。`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "输出版本信息并推出",
		},
	},
	Action: func(c *cli.Context) error {
		if c.Bool("version") {
			fmt.Println(PasteCMDVersion)
			return nil
		}
		return nil
	},
}
