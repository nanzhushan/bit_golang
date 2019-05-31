package main

import "fmt"

// 函数外定义
var f1 = func() string { // 无名函数变有名函数（定义）
	fmt.Print("fdsf")
	return "fdsaf" //
}() // () 是执行

func main() {
	fmt.Print(f1)

}
