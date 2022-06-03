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

	//We format the jsonData into byte so it can be transmitted
	var jsonData = []byte(`{
		"name": "morpheus",
		"job": "leader"
	}`)

	//Here we POST the JSON
	request, error := http.NewRequest("POST", httpurl, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, error := client.Do(request)
	
	if error != nil {
		
		//"Nicer" way to deal with the error
		panic(error)
	}
	
	defer response.Body.Close()
	
	body, _ := ioutil.ReadAll(response.Body)
	//Here we can ignore error because we have used panic()
	
	fmt.Println("response Body:", string(body))

}
