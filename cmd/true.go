package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// TrueCMDVersion true version
const TrueCMDVersion = "0.0.1"

// TrueCMD define `true` command
var TrueCMD = cli.Command{
	Name:        "true",
	Aliases:     []string{"TRUE"},
	Usage:       "什么都不做，成功执行空白进程",
	Description: "会返回一个状态码表示成功执行",
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
		return nil
	},
}
