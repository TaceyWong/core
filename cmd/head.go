package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// HeadCMDVersion head version
const HeadCMDVersion = "0.0.1"

// HeadCMD define `head` cmd
var HeadCMD = cli.Command{
	Name:    "head",
	Aliases: []string{"HEAD"},
	Usage:   "打印文本文件的最前10行到标准输出",
	Description: `Print the first 10 lines of each FILE to standard output.
	With more than one FILE, precede each with a header giving the file name.
	
	如果没有指定文件，或者文件为"-"，则从标准输入读取。
	
	NUM may have a multiplier suffix:
		b 512, kB 1000, K 1024, MB 1000*1000, M 1024*1024,
		GB 1000*1000*1000, G 1024*1024*1024, and so on for T, P, E, Z, Y.`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "输出版本信息并推出",
		}, &cli.IntFlag{
			Name:    "bytes",
			Aliases: []string{"c"},
			Usage:   "print the first `NUM` bytes of each file;with the leading '-', print all but the last NUM bytes of each file",
		}, &cli.IntFlag{
			Name:    "lines",
			Aliases: []string{"n"},
			Usage:   "print the first `NUM` lines instead of the first 10;with the leading '-', print all but the last NUM lines of each file",
		}, &cli.BoolFlag{
			Name:    "silent",
			Aliases: []string{"s"},
			Usage:   "不显示包含给定文件名的文件头",
		}, &cli.BoolFlag{
			Name:    "quiet",
			Aliases: []string{"q"},
			Usage:   "和--silen相同",
		}, &cli.BoolFlag{
			Name:    "verbose",
			Aliases: []string{"V"},
			Usage:   "总是显示包含给定文件名的文件头",
		}, &cli.BoolFlag{
			Name:    "zero-terminated",
			Aliases: []string{"z"},
			Usage:   "以 NUL 字符而非换行符作为行尾分隔符",
		},
	},
	Action: func(c *cli.Context) error {
		if c.Bool("version") {
			fmt.Println(HeadCMDVersion)
			return nil
		}
		return nil
	},
}
