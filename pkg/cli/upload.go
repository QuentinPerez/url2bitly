package cli

import "github.com/QuentinPerez/url2bitly/pkg/command"

type cmdUploadDefinition struct {
	command.CommandBase
}

var cmdUpload cmdUploadDefinition

func (up cmdUploadDefinition) GetName() string {
	return "upload"
}

func (up cmdUploadDefinition) Do(Args []string) error {
	return nil
}
