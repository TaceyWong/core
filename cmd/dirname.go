package cmd

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

// DirNameCMDVersion dirname version
const DirNameCMDVersion = "0.0.1"

// DirNameCMD define dirname cmd
var DirNameCMD = cli.Command{
	Name:      "dirname",
	Aliases:   []string{"DIRNAME"},
	UsageText: "core dirname [选项] 名称...",
	Usage:     "输出每个NAME的最后一个非斜杠组成部分，并删除尾随的斜杠",
	Description: `Output each NAME with its last non-slash component and trailing slashes
	removed; if NAME contains no /'s, output '.' (meaning the current directory).
	
	Examples:
		dirname /usr/bin/          -> "/usr"
		dirname dir1/str dir2/str  -> "dir1" followed by "dir2"
		dirname stdio.h            -> "."`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "输出版本信息并推出",
		}, &cli.BoolFlag{
			Name:    "zero",
			Aliases: []string{"z"},
			Usage:   "end each output line with NUL, not newline",
		},
	},
	Action: func(c *cli.Context) error {
		if c.Bool("version") {
			fmt.Println(DirNameCMDVersion)
			return nil
		}
		if c.Args().Len() == 0 {
			return errors.New("缺少操作对象")
		}
		end := "\n"
		if c.Bool("zero") {
			end = ""
		}
		for _, p := range c.Args().Slice() {
			fmt.Printf("%s%s", filepath.Dir(p), end)
		}
		return nil
	},
}
