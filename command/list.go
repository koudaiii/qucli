package command

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/koudaiii/dockerepos/quay"
	"strconv"
)

type ListCommand struct {
	Meta
}

func (c *ListCommand) Run(args []string) int {
	var repositoryColumns = []string{"NAME", "isPublic", "DESCRIPTION"}

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

	repositoryPrint := new(tabwriter.Writer)
	repositoryPrint.Init(os.Stdout, 0, 8, 1, '\t', 0)

	fmt.Fprintln(repositoryPrint, strings.Join(repositoryColumns, "\t"))

	for _, repos := range repositories.Items {
		fmt.Fprintln(repositoryPrint, strings.Join(
			[]string{"quay.io/" + repos.Namespace + "/" + repos.Name, strconv.FormatBool(repos.IsPublic), repos.Description}, "\t",
		))
	}
	repositoryPrint.Flush()
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
