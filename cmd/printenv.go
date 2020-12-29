package cmd

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

// PrintEnvCMDVersion printenv version no
var PrintEnvCMDVersion = "v0.0.1"

// PrintEnvCMD define printenv cli command
var PrintEnvCMD = cli.Command{
	Name:    "printenv",
	Aliases: []string{"PRINTENV"},
	Usage: `显示指定的环境变量的值。
	如果没有指定变量，则打印出所有变量的名称和值。`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "输出版本信息并推出",
		}, &cli.BoolFlag{
			Name:    "null",
			Aliases: []string{"0"},
			Usage:   `end each output line with NUL, not newline`,
		},
	},
	Action: printenvAction,
}

func printenvAction(c *cli.Context) error {
	if c.Bool("version") {
		fmt.Println(c.Command.Name, PrintEnvCMDVersion)
		return nil
	}
	end := "\n"
	if c.Bool("0") {
		end = ""
	}
	envs := os.Environ()
	for _, env := range envs {
		fmt.Printf("%s%s", env, end)
	}
	return nil

}
