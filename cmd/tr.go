package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// TrCMDVersion tr version
const TrCMDVersion = "0.0.1"

// TrCMD define `tr` cmd
var TrCMD = cli.Command{
	Name:    "tr",
	Aliases: []string{"TR"},
	Usage:   "translate or delete characters",
	Description: `Translate, squeeze, and/or delete characters from standard input,
	writing to standard output.
	
	SET 是一组字符串，一般都可按照字面含义理解。解析序列如下：

  \NNN  八进制值为NNN 的字符(1 至3 个数位)
  \\            反斜杠
  \a            终端鸣响
  \b            退格
  \f            换页
  \n            换行
  \r            回车
  \t            水平制表符
  \v            垂直制表符
  字符1-字符2   从字符1 到字符2 的升序递增过程中经历的所有字符
  [字符*]       在SET2 中适用，指定字符会被连续复制直到吻合设置1 的长度
  [字符*次数]   对字符执行指定次数的复制，若次数以 0 开头则被视为八进制数
  [:alnum:]     所有的字母和数字
  [:alpha:]     所有的字母
  [:blank:]     所有呈水平排列的空白字符
  [:cntrl:]     所有的控制字符
  [:digit:]     所有的数字
  [:graph:]     所有的可打印字符，不包括空格
  [:lower:]     所有的小写字母
  [:print:]     所有的可打印字符，包括空格
  [:punct:]     所有的标点字符
  [:space:]     所有呈水平或垂直排列的空白字符
  [:upper:]     所有的大写字母
  [:xdigit:]    所有的十六进制数
  [=字符=]      所有和指定字符相等的字符

Translation occurs if -d is not given and both SET1 and SET2 appear.
-t may be used only when translating.  SET2 is extended to length of
SET1 by repeating its last character as necessary.  Excess characters
of SET2 are ignored.  Only [:lower:] and [:upper:] are guaranteed to
expand in ascending order; used in SET2 while translating, they may
only be used in pairs to specify case conversion.  -s uses the last
specified SET, and occurs after translation or deletion.`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "output version information and exit",
		},
	},
	Action: func(c *cli.Context) error {
		if c.Bool("version") {
			fmt.Println(TrCMDVersion)
			return nil
		}
		return nil
	},
}
