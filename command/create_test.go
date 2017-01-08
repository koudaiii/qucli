package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestCreateCommand_implement(t *testing.T) {
	var _ cli.Command = &CreateCommand{}
}
