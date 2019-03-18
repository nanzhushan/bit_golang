package main
 
import (
    "fmt"
    "net/http"
)
 
func IndexHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "hello world")
}
 
func main() {
    http.HandleFunc("/", IndexHandler)
    err := http.ListenAndServe("0.0.0.0:8000", nil)
    if err != nil{
           fmt.Println(err)
        }
}
