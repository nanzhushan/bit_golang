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
	str, _ := cfg.GetValue(goconfig.DEFAULT_SECTION, "city")
	fmt.Println(str)
}
