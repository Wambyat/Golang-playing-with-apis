package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println(" starting ")
	//Message is stored after "http://localhost:8080/
	resp, err := http.Get("http://localhost:8080/noice")
	
	//If err occurs print that
	if err != nil {
		fmt.Println(err)
	}
	
	body, err := ioutil.ReadAll(resp.Body)
	
	//If err occurs print that	
	if err != nil {
		fmt.Println(err)
	}
	
	sb := string(body)
	fmt.Println(sb)
}
