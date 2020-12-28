package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// JoinCMDVersion join version
const JoinCMDVersion = "0.0.1"

// JoinCMD define `join` cmd
var JoinCMD = cli.Command{
	Name:    "join",
	Aliases: []string{"JOIN"},
	Usage:   "join lines of two files on a common field",
	Description: `For each pair of input lines with identical join fields, write a line to
	standard output.  The default join field is the first, delimited by blanks.
	
	When FILE1 or FILE2 (not both) is -, read standard input.
	
	Unless -t CHAR is given, leading blanks separate fields and are ignored,
	else fields are separated by CHAR.  Any FIELD is a field number counted
	from 1.  FORMAT is one or more comma or blank separated specifications,
	each being 'FILENUM.FIELD' or '0'.  Default FORMAT outputs the join field,
	the remaining fields from FILE1, the remaining fields from FILE2, all
	separated by CHAR.  If FORMAT is the keyword 'auto', then the first
	line of each file determines the number of fields output for each line.

	Important: FILE1 and FILE2 must be sorted on the join fields.
	E.g., use "sort -k 1b,1" if 'join' has no options,
	or use "join -t ''" if 'sort' has no options.
	Note, comparisons honor the rules specified by 'LC_COLLATE'.
	If the input is not sorted and some lines cannot be joined, a
	warning message will be given.`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "输出版本信息并推出",
		},
	},
	Action: func(c *cli.Context) error {
		if c.Bool("version") {
			fmt.Println(JoinCMDVersion)
			return nil
		}
		return nil
	},
}
