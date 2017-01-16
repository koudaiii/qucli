package command

import (
	"fmt"
	"strings"
	"os"

	"github.com/koudaiii/dockerepos/quay"
)

type DeleteCommand struct {
	Meta
}

func (c *DeleteCommand) Run(args []string) int {
	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, c.Help())
		os.Exit(1)
	}

	ss := strings.Split(args[0], "/")
	if len(ss) != 3 {
		fmt.Fprintln(os.Stderr, c.Help())
		os.Exit(1)
	}

	repos, err := quay.DeleteRepository(ss[1], ss[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, "Deleted! quay.io/%v/%v\n", repos.Namespace, repos.Name)
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
