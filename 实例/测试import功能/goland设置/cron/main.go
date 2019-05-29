package main

import (
	"fmt"

	cr "github.com/robfig/cron"   // 重命名包名,cron是包名，文件夹的名字是包名
	"cron/mylib"
)

// 定时执行get请求
func main() {
	var c = cr.New()                // 里面有定义队列
	c.AddFunc("*/5 * * * *", func() { // 每隔5s
		fmt.Print("ccc")
		mylib.Get()  // 要加包名，包名+函数名

	})
	c.Start()
	select {}
}
