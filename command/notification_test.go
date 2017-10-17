package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestAddNotificationCommand_implement(t *testing.T) {
	var _ cli.Command = &AddNotificationCommand{}
}

func TestDeleteNotificationCommand_implement(t *testing.T) {
	var _ cli.Command = &DeleteNotificationCommand{}
}

func TestTestNotificationCommand_implement(t *testing.T) {
	var _ cli.Command = &TestNotificationCommand{}
}
