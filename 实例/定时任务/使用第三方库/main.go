package main

import (
	"fmt"

	"github.com/robfig/cron"
)

func text() { // define task
	fmt.Print("hello")
	fmt.Print("\n")
}

func main() {
	var c = cron.New()                // 里面有定义channel
	c.AddFunc("*/5 * * * *", func() { // 每隔5s
		text() // exec task
	})
	c.Start()

	select {}   // 之前有定义队列 所以可以使用select

}
