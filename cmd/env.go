package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// EnvCMDVersion env version no
var EnvCMDVersion = "v0.0.1"

// EnvCMD define env cli command
var EnvCMD = cli.Command{
	Name:    "env",
	Aliases: []string{"ENV"},
	Usage:   "在修改后的环境中运行程序",
	Description: `Set each NAME to VALUE in the environment and run COMMAND
	单纯的 - 意味着 -i。如果没有命令，则打印结果环境。`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "output version information and exit",
		}, &cli.BoolFlag{
			Name:    "ignore-environment",
			Aliases: []string{"i"},
			Usage:   `start with an empty environment`,
		}, &cli.BoolFlag{
			Name:    "null",
			Aliases: []string{"0"},
			Usage:   `end each output line with NUL, not newline`,
		}, &cli.StringFlag{
			Name:    "unset",
			Aliases: []string{"u"},
			Usage:   `remove variable from the environment`,
		}, &cli.PathFlag{
			Name:    `chdir`,
			Aliases: []string{"C"},
			Usage:   `change working directory to DIR`,
		},
	},
	Action: envAction,
}

func envAction(c *cli.Context) error {
	if c.Bool("version") {
		fmt.Println(c.Command.Name, EnvCMDVersion)
		return nil
	}

	return nil

}
