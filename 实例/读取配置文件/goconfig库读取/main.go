package main

import (
	"fmt"
	"log"

	"github.com/Unknwon/goconfig"
)

func main() {
	cfg, err := goconfig.LoadConfigFile("app.conf")
	if err != nil {
		log.Fatal(err)
		return
	}
	// str, _ := cfg.GetValue(goconfig.DEFAULT_SECTION, "timeout") // 取键值
	// str, _ := cfg.GetValue("", "mysqld:timeout") // 可以省略
	str, _ := cfg.GetSection("mysqld") // 取分区所有数据
	fmt.Println(str)
	fmt.Println(str["timeout"])

}
