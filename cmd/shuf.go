package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// ShufCMDVersion shuf version
const ShufCMDVersion = "0.0.1"

// ShufCMD define `shuf` cmd
var ShufCMD = cli.Command{
	Name:    "shuf",
	Aliases: []string{"SHUF"},
	Usage:   "生成随机排列",
	Description: `Write a random permutation of the input lines to standard output.

	如果没有指定文件，或者文件为"-"，则从标准输入读取。`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "输出版本信息并退出",
		},
	},
	Action: func(c *cli.Context) error {
		if c.Bool("version") {
			fmt.Println(ShufCMDVersion)
			return nil
		}
		return nil
	},
}
