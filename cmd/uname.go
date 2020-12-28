package cmd

import (
	"fmt"
	"syscall"

	"github.com/urfave/cli/v2"
)

// UNameCMDVersion version number
const UNameCMDVersion = "0.0.1"

// UNameCMD define `uname` cli command
var UNameCMD = cli.Command{
	Name:        "uname",
	Aliases:     []string{"UNAME"},
	Usage:       "打印系统信息",
	Description: `Print certain system information.  With no OPTION, same as -s.`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "输出版本信息并推出",
		},
		&cli.BoolFlag{
			Name:    "machine",
			Aliases: []string{"m"},
			Usage:   "print the machine hardware name",
		},
	},
	Action: Uname,
}

func ctos(input *[65]int8) string {
	length := 0
	str := make([]byte, len(input))
	for ; length < len(input); length++ {
		if input[length] == 0 {
			break
		}
		str[length] = uint8(input[length])
	}
	return string(str[0:length])
}

// Uname  real action
func Uname(c *cli.Context) error {
	utsname := syscall.Utsname{}
	if err := syscall.Uname(&utsname); err != nil {
		return err
	}
	if c.Bool("version") {
		fmt.Println(c.Command.Name, UNameCMDVersion)
		return nil
	}
	if c.Bool("m") {
		fmt.Println(ctos(&utsname.Machine))
	}
	return nil
}
