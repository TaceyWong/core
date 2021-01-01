package cmd

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/urfave/cli/v2"
)

// MkTempCMDVersion mktemp version no
var MkTempCMDVersion = "v0.0.1"

// MkTempCMD define mktemp cli command
var MkTempCMD = cli.Command{
	Name:    "mktemp",
	Aliases: []string{"MKTEMP"},
	Usage:   "创建临时文件、临时目录",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "输出版本信息并推出",
		},
		&cli.BoolFlag{
			Name:    "directory",
			Aliases: []string{"d"},
			Usage:   "create a directory, not a file",
		},
		&cli.BoolFlag{
			Name:    "dry-run",
			Aliases: []string{"u"},
			Usage:   "do not create anything; merely print a name (unsafe)",
		},
		&cli.StringFlag{
			Name:    "suffix",
			Aliases: []string{"s"},
			Usage:   "append SUFF to TEMPLATE; SUFF must not contain a slash.  This option is implied if TEMPLATE does not end in X",
		},
		&cli.StringFlag{
			Name:    "tmpdir",
			Aliases: []string{"p"},
			Usage: `interpret TEMPLATE relative to DIR; if DIR is not specified, use $TMPDIR if set, else /tmp.  With this option, TEMPLATE must not be an absolute name; unlike  with  -t,  TEM‐
			PLATE may contain slashes, but mktemp creates only the final component`,
		},
		&cli.BoolFlag{
			Name:  "t",
			Usage: "interpret TEMPLATE as a single file name component, relative to a directory: $TMPDIR, if set; else the directory specified via -p; else /tmp [deprecated]",
		},
	},
	Action: mktemp,
}

func mktemp(c *cli.Context) error {
	if c.Bool("version") {
		fmt.Println(c.Command.Name, ArchCMDVersion)
		return nil
	}
	suffix := c.String("suffix")
	tempdir := c.String("temp-dir")
	if tempdir == "" {
		tempdir = "/tmp"
	}

	if c.Bool("directory") {
		dir, err := ioutil.TempDir(tempdir, suffix)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(dir)
		// defer os.RemoveAll(dir)
	} else {
		file, err := ioutil.TempFile(tempdir, suffix)
		if err != nil {
			log.Fatal(err)
		}
		// defer os.Remove(file.Name())
		fmt.Println(file.Name()) // For example "dir/myname.054003078.bat"
	}
	return nil

}
