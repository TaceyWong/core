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
	Aliases: []string{"machine_arch"},
	Usage:   "print machine hardware name (same as uname -m)",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "output version information and exit",
		},
	},
	Action: arch,
}

func arch(c *cli.Context) error {
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
