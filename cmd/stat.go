package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// StatCMDVersion stat version no
var StatCMDVersion = "v0.0.1"

// StatCMD define stat cli command
var StatCMD = cli.Command{
	Name:    "stat",
	Aliases: []string{"STAT"},
	Usage:   "Display file or file system status",
	Description: `The valid format sequences for files (without --file-system):

	%a   access rights in octal (note '#' and '0' printf flags)
	%A   access rights in human readable form
	%b   number of blocks allocated (see %B)
	%B   the size in bytes of each block reported by %b
	%C   SELinux security context string
	%d   device number in decimal
	%D   device number in hex
	%f   raw mode in hex
	%F   file type
	%g   group ID of owner
	%G   group name of owner
	%h   number of hard links
	%i   inode number
	%m   mount point
	%n   file name
	%N   quoted file name with dereference if symbolic link
	%o   optimal I/O transfer size hint
	%s   total size, in bytes
	%t   major device type in hex, for character/block device special files
	%T   minor device type in hex, for character/block device special files
	%u   user ID of owner
	%U   user name of owner
	%w   time of file birth, human-readable; - if unknown
	%W   time of file birth, seconds since Epoch; 0 if unknown
	%x   time of last access, human-readable
	%X   time of last access, seconds since Epoch
	%y   time of last data modification, human-readable
	%Y   time of last data modification, seconds since Epoch
	%z   time of last status change, human-readable
	%Z   time of last status change, seconds since Epoch
  
  Valid format sequences for file systems:
  
	%a   free blocks available to non-superuser
	%b   total data blocks in file system
	%c   total file nodes in file system
	%d   free file nodes in file system
	%f   free blocks in file system
	%i   file system ID in hex
	%l   maximum length of filenames
	%n   file name
	%s   block size (for faster transfers)
	%S   fundamental block size (for block counts)
	%t   file system type in hex
	%T   file system type in human readable form
  
  注意：您的shell 可能内置了自己的stat 程序版本，它会覆盖这里所提及的相应
  版本。请查阅您的shell 文档获知它所支持的选项。
  
  GNU coreutils online help: <http://www.gnu.org/software/coreutils/>
  请向<http://translationproject.org/team/zh_CN.html> 报告stat 的翻译错误
  Full documentation at: <http://www.gnu.org/software/coreutils/stat>
  or available locally via: info '(coreutils) stat invocation'
  ➜  core git:(master) ✗ stat core.go 
	文件：core.go
	大小：1749            块：8          IO 块：4096   普通文件
  设备：805h/2053d        Inode：17340381    硬链接：1
  权限：(0664/-rw-rw-r--)  Uid：( 1000/   tacey)   Gid：( 1000/   tacey)
  最近访问：2020-11-25 10:50:06.913339128 +0800
  最近更改：2020-11-25 10:50:06.857341765 +0800
  最近改动：2020-11-25 10:50:06.857341765 +0800
  创建时间：-
  ➜  core git:(master) ✗ stat --help
  用法：stat [选项]... 文件...
  Display file or file system status.
  
  必选参数对长短选项同时适用。
	-L, --dereference     follow links
	-f, --file-system     display file system status instead of file status
	-c  --format=FORMAT   use the specified FORMAT instead of the default;
							output a newline after each use of FORMAT
		--printf=FORMAT   like --format, but interpret backslash escapes,
							and do not output a mandatory trailing newline;
							if you want a newline, include \n in FORMAT
	-t, --terse           print the information in terse form
		--help            显示此帮助信息并退出
		--version         显示版本信息并退出
  
  The valid format sequences for files (without --file-system):
  
	%a   access rights in octal (note '#' and '0' printf flags)
	%A   access rights in human readable form
	%b   number of blocks allocated (see %B)
	%B   the size in bytes of each block reported by %b
	%C   SELinux security context string
	%d   device number in decimal
	%D   device number in hex
	%f   raw mode in hex
	%F   file type
	%g   group ID of owner
	%G   group name of owner
	%h   number of hard links
	%i   inode number
	%m   mount point
	%n   file name
	%N   quoted file name with dereference if symbolic link
	%o   optimal I/O transfer size hint
	%s   total size, in bytes
	%t   major device type in hex, for character/block device special files
	%T   minor device type in hex, for character/block device special files
	%u   user ID of owner
	%U   user name of owner
	%w   time of file birth, human-readable; - if unknown
	%W   time of file birth, seconds since Epoch; 0 if unknown
	%x   time of last access, human-readable
	%X   time of last access, seconds since Epoch
	%y   time of last data modification, human-readable
	%Y   time of last data modification, seconds since Epoch
	%z   time of last status change, human-readable
	%Z   time of last status change, seconds since Epoch
  
  Valid format sequences for file systems:
  
	%a   free blocks available to non-superuser
	%b   total data blocks in file system
	%c   total file nodes in file system
	%d   free file nodes in file system
	%f   free blocks in file system
	%i   file system ID in hex
	%l   maximum length of filenames
	%n   file name
	%s   block size (for faster transfers)
	%S   fundamental block size (for block counts)
	%t   file system type in hex
	%T   file system type in human readable form`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "输出版本信息并推出",
		}, &cli.BoolFlag{
			Name:    "dereference",
			Aliases: []string{"L"},
			Usage:   "follow links",
		}, &cli.BoolFlag{
			Name:    "file-system",
			Aliases: []string{"f"},
			Usage:   "display file system status instead of file status",
		}, &cli.StringFlag{
			Name:    "--format",
			Aliases: []string{"c"},
			Usage: `use the specified FORMAT instead of the default;
			output a newline after each use of FORMAT`,
		}, &cli.StringFlag{
			Name: "printf",
			Usage: `like --format, but interpret backslash escapes,
			and do not output a mandatory trailing newline;
			if you want a newline, include \n in FORMAT
			`,
		}, &cli.BoolFlag{
			Name:    "terse",
			Aliases: []string{"t"},
			Usage:   " print the information in terse form",
		},
	},
	Action: statAction,
}

func statAction(c *cli.Context) error {
	if c.Bool("version") {
		fmt.Println(c.Command.Name, ArchCMDVersion)
		return nil
	}

	return nil

}
