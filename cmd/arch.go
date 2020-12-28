package cmd

import (
	"fmt"
	"syscall"

	"github.com/urfave/cli/v2"
)

// ArchCMDVersion arch version no
var ArchCMDVersion = "v0.0.1"

// ArchCMD define arch cli command
var ArchCMD = cli.Command{
	Name:    "arch",
	Aliases: []string{"ARCH"},
	Usage:   "打印机器硬件名称 (同 uname -m)",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "输出版本信息并推出",
		},
	},
	Action: archAction,
}

func archAction(c *cli.Context) error {
	if c.Bool("version") {
		fmt.Println(c.Command.Name, ArchCMDVersion)
		return nil
	}
	utsname := syscall.Utsname{}
	if err := syscall.Uname(&utsname); err != nil {
		return err
	}
	fmt.Println(ctos(&utsname.Machine))
	return nil

}
