package main
import (
 "encoding/json"
 "fmt"
 "io/ioutil"
 "net/http"
)

type Response struct {
 ID string `json:"id"`
 Joke string `json:"joke"`
 Status int `json:"status"`
}
/*The API here returns a JSON object with the following fields:
ID: The ID of the joke.
Joke: The joke itself.
Status: The status of the request.
*/
func main() {
	fmt.Println(" starting ")
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://icanhazdadjoke.com/", nil)
	if err != nil {
	fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
	fmt.Print(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	fmt.Print(err.Error())
	}
	var responseObject Response
	json.Unmarshal(bodyBytes, &responseObject)
	fmt.Printf(" %v\n", responseObject)
}