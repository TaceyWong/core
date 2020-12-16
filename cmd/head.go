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
	Usage:   "Print the first 10 lines of each FILE to standard output.",
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
			Usage:   "output version information and exit",
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
