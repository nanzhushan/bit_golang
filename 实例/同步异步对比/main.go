package main

import (
 "fmt"
 "sync"
 "time"
)

func main () {
 var i int

 go func() {           // 启一个协程进行同步操作
  fmt.Println("同步开始",time.Now().Unix())
  for i = 1; i<= 10 ; i++ {    // for循环是同步的
   time.Sleep(1 * time.Second)
  }
  fmt.Println("同步结束",time.Now().Unix())
 }()

 go func() {            // 启动一个协程进行异步操作实验
  fmt.Println("异步开始",time.Now().Unix())
  wg := &sync.WaitGroup{}   // 计数器
  for i = 1; i<= 10 ; i++ {
   go func () {
    wg.Add(1)
    time.Sleep(1 * time.Second)
    wg.Done()   // 计算器减1
   } ()
  }
  wg.Wait()
  fmt.Println("异步结束",time.Now().Unix())
 }()

 time.Sleep(15*time.Second)
}
