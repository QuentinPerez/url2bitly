package cli

import (
	"os"

	"github.com/QuentinPerez/url2bitly/pkg/api"
	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

var bitlyAPI *api.API

func init() {
	// bitlyAPI := api.NewAPI()
}

func Start(c *cli.Context) {
	logrus.SetOutput(os.Stderr)
	if c.Bool("debug") {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}
	for _, cmd := range commands {
		if c.Bool(cmd.GetName()) {
			cmd.Do(c.Args())
		}
	}
	// if c.Bool("expand") {
	// 	url2bitly.Expand(c.Args())
	// 	return
	// }

	// if err := url2bitly.Start(); err != nil {
	// 	logrus.Fatal(err)
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
