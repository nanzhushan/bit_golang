package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	FN = "../app.conf"
)

//str配置文件中的key
func Get(str string) string {
	f, err := os.Open(FN)
	defer f.Close()
	if err != nil {
		os.Exit(-1)
	}
	buf := bufio.NewReader(f)
	for {
		line, _ := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		strs := strings.Split(line, "=")
		if strings.TrimSpace(strs[0]) == str {
			return strings.TrimSpace(strs[1])
		}
	}
	return ""
}

func main() {
	fmt.Println(Get("city"))
}
