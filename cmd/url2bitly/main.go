package main

import (
	"os"
	"path"

	"github.com/QuentinPerez/url2bitly"
	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

var VERSION string

func main() {
	app := cli.NewApp()
	app.Name = path.Base(os.Args[0])
	app.Author = "Quentin Perez"
	app.Email = "https://github.com/QuentinPerez/url2bitly"
	app.Version = VERSION
	app.Usage = "Convert your URLs to a bitly's URL"
	app.ArgsUsage = "URL"
	app.Action = action
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug, D",
			Usage: "Enable debug mode",
		},
		cli.BoolFlag{
			Name:  "expand, E",
			Usage: "Expand bitly's URL",
		},
	}

	app.Run(os.Args)
}

func action(c *cli.Context) {
	// setting up logrus
	logrus.SetOutput(os.Stderr)
	if c.Bool("debug") {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}
	url2bitly.Init()

	if c.Bool("expand") {
		url2bitly.Expand(c.Args())
		return
	}

	// if len(c.Args()) != 1 {
	// 	logrus.Fatalf("usage: %s [URL]", os.Args[0])
	// }
	url2bitly.Upload(c.Args())
}
