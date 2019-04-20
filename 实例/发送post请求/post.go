package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func httpPost() {
	postData := strings.NewReader(`{"name":"knight"}`)   // 使用strings库将其转换成io.Reader流 为传输做好准备
	res, err := http.Post("https://httpbin.org/post", "application/json", postData)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()   // 客户端读取完所有数据之后 把连接进行关闭
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)

	}
	fmt.Printf("%s", body)
}

func main() {
	httpPost()
}
