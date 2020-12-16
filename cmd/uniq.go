package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// UniqCMDVersion uniq version
const UniqCMDVersion = "0.0.1"

// UniqCMD define `uniq` cmd
var UniqCMD = cli.Command{
	Name:    "uniq",
	Aliases: []string{"UNIQ"},
	Usage: `Filter adjacent matching lines from INPUT (or standard input),
	writing to OUTPUT (or standard output)`,
	Description: `Note: 'uniq' does not detect repeated lines unless they are adjacent.
	You may want to sort the input first, or use 'sort -u' without 'uniq'.
	Also, comparisons honor the rules specified by 'LC_COLLATE'.`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "output version information and exit",
		},
	},
	Action: func(c *cli.Context) error {
		if c.Bool("version") {
			fmt.Println(TrueCMDVersion)
			return nil
		}
		return nil

	},
}
