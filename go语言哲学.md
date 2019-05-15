## go语言哲学
哲学思想： 少即是多，追求简洁优雅，就是说如果异常价值不大，就不会将异常加入到语言特性中。


## 和php,python,java对比,go和哪个语言最像
除开结构体之外和java最像

## go的异常捕获
Go语言追求简洁优雅，所以，Go语言不支持传统的 try…catch…finally 这种异常，因为Go语言的设计者们认为，将异常与控制结构混在一起会很容易使得代码变得混乱。因为开发者很容易滥用异常，甚至一个小小的错误都抛出一个异常。遇到真正的异常的情况下（比如除数为 0了）。才使用Go中引入的Exception处理：defer,panic,recover

```go
package main
import "fmt"

func main(){
    defer func(){           // 必须要先声明defer，否则不能捕获到panic异常
        fmt.Println("c")
        if err:=recover();err!=nil{
            fmt.Println(err)                // 这里的err其实就是panic传入的内容，55
        }
        fmt.Println("d")
    }()
    f()

}
 
func f(){
    fmt.Println("a")
    panic(55)
    fmt.Println("b")
    fmt.Println("f")

}
```

## go的错误处理
Golang中引入两个内置函数panic和recover来触发和终止异常处理流程，同时引入关键字defer来延迟执行defer后面的函数。

### 为什么go没有多进程和线程的概念？
因为go把操作线程和进程的函数进行了封闭,进行了隐藏,使开发者不能进行操作，所以golang操作不了线程和进程.

golang有goroutine调度器,go调度器能分配cpu和内存资源,工作在用户态


隐藏的原因是因为go有自己的调度器(go的调度器是很奢侈的事情，因为写调度器很难)，大部分语言都没有调度器的概念，比如java,python就没有

#### go的协程实现并发和java的多线程实现并发的异同？
1)go协程实现的并发是通过通讯(放入队列也就是channel实现的共享内存

2)java的多线程和线程池是通过 共享内存实现通讯(线程通讯)




