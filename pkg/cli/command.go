package cli

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/QuentinPerez/url2bitly/pkg/api"
	"github.com/QuentinPerez/url2bitly/pkg/command"
	"github.com/Sirupsen/logrus"
	"github.com/docker/docker/pkg/mflag"
)

var (
	ErrExitFailure = errors.New("exit 1")
	ErrExitSuccess = errors.New("exit 0")
)

type Command struct {
	// Exec executes the command
	Exec func(cmd *Command, args []string) error

	// Usage is the one-line usage message.
	UsageLine string

	// Description is the description of the command
	Description string

	// Help is the full description of the command
	Help string

	// Examples are some examples of the command
	Examples string

	// Flag is a set of flags specific to this command.
	Flag mflag.FlagSet

	// API is the interface used to communicate with Bitly's API
	API *api.API

	streams *command.Streams
}

// GetContext returns a standard context, with real stdin, stdout, stderr, a configured API and raw arguments
func (c *Command) GetContext(rawArgs []string) command.CommandContext {
	ctx := command.CommandContext{
		RawArgs: rawArgs,
		API:     c.API,
	}

	if c.streams != nil {
		ctx.Streams = *c.streams
	} else {
		ctx.Streams = command.Streams{
			Stdin:  os.Stdin,
			Stdout: os.Stdout,
			Stderr: os.Stderr,
		}
	}
	return ctx
}

func (c *Command) Streams() *command.Streams {
	if c.streams != nil {
		return c.streams
	}
	return &command.Streams{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

}

// Name returns the command's name
func (c *Command) Name() string {
	name := c.UsageLine
	i := strings.Index(name, " ")
	if i >= 0 {
		name = name[:i]
	}
	return name
}

// PrintUsage prints a full command usage
func (c *Command) PrintUsage() error {
	helpMessage, err := commandHelpMessage(c)
	if err != nil {
		logrus.Fatalf("%v", err)
	}
	fmt.Fprintf(c.Streams().Stdout, "%s\n", helpMessage)
	return ErrExitFailure
}

// PrintShortUsage prints a short command usage
func (c *Command) PrintShortUsage() error {
	fmt.Fprintf(c.Streams().Stderr, "usage: scw %s. See 'scw %s --help'.\n", c.UsageLine, c.Name())
	return ErrExitFailure
}

// Options returns a string describing options of the command
func (c *Command) Options() string {
	var options string
	visitor := func(flag *mflag.Flag) {
		name := strings.Join(flag.Names, ", -")
		var optionUsage string
		if flag.DefValue == "" {
			optionUsage = fmt.Sprintf("%s=\"\"", name)
		} else {
			optionUsage = fmt.Sprintf("%s=%s", name, flag.DefValue)
		}
		options += fmt.Sprintf("  -%-20s %s\n", optionUsage, flag.Usage)
	}
	c.Flag.VisitAll(visitor)
	if len(options) == 0 {
		return ""
	}
	return fmt.Sprintf("Options:\n\n%s", options)
}

// ExamplesHelp returns a string describing examples of the command
func (c *Command) ExamplesHelp() string {
	if c.Examples == "" {
		return ""
	}
	return fmt.Sprintf("Examples:\n\n%s", strings.Trim(c.Examples, "\n"))
}

var fullHelpTemplate = `
Usage: scw {{.UsageLine}}

{{.Help}}

{{.Options}}
{{.ExamplesHelp}}
`

func commandHelpMessage(cmd *Command) (string, error) {
	t := template.New("full")
	template.Must(t.Parse(fullHelpTemplate))
	var output bytes.Buffer
	err := t.Execute(&output, cmd)
	if err != nil {
		return "", err
	}
	return strings.Trim(output.String(), "\n"), nil
}
