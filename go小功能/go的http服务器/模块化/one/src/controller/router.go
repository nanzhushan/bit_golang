package controller

import (
	"net/http"
	"view"
)

func init()  {
	//
}

func Router()  {
	mux := http.NewServeMux()
	mux.HandleFunc("/", view.IndexHandler)
	mux.HandleFunc("/path",view.PathHandler)

	http.ListenAndServe("0.0.0.0:12345", mux)
}

