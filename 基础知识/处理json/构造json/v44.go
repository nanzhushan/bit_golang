package main

// 必须在函数体内赋值
import "fmt"

var b string

func test() {
	b = "aa"
	fmt.Print(b)

}

func main() {
	test()
}

