package command

import (
	"fmt"
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
	return fmt.Sprint("Create repository in Quay")
}

func (c *CreateCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
