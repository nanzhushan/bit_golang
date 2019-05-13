package main

import (
	"fmt"
	"time"
)

func main()  {
	var i = 3
	go func(a int) {
		fmt.Println(a)
		fmt.Println("1")

	}(i)
	fmt.Println("2")    
	time.Sleep(1* time.Second)    // 主协程里面一般进行sleep 而不马上退出 留时间给子协程运行
}



// 程序会优先执行主协程，当主线程执行完成之后 程序会立刻退出，没有多余的时间去执行子协程
// 如果在程序的最后让主线程休眠1秒钟，程序就有时间取执行子协程
