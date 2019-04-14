package main

import (
	"encoding/json"
	"log"
	"fmt"
)

func main()  {
	person:=make([]string,2)
	person[0]="aa"
	person[1]="bb"
	jsonStr, err := json.Marshal(person)
	if err!= nil{
		log.Fatal()
	}
	fmt.Printf("%s",jsonStr)
}
