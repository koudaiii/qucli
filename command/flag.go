package command

import (
	flag "github.com/spf13/pflag"
)

var (
	role       string
	visibility string
)

func FlagInit(args []string) error {
	flags := flag.NewFlagSet("dockerepos", flag.ExitOnError)

	flags.Usage = func() {
		flags.PrintDefaults()
	}

	flags.StringVar(&visibility, "visibility", "public", "visibility set to 'public' or 'private'. default: public")
	flags.StringVar(&role, "role", "read", "role to use for the user =  ['read', 'write', 'admin']. default: read")

	if err := flags.Parse(args[0:]); err != nil {
		return err
	}
	return nil
}
