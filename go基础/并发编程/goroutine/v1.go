package main

import (
	"fmt"
	"time"
)

func Hello()  {
	fmt.Print("hello...\n")
}

func main()  {
	for i:=1;i<100;i++ {
		go Hello()       // go 并发执行

	}
	time.Sleep(time.Duration(2)*time.Second)   // 休眠2秒
}
