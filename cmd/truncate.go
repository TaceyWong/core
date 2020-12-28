package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// TruncateCMDVersion truncate version
const TruncateCMDVersion = "0.0.1"

// TruncateCMD define `truncate` cmd
var TruncateCMD = cli.Command{
	Name:    "truncate",
	Aliases: []string{"TRUNCATE"},
	Usage:   "shrink or extend the size of a file to the specified size",
	Description: `A FILE argument that does not exist is created.

	If a FILE is larger than the specified size, the extra data is lost.
	If a FILE is shorter, it is extended and the extended part (hole)
	reads as zero bytes.
	
	The SIZE argument is an integer and optional unit (example: 10K is 10*1024).
	Units are K,M,G,T,P,E,Z,Y (powers of 1024) or KB,MB,... (powers of 1000).

	SIZE may also be prefixed by one of the following modifying characters:
	'+' extend by, '-' reduce by, '<' at most, '>' at least,
	'/' round down to multiple of, '%' round up to multiple of.`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "输出版本信息并推出",
		},
	},
	Action: func(c *cli.Context) error {
		if c.Bool("version") {
			fmt.Println(TruncateCMDVersion)
			return nil
		}
		return nil
	},
}
