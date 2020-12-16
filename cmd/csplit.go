package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// CSplitCMDVersion csplit version
const CSplitCMDVersion = "0.0.1"

// CSplitCMD define `csplit` cmd
var CSplitCMD = cli.Command{
	Name:    "csplit",
	Aliases: []string{"CSPLIT"},
	Usage:   "Output pieces of FILE separated by PATTERN(s) to files 'xx00', 'xx01', ...,",
	Description: `Output pieces of FILE separated by PATTERN(s) to files 'xx00', 'xx01', ...,
	and output byte counts of each piece to standard output.
	
	Read standard input if FILE is -
	
	Each PATTERN may be:
		INTEGER            copy up to but not including specified line number
		/REGEXP/[OFFSET]   copy up to but not including a matching line
		%REGEXP%[OFFSET]   skip to, but not including a matching line
		{INTEGER}          repeat the previous pattern specified number of times
		{*}                repeat the previous pattern as many times as possible

	A line OFFSET is a required '+' or '-' followed by a positive integer.`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "output version information and exit",
		},
	},
	Action: func(c *cli.Context) error {
		if c.Bool("version") {
			fmt.Println(CSplitCMDVersion)
			return nil
		}
		return nil
	},
}
