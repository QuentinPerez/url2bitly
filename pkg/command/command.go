package command

import "io"

type Streams struct {
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

type CommandBase struct {
	Streams

	Args []string
}

type Command interface {
	GetName() string
	Do(Args []string) error
}
