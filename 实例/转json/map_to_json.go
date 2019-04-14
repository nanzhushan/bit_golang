package main

import (
	"encoding/json"
	"log"
	"fmt"
)

// map转json
func main() {
	user := make(map[string]string)
	user["name"]="knight"
	user["age"]="18"
	jsonStr, err := json.Marshal(user)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Printf("%s",jsonStr)
	}
