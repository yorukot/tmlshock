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
				Usage:   "Set the string color see full color in(https://github.com/MHNightCat/tmlshock?tab=readme-ov-file#color)",
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
						Usage:   "Set the string color (https://github.com/MHNightCat/tmlshock?tab=readme-ov-file#color)",
					},
					&cli.StringFlag{
						Name:    "disable-hour",
						Aliases: []string{"dh"},
						Usage:   "Disable hour(true or false)",
					},
					&cli.StringFlag{
						Name:    "colon-color",
						Aliases: []string{"cc"},
						Usage:   "Set the colon color (https://github.com/MHNightCat/tmlshock?tab=readme-ov-file#color)",
					},
				},
			},
			{
				Name:    "timer",
				Aliases: []string{"t"},
				Usage:   "Start a timer",
				Action:  util.Timer,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "color",
						Aliases: []string{"c"},
						Usage:   "Set the string color (https://github.com/MHNightCat/tmlshock?tab=readme-ov-file#color)",
					},
					&cli.StringFlag{
						Name:    "hour",
						Aliases: []string{"hr"},
						Usage:   "Enter how many hours you want to count down",
					},
					&cli.StringFlag{
						Name:    "minute",
						Aliases: []string{"m", "min"},
						Usage:   "Enter how many minunts you want to count down",
					},
					&cli.StringFlag{
						Name:    "second",
						Aliases: []string{"s", "sec"},
						Usage:   "Enter how many seconds you want to count down",
					},
					&cli.StringFlag{
						Name:    "time",
						Aliases: []string{"t"},
						Usage:   "Enter how many time you want to count down(format: 00:00:00)",
					},
					&cli.StringFlag{
						Name:    "disable-hour",
						Aliases: []string{"dh"},
						Usage:   "Disable hour(true or false)",
					},
					&cli.StringFlag{
						Name:    "disable-millisecond",
						Aliases: []string{"dm"},
						Usage:   "Disable millisecond(true or false)",
					},
					&cli.StringFlag{
						Name:    "colon-color",
						Aliases: []string{"cc"},
						Usage:   "Set the colon color (https://github.com/MHNightCat/tmlshock?tab=readme-ov-file#color)",
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
						Usage:   "Set the clock date formate(2006/02/01)YYYY/MM/DD",
					},
					&cli.StringFlag{
						Name:    "colon-color",
						Aliases: []string{"cc"},
						Usage:   "Set the colon color (https://github.com/MHNightCat/tmlshock?tab=readme-ov-file#color)",
					},
					&cli.StringFlag{
						Name:    "hour-format",
						Aliases: []string{"hf"},
						Usage:   "Set the clock 24 hr or 12hr (type 24 or 12)",
					},
				},
			},
		},
	}
	app.Name = "tmlshock"
	app.Version = "1.2"
	app.Usage = "A terminal ttl clock timer and stopwatch build by golang "
	app.UsageText = "tmlshock [command] [option]"

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
