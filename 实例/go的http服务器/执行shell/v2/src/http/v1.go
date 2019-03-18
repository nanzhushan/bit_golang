package main

import (
	"net/http"
	"fmt"
	"xx"
)

func hello(w http.ResponseWriter, r *http.Request) {
	res := r.URL.Query()   // map 字典类型
	name := res["name"]
	//fmt.Print(reflect.TypeOf(name))   // 是一个字符串的数组
	if name[0] == "crm-web"{
		fmt.Print("1")
	}else if name[0] == "crm" {
		fmt.Print("2")
	}else if name[0] == "sc-crm" {
		fmt.Print("3")
	}else {
		strData := xx.ExecCommand("hostname")
		//fmt.Print("args not match")
		fmt.Fprint(w,strData)
	}

	fmt.Fprint(w,"hello")     // 前端返回输出

}
func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe("0.0.0.0:8000", nil)
}