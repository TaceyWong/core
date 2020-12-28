package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// TestCMDVersion yes version
const TestCMDVersion = "0.0.1"

// TestCMD define `test` cmd
var TestCMD = cli.Command{
	Name:    "test",
	Aliases: []string{"TEST"},
	Usage:   "check file types and compare values",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "输出版本信息并推出",
		},
	},
	Action: func(c *cli.Context) error {
		if c.Bool("version") {
			fmt.Println(TestCMDVersion)
			return nil
		}
		return nil
	},
}
