package main

import "fmt"

// 通道的作用就是在多线程之间传递数据


func main()  {
	ch:=make(chan  int)  // 创建一个无长度/无缓冲的channel，也就是同步的
	ch<- 1 //ordeer_id
	go func() {
		<-ch
		fmt.Println("11")
	}()
	fmt.Println("22")
}

/*
上面的程序会报错"fatal error: all goroutines are asleep - deadlock!"
因为我们的channel是无缓冲的，即同步的,赋值完成后来不及读取channel，程序就已经阻塞了。
这里介绍一个非常重要的概念：channel的机制是先进先出，如果你给channel赋值了，那么必须要读取它的值，不然就会造成阻塞
*/

// 通过通讯实现共享内存，通讯的过程就是进行数据共享