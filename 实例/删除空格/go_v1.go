package main

import (
	"fmt"
	"strings"
)

var b = "    fdsf x yz dff  "

func main() {
	b = strings.Replace(b, " ", "", -1)
	fmt.Println(b)

}
