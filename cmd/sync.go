package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// SyncCMDVersion sync version
const SyncCMDVersion = "0.0.1"

// SyncCMD define `sync` cmd
var SyncCMD = cli.Command{
	Name:    "sync",
	Aliases: []string{"SYNC"},
	Usage:   "Synchronize cached writes to persistent storage",
	Description: `If one or more files are specified, sync only them,
	or their containing file systems.`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "输出版本信息并推出",
		},
	},
	Action: func(c *cli.Context) error {
		if c.Bool("version") {
			fmt.Println(SyncCMDVersion)
			return nil
		}
		return nil
	},
}
