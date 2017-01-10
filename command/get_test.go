package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestGetCommand_implement(t *testing.T) {
	var _ cli.Command = &GetCommand{}
}
