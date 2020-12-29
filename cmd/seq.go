package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/urfave/cli/v2"
)

// SeqCMDVersion seq version no
var SeqCMDVersion = "v0.0.1"

// SeqCMD define seq cli command
var SeqCMD = cli.Command{
	Name:      "seq",
	Aliases:   []string{"SEQ"},
	UsageText: "core seq [选项]... 尾数\n\t core seq [选项]... 首数 尾数\n\t core seq [选项]... 首数 增量 尾数",
	Usage:     "打印数字序列(暂时只支持整数)",
	Description: `Print numbers from FIRST to LAST, in steps of INCREMENT.
	If FIRST or INCREMENT is omitted, it defaults to 1.  That is, an
	omitted INCREMENT defaults to 1 even when LAST is smaller than FIRST.
	The sequence of numbers ends when the sum of the current number and
	INCREMENT would become greater than LAST.
	FIRST, INCREMENT, and LAST are interpreted as floating point values.
	INCREMENT is usually positive if FIRST is smaller than LAST, and
	INCREMENT is usually negative if FIRST is greater than LAST.
	INCREMENT must not be 0; none of FIRST, INCREMENT and LAST may be NaN.
	FORMAT must be suitable for printing one argument of type 'double';
	it defaults to %.PRECf if FIRST, INCREMENT, and LAST are all fixed point
	decimal numbers with maximum precision PREC, and to %g otherwise.`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "输出版本信息并推出",
		}, &cli.StringFlag{
			Name:    "format",
			Aliases: []string{"f"},
			Usage:   "use printf style floating-point FORMAT（尚未实现）",
		}, &cli.StringFlag{
			Name:    "separator",
			Aliases: []string{"s"},
			Usage:   "use `STRING` to separate numbers",
			Value:   "\n",
		}, &cli.BoolFlag{
			Name:    "equal-width",
			Aliases: []string{"w"},
			Usage:   "equalize width by padding with leading zeroes（尚未实现）",
		},
	},
	Action: seqAction,
}

func seqAction(c *cli.Context) (err error) {
	if c.Bool("version") {
		fmt.Println(c.Command.Name, Base64CMDVersion)
		return nil
	}
	start := 1
	sep := 1
	end := 0
	switch c.Args().Len() {
	case 0:
		cli.ShowCommandHelp(c, c.Command.Name)
		return errors.New("缺少操作数")
	case 1:
		end, err = strconv.Atoi(c.Args().First())
		if err != nil {
			cli.ShowCommandHelp(c, c.Command.Name)
			return fmt.Errorf(`无效的参数 "%v"`, c.Args().First())
		}
	case 2:
		start, err = strconv.Atoi(c.Args().First())
		if err != nil {
			cli.ShowCommandHelp(c, c.Command.Name)
			return fmt.Errorf(`无效的参数 "%v"`, c.Args().First())
		}
		end, err = strconv.Atoi(c.Args().Get(1))
		if err != nil {
			cli.ShowCommandHelp(c, c.Command.Name)
			return fmt.Errorf(`无效的参数 "%v"`, c.Args().Get(1))
		}
	case 3:
		start, err = strconv.Atoi(c.Args().First())
		if err != nil {
			cli.ShowCommandHelp(c, c.Command.Name)
			return fmt.Errorf(`无效的参数 "%v"`, c.Args().First())

		}
		sep, err = strconv.Atoi(c.Args().Get(1))
		if err != nil {
			cli.ShowCommandHelp(c, c.Command.Name)
			return fmt.Errorf(`无效的参数 "%v"`, c.Args().Get(1))

		}
		end, err = strconv.Atoi(c.Args().Get(2))
		if err != nil {
			cli.ShowCommandHelp(c, c.Command.Name)
			return fmt.Errorf(`无效的参数 "%v"`, c.Args().Get(2))

		}
	default:
		cli.ShowCommandHelp(c, c.Command.Name)
		msg := fmt.Sprintf("额外的操作数 %v", c.Args().Slice()[3:])
		return errors.New(msg)

	}
	for i := start; i <= end; i += sep {
		fmt.Printf("%v", i)
		if i+sep <= end {
			fmt.Printf("%v", c.String("separator"))
		}
	}
	fmt.Println()
	return nil
}
