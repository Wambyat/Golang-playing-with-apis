package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

//Here we send a JSON to the API and get the response in a JSON

func main() {
	httpurl := "https://reqres.in/api/users"
	fmt.Println("API URL: ", httpurl)

	var jsonData = []byte(`{
		"name": "morpheus",
		"job": "leader"
	}`)

	request, error := http.NewRequest("POST", httpurl, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))

}
