package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// ChOwnCMDVersion chown version
const ChOwnCMDVersion = "0.0.1"

// ChOwnCMD define `chown` cmd
var ChOwnCMD = cli.Command{
	Name:    "chown",
	Aliases: []string{"CHOWN"},
	Usage:   "更改文件所有者和组",
	Description: `Change the owner and/or group of each FILE to OWNER and/or GROUP.
	With --reference, change the owner and group of each FILE to those of RFILE.
	
	Owner is unchanged if missing.  Group is unchanged if missing, but changed
	to login group if implied by a ':' following a symbolic OWNER.
	OWNER and GROUP may be numeric as well as symbolic.

	示例：
		chown root /u         将 /u 的属主更改为"root"。
		chown root:staff /u   和上面类似，但同时也将其属组更改为"staff"。
		chown -hR root /u     将 /u 及其子目录下所有文件的属主更改为"root"。`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "output version information and exit",
		},
	},
	Action: func(c *cli.Context) error {
		if c.Bool("version") {
			fmt.Println(ChOwnCMDVersion)
			return nil
		}
		return nil
	},
}
