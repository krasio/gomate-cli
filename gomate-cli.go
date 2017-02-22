package main

import (
	"fmt"
	"os"

	"github.com/krasio/gomate"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "gomate"
	app.Usage = "autocompelate like a boss"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "redis, r",
			Value:  "redis://localhost:6379/0",
			Usage:  "Redis connection string",
			EnvVar: "GOMATE_REDIS_URL",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:        "load",
			ShortName:   "l",
			Usage:       "Replaces collection specified by TYPE with items read from stdin in the JSON lines format.",
			Description: "load [TYPE] < path/to/data.json",
			Action: func(c *cli.Context) {
				kind := c.Args()[0]

				conn, _ := gomate.Connect(c.GlobalString("redis"))
				defer conn.Close()

				gomate.BulkLoad(kind, os.Stdin, conn)
			},
		},
		{
			Name:        "query",
			ShortName:   "q",
			Usage:       "Queries for items from collection specified by TYPE.",
			Description: "query [TYPE] [TERM]",
			Action: func(c *cli.Context) {
				kind := c.Args()[0]
				query := c.Args()[1]

				conn, _ := gomate.Connect(c.GlobalString("redis"))
				defer conn.Close()

				fmt.Printf("Query %s for \"%s\":\n", kind, query)

				matches := gomate.Query(kind, query, conn)

				fmt.Println()
				if len(matches) > 0 {
					for _, match := range matches {
						fmt.Printf("  %s\n", match.Term)
					}
				} else {
					fmt.Println("  Sorry, nothing found. ¯\\_(ツ)_/¯")
				}
				fmt.Println()
			},
		},
	}

	app.Run(os.Args)
}
