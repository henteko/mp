package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "mp"
	app.Version = Version
	app.Usage = ""
	app.Author = "henteko"
	app.Email = "henteko07@gmail.com"
	app.Commands = Commands

	app.Run(os.Args)
}
