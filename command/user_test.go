package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestAddUserCommand_implement(t *testing.T) {
	var _ cli.Command = &AddUserCommand{}
}

func TestDeleteUserCommand_implement(t *testing.T) {
	var _ cli.Command = &DeleteUserCommand{}
}
