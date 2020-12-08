package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// LsCPUCMDVersion lscpu version no
var LsCPUCMDVersion = "v0.0.1"

// LsCPUCMD define lscpu cli command
var LsCPUCMD = cli.Command{
	Name:    "lscpu",
	Aliases: []string{"LSCPU"},
	Usage:   "显示 CPU 架构信息",
	Description: `lscpu  gathers  CPU  architecture  information  from  sysfs,  /proc/cpuinfo  and any applicable architecture-specific
	libraries (e.g. librtas on Powerpc).  The command output can be optimized for parsing  or  for  easy  readability  by
	humans.   The  information includes, for example, the number of CPUs, threads, cores, sockets, and Non-Uniform Memory
	Access (NUMA) nodes.  There is also information about the CPU caches and cache sharing, family, model, bogoMIPS, byte
	order, and stepping.

	In virtualized environments, the CPU architecture information displayed reflects the configuration of the guest oper‐
	ating system which is typically different from the physical (host) system.  On architectures that support  retrieving
	physical topology information, lscpu also displays the number of physical sockets, chips, cores in the host system.

	Options  that  result  in  an  output table have a list argument.  Use this argument to customize the command output.
	Specify a comma-separated list of column labels to limit the output table to only the specified columns, arranged  in
	the specified order.  See COLUMNS for a list of valid column labels.  The column labels are not case sensitive.

	Not  all  columns are supported on all architectures.  If an unsupported column is specified, lscpu prints the column
	but does not provide any data for it.
	
	Available output columns:
           CPU  逻辑 CPU 数量
          CORE  逻辑核心数量
        SOCKET  逻辑(CPU)座数量
          NODE  逻辑 NUMA 节点数量
          BOOK  逻辑 book 数
        DRAWER  逻辑抽屉号
         CACHE  显示 CPU 间是如何共享缓存的
  POLARIZATION  虚拟硬件上的 CPU 调度模式
       ADDRESS  CPU 的物理地址
    CONFIGURED  显示超级监督(hypervisor)是否分配了 CPU
        ONLINE  显示 Linux 当前是否在使用该 CPU
        MAXMHZ  显示 CPU 的最大 MHz
        MINMHZ  显示 CPU 的最小 MHz`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "output version information and exit",
		}, &cli.BoolFlag{
			Name:    "all",
			Aliases: []string{"a"},
			Usage:   `同时打印在线和离线 CPU (-e 选项默认值)`,
		}, &cli.BoolFlag{
			Name:    "online",
			Aliases: []string{"b"},
			Usage:   `只打印在线 CPU (-p 选项默认值)`,
		}, &cli.BoolFlag{
			Name:    "offline",
			Aliases: []string{"c"},
			Usage:   `只打印离线 CPU`,
		}, &cli.BoolFlag{
			Name:    `json`,
			Aliases: []string{"J"},
			Usage:   `use JSON for default or extended format`,
		}, &cli.StringSliceFlag{
			Name:    `extended`,
			Aliases: []string{"e"},
			Usage:   `打印扩展的可读格式`,
		}, &cli.StringSliceFlag{
			Name:    `parse`,
			Aliases: []string{"p"},
			Usage:   `打印可解析格式`,
		}, &cli.PathFlag{
			Name:    `sysroot`,
			Aliases: []string{"s"},
			Usage:   `以指定目录作为系统根目录`,
		}, &cli.BoolFlag{
			Name:    `hex`,
			Aliases: []string{"x"},
			Usage:   `打印十六进制掩码而非 CPU 列表`,
		}, &cli.BoolFlag{
			Name:    `physical`,
			Aliases: []string{"y"},
			Usage:   `打印物理 ID 而非逻辑 ID`,
		},
	},
	Action: lsCPUAction,
}

func lsCPUAction(c *cli.Context) error {
	if c.Bool("version") {
		fmt.Println(c.Command.Name, EnvCMDVersion)
		return nil
	}

	return nil

}
