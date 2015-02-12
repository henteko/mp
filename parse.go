package main

import (
	"os"
	"os/exec"
	"bytes"
	"fmt"

	"howett.net/plist"
)

type smobileProvisioningFileParseBundleHeader struct {
	AppIDName string `plist:"AppIDName"`
	Name string `plist:"Name"`
	TeamName string `plist:"TeamName"`
	TimeToLive int `plist:"TimeToLive"`
	UUID string `plist:"UUID"`
	Version int `plist:"Version"`
}

func getPlistData(mobileprovisioningFilePath string) smobileProvisioningFileParseBundleHeader {
	out := getString(mobileprovisioningFilePath)
	return parse(out)
}

func getString(mobileProvisioningFilePath string) string {
	// fileの存在確認
	if _, err := os.Stat(mobileProvisioningFilePath); os.IsNotExist(err) {
		fmt.Println("no such file or directory: ", mobileProvisioningFilePath)
		os.Exit(1)
	}
	
	out, err := exec.Command("security", "cms", "-D", "-i", mobileProvisioningFilePath).Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return string(out)
}

func parse(mobileProvisioningFile string) smobileProvisioningFileParseBundleHeader {
	buf := bytes.NewReader([]byte(mobileProvisioningFile))
	var data smobileProvisioningFileParseBundleHeader
	decoder := plist.NewDecoder(buf)
	err := decoder.Decode(&data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	return data
}
