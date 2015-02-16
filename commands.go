package main

import (
	"log"
	"os"
	"fmt"

	"github.com/codegangsta/cli"
)

var Flags = []cli.Flag {
	cli.StringFlag{
		Name: "key, k",
		Usage: "get key value from mobileprovision file",
	},
	cli.BoolFlag{
		Name: "install, i",
		Usage: "install to mobileprovision file",
	},
	cli.StringFlag{
		Name: "list, l",
		Usage: "install list for mobileprovision files",
	},
}

var Action = doAction

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

func doAction(c *cli.Context) {
	// listの場合はmobileprovision fileはいらない
	// TODO: 空の場合はallにする
	if c.String("list") != "" {
		doInstallList(c.String("list"))
		os.Exit(0)
	}

	if len(c.Args()) <= 0 {
		fmt.Println("Please input *.mobileprovision file path")
		os.Exit(1)
	}
	mobileProvisioningFilePath := c.Args()[0]
	
	if c.Bool("install") {
		doInstall(mobileProvisioningFilePath)
	} else if c.String("key") != "" {
		doRead(mobileProvisioningFilePath, c.String("key"))
	}else {
		doDump(mobileProvisioningFilePath)
	}
}

func doInstall(mobileProvisioningFilePath string) {
	install(mobileProvisioningFilePath)
}

func doInstallList(mobileProvisioningName string) {
	installList(mobileProvisioningName)
}

func doRead(mobileProvisioningFilePath string, key string) {
	plist := getPlistData(mobileProvisioningFilePath)
	
	switch key {
	case "AppIDName":fmt.Println(plist.AppIDName)
	case "Name": fmt.Println(plist.Name)
	case "TeamName": fmt.Println(plist.TeamName)
	case "TimeToLive": fmt.Println(plist.TimeToLive)
	case "UUID": fmt.Println(plist.UUID)
	case "Version": fmt.Println(plist.Version)
	default: fmt.Println("Not support key...")
	}
}

func doDump(mobileProvisioningFilePath string) {
	plistString := getString(mobileProvisioningFilePath)
	fmt.Println(plistString)
}
