package main

import (
	"fmt"
	"strconv"
	"time"
)

var odrCh = make(chan string, 200)

func main()  {

	var i int
	for i = 1; i <= 100 ; i++ {
		order("odrjfiajfofjiaoj" + strconv.Itoa(i))    // strconv 用于整型转成字符串
	}

	go insert(odrCh)   // go关键字并发执行

	time.Sleep(10 * time.Second)

}

func order(odrId string) {     // 两种方式，逻辑代码，不传入channel
	odrCh<- odrId
	fmt.Println("下订单",odrId)
}

func insert(odrChan chan string) {
	for {
		select {
		case odrId := <-odrChan:
			fmt.Println("入库",odrId )
		}
	}
}