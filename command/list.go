package command

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"strconv"

	"github.com/koudaiii/qucli/quay"
)

type ListCommand struct {
	Meta
}

func (c *ListCommand) Run(args []string) int {
	var repositoryColumns = []string{"NAME", "isPublic", "DESCRIPTION"}

	if err := FlagInit(args); err != nil {
		fmt.Fprintln(os.Stderr, c.Help())
		os.Exit(1)
	}

	if len(subcommandArgs) < 1 {
		fmt.Fprintln(os.Stderr, c.Help())
		os.Exit(1)
	}

	repositories, err := quay.ListRepository(subcommandArgs[0], public, hostname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", err)
		os.Exit(1)
	}

	repositoryPrint := new(tabwriter.Writer)
	repositoryPrint.Init(os.Stdout, 0, 8, 1, '\t', 0)

	fmt.Fprintln(repositoryPrint, strings.Join(repositoryColumns, "\t"))

	for _, repos := range repositories.Items {
		if public == repos.IsPublic {
			fmt.Fprintln(repositoryPrint, strings.Join(
				[]string{hostname + "/" + repos.Namespace + "/" + repos.Name, strconv.FormatBool(repos.IsPublic), repos.Description}, "\t",
			))
		}
	}
	repositoryPrint.Flush()
	return 0
}

func (c *ListCommand) Synopsis() string {
	return fmt.Sprint("List repository and Permissions in Quay")
}

func (c *ListCommand) Help() string {
	helpText := `
qucli supported only Quay.io
Usage: list
  qucli list koudaiii
`
	return strings.TrimSpace(helpText)
}
