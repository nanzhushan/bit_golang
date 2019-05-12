package main

import (
	"time"
	"fmt"
)

func slowFunc()  {
	time.Sleep(time.Second * 2)
	fmt.Println("hello")
}

func main()  {
	go slowFunc()
	fmt.Println("world...")
	time.Sleep(time.Second * 3)
}