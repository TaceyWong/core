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
	Usage:   "打印文本文件的最后10行到标准输出",
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
			Usage:   "输出版本信息并推出",
		}, &cli.IntFlag{
			Name:    "bytes",
			Aliases: []string{"c"},
			Usage:   "output the last NUM bytes; or use -c +NUM to output starting with byte NUM of each file",
		}, &cli.BoolFlag{
			Name:    "follow",
			Aliases: []string{"f"},
			Usage:   "output appended data as the file grows'",
		}, &cli.BoolFlag{
			Name:    "follow-at",
			Aliases: []string{"x"},
			Usage:   "output appended data as the file grows; an absent option argument means '`descriptor`'",
		}, &cli.BoolFlag{
			Name:  "F",
			Usage: "same as --follow-at=name --retry",
		}, &cli.StringFlag{
			Name:    "lines",
			Aliases: []string{"n"},
			Usage:   "output the last NUM lines, instead of the last 10;or use -n +NUM to output starting with line NUM",
		}, &cli.IntFlag{
			Name:  "max-unchanged-stats",
			Value: 5,
			Usage: "with --follow=name, reopen a FILE which has not" +
				"changed size after `N` (default 5) iterations" +
				"to see if it has been unlinked or renamed" +
				"(this is the usual case of rotated log files);" +
				"with inotify, this option is rarely useful",
		}, &cli.IntFlag{
			Name:  "pid",
			Usage: "with -f, terminate after process ID, PID dies",
		}, &cli.BoolFlag{
			Name:    "quiet",
			Aliases: []string{"q"},
			Usage:   "never output headers giving file names",
		}, &cli.BoolFlag{
			Name:  "silent",
			Usage: "same as --quiet",
		}, &cli.BoolFlag{
			Name:  "retry",
			Usage: "keep trying to open a file if it is inaccessible",
		}, &cli.Float64Flag{
			Name:    "sleep-interval",
			Aliases: []string{"s"},
			Value:   1.0,
			Usage: "with -f, sleep for approximately `N` seconds between iterations" +
				"with inotify and --pid=P, check process P at least once every N seconds",
		}, &cli.BoolFlag{
			Name:    "verbose",
			Aliases: []string{"V"},
			Usage:   "always output headers giving file names",
		}, &cli.BoolFlag{
			Name:    "zero-terminated",
			Aliases: []string{"z"},
			Usage:   "line delimiter is NUL, not newline",
		},
	},
	Action: func(c *cli.Context) (err error) {
		if c.Bool("version") {
			fmt.Println(TailCMDVersion)
			return nil
		}
		fmt.Println(c.FlagNames())
		return err
	},
}
