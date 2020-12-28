package cmd

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

// FalseCMDVersion false version
const FalseCMDVersion = "0.0.1"

// FalseCMD define `false` command
var FalseCMD = cli.Command{
	Name:        "false",
	Aliases:     []string{"FALSE"},
	Usage:       "返回布尔False",
	Description: "遵照IEEE Std 1003.2-1992 (``POSIX.2'') 总以非零退出码退出,",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "输出版本信息并推出",
		},
	},
	Action: func(c *cli.Context) error {
		if c.Bool("version") {
			fmt.Println(TrueCMDVersion)
			return nil
		}
		os.Exit(1)
		return nil
	},
}
