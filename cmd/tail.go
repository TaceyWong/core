package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// TailCMDVersion tail version
const TailCMDVersion = "0.0.1"

// TailCMD define `tail` command
var TailCMD = cli.Command{
	Name:    "tail",
	Aliases: []string{"TAIL"},
	Usage:   "Print the last 10 lines of each FILE to standard output",
	Description: `With more than one FILE, precede each with a header giving the file name.
	如果没有指定文件，或者文件为"-"，则从标准输入读取。
	
	NUM may have a multiplier suffix:
b 512, kB 1000, K 1024, MB 1000*1000, M 1024*1024,
GB 1000*1000*1000, G 1024*1024*1024, and so on for T, P, E, Z, Y.

如果您希望即时追查一个文件的有效名称而非描述内容(例如循环日志)，默认
的程序动作并不如您所愿。在这种场合可以使用--follow=name 选项，它会使
tail 定期追踪打开给定名称的文件，以确认它是否被删除或被其它某些程序重新创建过。`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "output version information and exit",
		},
	},
	Action: func(c *cli.Context) error {
		if c.Bool("version") {
			fmt.Println(TailCMDVersion)
			return nil
		}
		return nil
	},
}
