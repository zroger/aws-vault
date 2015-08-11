package main

import (
	"log"
	"os"

	"github.com/99designs/aws-vault/command"
	"github.com/99designs/aws-vault/keyring"
	"github.com/mitchellh/cli"
)

var (
	Version string
)

func main() {
	ui := &cli.BasicUi{
		Writer: os.Stdout,
		Reader: os.Stdin,
	}

	k := keyring.DefaultKeyring

	c := cli.NewCLI("aws-vault", Version)
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"add": func() (cli.Command, error) {
			return &command.AddCommand{
				Ui:      ui,
				Keyring: k,
			}, nil
		},
		"rm": func() (cli.Command, error) {
			return &command.RemoveCommand{
				Ui:      ui,
				Keyring: k,
			}, nil
		},
		"exec": func() (cli.Command, error) {
			return &command.ExecCommand{
				Ui:      ui,
				Keyring: k,
			}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
