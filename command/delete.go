package command

import (
	"fmt"
	"os"
	"strings"

	"github.com/koudaiii/qcli/quay"
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

	err := quay.DeleteRepository(ss[1], ss[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, "Deleted! quay.io/%v/%v\n", ss[1], ss[2])
	return 0
}

func (c *DeleteCommand) Synopsis() string {
	return fmt.Sprint("Delete repository in Quay")
}

func (c *DeleteCommand) Help() string {
	helpText := `
qcli supported only Quay.io
Usage: delete
  qcli delete quay.io/koudaiii/qcli
`
	return strings.TrimSpace(helpText)
}
