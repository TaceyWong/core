
## <font color=red> 只是用来练习基础编程, 不要将其用于真实环境 </font>

## <font color=red>Linux Only,Mac is not</font>



# core


![License](https://img.shields.io/badge/license-GPL3.0-brightgreen.svg?style=flat-square)
![Go Report Card](https://goreportcard.com/badge/github.com/TaceyWong/core)

Coreutils(git://git.sv.gnu.org/coreutils)的Go实现

```sh

   ______   ____     ____     ______
  / ____/  / __ \   / __ \   / ____/
 / /      / / / /  / /_/ /  / __/
/ /___   / /_/ /  / _, _/  / /___
\____/   \____/  /_/ |_|  /_____/

Please type "ctools -h/--help" for the help of usage
```

```sh
NAME:
   core - Clone of GNU Coreutils

USAGE:
   core [global options] command [command options] [arguments...]

VERSION:
   0.0.1

AUTHOR:
   Tacey Wong <xinyong.wang@qq.com>

COMMANDS:
   arch, ARCH          打印机器硬件名称 (同 uname -m)
   base32, BASE32      base32编码/解码数据并打印到标准输出
   base64, BASE64      base64编码/解码数据并打印到标准输出
   basename, BASENAME  移除文件名的目录和后缀
   cat, CAT            连接所有指定文件并将结果写到标准输出
   chgrp, CHGRP        更改组所有权
   chown, CHOWN        更改文件所有者和组
   chroot, CHROOT      以指定的新根为运行指定命令时的的根目录。
   cksum, CKSUM        输出每个文件的 CRC 校验值和字节统计。
   comm, COMM          逐行比较已排序的文件文件1 和文件2
   cp, CP              拷贝复制文件
   csplit, CSPLIT      将用PATTERN分隔的FILE输出到文件“ xx00”，“ xx01……
   cut, CUT            将每个文件中选定行的部分打印到标准输出
   date, DATE          打印或设置系统日期和时间
   dd, DD              转换和拷贝文件
   df, DF              显示文件系统信息
   dirname, DIRNAME    输出每个NAME的最后一个非斜杠组成部分，并删除尾随的斜杠
   du, DU              递归地汇总目录或文件的磁盘使用情况
   env, ENV            在修改后的环境中运行程序
   expand, EXPAND      制表符转换为空格
   expr, EXPR          将表达式的值列印到标准输出，分隔符下面的空行可提升算式优先级。
   factor, FACTOR      输出每个指定的数字的素因子，如果没有在命令行中指定则从标准输入读取。
   false, FALSE        返回布尔False
   groups, GROUPS      显示每个输入的用户名所在的全部组，
                       如果没有指定用户名则默认为当前进程用户(当用户组数据库发生变更时可能导致差异)。
   head, HEAD          打印文本文件的最前10行到标准输出
   hostid, HOSTID      打印当前主机的数字标识符
   hostname, HOSTNAME  显示或设置系统主机名
   id, ID              打印指定USER或当前用户（USER省略时）的用户和组信息
   join, JOIN          在一个公共字段上连接两个文件的行
   logname, LOGNAME    显示当前用户的名称
   lscpu, LSCPU        显示 CPU 架构信息
   mktemp, MKTEMP      创建临时文件、临时目录
   mv, MV              文件重命名或文件剪切
   nice, NICE          修改调度优先级运行程序
   nl, NL              文本文件诸行添加数字编号打印到标准输出
   nproc, NPROC        打印可用的处理器单元数
   numfmt, NUMFMT      将数字转换为人类可读的字符串
   paste, PASTE        合并文件文本行
   pathchk, PATHCHK    检查文件名是否有效或可移植
   printenv, PRINTENV  显示指定的环境变量的值。
                       如果没有指定变量，则打印出所有变量的名称和值。
   ptx, PTX            生成文件内容的排列索引
   pwd, PWD            打印当前/工作目录名称
   realpath, REALPATH  打印已解析的路径[绝对地址]
   rm, RM              删除文件
   seq, SEQ            打印数字序列(暂时只支持整数)
   shred, SHRED        覆盖文件以隐藏其内容，并可选择将其删除
   shuf, SHUF          生成随机排列
   sort, SORT          排序文本文件的文本行
   split, SPLIT        文件分片
   stat, STAT          显示文件或文件系统状态
   sync, SYNC          同步缓存写到持久存储
   tail, TAIL          打印文本文件的最后10行到标准输出
   tee, TEE            标准输入复制到每个指定文件，并显示到标准输出。
   test, TEST          检查文件类型并比较值
   timeout, TIMEOUT    有时间限制地运行命令
   touch, TOUCH        更改文件时间戳
   tr, TR              转换或删除字符
   true, TRUE          什么都不做，成功执行空白进程
   truncate, TRUNCATE  缩小或扩展文件的大小到指定的大小
   tsort, TSORT        执行拓扑排序
   tty, TTY            显示出连接到当前标准输入的终端设备文件名。
   uname, UNAME        打印系统信息
   unexpand, UNEXPAND  空格转换成制表符
   uptime, UPTIME      显示系统运行了多长时间
   users, USERS        根据文件判断输出当前有谁正登录在系统上。
                       如果文件未予指定，则使用/var/run/utmp，/var/log/wtmp 是通用的相关文件
   yes, YES            重复输出一个字符串直到被kill

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)

COPYRIGHT:
   (c) 2020 - Forever Tacey Wong
```
## Main

**整个文件的输出**

+ [ ] `cat`: 连接文件并打印到标准输出
+ [ ] `tac`: 逆向连接文件并打印到标准输出
+ [ ] `nl`:  文件添加行号
+ [ ] `od`: dump files in octal and other formats
+ [x] `base32`: base32 编码/解码 数据并打印到标准输出
+ [x] `base64`: base64 编码/解码 数据并打印到标准输出
+ [ ] `basenc`: 

**格式化文件内容**

+ [ ] `fm`: `sudo apt install fmtools` : simple optimal text formatter
+ [ ] `pr`: 转换文本文件以便于打印
+ [ ] `fold`: wrap each input line to fit in specified width

**文件的部分输出**

+ [ ] `head`: 输出文件开始的部分
+ [ ] `tail`: 输出文件结束的部分
+ [ ] `split`: 文件切片
+ [ ] `csplit`: split a file into sections determined by context lines

**汇总文件**

+ [ ] `cksum`: 校验和计算文件中的字节数
+ [ ] `b2sum`: 计算并校验BLAKE2
+ [ ] `md5sum`: 计算并校验MD5
+ [ ] `sha1sum`: 计算并校验SHA1
+ [ ] `sha2 utilties`:

**对已排序文件进行操作**

+ [ ] `sort`: 文本文件按行排序
+ [ ] `shuf`: 生成随机排列
+ [ ] `uniq`: 移除重复的行
+ [ ] `comm`: 逐行比对两个已经排序的文件
+ [ ] `ptx`: 生成文件内容的排列索引
+ [ ] `tsort`: 执行拓扑排序

**文本字段操作**

+ [ ] `cut`: 从文件的每一行中删除部分
+ [ ] `paste`: 合并文件行
+ [ ] `join`: 在一个公共字段上连接两个文件的行

**字符操作**

+ [ ] `tr`: translate or delete characters
+ [ ] `expand`: 制表转换为空格
+ [ ] `unexpand`:  空格转换为制表

**目录列表**

+ [ ] `ls`: 列出目录清单
+ [ ] `dir`: 列出目录清单
+ [ ] `vdir`:  列出目录清单
+ [ ] `dircolors`: color setup for ls

**基础操作**

+ [ ] `cp`: 拷贝文件和目录
+ [ ] `dd`: 转换和拷贝一个文件
+ [ ] `install`:  拷贝文件和设置属性
+ [ ] `mv`:  移动和重命名文件
+ [ ] `rm`: 删除文件或目录
+ [ ] `shred`:  overwrite a file to hide its contents, and optionally delete it

**特殊文件类型**

+ [ ] `link`: call the link function to create a link to a file
+ [ ] `ln`:  make links between files
+ [ ] `mkdir`:  make directories
+ [ ] `mkfifo`: make FIFOs (named pipes)
+ [ ] `mknod`:  make block or character special files
+ [ ] `readlink`: print resolved symbolic links or canonical file names
+ [ ] `rmdir`: 移除空目录
+ [ ] `unlink`: call the unlink function to remove the specified file

**更改文件属性**

+ [ ] `chown`: change file owner and group
+ [ ] `chgrp`: change group ownership
+ [ ] `chmod`:  更改文件模式位
+ [ ] `touch`: 更改文件时间戳

**Disk usage**

+ [ ] `df`: 报告文件系统磁盘空间使用情况
+ [ ] `du`:  评估文件空间使用情况
+ [ ] `stat`: 显示文件或文件系统状态
+ [ ] `sync`: 同步缓存写到持久存储
+ [ ] `truncate`: shrink or extend the size of a file to the specified size 

**打印文本**

+ [ ] `echo`:  display a line of text
+ [ ] `printf`: format and print data
+ [x] `yes`: output a string repeatedly until killed

**条件**

+ [x] `false`: do nothing, unsuccessfully
+ [x] `true`: do nothing, successfully
+ [ ] `test`:  check file types and compare values
+ [ ] `expr`:  evaluate expressions

**重导**

+ [ ] `tee`:  read from standard input and write to standard output and files

**文件名操作**

+ [x] `basename`: 从文件名中删除目录和后缀
+ [ ] `dirname`: 从文件名中删除最后一部分
+ [ ] `pathchk`: check whether file names are valid or portable
+ [ ] `mktemp`: 创建一个临时文件或文件夹
+ [ ] `realpath`: print the resolved path

**工作上下文**

+ [ ] `pwd`:  print name of current/working directory
+ [ ] `stty`:  change and print terminal line settings
+ [x] `printenv`:  print all or part of environment
+ [ ] `tty`:  print the file name of the terminal connected to standard input

**用户信息**

+ [ ] `id`: print real and effective user and group IDs
+ [ ] `logname`: 打印用户的登录名
+ [ ] `whoami`:  打印有效的用户标识
+ [ ] `groups`: 打印用户所在的群组
+ [ ] `users`: - print the user names of users currently logged in to the current host
+ [ ] `who`: 显示谁登录了系统

**系统上下文**

+ [ ] `date`: 显示和设置系统日期和时间
+ [x] `arch`:  打印机器硬件名称 (和`uname -m`相同)
+ [ ] `nproc`: print the number of processing units available
+ [x] `uname`: print system information
+ [x] `hostname`: show or set the system's host name
+ [x] `hostid`:  print the numeric identifier for the current host
+ [x] `uptime`: 显示系统已经运行的时长


**SELinux 上下文**

+ [ ] `chcon`: 改变文件安全上下文
+ [ ] `runcon`: run command with specified security context

**修改命令调用**

+ [ ] `chroot`:run command or interactive shell with special root directory
+ [ ] `env`: run a program in a modified environment
+ [ ] `nice`: run a program with modified scheduling priority
+ [ ] `nohup`: run a command immune to hangups, with output to a non-tty
+ [ ] `stdbuf`: Run COMMAND, with modified buffering operations for its standard streams.
+ [ ] `timeout`:  run a command with a time limit

**进程控制**

+ [ ] `kill`: 给进程发送一个停止运行的信号

**延时**

+ [ ] `sleep`:  延迟指定长的时间


**数字操作**

+ [ ] `factor`: factor numbers 
+ [ ] `numfmt`: Convert numbers from/to human-readable strings
+ [x] `seq`: 打印一个数字序列


## 引用参考

+ https://www.gnu.org/software/coreutils/
+ http://www.maizure.org/projects/decoded-gnu-coreutils/
+ https://en.wikipedia.org/wiki/List_of_GNU_Core_Utilities_commands




















































































