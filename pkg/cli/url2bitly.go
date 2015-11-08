package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/QuentinPerez/url2bitly/pkg/command"
	"github.com/docker/docker/pkg/mflag"
)

var (
	flDebug = mflag.Bool([]string{"D", "-debug"}, false, "Enable debug mode")
)

func Start(rawArgs []string, streams *command.Streams) (int, error) {
	if streams == nil {
		streams = &command.Streams{
			Stdin:  os.Stdin,
			Stdout: os.Stdout,
			Stderr: os.Stderr,
		}
	}
	if err := flag.CommandLine.Parse(rawArgs); err != nil {
		return 1, fmt.Errorf("Start.Parse: %v", err)
	}
	if *flDebug {
		os.Setenv("DEBUG", "1")
	}
	args := flag.Args()
	if len(args) < 1 {
		// CmdHelp.Exec(CmdHelp, []string{})
		return 1, nil
	}
	// name := args[0]
	args = args[1:]
	for _, cmd := range commands {
		cmd.streams = streams
	}
	return 0, nil

	// logrus.SetOutput(os.Stderr)
	// if c.Bool("debug") {
	// 	logrus.SetLevel(logrus.DebugLevel)
	// } else {
	// 	logrus.SetLevel(logrus.InfoLevel)
	// }
	// for _, cmd := range commands {
	// if c.Bool(cmd.Name()) {
	// 	cmd.Exec(cmd, c.Args())
	// }
	// }
}

// func Upload(urls []string) {
// 	for _, url := range urls {
// 		if upload, err := bitlyAPI.GetLinkSave(url); err != nil {
// 			logrus.Error(err)
// 		} else {
// 			fmt.Printf("%-16s %s\n", "Link:", upload.Data.LinkSave.Link)
// 			fmt.Printf("%-16s %s\n", "AggreagateLink:", upload.Data.LinkSave.AggregateLink)
// 			fmt.Printf("%-16s %s\n", "LongURL:", upload.Data.LinkSave.LongURL)
// 		}
// 	}
// }
//
// func Expand(urls []string) {
// 	for _, url := range urls {
// 		if expand, err := bitlyAPI.GetExpand(url); err != nil {
// 			logrus.Error(err)
// 		} else {
// 			if expand.Data.Expand[0].Error == "" {
// 				fmt.Printf("%-12s %s\n", "ShortURL:", expand.Data.Expand[0].ShortURL)
// 				fmt.Printf("%-12s %s\n", "LongURL:", expand.Data.Expand[0].LongURL)
// 				fmt.Printf("%-12s %s\n", "UserHash:", expand.Data.Expand[0].UserHash)
// 				fmt.Printf("%-12s %s\n", "GlobalHash:", expand.Data.Expand[0].GlobalHash)
// 			} else {
// 				fmt.Printf("%-12s %s\n", "Error:", expand.Data.Expand[0].Error)
// 			}
// 		}
// 	}
// }
