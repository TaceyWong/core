package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// FactorCMDVersion factor version
const FactorCMDVersion = "0.0.1"

// FactorCMD define `factor` cmd
var FactorCMD = cli.Command{
	Name:    "factor",
	Aliases: []string{"FACTOR"},
	Usage:   "输出每个指定的数字的素因子，如果没有在命令行中指定则从标准输入读取。",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "output version information and exit",
		},
	},
	Action: func(c *cli.Context) error {
		if c.Bool("version") {
			fmt.Println(FactorCMDVersion)
			return nil
		}
		return nil
	},
}
