package cmd

import (
	"errors"
	"fmt"
	"os"
	"syscall"
	"time"

	"github.com/urfave/cli/v2"
)

// TouchCMDVersion touch version no
var TouchCMDVersion = "v0.0.1"

// TouchCMD define touch cli command
var TouchCMD = cli.Command{
	Name:      "touch",
	Aliases:   []string{"TOUCH"},
	UsageText: "core touch [选项]... 文件...",
	Usage:     `更改文件时间戳`,
	Description: `Update the access and modification times of each FILE to the current time.
	A FILE argument that does not exist is created empty, unless -c or -h
	is supplied.
	
	A FILE argument string of - is handled specially and causes touch to
	change the times of the file associated with standard output.
	
	请注意，-d 和-t 选项可接受不同的时间/日期格式。`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "输出版本信息并推出",
		}, &cli.BoolFlag{
			Name:    "access",
			Aliases: []string{"a"},
			Usage:   `只更改访问时间`,
		}, &cli.BoolFlag{
			Name:    "no-create",
			Aliases: []string{"c"},
			Usage:   `不创建任何文件`,
		}, &cli.StringFlag{
			Name:    "date",
			Aliases: []string{"d"},
			Usage:   `使用指定字符串表示时间而非当前时间`,
		}, &cli.BoolFlag{
			Name:    "no-dereference",
			Aliases: []string{"nd"},
			Usage: `会影响符号链接本身，而非符号链接所指示的目的地
			(当系统支持更改符号链接的所有者时，此选项才有用)[尚未实现]`,
		}, &cli.BoolFlag{
			Name:    "modify",
			Aliases: []string{"m"},
			Usage:   `只更改修改时间`,
		}, &cli.PathFlag{
			Name:    "reference",
			Aliases: []string{"r"},
			Usage:   `use this file's times instead of current time`,
		}, &cli.StringFlag{
			Name: "time",
			Usage: `change the specified time:
			WORD is access, atime, or use: equivalent to -a
			WORD is modify or mtime: equivalent to -m 【尚未实现】`,
		}, &cli.BoolFlag{
			Name:    "STAMP",
			Aliases: []string{"s"},
			Usage:   `use [[CC]YY]MMDDhhmm[.ss] instead of current time【尚未实现】`,
		},
	},
	Action: touchAction,
}

func touchAction(c *cli.Context) (err error) {
	if c.Bool("version") {
		fmt.Println(c.Command.Name, TouchCMDVersion)
		return nil
	}
	if c.Args().Len() == 0 {
		cli.ShowCommandHelp(c, c.Command.Name)
		return errors.New("缺少了文件操作数")
	}

	atime := time.Now()
	mtime := time.Now()

	if r := c.Path("reference"); r != "" {
		atime, mtime, _, err = StatTimes(r)
		if err != nil {
			return err
		}
	}
	for _, f := range c.Args().Slice() {
		if _, e := os.Stat(f); e != nil {
			if c.Bool("no-create") {
				continue
			} else {
				_, err = os.Create(f)
				if err != nil {
					return err
				}
			}
		}
		if c.Bool("modify") {
			_, mtime, _, err = StatTimes(f)
			if err != nil {
				return err
			}
		}
		err = os.Chtimes(f, atime, mtime)
		if err != nil {
			return err
		}
	}

	return nil

}

func StatTimes(name string) (atime, mtime, ctime time.Time, err error) {
	fi, err := os.Stat(name)
	if err != nil {
		return
	}
	mtime = fi.ModTime()
	stat := fi.Sys().(*syscall.Stat_t)
	atime = time.Unix(int64(stat.Atim.Sec), int64(stat.Atim.Nsec))
	ctime = time.Unix(int64(stat.Ctim.Sec), int64(stat.Ctim.Nsec))
	return
}
