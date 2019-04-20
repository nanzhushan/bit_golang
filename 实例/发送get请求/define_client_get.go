package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	cs := &http.Client{} // 不使用nt/http的快捷方法get,而创建一个http客户端
	// postData := strings.NewReader(`{"name":"knight"}`)
	// resquest, err := http.NewRequest("Post", "https://httpbin.org/post", postData)
	resquest, err := http.NewRequest("GET", "https://ifconfig.co", nil) // 发出get请求
	if err != nil {
		log.Fatal(err)
	}
	respose, err := cs.Do(resquest) // 使用do方法发送请求并处理响应
	defer respose.Body.Close()
	body, err := ioutil.ReadAll(respose.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)

}
