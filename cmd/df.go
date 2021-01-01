package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// DfCMDVersion df version no
var DfCMDVersion = "v0.0.1"

// DfCMD define df cli command
var DfCMD = cli.Command{
	Name:      "df",
	Aliases:   []string{"DF"},
	UsageText: "core df [选项]... [文件]...",
	Usage:     `显示文件系统信息`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "输出版本信息并推出",
		}, &cli.BoolFlag{
			Name:    "all",
			Aliases: []string{"a"},
			Usage:   "include pseudo, duplicate, inaccessible file systems",
		}, &cli.StringFlag{
			Name:    "block-size",
			Aliases: []string{"B"},
			Usage: `scale sizes by SIZE before printing them; e.g.,
			'-BM' prints sizes in units of 1,048,576 bytes;
			see SIZE format below`,
		}, &cli.BoolFlag{
			Name:    "human-readable",
			Aliases: []string{"hu"},
			Usage:   "print sizes in powers of 1024 (e.g., 1023M)",
		}, &cli.BoolFlag{
			Name:    "si",
			Aliases: []string{"H"},
			Usage:   "print sizes in powers of 1000 (e.g., 1.1G)",
		}, &cli.BoolFlag{
			Name:    "inodes",
			Aliases: []string{"i"},
			Usage:   "显示inode 信息而非块使用量",
		}, &cli.BoolFlag{
			Name:    "local",
			Aliases: []string{"l"},
			Usage:   "只显示本机的文件系统",
		}, &cli.BoolFlag{
			Name:  "no-sync",
			Value: true,
			Usage: "取得使用量数据前不进行同步动作",
		}, &cli.BoolFlag{
			Name:    "portability",
			Aliases: []string{"P"},
			Usage:   "use the POSIX output format",
		}, &cli.BoolFlag{
			Name:  "sync",
			Usage: "invoke sync before getting usage info",
		}, &cli.BoolFlag{
			Name: "total",
			Usage: ` elide all entries insignificant to available space,
			and produce a grand total`,
		}, &cli.StringFlag{
			Name:    "type",
			Aliases: []string{"t"},
			Usage:   "imit listing to file systems of type TYPE",
		}, &cli.BoolFlag{
			Name:    "print-type",
			Aliases: []string{"T"},
			Usage:   " print file system type",
		}, &cli.StringFlag{
			Name:    "exclude-type",
			Aliases: []string{"x"},
			Usage:   "limit listing to file systems not of type TYPE",
		},
	},
	Action: dfAction,
}

func dfAction(c *cli.Context) error {
	if c.Bool("version") {
		fmt.Println(c.Command.Name, DfCMDVersion)
		return nil
	}
	f := "%10s %10s %10s %10s %10s %10s\n"
	fmt.Printf(f, "文件系统", "1K-块", "已用", "可用", "已用%", "挂载点")
	//  syscall.Statfs_t
	return nil

}
