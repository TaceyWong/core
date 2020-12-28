package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// PathChkCMDVersion pathchk version
const PathChkCMDVersion = "0.0.1"

// PathChkCMD define `pathchk` cmd
var PathChkCMD = cli.Command{
	Name:    "pathchk",
	Aliases: []string{"PATHCHK"},
	Usage:   "check whether file names are valid or portable",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "输出版本信息并推出",
		},
	},
	Action: func(c *cli.Context) error {
		if c.Bool("version") {
			fmt.Println(PathChkCMDVersion)
			return nil
		}
		return nil
	},
}
