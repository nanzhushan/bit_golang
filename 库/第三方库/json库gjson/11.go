package main

import (
	"github.com/tidwall/gjson"
	"fmt"
)

const jsonstr = `{"name":{"first":"xiao","last":"zhou"},"age":30}`

func main() {
	value := gjson.Get(jsonstr, "name.last")
	aa:= gjson.Get(jsonstr,"age")
	fmt.Print(aa)

	println(value.String())
}