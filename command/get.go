package command

import (
	"fmt"
	"os"
	"strings"

	"github.com/koudaiii/dockerepos/quay"
)

type GetCommand struct {
	Meta
}

func (c *GetCommand) Run(args []string) int {
	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, c.Help())
		os.Exit(1)
	}


	ss := strings.Split(args[0], "/")
	if len(ss) != 3 {
		fmt.Fprintln(os.Stderr, c.Help())
		os.Exit(1)
	}

	resp, err := quay.GetRepository(ss[1], ss[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, "repository: quay.io/%v/%v\n", resp.Name, resp.Namespace)
	return 0
}

func (c *GetCommand) Synopsis() string {
	return fmt.Sprint("Get repository and Permissions in Quay")
}

func (c *GetCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
