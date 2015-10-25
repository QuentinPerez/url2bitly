package main

import (
	"os"
	"path"

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
	}

	app.Run(os.Args)
}

func action(c *cli.Context) {
	if len(c.Args()) != 1 {
		logrus.Fatalf("usage: %s [URL]", os.Args[0])
	}

	// setting up logrus
	logrus.SetOutput(os.Stderr)
	if c.Bool("debug") {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

}