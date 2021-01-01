package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// SortCMDVersion sort version
const SortCMDVersion = "0.0.1"

// SortCMD define `sort` cmd
var SortCMD = cli.Command{
	Name:    "sort",
	Aliases: []string{"SORT"},
	Usage:   "排序文本文件的文本行",
	Description: `Write sorted concatenation of all FILE(s) to standard output.

	如果没有指定文件，或者文件为"-"，则从标准输入读取。
	
	KEYDEF is F[.C][OPTS][,F[.C][OPTS]] for start and stop position, where F is a
field number and C a character position in the field; both are origin 1, and
the stop position defaults to the line's end.  If neither -t nor -b is in
effect, characters in a field are counted from the beginning of the preceding
whitespace.  OPTS is one or more single-letter ordering options [bdfgiMhnRrV],
which override global ordering options for that key.  If no key is given, use
the entire line as the key.  Use --debug to diagnose incorrect key usage.

SIZE may be followed by the following multiplicative suffixes:
% 1% of memory, b 1, K 1024 (default), and so on for M, G, T, P, E, Z, Y.

*** WARNING ***
The locale specified by the environment affects sort order.
Set LC_ALL=C to get the traditional sort order that uses
native byte values.`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "输出版本信息并推出",
		},
	},
	Action: func(c *cli.Context) error {
		if c.Bool("version") {
			fmt.Println(SortCMDVersion)
			return nil
		}
		return nil
	},
}
