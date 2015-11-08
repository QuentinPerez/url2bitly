package main

import (
	"os"

	"github.com/QuentinPerez/url2bitly/pkg/cli"
	"github.com/Sirupsen/logrus"
)

var VERSION string

func main() {
	ec, err := cli.Start(os.Args[1:], nil)
	if err != nil {
		logrus.Fatalf("%s", err)
	}
	os.Exit(ec)
}
