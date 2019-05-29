package main

import (
	"fmt"
	"time"
)

func main() {
	var ch chan int // 定义无缓冲通道
	var ticker = time.NewTicker(time.Second * 2)

	go func() {
		for range ticker.C {
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		}
		ch <- 1 // 往通道写入数据
	}()
	<-ch // 从通道读数据
}
