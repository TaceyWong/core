package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// YesCMDVersion yes version
const YesCMDVersion = "0.0.1"

// YesCMD define `yes` cmd
var YesCMD = cli.Command{
	Name:    "yes",
	Aliases: []string{"YES"},
	Usage:   "重复输出一个字符串直到被kill",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "输出版本信息并推出",
		},
	},
	Action: func(c *cli.Context) error {
		if c.Bool("version") {
			fmt.Println(TrueCMDVersion)
			return nil
		}
		strContent := "y"
		if c.Args().Get(0) != "" {
			strContent = c.Args().Get(0)
		}
		for {
			fmt.Println(strContent)
		}
	},
}
