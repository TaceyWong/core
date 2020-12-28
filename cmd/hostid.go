package cmd

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

// HostidCMDVersion hostid version no
var HostidCMDVersion = "v0.0.1"

// HostidCMD define hostid cli command
var HostidCMD = cli.Command{
	Name:    "hostid",
	Aliases: []string{"HOSTID"},
	Usage:   "print the numeric identifier for the current host",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "输出版本信息并推出",
		},
	},
	Action: hostid,
}

func hostid(c *cli.Context) error {
	if c.Bool("version") {
		fmt.Println(c.Command.Name, ArchCMDVersion)
		return nil
	}
	// https://github.com/MaiZure/coreutils-8.3/blob/master/src/hostid.c
	// `man gethostid` -> It is impossible to ensure that the identifier is globally unique.
	// -> https://code.woboq.org/userspace/glibc/sysdeps/unix/sysv/linux/gethostid.c.html
	hi := os.Hostname // temp instead
	hn, err := hi()
	if err != nil {
		return err
	}
	fmt.Println(hn)
	return nil
}
