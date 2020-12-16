package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// ShredCMDVersion shred version
const ShredCMDVersion = "0.0.1"

// ShredCMD define `shred` cmd
var ShredCMD = cli.Command{
	Name:    "shred",
	Aliases: []string{"SHRED"},
	Usage:   "overwrite a file to hide its contents, and optionally delete it",
	Description: `Overwrite the specified FILE(s) repeatedly, in order to make it harder
	for even very expensive hardware probing to recover the data.
	
	If FILE is -, shred standard output.
	
	Delete FILE(s) if --remove (-u) is specified.  The default is not to remove
the files because it is common to operate on device files like /dev/hda,
and those files usually should not be removed.
The optional HOW parameter indicates how to remove a directory entry:
'unlink' => use a standard unlink call.
'wipe' => also first obfuscate bytes in the name.
'wipesync' => also sync each obfuscated byte to disk.
The default mode is 'wipesync', but note it can be expensive.

警告：请注意使用shred 时有一个很重要的条件：
文件系统会在原来的位置覆盖指定的数据。传统的文件系统符合此条件，但许多现代
的文件系统都不符合条件。以下是会令shred 无效或不担保一定有效的文件系统的
例子：

* 有纪录结构或是日志式文件系统，如AIX 及Solaris 使用的文件系统 (以及
   JFS、ReiserFS、XFS、Ext3 等)

* 会重复写入数据，及即使一部份写入动作失败后仍可继续的文件系统，如使用
   RAID 的文件系统

* 会不时进行快照记录的文件系统，像Network Applicance 的NFS 服务器

* 文件系统是存放于缓存位置，比如NFS 第三版用户端

* 压缩文件系统

在Ext3 文件系统中，以上免责声明仅适用于启用了data=journal 模式的情况，
此时文件日志记录了附加的元数据 shred 的作用将受到影响。在data=ordered(默认)
或data=writeback 模式下shred 仍然有效。
Ext3 日志模式可通过向/etc/fstab 的挂载选项中添加data=something 进行设置，
您可以查看mount 的man 页面以获得详细信息。

另外，文件系统备份和远程镜像可能会
包含不能被删除的文件副本，这将会
允许碎片文件被恢复。`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "output version information and exit",
		},
	},
	Action: func(c *cli.Context) error {
		if c.Bool("version") {
			fmt.Println(ShredCMDVersion)
			return nil
		}
		return nil
	},
}
