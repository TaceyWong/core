package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

// RealPathCMDVersion realpath version
const RealPathCMDVersion = "0.0.1"

// RealPathCMD define `realpath` cmd
var RealPathCMD = cli.Command{
	Name:      "realpath",
	Aliases:   []string{"REALPATH"},
	UsageText: "core realpath [选项]... 文件...",
	Usage:     "打印已解析的路径[绝对地址]",
	Description: `Print the resolved absolute file name;
	all but the last component must exist
	相较于原版，core realpath不进行文件是否存在的校验
	`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "输出版本信息并推出",
		}, &cli.BoolFlag{
			Name:    "canonicalize-existing",
			Aliases: []string{"e"},
			Usage:   "all components of the path must exist",
		},
		&cli.BoolFlag{
			Name:    "canonicalize-missin",
			Aliases: []string{"m"},
			Usage:   "no path components need exist or be a directory",
		},
		&cli.BoolFlag{
			Name:    "logical",
			Aliases: []string{"L"},
			Usage:   "resolve '..' components before symlinks",
		},
		&cli.BoolFlag{
			Name:    "physical",
			Aliases: []string{"P"},
			Usage:   "resolve symlinks as encountered (default)",
		},
		&cli.BoolFlag{
			Name:    "quiet",
			Aliases: []string{"q"},
			Usage:   "suppress most error messages",
		}, &cli.PathFlag{
			Name:  "relative-to",
			Usage: "print the resolved path relative to `DIR`",
		}, &cli.PathFlag{
			Name:  "relative-base",
			Usage: "print absolute paths unless paths below `DIR`",
		}, &cli.BoolFlag{
			Name:    "strip",
			Aliases: []string{"s"},
			Usage:   "don't expand symlinks",
		}, &cli.BoolFlag{
			Name:  "no-symlinks",
			Usage: "the same as --strip",
		}, &cli.BoolFlag{
			Name:    "zero",
			Aliases: []string{"z"},
			Usage:   "end each output line with NUL, not newline",
		},
	},
	Action: func(c *cli.Context) (err error) {
		if c.Bool("version") {
			fmt.Println(RealPathCMDVersion)
			return nil
		}
		if c.Args().Len() == 0 {
			cli.ShowCommandHelp(c, c.Command.Name)
			return fmt.Errorf("%s: 缺少操作数", c.Command.Name)
		}
		base := "/"
		end := "\n"
		if b := c.Path("relative-to"); b != "" {
			base = b
		}
		if c.Bool("zero") {
			end = ""
		}
		for _, p := range c.Args().Slice() {
			abs, _ := filepath.Abs(p)

			result, err := filepath.Rel(base, abs)
			if err != nil {
				return err
			}
			fmt.Printf("%s%s", result, end)
		}
		return err
	},
}
