package view

import (
	"net/http"
	"io"
)

func  init()  {
	//
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, world.....\n")
}
