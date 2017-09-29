package command

import (
	"fmt"
	"os"
	"strings"

	"github.com/koudaiii/qucli/quay"
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
	if len(ss) != 2 {
		fmt.Fprintln(os.Stderr, c.Help())
		os.Exit(1)
	}

	err := quay.DeleteRepository(ss[0], ss[1], hostname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, "Deleted! quay.io/%v/%v\n", ss[0], ss[1])
	return 0
}

func (c *DeleteCommand) Synopsis() string {
	return fmt.Sprint("Delete repository in Quay")
}

func (c *DeleteCommand) Help() string {
	helpText := `
qucli supported only Quay.io
Usage: delete
  qucli delete koudaiii/qucli
`
	return strings.TrimSpace(helpText)
}
