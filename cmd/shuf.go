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
		}, &cli.BoolFlag{
			Name:    "echo",
			Aliases: []string{"e"},
			Usage:   "treat each ARG as an input line",
		}, &cli.StringFlag{
			Name:    "input-range",
			Aliases: []string{"i"},
			Usage:   "treat each number LO through HI as an input line(`LO-HI`)",
		}, &cli.PathFlag{
			Name:    "output",
			Aliases: []string{"o"},
			Usage:   "write result to `FILE` instead of standard output",
		}, &cli.PathFlag{
			Name:  "random-source",
			Usage: "get random bytes from `FILE`",
		}, &cli.BoolFlag{
			Name:    "repeat",
			Aliases: []string{"r"},
			Usage:   "output lines can be repeated",
		}, &cli.BoolFlag{
			Name:    "zero-terminated",
			Aliases: []string{"z"},
			Usage:   "line delimiter is NUL, not newline",
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
