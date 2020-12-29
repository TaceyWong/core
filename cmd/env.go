package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/urfave/cli/v2"
)

// EnvCMDVersion env version no
var EnvCMDVersion = "v0.0.1"

// EnvCMD define env cli command
var EnvCMD = cli.Command{
	Name:      "env",
	Aliases:   []string{"ENV"},
	UsageText: "core env [选项]... [-] [名称=值]... [命令 [参数]...]",
	Usage:     "在修改后的环境中运行程序",
	Description: `Set each NAME to VALUE in the environment and run COMMAND
	单纯的 - 意味着 -i。如果没有命令，则打印结果环境。`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "输出版本信息并推出",
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

func envAction(c *cli.Context) (err error) {
	if c.Bool("version") {
		fmt.Println(c.Command.Name, EnvCMDVersion)
		return nil
	}
	end := "\n"
	if c.Bool("0") {
		end = ""
	}
	// flag i
	if c.Bool("ignore-environment") {
		os.Clearenv()
	}
	// flag unset
	if unset := c.String("unset"); unset != "" {
		os.Unsetenv(unset)
	}

	targetCMD := ""
	argsS := c.Args().Slice()
	for i, arg := range argsS {
		if j := strings.Index(arg, "="); j != -1 {
			os.Setenv(arg[:j], arg[j+1:len(arg)])
		} else {
			targetCMD = strings.Join(argsS[i:len(argsS)], " ")
			break
		}
	}
	if targetCMD == "" {
		PrintENV(end)
		return err
	}

	cwd, _ := os.Getwd()
	chDir := cwd
	if chdir := c.Path("chdir"); chdir != "" {
		if targetCMD == "" {
			cli.ShowCommandHelp(c, c.Command.Name)
			return errors.New("must specify command with --chdir (-C)")
		}
		if m, e := os.Stat(chdir); e != nil || !m.IsDir() {
			cli.ShowCommandHelp(c, c.Command.Name)
			return fmt.Errorf("cannot change directory to '%s': 没有那个文件或目录", chdir)
		}
		chDir = chdir
	}
	defer func(cwd string) {
		recover()
		os.Chdir(cwd)
	}(cwd)
	os.Chdir(chDir)
	// 缺少which操作
	cmd := exec.Command("bash", "-c", targetCMD)
	// PrintENV("\n")
	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Run()
	return nil

}
