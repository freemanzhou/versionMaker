# versionMaker
程序版本号自动生成工具

## 程序功能
+ 使用`make`命令自动编译Go语言程序。
+ 会生成两个程序。他们唯一的区别是，程序名称是否带版本号。
+ `BuildHistory.json`会记录同一个版本号所有的编译次数。
+ 通过`version`或`v`参数，查看程序的版本信息。
```bash
versionMaker git:(master)  ./versionMaker v
Version:  0.0.1.1
Time:     2016-12-12T15:02:44+0800
GitHash:  1b9ab3594a19fc8fbc6c0356e81f37a470f9d643
```
## 使用方法

1. 下载本源代码，在其目录，使用`go build`命令生成`versionMaker`程序。
2. 复制`versionMaker`和`Makefile`文件到目标程序目录。
3. 修改`Makefile`中第2行的`BINARY`变量设置成目标程序的名称。
5. 【选做】在目标程序目录，添加`Version`文件，添加自己的主版本号信息。不做的话，会自动生成`Version`文件，并设置主版本号为`0.0.0`。
6. 在目标程序中，添加下方的代码。
7. 使用`make`命令，自动编译Go语言程序。
```go
var (
	//Version 版本号
	Version string
	//BuildTime 编译时间
	BuildTime string
	//GitHash 当前的Git Hash码
	GitHash string
	//BuildNumber 编译次数
	BuildNumber string
)

func main() {
	if len(os.Args) > 1 && (os.Args[1] == "version" || os.Args[1] == "v") {
		fmt.Println("Version: ", Version+"."+BuildNumber)
		fmt.Println("Time:    ", BuildTime)
		fmt.Println("GitHash: ", GitHash)
       return 
	}
}
```