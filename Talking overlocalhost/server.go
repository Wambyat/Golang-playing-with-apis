package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", returnthings)
	http.ListenAndServe(":8080", nil)
}

func returnthings(w http.ResponseWriter, r *http.Request) {
	
	fmt.Fprintf(w, "Hello %s", r.URL.Path[1:])
	fmt.Println(w)
	fmt.Println(r)
}
