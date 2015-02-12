package main

import (
	"os"
	"os/exec"
	"os/user"
	"fmt"
)


func install(mobileProvisioningFilePath string) {
	plist := getPlistData(mobileProvisioningFilePath)
	cpFileName := getCpFileName(plist)

	err := exec.Command("cp", mobileProvisioningFilePath, cpFileName).Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Installed to ", cpFileName)
}

func getUserHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	return usr.HomeDir
}

const installPath = "/Library/MobileDevice/Provisioning Profiles"
func getCpFileName(plist smobileProvisioningFileParseBundleHeader) string {
	return getUserHomeDir() + installPath + "/" + plist.UUID + ".mobileprovision"
}
