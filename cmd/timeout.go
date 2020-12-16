package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// TimeoutCMDVersion timeout version
const TimeoutCMDVersion = "0.0.1"

// TimeoutCMD define `timeout` cmd
var TimeoutCMD = cli.Command{
	Name:    "timeout",
	Aliases: []string{"TIMEOUT"},
	Usage:   "有时间限制地运行命令",
	Description: `DURATION is a floating point number with an optional suffix:
	's' for seconds (the default), 'm' for minutes, 'h' for hours or 'd' for days.
	
	If the command times out, and --preserve-status is not set, then exit with
	status 124.  Otherwise, exit with the status of COMMAND.  If no signal
	is specified, send the TERM signal upon timeout.  The TERM signal kills
	any process that does not block or catch that signal.  It may be necessary
	to use the KILL (9) signal, since this signal cannot be caught, in which
	case the exit status is 128+9 rather than 124.`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "output version information and exit",
		},
	},
	Action: func(c *cli.Context) error {
		if c.Bool("version") {
			fmt.Println(TimeoutCMDVersion)
			return nil
		}
		return nil
	},
}
