package main


//  curl -is http://127.0.0.1:8000?name=knight
import (
	"net/http"
	"fmt"
)

func hello(w http.ResponseWriter, r *http.Request) {
	for k := range  r.URL.Query(){
		fmt.Print(k)
	}

}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe("0.0.0.0:8000", nil)
}