package cmd

import (
	"fmt"
	"syscall"

	"github.com/urfave/cli/v2"
)

// UNameCMDVersion version number
const UNameCMDVersion = "0.0.1"

// UNameCMD define `uname` cli command
var UNameCMD = cli.Command{
	Name:        "uname",
	Aliases:     []string{"UNAME"},
	Usage:       "打印系统信息",
	Description: `Print certain system information.  With no OPTION, same as -s.`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"V"},
			Usage:   "输出版本信息并推出",
		},
		&cli.BoolFlag{
			Name:    "all",
			Aliases: []string{"a"},
			Usage:   "以如下次序输出所有信息。其中若-p 和 -i 的探测结果不可知则被省略：",
		}, &cli.BoolFlag{
			Name:    "kernel-name",
			Aliases: []string{"s"},
			Usage:   "输出内核名称",
		}, &cli.BoolFlag{
			Name:    "nodename",
			Aliases: []string{"n"},
			Usage:   "输出网络节点上的主机名",
		}, &cli.BoolFlag{
			Name:    "kernel-release",
			Aliases: []string{"r"},
			Usage:   "输出内核发行号",
		}, &cli.BoolFlag{
			Name:    "kernel-version",
			Aliases: []string{"v"},
			Usage:   "输出内核版本",
		}, &cli.BoolFlag{
			Name:    "machine",
			Aliases: []string{"m"},
			Usage:   "输出机器硬件名称",
		}, &cli.BoolFlag{
			Name:    "processor",
			Aliases: []string{"p"},
			Usage:   "输出处理器类型(non-portable)",
		}, &cli.BoolFlag{
			Name:    "hardware-platform",
			Aliases: []string{"i"},
			Usage:   "输出硬件平台(non-portable)",
		}, &cli.BoolFlag{
			Name:    "operating-system",
			Aliases: []string{"o"},
			Usage:   "输出操作系统",
		}, &cli.BoolFlag{
			Name:  "pretty",
			Usage: "美化输出",
		},
	},
	Action: Uname,
}

func ctos(input *[65]int8) string {
	length := 0
	str := make([]byte, len(input))
	for ; length < len(input); length++ {
		if input[length] == 0 {
			break
		}
		str[length] = uint8(input[length])
	}
	return string(str[0:length])
}

type utsNameStr struct {
	Sysname    string
	Nodename   string
	Release    string
	Version    string
	Machine    string
	Domainname string
}

// Uname  real action
func Uname(c *cli.Context) error {

	utsname := syscall.Utsname{}
	if err := syscall.Uname(&utsname); err != nil {
		return err
	}
	if c.Bool("version") {
		fmt.Println(c.Command.Name, UNameCMDVersion)
		return nil
	}

	u := utsNameStr{
		Sysname:    ctos(&utsname.Sysname),
		Nodename:   ctos(&utsname.Nodename),
		Release:    ctos(&utsname.Release),
		Version:    ctos(&utsname.Version),
		Machine:    ctos(&utsname.Machine),
		Domainname: ctos(&utsname.Domainname),
	}
	outList := []string{}
	outValueList := []string{}
	result := ""
	flagsNum := c.NumFlags()
	for _, n := range c.FlagNames() {
		if n == "pretty" {
			flagsNum--
			break
		}
	}
	if c.Bool("all") {
		result = fmt.Sprintf("%s %s %s %s %s %s %s %s",
			u.Sysname, u.Nodename, u.Release, u.Version,
			u.Machine, u.Machine, u.Machine, u.Sysname)
		outList = []string{"Operating-System", "NodeName", "Kernel-Release",
			"Kernel-Version", "Machine", "Processor", "Hardware-Platform",
			"Operating-System"}
		outValueList = []string{u.Sysname, u.Nodename, u.Release, u.Version,
			u.Machine, u.Machine, u.Machine, u.Sysname}
	} else if flagsNum == 0 {
		result = u.Sysname
		outList = append(outList, "Operating-System")
		outValueList = append(outValueList, u.Sysname)
	} else {
		if c.Bool("kernel-name") {
			result += u.Machine
			outList = append(outList, "Kernel-Name")
			outValueList = append(outValueList, u.Machine)
		}
		if c.Bool("nodename") {
			result += u.Nodename
			outList = append(outList, "NodeName")
			outValueList = append(outValueList, u.Nodename)
		}
		if c.Bool("kernel-release") {
			result += u.Release
			outList = append(outList, "Kernel-Release")
			outValueList = append(outValueList, u.Release)
		}
		if c.Bool("kernel-version") {
			result += u.Version
			outList = append(outList, "Kernel-Version")
			outValueList = append(outValueList, u.Version)
		}
		if c.Bool("machine") {
			result += u.Machine
			outList = append(outList, "Machine")
			outValueList = append(outValueList, u.Machine)
		}
		if c.Bool("processor") {
			result += u.Machine
			outList = append(outList, "Processor")
			outValueList = append(outValueList, u.Machine)
		}
		if c.Bool("hardware-platform") {
			result += u.Machine
			outList = append(outList, "Hardware-Platform")
			outValueList = append(outValueList, u.Machine)
		}
		if c.Bool("operating-system") {
			result += u.Sysname
			outList = append(outList, "Operating-System")
			outValueList = append(outValueList, u.Sysname)
		}

	}

	if c.Bool("pretty") {
		for index, name := range outList {
			fmt.Println(name, ":", outValueList[index])
		}
	} else {
		fmt.Println(result)
	}
	return nil
}
