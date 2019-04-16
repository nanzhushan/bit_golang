package main

import (
	"fmt"
	"encoding/json"
	"log"
)

type Student struct {
	Name string
	City string
}

func main()  {
	str := `{"name":"tom","city":"cs"}`    // 必须和上面的结构体一一对应
	beta := []byte(str)
	p := Student{}
	err := json.Unmarshal(beta,&p)
	if err!= nil{
		log.Fatal(err)
	}
	fmt.Print(p)
	//fmt.Println(reflect.TypeOf(p))

}