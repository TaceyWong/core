package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// CommCMDVersion comm version no
var CommCMDVersion = "v0.0.1"

// CommCMD define comm cli command
var CommCMD = cli.Command{
	Name:    "comm",
	Aliases: []string{"COMM"},
	Usage:   "逐行比较已排序的文件文件1 和文件2",
	Description: `When FILE1 or FILE2 (not both) is -, read standard input.

	如果不附带选项，程序会生成三列输出。第一列包含文件1 特有的行，第二列包含 文件2 特有的行，
	而第三列包含两个文件共有的行。
	
Note, comparisons honor the rules specified by 'LC_COLLATE'.

示例：
  comm -12 文件1 文件2  只打印在文件1 和文件2 中都有的行
  comm -3  文件1 文件2  打印在文件1 中有，而文件2 中没有的行。反之亦然。`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "输出版本信息并推出",
		},
	},
	Action: commAction,
}

func commAction(c *cli.Context) error {
	if c.Bool("version") {
		fmt.Println(c.Command.Name, CommCMDVersion)
		return nil
	}
	return nil

}
