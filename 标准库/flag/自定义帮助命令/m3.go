package main

import (
	"flag"
	"fmt"
)


// flag.String 是设置传递参数的值类型
func main()  {
	flag.Usage = func() {
		var useageText = `nginx version: nginx/1.10.0 
Usage: nginx -h/v/t`
		fmt.Print(useageText)
	}
	var m = flag.String("m"," aaa","hahaha")   // 看源码是个指针,value是默认值
	var host = flag.String("host"," bbb","hahaha")   // 看源码是个指针,value是默认值
	flag.Parse()
	fmt.Print("cmd is:",* m)
	fmt.Print("host is:",* host)
}

