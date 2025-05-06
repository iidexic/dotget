package main

import "github.com/urfave/cli/v3"

func getCmdLine() *cli.Command {

	cmd := &cli.Command{
		Commands: []*cli.Command{
			commandADD(),
		},
	}
	return cmd
}

func commandADD() *cli.Command {
	add := &cli.Command{
		Name:    "add",
		Aliases: []string{"a"},
		Usage:   "add a config to a registry", //catalog? collection?
	}
	return add
}

func addToReg() {

}
