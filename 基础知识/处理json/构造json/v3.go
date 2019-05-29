package main

import (
	"fmt"
)

type person struct {
	name string
	sex  string
}

var p person

func main() {
	p.name = "tom"
	p.sex = "man"
	fmt.Printf(p.name)
}
