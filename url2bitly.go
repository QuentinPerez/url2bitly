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
}

func Expand(urls []string) {
	for _, url := range urls {
		if expand, err := bitlyAPI.GetExpand(url); err != nil {
			logrus.Warn(err)
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
