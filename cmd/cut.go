package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// CutCMDVersion cut version
const CutCMDVersion = "0.0.1"

// CutCMD define `cut` cmd
var CutCMD = cli.Command{
	Name:        "cut",
	Aliases:     []string{"CUT"},
	Usage:       "将每个文件中选定行的部分打印到标准输出",
	Description: `如果没有指定文件，或者文件为"-"，则从标准输入读取`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "输出版本信息并推出",
		},
	},
	Action: func(c *cli.Context) error {
		if c.Bool("version") {
			fmt.Println(CutCMDVersion)
			return nil
		}
		return nil
	},
}
