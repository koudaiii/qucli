package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestAddTeamCommand_implement(t *testing.T) {
	var _ cli.Command = &AddTeamCommand{}
}

func TestDeleteTeamCommand_implement(t *testing.T) {
	var _ cli.Command = &DeleteTeamCommand{}
}
