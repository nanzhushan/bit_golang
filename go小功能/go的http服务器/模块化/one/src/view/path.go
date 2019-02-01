package view

import (
	"net/http"
	"fmt"
)

func  init()  {
	//
}

func PathHandler(w http.ResponseWriter, r *http.Request) {
	//io.WriteString(w, "this is path html")
	fmt.Print("this is path html")
}
