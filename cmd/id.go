package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// IDCMDVersion id version no
var IDCMDVersion = "v0.0.1"

// IDCMD define id cli command
var IDCMD = cli.Command{
	Name:    "id",
	Aliases: []string{"ID"},
	Usage: `Print user and group information for the specified USER,
	or (when USER omitted) for the current user.`,
	Description: `如果不附带任何选项，程序会显示一些可供识别用户身份的有用信息。`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "output version information and exit",
		}, &cli.BoolFlag{
			Name:    "context",
			Aliases: []string{"Z"},
			Usage:   `print only the security context of the process`,
		}, &cli.BoolFlag{
			Name:    "group",
			Aliases: []string{"g"},
			Usage:   `print only the effective group ID`,
		}, &cli.BoolFlag{
			Name:    "groups",
			Aliases: []string{"G"},
			Usage:   `print all group IDs`,
		}, &cli.BoolFlag{
			Name:    "name",
			Aliases: []string{"n"},
			Usage:   `print a name instead of a number, for -ugG`,
		}, &cli.BoolFlag{
			Name:    "real",
			Aliases: []string{"r"},
			Usage:   ` print the real ID instead of the effective ID, with -ugG`,
		}, &cli.BoolFlag{
			Name:    "user",
			Aliases: []string{"u"},
			Usage:   `print only the effective user ID`,
		}, &cli.BoolFlag{
			Name:    "zero",
			Aliases: []string{"z"},
			Usage: `delimit entries with NUL characters, not whitespace;
			not permitted in default format`,
		},
	},
	Action: idAction,
}

func idAction(c *cli.Context) error {
	if c.Bool("version") {
		fmt.Println(c.Command.Name, IDCMDVersion)
		return nil
	}
	return nil

}
