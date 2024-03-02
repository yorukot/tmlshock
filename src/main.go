package main

import (
	"log"
	"os"

	"github.com/MHNightCat/tshock/util"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "color",
				Aliases: []string{"c"},
				Usage:   "Set the string color see full color in()",
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "stopwatch",
				Aliases: []string{"s"},
				Usage:   "Start a stopwatch",
				Action:  util.Stopwatch,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "color",
						Aliases: []string{"c"},
						Usage:   "Set the string color",
					},
				},
			},
			{
				Name:    "clock",
				Aliases: []string{"c"},
				Usage:   "Start a clock",
				Action:  util.Clock,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "color",
						Aliases: []string{"c"},
						Usage:   "Set the string color",
					},
					&cli.StringFlag{
						Name:    "second",
						Aliases: []string{"s", "sec"},
						Value:   "true",
						Usage:   "Set the clock with second(true or false)",
					},

					&cli.StringFlag{
						Name:    "date",
						Aliases: []string{"d"},
						Value:   "false",
						Usage:   "Set the clock with date(true or false)",
					},
					&cli.StringFlag{
						Name:    "dateformate",
						Aliases: []string{"df"},
						Value:   "2006/02/01",
						Usage:   "Set the clock date formate",
					},
				},
			},
		},
	}
	app.Name = "tshock"
	app.Version = "0.5"
	app.Usage = "a beauty clock and stopwatch"
	app.UsageText = "tshock [command] [option]"

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
