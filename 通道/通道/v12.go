package main

import (
	"fmt"
	"time"
)


func main()  {
	ch:=make(chan  int,2)  // 创建有缓冲的channel，类型为int
	ch<- 11
	ch<- 33
	//ch<- 33
	go func() {
		for {
			select {
			// select，下只能有case 代码块只能有通道
			 case v := <- ch:
			 	fmt.Println(v)
			}
		}
		//v:= <-ch // 读数据
		//v=<- ch   // 已经声明过了就直接等号即可
		//fmt.Println(v)
	}()
	fmt.Println("22")
	time.Sleep(2*time.Second)
}


/*
	总结:
		1. 通道只有一个向左的箭头,写入和读取都是一个箭头
		2. 读取的过程就是赋值的过程

*/



