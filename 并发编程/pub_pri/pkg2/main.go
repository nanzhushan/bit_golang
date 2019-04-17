package main

import (
	"fmt"
	"pub_pri/pkg1"
)
func main()  {
	fmt.Println(pkg1.Public)

}

// 总结: 函数和变量大写是全局的，小写是局部的，局包的变量包可见。全局变量项目可见