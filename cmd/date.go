package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/urfave/cli/v2"
)

// DateCMDVersion date version
const DateCMDVersion = "0.0.1"
const layout = "2006年 01月 02日 Monday 15:04:05 CST"

var weekMap = map[string]string{"Monday": "星期一", "Tuesday": "星期二", "Wednesday": "星期三", "Thursday": "星期四",
	"Friday": "星期五", "Saturday": "星期六", "Sunday": "星期日",
}

// DateCMD define `date` cmd
var DateCMD = cli.Command{
	Name:    "date",
	Aliases: []string{"DATE"},
	Usage:   "打印或设置系统日期和时间",
	Description: `Display the current time in the given FORMAT, or set the system date.
	给定的格式FORMAT 控制着输出，解释序列如下：

	%%    一个文字的 %
	%a    当前locale 的星期名缩写(例如： 日，代表星期日)
	%A    当前locale 的星期名全称 (如：星期日)
	%b    当前locale 的月名缩写 (如：一，代表一月)
	%B    当前locale 的月名全称 (如：一月)
	%c    当前locale 的日期和时间 (如：2005年3月3日 星期四 23:05:25)
	%C    世纪；比如 %Y，通常为省略当前年份的后两位数字(例如：20)
	%d    按月计的日期(例如：01)
	%D    按月计的日期；等于%m/%d/%y
	%e    按月计的日期，添加空格，等于%_d
	%F    完整日期格式，等价于 %Y-%m-%d
	%g    ISO-8601 格式年份的最后两位 (参见%G)
	%G    ISO-8601 格式年份 (参见%V)，一般只和 %V 结合使用
	%h    等于%b
	%H    小时(00-23)
	%I    小时(00-12)
	%j    按年计的日期(001-366)
	%k   hour, space padded ( 0..23); same as %_H
	%l   hour, space padded ( 1..12); same as %_I
	%m   month (01..12)
	%M   minute (00..59)
	%n   a newline
	%N   nanoseconds (000000000..999999999)
	%p   locale's equivalent of either AM or PM; blank if not known
	%P   like %p, but lower case
	%q   quarter of year (1..4)
	%r   locale's 12-hour clock time (e.g., 11:11:04 PM)
	%R   24-hour hour and minute; same as %H:%M
	%s   seconds since 1970-01-01 00:00:00 UTC
	%S    秒(00-60)
	%t    输出制表符 Tab
	%T    时间，等于%H:%M:%S
	%u    星期，1 代表星期一
	%U    一年中的第几周，以周日为每星期第一天(00-53)
	%V    ISO-8601 格式规范下的一年中第几周，以周一为每星期第一天(01-53)
	%w    一星期中的第几日(0-6)，0 代表周一
	%W    一年中的第几周，以周一为每星期第一天(00-53)
	%x    当前locale 下的日期描述 (如：12/31/99)
	%X    当前locale 下的时间描述 (如：23:13:48)
	%y    年份最后两位数位 (00-99)
	%Y    年份
	%z +hhmm              数字时区(例如，-0400)
	%:z +hh:mm            数字时区(例如，-04:00)
	%::z +hh:mm:ss        数字时区(例如，-04:00:00)
	%:::z                 数字时区带有必要的精度 (例如，-04，+05:30)
	%Z                    按字母表排序的时区缩写 (例如，EDT)
  
  默认情况下，日期的数字区域以0 填充。
  The following optional flags may follow '%':
  
	-  (hyphen) do not pad the field
	_  (underscore) pad with spaces
	0  (zero) pad with zeros
	^  use upper case if possible
	#  use opposite case if possible
  
  在任何标记之后还允许一个可选的域宽度指定，它是一个十进制数字。
  作为一个可选的修饰声明，它可以是E，在可能的情况下使用本地环境关联的
  表示方式；或者是O，在可能的情况下使用本地环境关联的数字符号。
  
  Examples:
  Convert seconds since the epoch (1970-01-01 UTC) to a date
	$ date --date='@2147483647'
  
  Show the time on the west coast of the US (use tzselect(1) to find TZ)
	$ TZ='America/Los_Angeles' date
  
  Show the local time for 9AM next Friday on the west coast of the US
	$ date --date='TZ="America/Los_Angeles" 09:00 next Fri'`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "输出版本信息并推出",
		}, &cli.StringFlag{
			Name:    "date",
			Aliases: []string{"d"},
			Usage:   "display time described by `STRING`, not 'now'",
		}, &cli.BoolFlag{
			Name:  "debug",
			Usage: "annotate the parsed date, and warn about questionable usage to stderr",
		}, &cli.PathFlag{
			Name:    "file",
			Aliases: []string{"f"},
			Usage:   "like --date; once for each line of `DATEFILE`",
		}, &cli.BoolFlag{
			Name:    "rfc-email",
			Aliases: []string{"R"},
			Usage:   "output date and time in RFC 5322 format. Example: Mon, 14 Aug 2006 02:34:56 -0600",
		}, &cli.StringFlag{
			Name:  "rfc-3339",
			Usage: "output date/time in RFC 3339 format. FMT='date', 'seconds', or 'ns' for date and time to the indicated precision. Example: 2006-08-14 02:34:56-06:00",
		}, &cli.StringFlag{
			Name:    "reference",
			Aliases: []string{"r"},
			Usage:   "display the last modification time of `FILE`",
		}, &cli.StringFlag{
			Name:    "set",
			Aliases: []string{"s"},
			Usage:   "set time described by `STRING`",
		}, &cli.BoolFlag{
			Name:    "utc",
			Aliases: []string{"u"},
			Usage:   "print or set Coordinated Universal Time (UTC)",
		}, &cli.BoolFlag{
			Name:  "universal",
			Usage: "the same as --utc",
		},
	},
	Action: func(c *cli.Context) error {
		if c.Bool("version") {
			fmt.Println(DateCMDVersion)
			return nil
		}
		t := time.Now()
		local := t.Format(layout)
		w := t.Weekday().String()
		fmt.Println(strings.ReplaceAll(local, w, weekMap[w]))
		return nil
	},
}
