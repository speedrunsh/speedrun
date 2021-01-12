package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "speedrun",
		Authors: []*cli.Author{{
			Name:  "Dawid Pogorzelski",
			Email: "dawid@pogorzelski.dev",
		}},
		Before: loadConfig,
		Commands: []*cli.Command{
			{
				Name:   "init",
				Usage:  "Initialize speedrun",
				Action: initialize,
			},
			{
				Name:  "run",
				Usage: "Runs a command on remote servers",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "filter", Usage: "gcloud resource filter"},
					&cli.BoolFlag{Name: "only-failures", Usage: "print only failures and errors"},
					&cli.BoolFlag{Name: "private-ip", Usage: "connect to private IPs instead of public ones"},
				},
				Before: configInitialized,
				Action: run,
				UsageText: "speedrun run [command options] <command to send>\n\n" +
					"EXAMPLES:\n" +
					"   speedrun run \"uname -r\"\n" +
					"   speedrun run --only-failures --filter \"labels.foo = bar AND labels.environment = staging\" \"uname -r\"",
			},
			{
				Name:  "key",
				Usage: "Manage ssh keys",
				Subcommands: []*cli.Command{
					{
						Name:   "new",
						Usage:  "Create a new ssh key",
						Before: configInitialized,
						Action: createKey,
					},
					{
						Name:   "show",
						Usage:  "Show current ssh key",
						Before: configInitialized,
						Action: showKey,
					},
					{
						Name:  "set",
						Usage: "Set key in the project's metadata",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "gcloud resource filter"},
						},
						Before: configInitialized,
						Action: setKey,
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
