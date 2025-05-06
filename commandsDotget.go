package main

import "github.com/urfave/cli/v3"

func makeCommands() []*cli.Command {
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
		// NOTE: can flags take arguments?
		&cli.BoolFlag{Name: "envVar", Aliases: []string{"e"}, Usage: "store config location in an environment variable instead of file in exe dir"},
		&cli.StringFlag{Name: "dir", Aliases: []string{"d"}, Usage: "specify a main directory to store dotget's config"},
		&cli.StringFlag{Name: "path"},
		&cli.BoolFlag{Name: "notEmpty", Aliases: []string{"n"}, Usage: "allow config to be placed in an existing non-empty directory"},
		&cli.BoolFlag{Name: "ForceOverwriteConfig", Usage: "Force write config, will overwrite any other file at destination path"},
	}

	//# Command set
	cset := &cli.Command{
		Name:                   "set",
		UseShortOptionHandling: true,
		Commands: []*cli.Command{
			{Name: "make", Usage: "Make a new set for app configs"},
			{Name: "switch", Usage: "Switch to modifying/targeting a different set"},
			{Name: "merge", Usage: "Merge two or more sets into one"},
		},
	}
	cset.Flags = []cli.Flag{}
	cmds := []*cli.Command{cinit, cset}
	return cmds
}
