package main

import (
	"log"
	"os"

	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	commandInstall,
	commandRead,
	commandDump,
}

var commandInstall = cli.Command{
	Name:  "install",
	Usage: "",
	Description: `
`,
	Action: doInstall,
}

var commandRead = cli.Command{
	Name:  "read",
	Usage: "",
	Description: `
`,
	Action: doRead,
}

var commandDump = cli.Command{
	Name:  "dump",
	Usage: "",
	Description: `
`,
	Action: doDump,
}

func debug(v ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		log.Println(v...)
	}
}

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func doInstall(c *cli.Context) {
}

func doRead(c *cli.Context) {
}

func doDump(c *cli.Context) {
}
