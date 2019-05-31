package main

import "fmt"

func main() {
	// f1()

	var f1 = func() string { // 无名函数变有名函数（定义）
		fmt.Print("fdsf")
		return "fdsaf" //
	}() // () 是执行
	fmt.Print(f1)

	// f1()

}
