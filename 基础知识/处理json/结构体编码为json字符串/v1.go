package main

import(
	"fmt"
	"encoding/json"
)
type User struct{
	Name string
	Age int
}

func main() {
	user := User{
		Name:"tom",
		Age:3,
	}
	b, _ := json.Marshal(user)   //结构体编码成json字符串
	//if err != nil {
	//	fmt.Println("json Marshal fail:", err)
	//}
	fmt.Println(string(b))
}