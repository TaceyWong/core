package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// PtxCMDVersion ptx version
const PtxCMDVersion = "0.0.1"

// PtxCMD define `ptx` cmd
var PtxCMD = cli.Command{
	Name:    "ptx",
	Aliases: []string{"PTX"},
	Usage:   "生成文件内容的排列索引",
	Description: `Output a permuted index, including context, of the words in the input files.

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
			fmt.Println(PtxCMDVersion)
			return nil
		}
		return nil
	},
}
