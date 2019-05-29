package main

import (
	"fmt"
)

type person struct {
	name string
	sex  string
}

var p = person{"annie", "man"} // 函数外赋值必须这样

func main() {
	fmt.Printf(p.name)
}
