package main

import (
	"github.com/dynport/gossh"
	"log"
)

func MakeLogger(prefix string) gossh.Writer {
	return func(args ...interface{}) {
		log.Println((append([]interface{}{prefix}, args...))...)
	}
}

func main() {
	client := gossh.New("192.168.0.101", "root")
	client.SetPassword("root")
	client.DebugWriter = MakeLogger("DEBUG")
	client.InfoWriter = MakeLogger("INFO ")
	client.ErrorWriter = MakeLogger("ERROR")

	defer client.Close()
	rsp, e := client.Execute("hostname")   // 远程执行shell命令
	if e != nil {
		client.ErrorWriter(e.Error())
	}else {
		client.InfoWriter(rsp.String())
	}
}