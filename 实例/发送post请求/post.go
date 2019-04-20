package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func httpPost() {
	postData := strings.NewReader(`{"name":"knight"}`)
	res, err := http.Post("https://httpbin.org/post", "application/json", postData)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)

	}
	fmt.Printf("%s", body)
}

func main() {
	httpPost()
}
