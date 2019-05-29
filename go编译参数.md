## 编译参数
[go build命令完全攻略-网址介绍](http://c.biancheng.net/view/120.html)

* 1)编译同目录的多个源码文件时,可以在go build 的后面提供多个文件名,
go build会编译这些源码,输出可执行文件go build+文件列表的格式如下:
```go
go build file1.go file2.go……

go build -h  // 查看编译参数帮助文档
```
* 2)如果源码中没有依赖 GOPATH 的包引用，那么这些源码可以使用无参数 go build

*3)“go build+包” 在设置GOPATH后,可以直接根据包名进行编译,即便包内文件被增（加）删（除）也不影响编译指令
原文件:
```go
package main
import (
    "chapter11/goinstall/mypkg"
    "fmt"
)
func main() {
    mypkg.CustomPkgFunc()
    fmt.Println("hello world")
}
```
编译参数:
```go
$ export GOPATH=/home/davy/golangbook/code
$ go build -o main chapter11/goinstall
$ ./goinstall
call CustomPkgFunc
hello world
```

