package cmd

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

// HostnameCMDVersion hostname version no
var HostnameCMDVersion = "v0.0.1"

// HostnameCMD define hostname cli command
var HostnameCMD = cli.Command{
	Name:    "hostname",
	Aliases: []string{"HOSTNAME"},
	Usage:   "显示或设置系统主机名",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "输出版本信息并推出",
		},
	},
	Action: hostname,
}

func hostname(c *cli.Context) error {
	if c.Bool("version") {
		fmt.Println(c.Command.Name, ArchCMDVersion)
		return nil
	}
	// /proc/sys/kernel/hostname | android -> localhost,
	// more detail golang-source-code
	hn, err := os.Hostname()
	if err != nil {
		return err
	}
	fmt.Println(hn)
	return nil
}
