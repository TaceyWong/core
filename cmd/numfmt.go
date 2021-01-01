package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// NumFmtCMDVersion numfmt version
const NumFmtCMDVersion = "0.0.1"

// NumFmtCMD define `numfmt` cmd
var NumFmtCMD = cli.Command{
	Name:    "numfmt",
	Aliases: []string{"NUMFMT"},
	Usage:   "将数字转换为人类可读的字符串",
	Description: `Reformat NUMBER(s), or the numbers from standard input if none are specified.
	
	UNIT options:
  none       no auto-scaling is done; suffixes will trigger an error
  auto       accept optional single/two letter suffix:
               1K = 1000,
               1Ki = 1024,
               1M = 1000000,
               1Mi = 1048576,
  si         accept optional single letter suffix:
               1K = 1000,
               1M = 1000000,
               ...
  iec        accept optional single letter suffix:
               1K = 1024,
               1M = 1048576,
               ...
  iec-i      accept optional two-letter suffix:
               1Ki = 1024,
               1Mi = 1048576,
               ...

FIELDS supports cut(1) style field ranges:
  N    N'th field, counted from 1
  N-   from N'th field, to end of line
  N-M  from N'th to M'th field (inclusive)
  -M   from first to M'th field (inclusive)
  -    all fields
Multiple fields/ranges can be separated with commas

FORMAT must be suitable for printing one floating-point argument '%f'.
Optional quote (%'f) will enable --grouping (if supported by current locale).
Optional width value (%10f) will pad output. Optional zero (%010f) width
will zero pad the number. Optional negative values (%-10f) will left align.
Optional precision (%.1f) will override the input determined precision.

Exit status is 0 if all input numbers were successfully converted.
By default, numfmt will stop at the first conversion error with exit status 2.
With --invalid='fail' a warning is printed for each conversion error
and the exit status is 2.  With --invalid='warn' each conversion error is
diagnosed, but the exit status is 0.  With --invalid='ignore' conversion
errors are not diagnosed and the exit status is 0.

Examples:
  $ numfmt --to=si 1000
            -> "1.0K"
  $ numfmt --to=iec 2048
           -> "2.0K"
  $ numfmt --to=iec-i 4096
           -> "4.0Ki"
  $ echo 1K | numfmt --from=si
           -> "1000"
  $ echo 1K | numfmt --from=iec
           -> "1024"
  $ df -B1 | numfmt --header --field 2-4 --to=si
  $ ls -l  | numfmt --header --field 5 --to=iec
  $ ls -lh | numfmt --header --field 5 --from=iec --padding=10
  $ ls -lh | numfmt --header --field 5 --from=iec --format %10f
	`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "输出版本信息并推出",
		},
	},
	Action: func(c *cli.Context) error {
		if c.Bool("version") {
			fmt.Println(NumFmtCMDVersion)
			return nil
		}
		return nil
	},
}
