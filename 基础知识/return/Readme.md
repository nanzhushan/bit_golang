## return
return用于在函数内部，退出函数执行过程,使用场景如下:

* 退出函数执行过程，不指定返回值
通常有两种情况不需要指定返回值退出函数执行过程
	- 第一是：函数没有返回值
	- 第二是：函数返回值有变量名，不需要显示的指定返回值
```
package main

import (
	"fmt"
)

func say(flag bool) {
	if !flag {
		fmt.Println("false")
		return
	}
	fmt.Println("没有返回值")
}

func getStatus() (num int) {
	// num是在返回值中定义的变量
	num = 100
	return     // 因为函数中已经定义了返回类型
}

func main() {
	say(true)
	say(false)
	fmt.Println(getStatus())
}

```
上边的示例代码中，say是一个没有返回值的函数，这个函数中的return就是一个不带返回值的退出。
getStatus是带一个返回值的函数，由于在返回值中定义了变量,所以,在函数退出时,可以不用显示的在return后边指定函数返回值，函数调用结束后，自动将之前定义的返回值变量作为这个函数的返回结果

* 退出函数执行过程，并指定返回值
	- 退出函数执行过程，并指定返回值(eg1)
	- 在返回值中定义变量，而return其他结果(eg2)

eg1:
```
package main

import (
	"fmt"
)

// getMsg函数需要返回一个string类型值
func getMsg() string {
	return "hello"  // 当函数有返回值时,如果返回值没有定义变量,那么一定要使用return加上返回值退出函数
}

func main() {
	fmt.Println(getMsg())
}


```

eg2:
```
package main

import (
	"fmt"
)

func getMsg() (msg string) {
	msg = "world"
	return "hello"  // 也可以指定另外的变量作为返回结果
}

func main() {
	fmt.Println(getMsg())
}

```	