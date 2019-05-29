package main

import "fmt"

// 定义map字典
var Person = make(map[string]string)

// Person["name"] = "tom"   // 定义可以在函数外,赋值必须在函数内

func main() {
	Person["name"] = "tom"
	Person["sex"] = "man"
	Person["city"] = "cs"

	fmt.Print(Person)
}
