package api

import (
	"errors"
	"os"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/parnurzeal/gorequest"
)

var apiURL = "api-ssl.bitly.com"

type API struct {
	endpoint string
	token    string
	debug    bool
}

func printErrors(errs []error) error {
	for _, err := range errs {
		logrus.Error(err)
	}
	return errors.New("Error(s) has occured")
}

func (api *API) GetResponse(url string) ([]byte, error) {
	request := gorequest.New().Get(strings.Join([]string{apiURL, url}, "/"))
	if api.debug {
		request = request.SetDebug(true)
	}
	_, body, errs := request.EndBytes()

	if len(errs) != 0 {
		return nil, printErrors(errs)
	}
	return body, nil
}

func NewAPI() *API {
	if userToken := os.Getenv("URL2BITLY_TOKEN"); userToken == "" {
		logrus.Fatal("Please set your token:\n\texport URL2BITLY_TOKEN=YOUR_TOKEN")
	} else {
		return &API{
			token: userToken,
			debug: os.Getenv("VERBOSE_API") != "",
		}
	}
	return nil
}
