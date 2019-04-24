package main

import (
	"fmt"
	"time"
)

func main()  {    // 主协程
	var i int

	for i = 1; i <= 5; i++ {
		go func() {                      // 加go关键字之后是5个子协程
			fmt.Println(i)
		}()
	}

	time.Sleep(3*time.Second)
}

/*
调度器分为系统调度器(内核态和用户态)和golang调度器,单核cpu是伪并发,我们现在考虑的是多核情况
上面的代码我们可以看出总共有6个协程, 加了go关键字之后 有公平的调度权利,可以先调度主协程 也可以先调度子协程
所以上面的代码会出现这个情况
*/

/*
总结:
  1. go关键字和主协程有平等的调度权利
  2. 当主协程结束之后，子协程也不存在了
  3. golang没有文件的概念只有包的概念,不通的包是不通的作用域
*/