package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

var (
	//Version 版本号
	Version string
	//BuildTime 编译时间
	BuildTime string
	//GitHash 当前的Git Hash码
	GitHash string
	//BuildNumber 编译次数
	BuildNumber string

	buildNumberFile = ".BuildNumber"
	versionFile     = "BuildVersion"
	versionFileOld  = ".BuildVersion.Old"
	defultVersion   = "0.0.0"
)

func main() {
	if len(os.Args) > 1 && (os.Args[1] == "version" || os.Args[1] == "v") {
		fmt.Println("Version: ", Version+"."+BuildNumber)
		fmt.Println("Time:    ", BuildTime)
		fmt.Println("GitHash: ", GitHash)
		return
	}
	setVersion()
}

func readBuildNumber() int {
	raw, _ := ioutil.ReadFile(buildNumberFile)
	buildNumber, _ := strconv.Atoi(string(raw))
	//以上如果产生err的话，则buildNumber为0
	return buildNumber
}

func setVersion() {
	buildNumber := readBuildNumber()
	version, err := ioutil.ReadFile(versionFile)
	if err != nil {
		version = []byte(defultVersion)
		fmt.Printf("读取%s文件时出错，设置主版本号为“%s”\n", versionFile, defultVersion)
		ioutil.WriteFile(versionFile, []byte(defultVersion), 0777)
	}
	versionOld, err := ioutil.ReadFile(versionFileOld)
	if err != nil || versionOld == nil {
		versionOld = []byte("")
	}

	if string(version) != string(versionOld) {
		ioutil.WriteFile(versionFileOld, version, 0777)
		buildNumber = 0
	} else {
		buildNumber++
	}
	ioutil.WriteFile(buildNumberFile, []byte(strconv.Itoa(buildNumber)), 0777)
}
