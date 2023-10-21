package cmd

import (
	"github.com/Masamerc/cli-skeletons/cli-cli/pkg"
	"github.com/urfave/cli/v2"
)

// NewRootCmd creates a new root command
func NewRootCmd() *cli.App {
	app := &cli.App{
		Name:   "joke",
		Usage:  "Get a random joke",
		Action: pkg.PrintJoke,
	}
	app.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:    "upper",
			Aliases: []string{"u"},
			Usage:   "Convert joke to uppercase",
		},
	}
	return app
}
