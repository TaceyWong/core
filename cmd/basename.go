package cmd

import (
	"fmt"
	"path"
	"strings"

	"github.com/urfave/cli/v2"
)

// BaseNameCMDVersion basename version no
var BaseNameCMDVersion = "v0.0.1"

// BaseNameCMD define basename cli command
var BaseNameCMD = cli.Command{
	Name:    "basename",
	Aliases: []string{"BASENAME"},
	Usage:   "移除文件名的目录和后缀",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "output this help information and exit",
		}, &cli.BoolFlag{
			Name:    "multiple",
			Aliases: []string{"a"},
			Usage:   "support multiple arguments and treat each as a NAME",
		}, &cli.StringFlag{
			Name:    "suffix",
			Aliases: []string{"s"},
			Value:   "/",
			Usage:   "remove a trailing SUFFIX; implies -a",
		}, &cli.BoolFlag{
			Name:    "zero",
			Aliases: []string{"z"},
			Usage:   "end each output line with NUL, not newline",
		},
	},
	Action: baseNameAction,
}

func baseNameAction(c *cli.Context) error {
	if c.Bool("version") {
		fmt.Println(c.Command.Name, Base64CMDVersion)
		return nil
	}
	target := c.Args().Get(0)
	if c.String("suffix") != "/" {
		target = strings.TrimSuffix(target, c.String("suffix"))
	}
	result := path.Base(target)
	if c.Bool("zero") {
		print(result)
	}else{
		println(path.Base(target))
	}
	return nil
}
