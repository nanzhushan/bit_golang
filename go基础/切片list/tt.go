package main

import "fmt"

func main()  {
	m1 := make(map[string]string)
	m1["a"]="aa"
	m1["b"]="bb"
	m1["c"]="cc"
	fmt.Print(m1)

	for k,v := range(m1){
		fmt.Print(k,v)
	}
}
