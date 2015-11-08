// Copyright (C) 2015 Scaleway. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.md file.

// Package commands contains the workflows behind each commands of the CLI (run, attach, start, exec, commit, ...)
package command

import (
	"io"

	"github.com/QuentinPerez/url2bitly/pkg/api"
)

// Streams is used to redirects the streams
type Streams struct {
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

// CommandContext is passed to all commands and contains streams, environment, api and arguments
type CommandContext struct {
	Streams

	RawArgs []string
	API     *api.API
}
