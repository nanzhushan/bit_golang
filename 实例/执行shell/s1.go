package main

import (
	"os/exec"
	"log"
	"fmt"
)

func main()  {
	command := "hostname"
	cmd := exec.Command(command)
	bytes,err := cmd.Output()
	if err != nil {
		log.Println(err)
	}
	resp := string(bytes)
	//log.Println(resp)
	fmt.Printf(resp)
}

