package cmd

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

// Base64CMDVersion base64 version no
var Base64CMDVersion = "v0.0.1"

// Base84CMD define base64 cli command
var Base64CMD = cli.Command{
	Name:    "base64",
	Aliases: []string{"BASE64"},
	Usage:   "encode or decode FILE, or standard input, to standard output",
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
	Action: base64Action,
}

//https://www.cnblogs.com/apocelipes/p/10199618.html
// 在判断文件是否存在前应该先判断自己对文件及其路径是否有访问权限

func base64Action(c *cli.Context) error {
	if c.Bool("version") {
		fmt.Println(c.Command.Name, Base64CMDVersion)
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
		decodeBytes, err := base64.StdEncoding.DecodeString(string(f))
		if err !=nil{
			log.Fatalln(err)
		}
		println(string(decodeBytes))
	} else {
		encodeString := base64.StdEncoding.EncodeToString(f)
		println(encodeString)
	}

	return nil

}
