package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// NiceCMDVersion nice version
const NiceCMDVersion = "0.0.1"

// NiceCMD define `nice` cmd
var NiceCMD = cli.Command{
	Name:    "nice",
	Aliases: []string{"NICE"},
	Usage:   "修改调度优先级运行程序",
	Description: `Run COMMAND with an adjusted niceness, which affects process scheduling.
	With no COMMAND, print the current niceness.  Niceness values range from
	-20 (most favorable to the process) to 19 (least favorable to the process).
	
	`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "输出版本信息并推出",
		}, &cli.IntFlag{
			Name:    "adjustment",
			Aliases: []string{"n"},
			Usage:   "add integer `N` to the niceness",
			Value:   10,
		},
	},
	Action: func(c *cli.Context) error {
		if c.Bool("version") {
			fmt.Println(NiceCMDVersion)
			return nil
		}
		return nil
	},
}
