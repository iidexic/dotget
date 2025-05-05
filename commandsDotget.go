package main

import "github.com/urfave/cli/v3"

func commandInit() {
	cinit := &cli.Command{
		Name:                   "init",
		UseShortOptionHandling: true,
		Usage: `
    initialize dotget's config. Creates a config file in a provided or standard location.
    By default, this also creates a file in the executable's dir pointing to the config location.
    `,
		Category: "initialize",
	}
	cinit.Flags = []cli.Flag{
		// WARNING: DIR/PATH FLAGS MAY NEED TO BE ARGUMENTS OR SUBCOMMANDS.
		// I do not know if flags themselves support arguments like this?
		// If not just add more usage details of init function with arg callout
		&cli.BoolFlag{Name: "envVar", Aliases: []string{"e"}, Usage: "store config location in an environment variable instead of file in exe dir"},
		&cli.StringFlag{Name: "dir", Aliases: []string{"d"}, Usage: "specify a main directory to store dotget's config"},
		&cli.StringFlag{Name: "path"},
		&cli.BoolFlag{Name: "notEmpty", Aliases: []string{"n"}, Usage: "allow config to be placed in an existing non-empty directory"},
		&cli.BoolFlag{Name: "ForceOverwriteConfig", Usage: "Force write config, will overwrite any other file at destination path"},
	}
}
