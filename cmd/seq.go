package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// SeqCMDVersion seq version no
var SeqCMDVersion = "v0.0.1"

// SeqCMD define seq cli command
var SeqCMD = cli.Command{
	Name:    "seq",
	Aliases: []string{"SEQ"},
	Usage:   "打印数字序列",
	Description: `Print numbers from FIRST to LAST, in steps of INCREMENT.
	If FIRST or INCREMENT is omitted, it defaults to 1.  That is, an
	omitted INCREMENT defaults to 1 even when LAST is smaller than FIRST.
	The sequence of numbers ends when the sum of the current number and
	INCREMENT would become greater than LAST.
	FIRST, INCREMENT, and LAST are interpreted as floating point values.
	INCREMENT is usually positive if FIRST is smaller than LAST, and
	INCREMENT is usually negative if FIRST is greater than LAST.
	INCREMENT must not be 0; none of FIRST, INCREMENT and LAST may be NaN.
	FORMAT must be suitable for printing one argument of type 'double';
	it defaults to %.PRECf if FIRST, INCREMENT, and LAST are all fixed point
	decimal numbers with maximum precision PREC, and to %g otherwise.`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "输出版本信息并推出",
		}, &cli.StringFlag{
			Name:    "format",
			Aliases: []string{"f"},
			Usage:   "use printf style floating-point FORMAT",
		}, &cli.StringFlag{
			Name:    "separator",
			Aliases: []string{"s"},
			Usage:   "use STRING to separate numbers (default: \\n)",
		}, &cli.BoolFlag{
			Name:    "equal-width",
			Aliases: []string{"w"},
			Usage:   "equalize width by padding with leading zeroes",
		},
	},
	Action: seqAction,
}

func seqAction(c *cli.Context) error {
	if c.Bool("version") {
		fmt.Println(c.Command.Name, Base64CMDVersion)
		return nil
	}
	return nil
}
