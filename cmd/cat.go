package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// CatCMDVersion cat version no
var CatCMDVersion = "v0.0.1"

// CatCMD define cat cli command
var CatCMD = cli.Command{
	Name:    "cat",
	Aliases: []string{"CAT"},
	Usage:   "连接所有指定文件并将结果写到标准输出",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "显示版本信息并退出",
		}, &cli.BoolFlag{
			Name:    "show-all",
			Aliases: []string{"A"},
			Usage:   "等同于-vET",
		}, &cli.BoolFlag{
			Name:    "number-nonblank",
			Aliases: []string{"b"},
			Usage:   "number nonempty output lines, overrides -n",
		}, &cli.BoolFlag{
			Name:    "show-ends",
			Aliases: []string{"E"},
			Usage:   "display $ at end of each line",
		}, &cli.BoolFlag{
			Name:    "number",
			Aliases: []string{"n"},
			Usage:   "number all output lines",
		},
	},
	Action: catAction,
}

func catAction(c *cli.Context) error {
	if c.Bool("version") {
		fmt.Println(c.Command.Name, Base64CMDVersion)
		return nil
	}
	target := c.Args().Get(0)
	println(target)
	return nil
}
