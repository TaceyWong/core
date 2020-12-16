package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// ChGrpCMDVersion chgrp version
const ChGrpCMDVersion = "0.0.1"

// ChGrpCMD define `chgrp` cmd
var ChGrpCMD = cli.Command{
	Name:    "chgrp",
	Aliases: []string{"CHGRP"},
	Usage:   "更改组所有权",
	Description: `Change the group of each FILE to GROUP.
	With --reference, change the group of each FILE to that of RFILE.
	
	示例：
  		chgrp staff /u            将 /u 的属组更改为"staff"。
  		chgrp -hR staff /u    将 /u 及其子目录下所有文件的属组更改为"staff"。
	`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "output version information and exit",
		},
	},
	Action: func(c *cli.Context) error {
		if c.Bool("version") {
			fmt.Println(ChGrpCMDVersion)
			return nil
		}
		return nil
	},
}
