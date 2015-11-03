package url2bitly

import (
	"fmt"

	"github.com/QuentinPerez/url2bitly/pkg/api"
	"github.com/Sirupsen/logrus"
)

var bitlyAPI *api.API

func Init() {
	bitlyAPI = api.NewAPI()
}

func Upload(urls []string) {
	for _, url := range urls {
		if upload, err := bitlyAPI.GetLinkSave(url); err != nil {
			logrus.Error(err)
		} else {
			fmt.Printf("%-16s %s\n", "Link:", upload.Data.LinkSave.Link)
			fmt.Printf("%-16s %s\n", "AggreagateLink:", upload.Data.LinkSave.AggregateLink)
			fmt.Printf("%-16s %s\n", "LongURL:", upload.Data.LinkSave.LongURL)
		}
	}
}

func Expand(urls []string) {
	for _, url := range urls {
		if expand, err := bitlyAPI.GetExpand(url); err != nil {
			logrus.Error(err)
		} else {
			if expand.Data.Expand[0].Error == "" {
				fmt.Printf("%-12s %s\n", "ShortURL:", expand.Data.Expand[0].ShortURL)
				fmt.Printf("%-12s %s\n", "LongURL:", expand.Data.Expand[0].LongURL)
				fmt.Printf("%-12s %s\n", "UserHash:", expand.Data.Expand[0].UserHash)
				fmt.Printf("%-12s %s\n", "GlobalHash:", expand.Data.Expand[0].GlobalHash)
			} else {
				fmt.Printf("%-12s %s\n", "Error:", expand.Data.Expand[0].Error)
			}
		}
	}
}
