package main

import (
	"fmt"
	"flag"
)

func init()  {

}


func main() {
	namePtr := flag.String("name", "username", "姓名")
	agePtr := flag.Int("age", 18, "年龄")

	var email string
	flag.StringVar(&email, "email", "qq@qq.com", "邮箱")

	flag.Parse()

	//args := flag.Args()
	fmt.Println("name:", *namePtr)
	fmt.Println("age:", *agePtr)
	fmt.Println("email:", email)
	//fmt.Println("args:", args)
}