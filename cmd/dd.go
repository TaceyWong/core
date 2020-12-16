package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// DDCMDVersion dd version
const DDCMDVersion = "0.0.1"

// DDCMD define `dd` cmd
var DDCMD = cli.Command{
	Name:    "dd",
	Aliases: []string{"DD"},
	Usage:   "转换和拷贝文件",
	Description: `Copy a file, converting and formatting according to the operands.

	bs=BYTES        read and write up to BYTES bytes at a time (default: 512);
					overrides ibs and obs
	cbs=BYTES       convert BYTES bytes at a time
	conv=CONVS      convert the file as per the comma separated symbol list
	count=N         copy only N input blocks
	ibs=BYTES       read up to BYTES bytes at a time (default: 512)
	if=FILE         read from FILE instead of stdin
	iflag=FLAGS     read as per the comma separated symbol list
	obs=BYTES       write BYTES bytes at a time (default: 512)
	of=FILE         write to FILE instead of stdout
	oflag=FLAGS     write as per the comma separated symbol list
	seek=N          skip N obs-sized blocks at start of output
	skip=N          skip N ibs-sized blocks at start of input
	status=LEVEL    The LEVEL of information to print to stderr;
					'none' suppresses everything but error messages,
					'noxfer' suppresses the final transfer statistics,
					'progress' shows periodic transfer statistics
  
  N and BYTES may be followed by the following multiplicative suffixes:
  c =1, w =2, b =512, kB =1000, K =1024, MB =1000*1000, M =1024*1024, xM =M,
  GB =1000*1000*1000, G =1024*1024*1024, and so on for T, P, E, Z, Y.
  
  Each CONV symbol may be:
  
	ascii     from EBCDIC to ASCII
	ebcdic    from ASCII to EBCDIC
	ibm       from ASCII to alternate EBCDIC
	block     pad newline-terminated records with spaces to cbs-size
	unblock   replace trailing spaces in cbs-size records with newline
	lcase     change upper case to lower case
	ucase     change lower case to upper case
	sparse    try to seek rather than write the output for NUL input blocks
	swab      swap every pair of input bytes
	sync      pad every input block with NULs to ibs-size; when used
			  with block or unblock, pad with spaces rather than NULs
	excl          fail if the output file already exists
	nocreat       do not create the output file
	notrunc       不截断输出文件
	noerror       读取数据发生错误后仍然继续
	fdatasync     结束前将输出文件数据写入磁盘
	fsync 类似上面，但是元数据也一同写入
  
  FLAG 符号可以是：
  
	append        追加模式(仅对输出有意义；隐含了conv=notrunc)
	direct        使用直接I/O 存取模式
	directory     除非是目录，否则 directory 失败
	dsync         使用同步I/O 存取模式
	sync          与上者类似，但同时也对元数据生效
	fullblock     为输入积累完整块(仅iflag)
	nonblock      使用无阻塞I/O 存取模式
	noatime       不更新存取时间
	nocache   Request to drop cache.  See also oflag=sync
	noctty        不根据文件指派控制终端
	nofollow      不跟随链接文件
	count_bytes  treat 'count=N' as a byte count (iflag only)
	skip_bytes  treat 'skip=N' as a byte count (iflag only)
	seek_bytes  treat 'seek=N' as a byte count (oflag only)
  
  Sending a USR1 signal to a running 'dd' process makes it
  print I/O statistics to standard error and then resume copying.`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "output version information and exit",
		},
	},
	Action: func(c *cli.Context) error {
		if c.Bool("version") {
			fmt.Println(DDCMDVersion)
			return nil
		}
		return nil
	},
}
