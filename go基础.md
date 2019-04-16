
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


## 变量

```
var aa int32 = 10   // 第一种方式定义
var aa = "hello"   // 第二种方式定义，根据值自定判断类型
fmt.Print("type:",reflect.TypeOf(aa))   // 判断类型, java 中是instanceof

var aa int32 = 10
_=aa                    // 定义变量而不使用，需要定义 空标识符

name:="haha"  // 第三种方式，这种智能在函数体内使用，不能定义为全局变量

var (                   // 多变量定义
    name1 type1
    name2 type2
)
```

## 常量
```
显式类型定义： const b string = "abc"
隐式类型定义： const b = "abc"
```

## 数组
```
// 一维数组
var team [3]string
team[0] = "hammer"
team[1]="fdsaf"
print(team[0])
```
### 结构体
定义结构体 ，数据类型可以不同，但是数组必须数据类型相同
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
	package main
	import "fmt"
	func main()  {
		var a[2] string		// 像java一样要定义数组长度
		a[0]="hello"
		a[1]="dd"
		fmt.Print(a)

		// slice 是个有序的值,可以不定义长度.append 可以添加元素，类似py的list
		p:=[]int{2,3,5,7}
		fmt.Print(p)
		fmt.Print(len(p))
	}
	
定义数组的三种方式
```
func main()  {
	// 定义有序数组
	var aa = [2]string{"fdaf","fdaf"}   // 数组初始化,定义好长度
	var bb =[...]string{"cc","rr"}    // 定义为可变长度
	var cc = [3]string{0:"tom",1:"18"}    //  第三种方式定义
	print(aa[0])
	print(len(bb))
	print("----",cc[0])

}
```
## map键值对(和java的hash map以及map类似)
定义: map[KeyType]ValueType
```
m := make(map[string]string)   // 定义键和值的类型
m["name"] =  "knight"
fmt.Print(m["name"])
```


## 编译安装
* go install 是建立在 GOPATH 上的，无法在独立的目录里使用 go install。
* GOPATH 下的 bin 目录放置的是使用 go install 生成的可执行文件，可执行文件的名称来自于编译时的包名。
* go install 输出目录始终为 GOPATH 下的 bin 目录，无法使用-o附加参数进行自定义。
* GOPATH 下的 pkg 目录放置的是编译期间的中间文件。

## 循环语句
```
func main() {
    for i := 0; i < 5; i++ {
        fmt.Printf("This is the %d iteration\n", i)
    }
}
```
## 函数
1) 传参(如果定义了返回类型需要使用return，如果没有定义返回类型就不需要)
```
func main()  {
	Greting("uu","77")
}

func Greting(x string,y string) string{   // 定义传入的参数类型以及返回类型
	print("庆祝 ",x,y)
	return x+y
}
```
或者
```
func main()  {
	Greting("uu","77")
}

func Greting(x string,y string){   // 定义传入的参数类型以及返回类型
	print("庆祝 ",x,y)
}
```
2)传入可变参数
```
func main()  {
	Greting("uu","77")
}
func Greting(x ...string) string{   // 传入可变参数
	print(x[0],x[1])         // 默认是数组进行传递
	//print(x)
	return "haha"
}
```
3)defer(推迟)关键字
```

func main()  {
	Greting("uu","77")
}
func Greting(x ...string) string{   // 传入可变参数
	print(x[0],x[1],len(x))    // len 返回长度
	defer lai()     // defer 用于return之后再执行一些语句类似java异常处理中的finally,必须写在retrun之前
	return "haha"

}
func lai()  {
	print("你们先来....")
}
```
## 计算时间
```
func main()  {
	start := time.Now()
	time.Sleep(3 * time.Second)   // sleep  3秒
	end:= time.Now()
	time_all:=end.Sub(start)
	Greting("uu","77")
	print("总耗时",time_all.Seconds())
}
```

## 转换成json
```
package main

import (
	"encoding/json"
	"fmt"
)

// 定义address结构体
type Address struct {
	city string
	age string
}

// 定义结构体 ，数据类型可以不同，但是数组必须数据类型相同
type VCard  struct {
	FirstName string
	LastName  string
	Remark    string
	Address  []*Address
}

func main()  {
	//s1:=new(Address)		// 指针初始化 定义指向结构体的指针,可以通过s1.age="" 进行赋值
	s2 := &Address{"cs","18"}	// 传统方法初始化
	s3 := &Address{"sh","19"}	//传统方法初始化

	vc := VCard{"Jan", "Kersschot",  "haha",[]*Address{s2,s3}}
	js, _ := json.Marshal(vc)	// 转json
	fmt.Printf("JSON format: %s", js)
}
```
map 转json
```
package main
import (
	"encoding/json"
	"fmt"
)

// map转json
func main()  {
	s := []map[string]interface{}{}
	m1 := map[string]interface{}{"name": "John", "age": 10}
	m2 := map[string]interface{}{"name": "tom", "age": 19}
	s = append(s, m1,m2)   // 定义指针指向map
	fmt.Print(m2)
	// 转json
	b,err := json.Marshal(s)
	if err != nil{              // 必须要判断是不是能正常转换，不然程序无法编译
		print("fail",err)
	}
	fmt.Print(string(b))
}
```
## 异常处理
```
b,err := json.Marshal(s)
if err != nil{              // 类似java和py中的try...expect..., 其中 defer关键字就是异常处理中的finally
	print("fail",err)
}
``` 
### 表示字符串
* 双引号: 可以表示字符串
* 反引号: 表示原生字符串，比如引号括号等

