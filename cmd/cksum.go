package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// CkSumCMDVersion cksum version
const CkSumCMDVersion = "0.0.1"

// CkSumCMD define `cksum` cmd
var CkSumCMD = cli.Command{
	Name:      "cksum",
	Aliases:   []string{"CKSUM"},
	UsageText: "core cksum [文件]...\ncore cksum [选项]",
	Usage:     "输出每个文件的 CRC 校验值和字节统计。",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "输出版本信息并推出",
		},
	},
	Action: func(c *cli.Context) error {
		if c.Bool("version") {
			fmt.Println(CkSumCMDVersion)
			return nil
		}
		return nil
	},
}
