package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// TtyCMDVersion tty version no
var TtyCMDVersion = "v0.0.1"

// TtyCMD define tty cli command
var TtyCMD = cli.Command{
	Name:    "tty",
	Aliases: []string{"TTY"},
	Usage:   `显示出连接到当前标准输入的终端设备文件名。`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "输出版本信息并推出",
		}, &cli.BoolFlag{
			Name:    "silent",
			Aliases: []string{"s"},
			Usage:   "什么也不显示，只返回退出状态值",
		}, &cli.BoolFlag{
			Name:  "quiet",
			Usage: "同--silent",
		},
	},
	Action: ttyAction,
}

func ttyAction(c *cli.Context) error {
	if c.Bool("version") {
		fmt.Println(c.Command.Name, TtyCMDVersion)
		return nil
	}
	return nil

}
