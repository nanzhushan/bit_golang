### 官方文档
[标准库](http://docscn.studygolang.com/pkg/)

[官网en](https://golang.org/)

###### 官网的正确解读
链接如右边: https://golang.google.cn/pkg/os/#Hostname

```
func Hostname
	func Hostname() (name string, err error)
		Hostname returns the host name reported by the kernel.
```
根据示例如下:
```
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	name, err := os.Hostname()  // Hostname函数返回一个字符串和error，所以要判断是否有error
	//fmt.Println(os.Hostname())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)
}

```


#### golang入门文档如下:
-----
[the-way-to-go_ZH_CN](https://github.com/nanzhushan/the-way-to-go_ZH_CN/blob/master/eBook/directory.md)

[go语言圣经](https://books.studygolang.com/gopl-zh/)

[web框架推荐gin](https://gin-gonic.com/zh-cn/docs/)


从php转过来的一定很喜欢gin
-----

##对比java/python/php  确实go还有有一些不足的地方
* 导入自定义的要放到gopath目录
* .....

### go使用场景
* 高并发处理，类似于c,支持协程，为并发而生
* 高性能的http服务器
* 代码加密
* 变异速度快
* 执行速度直逼c和c++,稍微超过java

### py为什么有全局锁
* 全局锁，同一时刻只准许一个线程在跑，这样的目的是为了防止多个线程访问同一资源排斥的问题，所有就有互斥锁的概念
* php 天生不支持线程，官方也不建议
* 全局锁是解释性语言的通病，所以解释性语言执行效率比编译型语言效率低

解释性语言有perl,python,php ,编译型语言有 java,golang

## go 引入自定义的包(只能放到gopath中)

    1) 要想引入自定义的包,需要自定义包的路径为GOPATH路径中
    2) golang和C或php不一样,不会自动查找当前路径下的文件,必须先在$GOAPTH里添加自己工程的路径
	3) 自定义包里面对外提供的API函数，首字母必须大写
	4) 通常情况下,import的包都是相对$GOPATH/src目录引入的,比如从github上面clone下来的项目,直接放到$GOPATH/src目录下
		
如何设置gopath变量了？
```
set GOPATH=D:\Go_res\v1     // windows设置方法
export PATH=$PATH:/usr/local/php/bin    //linux 设置方法   

```

## golang是面向什么的语言
golang是面向接口的语言，和传统面向对象的语言不同，面向对象的 封装多态继承在golang中都是通过接口去实现的

#### RPC调用和http调用区别在哪里？
RPC主要是基于TCP/IP协议,工作在传输层
而HTTP服务主要是基于HTTP协议,工作在应用层


### 等号和赋值操作什么时候用
* `:=`  给某变量的第一次赋值,初始化
* `=`   变量的非第一次赋值
* `==`  等于操作符

### 什么时候需要判断error？
返回值中基本都会包含err信息,就是我们需要从这个返回值取出数据，都要判断err!=nil

### go一定要缩进吗？
go语言是通过花括号进行区分代码块的,先不要缩进,但是最好缩进那样代码更加美观

### go中下划线的含义？
* 下划线意思是忽略这个变量
```
比如os.Open，返回值为*os.File，error
普通写法是f,err := os.Open(xxxxxxx)
如果此时不需要知道返回的错误值
就可以用f, _ := os.Open(xxxxxx)
如此则忽略了error变量
```
* 占位符
```
意思是那个位置本应赋给某个值，但是咱们不需要这个值，所以就把该值赋给下划线，意思是丢掉不要，
这样编译器可以更好的优化，任何类型的单个值都可以丢给下划线
```

