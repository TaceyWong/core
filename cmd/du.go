package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// DuCMDVersion du version no
var DuCMDVersion = "v0.0.1"

// DuCMD define du cli command
var DuCMD = cli.Command{
	Name:    "du",
	Aliases: []string{"DU"},
	Usage:   `递归地汇总目录或文件的磁盘使用情况`,
	Description: `所显示的数值是来自 --block-size、DU_BLOCK_SIZE、BLOCK_SIZE 
	及 BLOCKSIZE 环境变量中第一个可用的 SIZE 单位。
	否则，默认单位是 1024 字节(或是 512，若设定 POSIXLY_CORRECT 的话)。
	
	The SIZE argument is an integer and optional unit (example: 10K is 10*1024).
	Units are K,M,G,T,P,E,Z,Y (powers of 1024) or KB,MB,... (powers of 1000)`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "输出版本信息并推出",
		}, &cli.BoolFlag{
			Name:    "null",
			Aliases: []string{"0"},
			Usage:   "end each output line with NUL, not newline",
		}, &cli.BoolFlag{
			Name:    "all",
			Aliases: []string{"a"},
			Usage:   "write counts for all files, not just directories",
		}, &cli.Int64Flag{
			Name: "apparent-size",
			Usage: `print apparent sizes, rather than disk usage; although
			the apparent size is usually smaller, it may be
			larger due to holes in ('sparse') files, internal
			fragmentation, indirect blocks, and the like`,
		}, &cli.StringFlag{
			Name:    "block-size",
			Aliases: []string{"B"},
			Usage: `scale sizes by SIZE before printing them; e.g.,
			'-BM' prints sizes in units of 1,048,576 bytes;
			see SIZE format below`,
		}, &cli.BoolFlag{
			Name:    "bytes",
			Aliases: []string{"b"},
			Usage:   `equivalent to '--apparent-size --block-size=1'`,
		}, &cli.BoolFlag{
			Name:    "total",
			Aliases: []string{"c"},
			Usage:   ` produce a grand total`,
		}, &cli.StringSliceFlag{
			Name:    "dereference-args",
			Aliases: []string{"D"},
			Usage: `dereference only symlinks that are listed on the
			command line`,
		}, &cli.Int64Flag{
			Name:    "max-depth",
			Aliases: []string{"d"},
			Usage: `print the total for a directory (or file, with --all)
			only if it is N or fewer levels below the command
			line argument;  --max-depth=0 is the same as
			--summarize`,
		}, &cli.StringFlag{
			Name: "files0-from",
			Usage: `summarize disk usage of the
			NUL-terminated file names specified in file F;
			if F is -, then read names from standard input`,
		}, &cli.BoolFlag{
			Name:    "human-readable",
			Aliases: []string{"hu"},
			Usage:   `print sizes in human readable format (e.g., 1K 234M 2G)`,
		}, &cli.BoolFlag{
			Name:  "inodes",
			Usage: `list inode usage information instead of block usage`,
		}, &cli.BoolFlag{
			Name:    "dereference",
			Aliases: []string{"L"},
			Usage:   `dereference all symbolic links`,
		}, &cli.BoolFlag{
			Name:    "count-links",
			Aliases: []string{"l"},
			Usage:   `count sizes many times if hard linked`,
		}, &cli.BoolFlag{
			Name:    "block-size-m",
			Aliases: []string{"m"},
			Usage:   `like --block-size=1M`,
		}, &cli.BoolFlag{
			Name:    "block-size-k",
			Aliases: []string{"k"},
			Usage:   `like --block-size=1K`,
		}, &cli.BoolFlag{
			Name:    "no-dereference",
			Aliases: []string{"P"},
			Usage:   `don't follow any symbolic links (this is the default)`,
		}, &cli.BoolFlag{
			Name:    "separate-dirs",
			Aliases: []string{"S"},
			Usage:   `for directories do not include size of subdirectories`,
		}, &cli.BoolFlag{
			Name:  "si",
			Usage: `like -hu, but use powers of 1000 not 1024`,
		}, &cli.BoolFlag{
			Name:    "summarize",
			Aliases: []string{"s"},
			Usage:   `display only a total for each argument`,
		}, &cli.Int64Flag{
			Name:    "threshold",
			Aliases: []string{"t"},
			Usage: `exclude entries smaller than SIZE if positive,
			or entries greater than SIZE if negative`,
		}, &cli.BoolFlag{
			Name: "time",
			Usage: `show time of the last modification of any file in the
			directory, or any of its subdirectorie`,
		}, &cli.BoolFlag{
			Name: "time-word",
			Usage: `show time as WORD instead of modification time:
			atime, access, use, ctime or status`,
		}, &cli.StringFlag{
			Name: "time-style",
			Usage: `show times using STYLE, which can be:
			full-iso, long-iso, iso, or +FORMAT;
			FORMAT is interpreted like in 'date'`,
		}, &cli.StringFlag{
			Name:    "exclude-from",
			Aliases: []string{"X"},
			Usage:   `exclude files that match any pattern in FILE`,
		}, &cli.StringFlag{
			Name:  "exclude",
			Usage: `exclude files that match PATTERN`,
		}, &cli.BoolFlag{
			Name:    "one-file-system",
			Aliases: []string{"x"},
			Usage:   `kip directories on different file systems`,
		},
	},
	Action: duAction,
}

func duAction(c *cli.Context) error {
	if c.Bool("version") {
		fmt.Println(c.Command.Name, DuCMDVersion)
		return nil
	}
	return nil

}
