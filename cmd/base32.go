package cmd

import (
	"encoding/base32"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

// Base32CMDVersion base32 version no
var Base32CMDVersion = "v0.0.1"

// Base32CMD define base32 cli command
var Base32CMD = cli.Command{
	Name:    "base32",
	Aliases: []string{"BASE32"},
	Usage:   "base32编码/解码数据并打印到标准输出",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "output this help information and exit",
		}, &cli.BoolFlag{
			Name:    "decode",
			Aliases: []string{"d"},
			Usage:   "decode data",
		}, &cli.BoolFlag{
			Name:    "ignore-garbag",
			Aliases: []string{"i"},
			Usage:   "ignore non-char while decoding",
		}, &cli.Int64Flag{
			Name:    "wrap",
			Aliases: []string{"w"},
			Value:   76,
			Usage:   "在指定的字符数后自动换行(默认为76)，0 为禁用自动换行",
		},
	},
	Action: base32Action,
}

func base32Action(c *cli.Context) error {
	if c.Bool("version") {
		fmt.Println(c.Command.Name, Base32CMDVersion)
		return nil
	}
	target := c.Args().Get(0)
	s, err := os.Lstat(target)
	// err := syscall.Access(path, syscall.F_OK)
	if os.IsNotExist(err) {
		log.Fatalln("没有那个文件或目录")
	}
	if s.IsDir() {
		log.Fatalln("读取错误: 是一个目录")
	}
	f, err := ioutil.ReadFile(target)
	if err != nil {
		log.Fatalln(err)
	}
	if c.Bool("decode") {
		decodeBytes, err := base32.StdEncoding.DecodeString(string(f))
		if err != nil {
			log.Fatalln(err)
		}
		println(string(decodeBytes))
	} else {
		encodeString := base32.StdEncoding.EncodeToString(f)
		println(encodeString)
	}

	return nil

}
