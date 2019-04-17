package main

import (
	"fmt"
	"strconv"
	"time"
)

var odrCh = make(chan string, 200)   // 定义通道长度为200

func main()  {

	var i int
	for i = 1; i <= 100 ; i++ {
		order("odrjfiajfofjiaoj" + strconv.Itoa(i))   // 模拟将整形转成字符串
	}

	go insert(odrCh)      // 并发入库

	time.Sleep(10 * time.Second)

}

func order(odrId string) {      // 逻辑处理进行下单操作
	odrCh<- odrId
	fmt.Println("下订单",odrId)
}

func insert(odrChan chan string) {    // 处理逻辑入库
	for {
		select {
		case odrId := <-odrChan:    // 从通道读取数据
			fmt.Println("入库",odrId )
		}
	}
}