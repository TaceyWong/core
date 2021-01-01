package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// SplitCMDVersion split version
const SplitCMDVersion = "0.0.1"

// SplitCMD define `split` cmd
var SplitCMD = cli.Command{
	Name:      "split",
	Aliases:   []string{"SPLIT"},
	UsageText: "core split [OPTION]... [FILE [PREFIX]]",
	Usage:     "文件分片",
	Description: `Output pieces of FILE to PREFIXaa, PREFIXab, ...;
	default size is 1000 lines, and default PREFIX is 'x'.
	
	如果没有指定文件，或者文件为"-"，则从标准输入读取。
	
	The SIZE argument is an integer and optional unit (example: 10K is 10*1024).
Units are K,M,G,T,P,E,Z,Y (powers of 1024) or KB,MB,... (powers of 1000).

CHUNKS may be:
  N       split into N files based on size of input
  K/N     output Kth of N to stdout
  l/N     split into N files without splitting lines/records
  l/K/N   output Kth of N to stdout without splitting lines/records
  r/N     like 'l' but use round robin distribution
  r/K/N   likewise but only output Kth of N to stdout`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "输出版本信息并推出",
		}, &cli.IntFlag{
			Name:    "suffix-lengt",
			Aliases: []string{"a"},
			Usage:   "generate suffixes of length `N`",
			Value:   2,
		}, &cli.IntFlag{
			Name:    "bytes",
			Aliases: []string{"b"},
			Usage:   "put `SIZE` bytes per output file",
		}, &cli.IntFlag{
			Name:    "line-bytes",
			Aliases: []string{"C"},
			Usage:   "put at most `SIZE` bytes of records per output file ",
		}, &cli.BoolFlag{
			Name:  "d",
			Usage: "use numeric suffixes starting at 0, not alphabetic",
		}, &cli.IntFlag{
			Name:  "numeric-suffixes",
			Usage: "same as -d, but allow setting the start value",
		}, &cli.BoolFlag{
			Name:  "x",
			Usage: "use hex suffixes starting at 0, not alphabetic",
		}, &cli.StringFlag{
			Name:  "hex-suffixes",
			Usage: "same as -x, but allow setting the start value",
		}, &cli.BoolFlag{
			Name:    "elide-empty-files",
			Aliases: []string{"e"},
			Usage:   "do not generate empty output files with '-n'",
		}, &cli.StringFlag{
			Name:  "filter",
			Usage: "write to shell COMMAND; file name is $FILE",
		}, &cli.IntFlag{
			Name:    "lines",
			Aliases: []string{"l"},
			Usage:   "put `NUMBER` lines/records per output file",
		}, &cli.IntFlag{
			Name:    "numbers",
			Aliases: []string{"n"},
			Usage:   "generate `CHUNKS` output files; see explanation below",
		}, &cli.StringFlag{
			Name:    "separator",
			Aliases: []string{"t"},
			Usage:   "use `SEP` instead of newline as the record separator;`\\0` (zero) specifies the NUL character",
		}, &cli.BoolFlag{
			Name:    "unbuffered",
			Aliases: []string{"u"},
			Usage:   "immediately copy input to output with '-n r/...'",
		}, &cli.BoolFlag{
			Name:  "verbose",
			Usage: "在每个输出文件打开前输出文件特征",
		},
	},
	Action: func(c *cli.Context) error {
		if c.Bool("version") {
			fmt.Println(SplitCMDVersion)
			return nil
		}
		return nil
	},
}
