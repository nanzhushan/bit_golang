package main

import (
	"flag"
	"fmt"
	"os"
)

// 实际中应该用更好的变量名
var (
	h bool
	v, V bool
	t, T bool
)

func init() {
	flag.BoolVar(&h, "h", false, "this help")
	flag.BoolVar(&v, "v", false, "show version and exit")
	flag.BoolVar(&V, "V", false, "show version and configure options then exit")

	flag.BoolVar(&t, "t", false, "test configuration and exit")
	flag.BoolVar(&T, "T", false, "test configuration, dump it and exit")

	// 改变默认的 Usage，flag包中的Usage 其实是一个函数类型。这里是覆盖默认函数实现，具体见后面Usage部分的分析
	flag.Usage = usage
}

func main() {
	flag.Parse()

	if h {
		flag.Usage()
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, `nginx version: nginx/1.10.0 
Usage: nginx -h/v/t`)
	fmt.Print("\n")

	flag.PrintDefaults()
}