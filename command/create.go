package command

import (
	"strings"
)

type CreateCommand struct {
	Meta
}

func (c *CreateCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *CreateCommand) Synopsis() string {
	return ""
}

func (c *CreateCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
