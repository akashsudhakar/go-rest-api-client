package main

import (
	"encoding/json"
	"fmt"
	"go-rest-api-client/model"
	"io/ioutil"
	"net/http"
)

func main() {
	callGetApi()
}

func callGetApi() {
	httpResponse, error := http.Get("http://localhost:8080/persons")
	if error != nil {
		fmt.Println("Exception while making rest call")
	}
	defer httpResponse.Body.Close()
	bodyBytes, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		fmt.Println("Exception while parsing response body")
	}
	var respStruct model.Response
	json.Unmarshal(bodyBytes, &respStruct)

	fmt.Println(len(respStruct.Persons))

	printElements(respStruct)
}

func printElements(respStruct model.Response) {
	for _, person := range respStruct.Persons {
		fmt.Println(person.Id, person.FirstName, person.LastName)
	}
}
