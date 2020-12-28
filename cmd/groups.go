package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// GroupsCMDVersion groups version no
var GroupsCMDVersion = "v0.0.1"

// GroupsCMD define groups cli command
var GroupsCMD = cli.Command{
	Name:    "groups",
	Aliases: []string{"GROUPS"},
	Usage: `显示每个输入的用户名所在的全部组，
	如果没有指定用户名则默认为当前进程用户(当用户组数据库发生变更时可能导致差异)。`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "输出版本信息并推出",
		},
	},
	Action: groupsAction,
}

func groupsAction(c *cli.Context) error {
	if c.Bool("version") {
		fmt.Println(c.Command.Name, GroupsCMDVersion)
		return nil
	}
	return nil

}
