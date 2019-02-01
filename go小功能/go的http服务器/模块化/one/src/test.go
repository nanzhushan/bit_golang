package main

import (
	"os/exec"
	"bytes"
	"log"
	"fmt"
)

// 定义公共函数
func exec_shell(s string) {
	cmd := exec.Command("/bin/bash", "-c", s)
	var out bytes.Buffer

	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", out.String())
}

func main()  {
	exec_shell("hostname")

}