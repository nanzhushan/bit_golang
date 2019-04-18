package main

import (
	"fmt"
)

// 生产者生产消息
func produce(p chan<- int) {
	for i := 0; i < 10; i++ {
		p <- i
		fmt.Println("send:", i)
	}
}

// 消费者消费消息
func consumer(c <-chan int) {
	for i := 0; i < 10; i++ {
		v := <-c
		fmt.Println("receive:", v)
	}
}

func main() {
	stopCh := make(chan struct{})
	ch := make(chan int)
	go produce(ch)
	go consumer(ch)

	//var aa string ="ccc"

	stopCh <- struct{}{} // 空struct也是有内容

	<-stopCh // 等stopCh数据,没有数据下面的执行不了,  场景:用于阻塞协程
	fmt.Println("fdsaf")
	//time.Sleep(1 * time.Second)
}

//channel作用: 1 通讯 2 阻塞
