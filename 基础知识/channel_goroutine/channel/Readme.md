### 通道的类型(4种类型)
#### 1)无缓冲的通道
```
ch:=make(chan int)
或者
ch:=make(chan int,0)
```
无缓冲的通道指的是通道的大小为0，也就是说，这种类型的通道在接收前没有能力保存任何值，

它要求发送goroutine和接收goroutine同时准备好，才可以完成发送和接收操作

发送goroutine和接收gouroutine必须是同步的，同时准备后，如果没有同时准备好的话，先执行的操作就会阻塞等待，
直到另一个相对应的操作准备好为止,这种无缓冲的通道我们也称之为 同步通道
```
func main() {
	ch := make(chan int)

	go func() {
		var sum int = 0
		for i := 0; i < 10; i++ {
			sum += i
		}
		ch <- sum
	}()
	
	fmt.Println(<-ch)
}
```
解析: 只有当计算和的goroutine完成后,并且发送到 ch 通道的操作准备好后,同时 <-ch 就会接收计算好的值,然后打印出来。

或者这样写：
```go
package main

import (
	"fmt"
)

func main() {
	ch := make(chan int) // 没有能力保存任何值
	// go func() {
	// 	ch <- 1
	// }()
	ch <- 1 // 不加go关键字就不是和print进行同一个时刻调度的
	ch <- 2
	// time.Sleep(time.Second * 1)
	// msg := <-ch
	fmt.Print(<-ch)

}

```

#### 2)有缓冲通道
```go
ch:=make(chan int,3)

ch <- 2 //发送数值2给这个通道
x:=<-ch //从通道里读取值，并把读取的值赋值给x变量
<-ch //从通道里读取值
```
还可以使用内置的 close 函数关闭
```go
close(ch)
```
如果一个通道被关闭了,我们就不能往这个通道里发送数据了,如果发送的话,会引起painc异常
但是,我们还可以接收通道里的数据,如果通道里没有数据的话,接收的数据是 nil

#### 3) 管道
```
package main

import "fmt"

func main() {
	one := make(chan int)
	two := make(chan int)
	go func() {         //必须加go关键字和其他协程同一个时刻调度
		one <- 100 // 往one channel发送数据
	}()

	go func() {
		v := <-one // 从one channel接收数据
		two <- v   // 从one接收到的数据发送到two 通道
	}()
	fmt.Println(<-two) // 打印从two通道读取的数据
}

```
解析: 定义两个通道 one和two,然后按照顺序，先把100发送给通道 one,然后用另外一个goroutine从 one 接收值,
再发送给通道 two,最终在主goroutine里等着接收打印 two 通道里的值,这就类似于一个管道的操作
把通道 one 的输出,当成通道 two 的输入,类似于接力赛一样

#### 4) 单向通道
单向通道 是指通道只能发送(写)或者只能接收(读)
```go
var send chan<- int //只能发送
var receive <-chan int //只能接收
```
单向通道应用于函数或者方法的参数比较多,例如:
```go
func counter(out chan<- int) {

}
```
只能进行发送操作,防止误操作,使用了接收操作,如果使用了接收操作,在编译的时候就会报错的

