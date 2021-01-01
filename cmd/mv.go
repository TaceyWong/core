package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// MVCMDVersion mv version no
var MVCMDVersion = "v0.0.1"

// MVCMD define mv cli command
var MVCMD = cli.Command{
	Name:    "mv",
	Aliases: []string{"MV"},
	Usage:   "文件重命名或文件剪切",
	Description: `The backup suffix is '~', unless set with --suffix or SIMPLE_BACKUP_SUFFIX.
	The version control method may be selected via the --backup option or through
	the VERSION_CONTROL environment variable.  Here are the values:
	
	  none, off       不进行备份(即使使用了--backup 选项)
	  numbered, t     备份文件加上数字进行排序
	  existing, nil   若有数字的备份文件已经存在则使用数字，否则使用普通方式备份
	  simple, never   永远使用普通方式备份`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "输出版本信息并推出",
		},
	},
	Action: mvAction,
}

func mvAction(c *cli.Context) error {
	if c.Bool("version") {
		fmt.Println(c.Command.Name, MVCMDVersion)
		return nil
	}
	return nil

}
