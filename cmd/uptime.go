package cmd

// https://github.com/shirou/gopsutil/blob/master/host/host_linux.go

import (
	"fmt"

	"golang.org/x/sys/unix"

	"github.com/urfave/cli/v2"
)

// UptimeCMDVersion uptime version no
var UptimeCMDVersion = "v0.0.1"

// UptimeCMD define uptime cli command
var UptimeCMD = cli.Command{
	Name:    "uptime",
	Aliases: []string{"UPTIME"},
	Usage:   "显示系统运行了多长时间",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "output this help information and exit",
		}, &cli.BoolFlag{
			Name:    "pretty",
			Aliases: []string{"p"},
			Usage:   "show uptime in pretty format",
		}, &cli.StringFlag{
			Name:    "since",
			Aliases: []string{"s"},
			Usage:   "system up since",
		},
	},
	Action: uptimeAction,
}

func uptimeAction(c *cli.Context) error {
	if c.Bool("version") {
		fmt.Println(c.Command.Name, Base64CMDVersion)
		return nil
	}
	sysinfo := &unix.Sysinfo_t{}
	if err := unix.Sysinfo(sysinfo); err != nil {
		return err
	}
	println(sysinfo.Uptime)
	return nil
}
