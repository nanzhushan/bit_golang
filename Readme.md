## beego
[beego官方文档](https://beego.me/docs/intro/)

## go
[go英文官方文档](https://golang.org/doc/)

[go中文文档](http://docscn.studygolang.com/doc/)

[Golang标准库文档](https://studygolang.com/pkgdoc)

## 打包
```
uname -r
env GOOS=linux GOARCH=386 go build main.go
```
根据系统类型然后进行打包

##基础语法
### 
* go 是强类型语言(比如docker一样，资源使用能限制那么死)

### 函数
* Go语言最少有个 main() 函数
* 函数内定义的变量称为局部变量
* 函数外定义的变量称为全局变量
* 函数定义中的变量称为形式参数

## 数组
```
// 一维数组
var team [3]string
team[0] = "hammer"
team[1]="fdsaf"
print(team[0])
```

