package main

import (
	"net/http"
	"log"
)

func sayBype(w http.ResponseWriter,r *http.Request)  {
	w.Write([]byte("hello wolrd"))
}
func main()  {
	mux := http.NewServeMux()
	mux.Handle("/",&aa{})
	mux.HandleFunc("/bye",sayBype)
	log.Println(".....")
	log.Fatal(http.ListenAndServe(":4000",mux))
}


type aa struct {}

func (*aa) ServeHTTP(w http.ResponseWriter,r *http.Request)  {
	w.Write([]byte("aaaa"))
}



