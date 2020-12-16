package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// DirNameCMDVersion dirname version
const DirNameCMDVersion = "0.0.1"

// DirNameCMD define dirname cmd
var DirNameCMD = cli.Command{
	Name:    "dirname",
	Aliases: []string{"DIRNAME"},
	Usage:   "Output each NAME with its last non-slash component and trailing slashes removed",
	Description: `Output each NAME with its last non-slash component and trailing slashes
	removed; if NAME contains no /'s, output '.' (meaning the current directory).
	
	Examples:
		dirname /usr/bin/          -> "/usr"
		dirname dir1/str dir2/str  -> "dir1" followed by "dir2"
		dirname stdio.h            -> "."`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "output version information and exit",
		},
	},
	Action: func(c *cli.Context) error {
		if c.Bool("version") {
			fmt.Println(DirNameCMDVersion)
			return nil
		}
		return nil
	},
}
