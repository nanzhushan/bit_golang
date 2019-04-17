package main

import "fmt"

func main()  {
	go fmt.Println("1")  // 使用go关键字加一个函数就可以创建一个携程
	fmt.Println("2")

}

