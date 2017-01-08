package main

import (
	"github.com/koudaiii/dockerepos/command"
	"github.com/mitchellh/cli"
)

func Commands(meta *command.Meta) map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		"add-team": func() (cli.Command, error) {
			return &command.AddTeamCommand{
				Meta: *meta,
			}, nil
		},
		"delete-team": func() (cli.Command, error) {
			return &command.DeleteTeamCommand{
				Meta: *meta,
			}, nil
		},
		"add-user": func() (cli.Command, error) {
			return &command.AddUserCommand{
				Meta: *meta,
			}, nil
		},
		"delete-user": func() (cli.Command, error) {
			return &command.DeleteUserCommand{
				Meta: *meta,
			}, nil
		},
		"create": func() (cli.Command, error) {
			return &command.CreateCommand{
				Meta: *meta,
			}, nil
		},
		"destroy": func() (cli.Command, error) {
			return &command.DestroyCommand{
				Meta: *meta,
			}, nil
		},

		"version": func() (cli.Command, error) {
			return &command.VersionCommand{
				Meta:     *meta,
				Version:  Version,
				Revision: GitCommit,
				Name:     Name,
			}, nil
		},
	}
}
