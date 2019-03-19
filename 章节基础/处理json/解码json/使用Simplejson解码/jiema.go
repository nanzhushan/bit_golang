package main

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"bytes"
)
func main() {
	str := `{"name":"tom","position":"beijing","data":[{"ip":1,"num":11},{"ip":2,"num":22}]}`
	buf :=bytes.NewBuffer([]byte(str))
	js,_:=simplejson.NewFromReader(buf)
	fmt.Println(js.Get("name").String())

	res,_:= simplejson.NewJson([]byte(str))
	fmt.Println(res.Get("name").String())

}