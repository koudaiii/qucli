package command

import (
	"strings"
	"fmt"
)

type DeleteCommand struct {
	Meta
}

func (c *DeleteCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *DeleteCommand) Synopsis() string {
	return fmt.Sprint("Delete repository in Quay")
}

func (c *DeleteCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
