package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// TeeCMDVersion tee version
const TeeCMDVersion = "0.0.1"

// TeeCMD define `tee` cmd
var TeeCMD = cli.Command{
	Name:    "tee",
	Aliases: []string{"TEE"},
	Usage:   "标准输入复制到每个指定文件，并显示到标准输出。",
	Description: `MODE determines behavior with write errors on the outputs:
	'warn'         diagnose errors writing to any output
	'warn-nopipe'  diagnose errors writing to any output not a pipe
	'exit'         exit on error writing to any output
	'exit-nopipe'  exit on error writing to any output not a pipe
  The default MODE for the -p option is 'warn-nopipe'.
  The default operation when --output-error is not specified, is to
  exit immediately on error writing to a pipe, and diagnose errors
  writing to non pipe outputs.`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "output version information and exit",
		},
	},
	Action: func(c *cli.Context) error {
		if c.Bool("version") {
			fmt.Println(TeeCMDVersion)
			return nil
		}
		return nil
	},
}
