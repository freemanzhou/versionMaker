package main

import (
	"encoding/json"
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

	buildHistory        = "BuildHistory.json"
	buildNumberFileName = "BuildNumber"
	buildVersion        = "Version"
	defaultVersion      = "0.0.0"
)

func main() {
	if len(os.Args) > 1 && (os.Args[1] == "version" || os.Args[1] == "v") {
		fmt.Println("Version: ", Version+"."+BuildNumber)
		fmt.Println("Time:    ", BuildTime)
		fmt.Println("GitHsh:  ", GitHash)
		return
	}
	makeBuildNumberFile()
}

func makeBuildNumberFile() {
	version := readVersion(buildVersion, defaultVersion)
	buildNumberMap := readBuildNumberMap(buildHistory)

	//先保存编译次数文件，再增加编译次数
	//所以，json文件保存的是下一次编译的次数
	saveBuildNumberFile(buildNumberMap[version], buildNumberFileName)
	buildNumberMap[version]++

	saveBuildNumberMap(buildNumberMap, buildHistory)
}

//获取主版本号的信息
//如果没有保存主版本号信息的文件，就自动生成一个
func readVersion(filename, defaultVersion string) string {
	version, err := ioutil.ReadFile(filename)
	if err != nil {
		version = []byte(defaultVersion)
		fmt.Printf("%s不存在，或者读取%s文件时出错，设置主版本号为“%s”。\n", filename, filename, defaultVersion)
		ioutil.WriteFile(filename, version, 0777)
	}
	return string(version)
}

//每个主版本号的编译次数，都保存在`BuildHistory.json`当中。
func readBuildNumberMap(filename string) map[string]int {
	buildNumberMap := map[string]int{}

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("%s不存在，或者读取%s失败，正在将其置零。\n", filename, filename)
		return buildNumberMap
	}

	if err := json.Unmarshal(bytes, &buildNumberMap); err != nil {
		fmt.Println("转换Json文件失败，正在将其置零。")
	}

	return buildNumberMap
}

//在相应的主版本号的编译次数++后，需要再把编译记录保存到json文件
func saveBuildNumberMap(bmap map[string]int, filename string) {
	bytes, err := json.Marshal(bmap)
	if err != nil {
		fmt.Printf("转换Json失败，不保存%s文件\n", filename)
		return
	}

	ioutil.WriteFile(filename, bytes, 0777)
}

//把当前编译次数保存到文件中，以便makefile读取。
func saveBuildNumberFile(number int, filename string) {
	if err := ioutil.WriteFile(filename, []byte(strconv.Itoa(number)), 0777); err != nil {
		fmt.Println("无法保存BuildNumber.")
	}
}
