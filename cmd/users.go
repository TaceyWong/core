package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// UsersCMDVersion users version no
var UsersCMDVersion = "v0.0.1"

// UsersCMD define users cli command
var UsersCMD = cli.Command{
	Name:    "users",
	Aliases: []string{"USERS"},
	Usage: `根据文件判断输出当前有谁正登录在系统上。
	如果文件未予指定，则使用/var/run/utmp，/var/log/wtmp 是通用的相关文件`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "output version information and exit",
		},
	},
	Action: usersAction,
}

func usersAction(c *cli.Context) error {
	if c.Bool("version") {
		fmt.Println(c.Command.Name, UsersCMDVersion)
		return nil
	}
	return nil

}
