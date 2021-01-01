package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// CPCMDVersion cp version
const CPCMDVersion = "0.0.1"

// CPCMD define `cp` cmd
var CPCMD = cli.Command{
	Name:      "cp",
	Aliases:   []string{"CP"},
	UsageText: "core cp [选项]... [-T] 源文件 目标文件]\ncore cp [选项]... 源文件... 目录\ncore cp [选项]... -t 目录 源文件...",
	Usage:     "拷贝复制文件",
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
			Usage:   "输出版本信息并推出",
		}, &cli.BoolFlag{
			Name:    "archive",
			Aliases: []string{""},
			Usage:   "等于-dR --preserve=all",
		}, &cli.BoolFlag{
			Name:  "attributes-only",
			Usage: "仅复制属性而不复制数据",
		}, &cli.BoolFlag{
			Name:  "backup",
			Usage: "为每个已存在的目标文件创建备份",
		}, &cli.BoolFlag{
			Name:  "b",
			Usage: "类似--backup 但不接受参数",
		}, &cli.BoolFlag{
			Name:  "copy-contents",
			Usage: "在递归处理是复制特殊文件内容",
		}, &cli.BoolFlag{
			Name:  "d",
			Usage: "等于--no-dereference --preserve=links",
		}, &cli.BoolFlag{
			Name:    "force",
			Aliases: []string{"f"},
			Usage:   "if an existing destination file cannot be opened, remove it and try again (this option is ignored when the -n option is also used)",
		}, &cli.BoolFlag{
			Name:    "interactive",
			Aliases: []string{"i"},
			Usage:   "prompt before overwrite (overrides a previous -n option)",
		}, &cli.BoolFlag{
			Name:  "H",
			Usage: "follow command-line symbolic links in SOURCE",
		}, &cli.BoolFlag{
			Name:    "link",
			Aliases: []string{"l"},
			Usage:   "hard link files instead of copying",
		}, &cli.BoolFlag{
			Name:    "dereference",
			Aliases: []string{"L"},
			Usage:   "always follow symbolic links in SOURCE",
		}, &cli.BoolFlag{
			Name:    "no-clobber",
			Aliases: []string{"n"},
			Usage:   "不要覆盖已存在的文件(使前面的 -i 选项失效)",
		}, &cli.BoolFlag{
			Name:    "no-dereferenc",
			Aliases: []string{"P"},
			Usage:   "不跟随源文件中的符号链接",
		}, &cli.BoolFlag{
			Name:  "p",
			Usage: "等于--preserve=模式,所有权,时间戳",
		}, &cli.StringSliceFlag{
			Name:  "preserve",
			Usage: "保持指定的`属性列表`(默认：模式,所有权,时间戳)，如果可能保持附加属性：环境、链接、xattr 等",
		}, &cli.StringSliceFlag{
			Name:  "sno-preserve",
			Usage: "不保留指定的文件`属性列表`",
		}, &cli.BoolFlag{
			Name:  "parents",
			Usage: "复制前在目标目录创建来源文件路径中的所有目录",
		}, &cli.BoolFlag{
			Name:    "recursive",
			Aliases: []string{"r", "R"},
			Usage:   "递归复制目录及其子目录内的所有内容",
		}, &cli.StringFlag{
			Name:  "reflink",
			Usage: "控制克隆/CoW 副本。请查看下面的内容。",
		}, &cli.BoolFlag{
			Name:  "remove-destination",
			Usage: "尝试打开目标文件前先删除已存在的目的地文件 (相对于 --force 选项)",
		}, &cli.StringFlag{
			Name:  "sparse",
			Usage: "控制创建稀疏文件的方式",
		}, &cli.BoolFlag{
			Name:  "strip-trailing-slashes",
			Usage: "删除参数中所有源文件/目录末端的斜杠",
		}, &cli.BoolFlag{
			Name:    "symbolic-link",
			Aliases: []string{"s"},
			Usage:   "只创建符号链接而不复制文件",
		}, &cli.StringFlag{
			Name:    "suffix",
			Aliases: []string{"S"},
			Usage:   "自行指定备份文件的`后缀`",
		}, &cli.PathFlag{
			Name:    "target-directory",
			Aliases: []string{"t"},
			Usage:   "将所有参数指定的源文件/目录复制至目标`目录`",
		}, &cli.BoolFlag{
			Name:    "no-target-directory",
			Aliases: []string{"T"},
			Usage:   "将目标目录视作普通文件",
		}, &cli.BoolFlag{
			Name:    "update",
			Aliases: []string{"u"},
			Usage:   "只在源文件比目标文件新，或目标文件不存在时才进行复制",
		}, &cli.BoolFlag{
			Name:  "verbose",
			Usage: "显示详细的进行步骤",
		}, &cli.BoolFlag{
			Name:    "one-file-system",
			Aliases: []string{"x"},
			Usage:   "不跨越文件系统进行操作",
		}, &cli.BoolFlag{
			Name:  "Z",
			Usage: "et SELinux security context of destination file to default type",
		}, &cli.StringFlag{
			Name:  "context",
			Usage: "like -Z, or if CTX is specified then set the SELinux or SMACK security context to `CTX`",
		},
	},
	Action: func(c *cli.Context) error {
		if c.Bool("version") {
			fmt.Println(CPCMDVersion)
			return nil
		}
		return nil
	},
	// UseShortOptionHandling: true,
}
