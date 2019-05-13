package main

import "fmt"

func main()  {    // 实际是主协程
	var i = 3
	go func(a int) {   // go + 函数就是一个协程--花括号抱起来的代码,可以并发执行
		fmt.Println(a)
		fmt.Println("1")

	}(i)
	fmt.Println("2")
}

// 程序会优先执行主协程，当主线程执行完成之后 程序会立刻退出，没有多余的时间去执行子协程
// 如果在程序的最后让主线程休眠1秒钟，程序就有时间取执行子协程
// golang 调度器 ->内核调度器