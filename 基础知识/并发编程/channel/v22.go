package main

import (
	"time"
	"fmt"
)
func slow(c chan string)  {    // 将通道作为参数传入，并定义了通道的内容
	time.Sleep(time.Second * 1)
	c <- "hello"
	c <- "world"
}
func main()  {
	c := make(chan string)   //  创建一个字符串通道 并赋值给变量c
	go slow(c)    // 协程执行slow函数
	msg:= <- c   // 声明变量msg用于接收通道的数据
	fmt.Println(msg)
}