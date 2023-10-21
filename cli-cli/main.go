package main

import (
	"log"
	"os"

	"github.com/Masamerc/cli-skeletons/cli-cli/cmd"
)

func main() {
	app := cmd.NewRootCmd()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
