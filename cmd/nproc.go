package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// NProcCMDVersion nproc version
const NProcCMDVersion = "0.0.1"

// NProcCMD define `nproc` cmd
var NProcCMD = cli.Command{
	Name:    "nproc",
	Aliases: []string{"NPROC"},
	Usage:   "print the number of processing units available",
	Description: `此数目可能小于实际工作数`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "output version information and exit",
		},
	},
	Action: func(c *cli.Context) error {
		if c.Bool("version") {
			fmt.Println(NProcCMDVersion)
			return nil
		}
		return nil
	},
}
