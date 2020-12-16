package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// CPCMDVersion cp version
const CPCMDVersion = "0.0.1"

// CPCMD define `cp` cmd
var CPCMD = cli.Command{
	Name:    "cp",
	Aliases: []string{"CP"},
	Usage:   "Copy SOURCE to DEST, or multiple SOURCE(s) to DIRECTORY.",
	Description: `默认情况下，源文件的稀疏性仅仅通过简单的方法判断，对应的目标文件目标文件也
	被为稀疏。这是因为默认情况下使用了--sparse=auto 参数。如果明确使用
	--sparse=always 参数则不论源文件是否包含足够长的0 序列也将目标文件创文
	建为稀疏件。
	使用--sparse=never 参数禁止创建稀疏文件。
	
	当指定了--reflink[=always] 参数时执行轻量化的复制，即只在数据块被修改的
	情况下才复制。如果复制失败或者同时指定了--reflink=auto，则返回标准复制模式。
	
	The backup suffix is '~', unless set with --suffix or SIMPLE_BACKUP_SUFFIX.
	The version control method may be selected via the --backup option or through
	the VERSION_CONTROL environment variable.  Here are the values:
	
	  none, off       不进行备份(即使使用了--backup 选项)
	  numbered, t     备份文件加上数字进行排序
	  existing, nil   若有数字的备份文件已经存在则使用数字，否则使用普通方式备份
	  simple, never   永远使用普通方式备份
	
	有一个特别情况：如果同时指定--force 和--backup 选项，而源文件和目标文件
	是同一个已存在的一般文件的话，cp 会将源文件备份。`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "output version information and exit",
		},
	},
	Action: func(c *cli.Context) error {
		if c.Bool("version") {
			fmt.Println(CPCMDVersion)
			return nil
		}
		return nil
	},
}
