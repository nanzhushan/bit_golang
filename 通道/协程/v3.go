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
	fmt.Println("2")    // 这是主协程
	time.Sleep(1* time.Second)
}

