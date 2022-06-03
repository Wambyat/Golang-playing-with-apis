package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/xuri/excelize/v2"
)

//Here we send a JSON to the API and get the response in a JSON
func main() {
	httpurl := "https://reqres.in/api/users"
	fmt.Println("API URL: ", httpurl)
	f, err := excelize.OpenFile("data.xlsx")
	if err != nil {
		fmt.Println("Error in opening file: ", err)
	}
	f.SetCellValue("Sheet1", "C1", "Output")

	//Extract values from excel sheet
	row,err:=f.GetRows("Sheet1")
	if err != nil {
		fmt.Println("Error in extracting number of rows: ", err)
	}
	
	//SI stores size of row
	SI:=len(row)

	for i := 2; i <= SI; i++ {
		name, err := f.GetCellValue("Sheet1", "A"+fmt.Sprint(i))
		if err != nil {
			fmt.Println("Error in getting cell value: ", err)
		}
		job, err := f.GetCellValue("Sheet1", "B"+fmt.Sprint(i))
		if err != nil {
			fmt.Println("Error in getting cell value: ", err)
		}
		
		var jsonData = []byte(`{
			"name": "` + name + `",
			"job": "` + job + `"
		}`)

		request, _ := http.NewRequest("POST", httpurl, bytes.NewBuffer(jsonData))
		request.Header.Set("Content-Type", "application/json; charset=UTF-8")

		client := &http.Client{}
		response, error := client.Do(request)
		if error != nil {
			panic(error)
		}
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		}

		a := string(body)
		b := strings.Index(a, "<pre>")
		if b != -1 {
			c := strings.Index(a, "</pre>")
			fmt.Println(a[b+5 : c])
		} else {
			f.SetCellValue("Sheet1", "C"+fmt.Sprint(i), a)
			fmt.Println(a)
		}

		f.SaveAs("data.xlsx")

	}

}
