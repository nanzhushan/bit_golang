package main

import(
	"fmt"
	"encoding/json"
)

type User struct{
	Name string
	Age int
	Role []string   // 字符串数组
	Skill map[string]string  // 定义key-value都为字符串的字典
}

//var play = make(map[string]string)   // 声明空字典
//play["python"] = "dic"
//play["php"] ="array"


func main() {
	var m1 = make(map[string]string)
	m1["python"] = "aa"
	m1["elixir"] = "bb"
	m1["ruby"] = "cc"

	user := User{
		Name:"tom",
		Age:3,
		Role:[]string{"admin","guest"},
		Skill:m1,
	}
	b, err := json.Marshal(user)   //结构体编码成json字符串
	if err != nil {
		fmt.Println("json Marshal fail....", err)
	}
	fmt.Println(string(b))
}