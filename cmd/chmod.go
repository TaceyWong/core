package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// ChModCMDVersion chmod version
const ChModCMDVersion = "0.0.1"

// ChModCMD define `chmod` cmd
var ChModCMD = cli.Command{
	Name:    "chmod",
	Aliases: []string{"CHMOD"},
	Usage:   "更改文件模式位",
	Description: `Change the mode of each FILE to MODE.
	With --reference, change the mode of each FILE to that of RFILE.
	
	Each MODE is of the form '[ugoa]*([-+=]([rwxXst]*|[ugo]))+|[-+=][0-7]+'.`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "output version information and exit",
		},
	},
	Action: func(c *cli.Context) error {
		if c.Bool("version") {
			fmt.Println(ChModCMDVersion)
			return nil
		}
		return nil
	},
}
