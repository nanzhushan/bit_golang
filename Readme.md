## beego
[beego官方文档](https://beego.me/docs/intro/)

## go 和java都是强类型语言，他们都很像
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

1)`:=` 结构替代var进行变量赋值不能使用在函数外。 

## 数组
```
// 一维数组
var team [3]string
team[0] = "hammer"
team[1]="fdsaf"
print(team[0])
```
### 结构体
结构体就是一个字段的集合
```
type aa struct {
	x int
	y int
}

func main()  {
	v := aa{1,4}
	v.x = 4    
	println(v.x)  // 通过点号访问
	
}
```
## 指针
```
 Go 具有指针。 指针保存了变量的内存地址。
类型 *T 是指向类型 T 的值的指针。其零值是 `nil`。

var p *int

& 符号会生成一个指向其作用对象的指针。

i := 42
p = &i

* 符号表示指针指向的底层的值。
fmt.Println(*p) // 通过指针 p 读取 i
*p = 21         // 通过指针 p 设置 i

这也就是通常所说的“间接引用”或“非直接引用”。
```
## 数组和slice
```
package main
import "fmt"
func main()  {
	var a[2] string   // 像java一样要定义数组长度
	a[0]="hello"
	a[1]="dd"
	fmt.Print(a)

	// slice 是个有序的值,可以不定义长度.append 可以添加元素，类似py的list
	p:=[]int{2,3,5,7}
	fmt.Print(p)
	fmt.Print(len(p))
}
```
## map键值对
定义: map[KeyType]ValueType
```
m := make(map[string]string)
m["name"] =  "knight"
fmt.Print(m["name"])
```
## 编译安装
* go install 是建立在 GOPATH 上的，无法在独立的目录里使用 go install。
* GOPATH 下的 bin 目录放置的是使用 go install 生成的可执行文件，可执行文件的名称来自于编译时的包名。
* go install 输出目录始终为 GOPATH 下的 bin 目录，无法使用-o附加参数进行自定义。
* GOPATH 下的 pkg 目录放置的是编译期间的中间文件。
