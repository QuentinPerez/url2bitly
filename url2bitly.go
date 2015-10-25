package url2bitly

import (
	"fmt"

	"github.com/QuentinPerez/url2bitly/pkg/api"
)

var bitlyAPI *api.API

func init() {
	bitlyAPI = api.NewAPI()
}

func Upload(urls []string) {
}

func Expand(urls []string) {
	for _, url := range urls {
		fmt.Println(url)
	}
}
