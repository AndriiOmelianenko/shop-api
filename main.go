package main

import (
	"os"

	"github.com/AndriiOmelianenko/shop-api/actions"
	"github.com/AndriiOmelianenko/shop-api/dao"
	"github.com/urfave/cli"
)

// main is the starting point to shop-api application.
func main() {
	app := cli.NewApp()
	app.Name = "shop-api"
	app.Usage = "Simple and lightweight Shop API"
	app.Version = "0.1.0"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "mongo",
			Usage:  "MongoDB address URL in format: mongodb://<user>:<password>@<address>:<port>",
			EnvVar: "MONGO",
			Value:  "mongodb://localhost:27017",
		},
		cli.StringFlag{
			Name:   "dbname",
			Usage:  "Database name for Shop",
			EnvVar: "DBNAME",
			Value:  "shop",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:      "daemon",
			ShortName: "d",
			Usage:     "Start API",
			Action:    actions.Serve,
		},
		{
			Name:      "seed",
			ShortName: "s",
			Usage:     "Seed the database with random values",
			Action:    dao.SeedDatabase,
		},
	}

	app.Run(os.Args)
}
