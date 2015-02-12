package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	cli.AppHelpTemplate = customAppHelpTemplate
	
	app := cli.NewApp()
	app.Name = "mp"
	app.Version = Version
	app.Usage = "mp is simple analyzer for .mobileprovision file."
	app.Author = "henteko"
	app.Email = "henteko07@gmail.com"
	app.Action = Action
	app.Flags = Flags

	app.Run(os.Args)
}

const customAppHelpTemplate = `NAME:
	{{.Name}} - {{.Usage}}
	
USAGE:
	{{.Name}} [global options] [mobileprovision file path]
	
EXAMPLE:
	$ {{.Name}} /path/to/example.mobilepovision
	$ {{.Name}} --key UUID /path/to/example.mobileprovision
	$ {{.Name}} --install /path/to/example.mobileprovision
	
VERSION:
	{{.Version}}{{if or .Author .Email}}
	
AUTHOR:{{if .Author}}
	{{.Author}}{{if .Email}} - <{{.Email}}>{{end}}{{else}}
	{{.Email}}{{end}}{{end}}
	
GLOBAL OPTIONS:
   {{range .Flags}}{{.}}
   {{end}}
`
