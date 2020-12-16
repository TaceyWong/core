package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// RealPathCMDVersion realpath version
const RealPathCMDVersion = "0.0.1"

// RealPathCMD define `realpath` cmd
var RealPathCMD = cli.Command{
	Name:    "realpath",
	Aliases: []string{"REALPATH"},
	Usage:   "打印已解析的路径[绝对地址]",
	Description: `Print the resolved absolute file name;
	all but the last component must exist`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "output version information and exit",
		},
	},
	Action: func(c *cli.Context) error {
		if c.Bool("version") {
			fmt.Println(RealPathCMDVersion)
			return nil
		}
		return nil
	},
}
