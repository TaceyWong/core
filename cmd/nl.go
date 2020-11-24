package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// NlCMDVersion nl version no
var NlCMDVersion = "v0.0.1"

// NlCMD define nl cli command
var NlCMD = cli.Command{
	Name:    "nl",
	Aliases: []string{"NL"},
	Usage:   "Write each FILE to standard output, with line numbers added.",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "output version information and exit",
		}, &cli.StringFlag{
			Name:    "body-numbering",
			Aliases: []string{"b"},
			Usage:   "use STYLE for numbering body lines",
		}, &cli.StringFlag{
			Name:    "section-delimiter",
			Aliases: []string{"d"},
			Usage:   "use CC for logical page delimiters",
		}, &cli.StringFlag{
			Name:    "footer-numbering",
			Aliases: []string{"f"},
			Usage:   "use STYLE for numbering footer lines",
		},
	},
	Action: nlAction,
}

func nlAction(c *cli.Context) error {
	if c.Bool("version") {
		fmt.Println(c.Command.Name, NlCMDVersion)
		return nil
	}
	target := c.Args().Get(0)
	println(target)

	return nil

}
