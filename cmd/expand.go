package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/urfave/cli/v2"
)

// ExpandCMDVersion expand version
const ExpandCMDVersion = "0.0.1"

// ExpandCMD define `expand` cmd
var ExpandCMD = cli.Command{
	Name:      "expand",
	Aliases:   []string{"EXPAND"},
	UsageText: "core expand [选项]... [文件]...",
	Usage:     "制表符转换为空格",
	Description: `Convert tabs in each FILE to spaces, writing to standard output.
	如果没有指定文件，或者文件为"-"，则从标准输入读取。`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "输出版本信息并推出",
		}, &cli.BoolFlag{
			Name:    "initial",
			Aliases: []string{"i"},
			Usage:   "do not convert tabs after non blanks",
		}, &cli.IntFlag{
			Name:    "tabs",
			Aliases: []string{"t"},
			Usage:   "have tabs `N` characters apart, not 8",
			Value:   8,
		},
	},
	Action: expandAction,
}

func expandAction(c *cli.Context) error {
	if c.Bool("version") {
		fmt.Println(ExpandCMDVersion)
		return nil
	}
	if c.Args().Len() == 0 {
		return errors.New("尚未实现从标准输入读取")
	}
	for _, p := range c.Args().Slice() {
		m, err := os.Stat(p)
		if err != nil {
			return err
		}
		if m.IsDir() {
			return fmt.Errorf("'%s'是目录而非文件", p)
		}
	}
	for _, p := range c.Args().Slice() {
		content, err := ioutil.ReadFile(p)
		if err != nil {
			return err
		}
		fmt.Println(string(content))

	}

	return nil
}

// Expand tab -> space
func Expand(text string) (result string) {
	return result
}
