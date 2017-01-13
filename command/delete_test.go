package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestDeleteCommand_implement(t *testing.T) {
	var _ cli.Command = &DeleteCommand{}
}
