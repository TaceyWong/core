package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// SplitCMDVersion split version
const SplitCMDVersion = "0.0.1"

// SplitCMD define `split` cmd
var SplitCMD = cli.Command{
	Name:    "split",
	Aliases: []string{"SPLIT"},
	Usage:   "split a file into pieces",
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
			Usage:   "output version information and exit",
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
