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
	Description: `
	By default, selects -v1 -i1 -l1 -sTAB -w6 -nrn -hn -bt -fn.
	CC are two delimiter characters used to construct logical page delimiters,
	a missing second character implies :.  Type \\ for \.  STYLE is one of:
	
	  a     对所有行编号
	  t     对非空行编号
	  n     不编行号
	  pBRE  只对符合正则表达式BRE 的行编号
	
	FORMAT 是下列之一:
	
	  ln    左对齐，空格不用0 填充
	  rn    右对齐，空格不用0 填充
	  rz    右对齐，空格用0 填充`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "输出版本信息并推出",
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
		}, &cli.StringFlag{
			Name:    "header-numbering",
			Aliases: []string{"he"},
			Usage:   "use STYLE for numbering header lines",
		}, &cli.Int64Flag{
			Name:    "line-increment",
			Aliases: []string{"i"},
			Usage:   "line number increment at each line",
		}, &cli.Int64Flag{
			Name:    "join-blank-lines",
			Aliases: []string{"l"},
			Usage:   "group of NUMBER empty lines counted as one",
		}, &cli.StringFlag{
			Name:    "number-format",
			Aliases: []string{"n"},
			Usage:   "insert line numbers according to FORMAT",
		}, &cli.BoolFlag{
			Name:    "no-renumber",
			Aliases: []string{"p"},
			Usage:   "do not reset line numbers for each section",
		}, &cli.StringFlag{
			Name:    "number-separator",
			Aliases: []string{"s"},
			Usage:   "add STRING after (possible) line number",
		}, &cli.Int64Flag{
			Name:    "starting-line-number",
			Aliases: []string{"sl"},
			Usage:   "first line number for each section",
		}, &cli.Int64Flag{
			Name:    "number-width",
			Aliases: []string{"w"},
			Usage:   "use NUMBER columns for line numbers",
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
