package main

import (
	"os"
	"os/exec"
	"os/user"
	"fmt"
	"io/ioutil"
	"regexp"
)

const installPath = "/Library/MobileDevice/Provisioning Profiles"

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

func installList(mobileProvisioningName string) {
	files, _ := ioutil.ReadDir(getInstallPath())
	for _, f := range files {
		if _, err := os.Stat(getInstallPath() + f.Name()); os.IsNotExist(err) {
			// fileの存在確認でなかったらスキップする
			continue
		}

		if isProvisioningFile(f.Name()) {
			plist := getPlistData(getInstallPath() + f.Name())
			if plist.Name == mobileProvisioningName {
				fmt.Println(f.Name())
			}
		}
	}
}

func getUserHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	return usr.HomeDir
}

func getCpFileName(plist smobileProvisioningFileParseBundleHeader) string {
	return getInstallPath() + plist.UUID + ".mobileprovision"
}

func getInstallPath() string{
	return getUserHomeDir() + installPath + "/"
}

func isProvisioningFile(fileName string) (b bool) {
	if m, _ := regexp.MatchString(".mobileprovision$", fileName); !m {
		return false
	}
	return true
}
