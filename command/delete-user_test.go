package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestDeleteUserCommand_implement(t *testing.T) {
	var _ cli.Command = &DeleteUserCommand{}
}
