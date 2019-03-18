package main

import (
	"xx"
	"fmt"
)

func main()  {
	strData := xx.ExecCommand("hostname")
	fmt.Println("Execute finished:" + strData)

}