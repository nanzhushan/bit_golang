package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/robfig/cron"
)

func Get() {
	cs := &http.Client{}
	resquest, err := http.NewRequest("GET", "https://ifconfig.co", nil) // 发出get请求
	if err != nil {
		log.Fatal(err)
	}
	respose, err := cs.Do(resquest) // 使用do方法发送请求并处理响应
	body, err := ioutil.ReadAll(respose.Body)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%s\n", body)
	fmt.Printf("%s", body)
}

// 定时执行get请求
func main() {
	var c = cron.New()                // 里面有定义队列
	c.AddFunc("*/5 * * * *", func() { // 每隔5s
		Get()
	})
	c.Start()
	select {}
}
