package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

// -------------------------------------------
// /= Example Functions to better understand.
// /= notes:
// /= commnd is just gonna be whatever the executable is named
// /= ALSO: please break this shit up these 8x nested brackets/tabs are brutal
// /==
// /= SOURCES: can set default values
// /-> by default can be either environment vars or text files
// /-> it is also possible to build/configure own sources
// /= SHORT OPTIONS:
// /-> set `UseShortOptionHandling` to true to allos use of short options (-a as opposed to --anus
// /-> these can only be either: a boolean flag OR a string flag specifically named "message"
// /? Q: What do aliases do besides this? probably can be used insntead of full flag name
// /= SUBCOMMANDS:
// /-> This example is using Commands: []*cli.Command{{name:, usage:, flags:, action:},{"}}
// /-> Found the explanation:
// /-> They operate like single-arguments that have their own functions
// /-> So basically, a nice qol, these could be built out of arguments but it would not be fun
// /-> `Category` can be added for programs with tons of subcommands (like git)
// /==
// -- need to learn how Help Text works, and error handling. For now just build shit
// -------------------------------------------
// /= ARGUMENTS:
// /-> use cli.Args() to get the arguments that were passed.
// /-> 1 arg = 1 word (space separated) coming after the .exe call.
// /-> look at cli.Args() to understand possibilities more.

type flagdata struct {
	str   []cli.StringFlag
	bool  []cli.BoolFlag
	int   []cli.IntFlag
	mutex []cli.MutuallyExclusiveFlags
}

type commandMaker struct {
	boolOptions map[string]bool

	//cli.StringMap stringmap used for tag name???

}

func cliFlagEX() {
	cmd := &cli.Command{
		UseShortOptionHandling: true, //note above
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "fart-style",
				Aliases: []string{"1", "fs"},          // can be used for shorthand not 100%
				Value:   "wet",                        //NOTE: The default value when flag not included
				Usage:   "the varietal of flatulence", //goes to help I believe
				// ok so the real ex uses "APP_LANG" here.
				Sources: cli.EnvVars("SET_FART", "FART_DEFAULT", "SYSTEM_FART"),
			},
			&cli.BoolFlag{Name: "RTX enabled", Aliases: []string{"r"}},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			name := "Bazinga"
			if cmd.NArg() > 0 {
				name = cmd.Args().Get(0)
			}
			if cmd.String("fart-style") == "wet" {
				fmt.Println("That was a wet ass fart!", name)
			} else {
				fmt.Println("That was a nice fart!", name)
			}
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func cliArgumentsEX() {
	cmd := &cli.Command{
		Action: func(ctx context.Context, cmd *cli.Command) error {
			fmt.Printf("Number of args : %d\n", cmd.Args().Len())
			var out string
			for i := 0; i < cmd.Args().Len(); i++ {
				out = out + fmt.Sprintf(" %v", cmd.Args().Get(i))
			}
			fmt.Printf("Hello%v", out)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
func cliSubcommandsEX() {
	cmd := &cli.Command{
		Commands: []*cli.Command{
			{
				Name:     "add",
				Aliases:  []string{"a"},
				Usage:    "add a task to the list",
				Category: "operation",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					fmt.Println("added task: ", cmd.Args().First())
					return nil
				},
			},
			{
				Name:    "complete",
				Aliases: []string{"c"},
				Usage:   "complete a task on the list",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					fmt.Println("completed task: ", cmd.Args().First())
					return nil
				},
			},
			{
				Name:    "template",
				Aliases: []string{"t"},
				Usage:   "options for task templates",
				Commands: []*cli.Command{
					{
						Name:  "add",
						Usage: "add a new template",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("new task template: ", cmd.Args().First())
							return nil
						},
					},
					{
						Name:  "remove",
						Usage: "remove an existing template",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("removed task template: ", cmd.Args().First())
							return nil
						},
					},
				},
			},
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
func cliBasicEX() { // the most basic, does nothing but print a line
	cmd := &cli.Command{
		Name:  "dotget",
		Usage: "config/dotfile backup",
		Action: func(context.Context, *cli.Command) error {
			fmt.Println("you did it great job yay woo yeah")
			return nil
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
