package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Sirupsen/logrus"
	"github.com/parnurzeal/gorequest"
)

var apiURL = "https://api-ssl.bitly.com"

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

func httpHandleError(answer *Answer) error {
	if answer.StatusCode != 200 {
		return fmt.Errorf("StatusCode %d: %s", answer.StatusCode, answer.StatusTxt)
	}
	return nil
}

func (api *API) GetResponse(resource ...string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", apiURL, filepath.Join(resource...))
	request := gorequest.New().Get(url)
	if api.debug {
		request = request.SetDebug(true)
	}
	_, body, errs := request.EndBytes()

	if len(errs) != 0 {
		return nil, printErrors(errs)
	}
	var answer Answer

	if err := json.Unmarshal(body, &answer); err != nil {
		return nil, err
	}
	if err := httpHandleError(&answer); err != nil {
		return nil, err
	}
	return body, nil
}

func NewAPI() *API {
	if userToken := os.Getenv("URL2BITLY_TOKEN"); userToken == "" {
		logrus.Fatal("Please set your token:\n\texport URL2BITLY_TOKEN=YOUR_TOKEN")
	} else {
		return &API{
			token: userToken,
			debug: os.Getenv("API_VERBOSE") != "",
		}
	}
	return nil
}
