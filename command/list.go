package command

import (
	"fmt"
	"os"
	"strings"

	"github.com/koudaiii/dockerepos/quay"
)

type ListCommand struct {
	Meta
}

func (c *ListCommand) Run(args []string) int {
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, c.Help())
		os.Exit(1)
	}

	if err := FlagInit(args); err != nil {
		fmt.Fprintln(os.Stderr, c.Help())
		os.Exit(1)
	}

	repositories, err := quay.ListRepository(args[0], public)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", err)
		os.Exit(1)
	}

	for _, repos := range repositories.Items {
		fmt.Fprintf(os.Stdout, "%v\tquay.io/%v/%v\n", repos.IsPublic, repos.Namespace, repos.Name)
	}
	return 0
}

func (c *ListCommand) Synopsis() string {
	return fmt.Sprint("List repository and Permissions in Quay")
}

func (c *ListCommand) Help() string {
	helpText := `
dockerepos supported only Quay.io
Usage: list
  dockerepos list
`
	return strings.TrimSpace(helpText)
}
