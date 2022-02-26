package main

import (
	"encoding/json"
	"fmt"
)

func GetMyData() (string, error) {
	return `{ "name": "ricky", "age": 99 }`, nil
}

type MyData struct {
	Age int
}

func main() {

	payload, err := GetMyData()
	if err != nil {
		// error occurred
	}
	fmt.Println("before", payload)

	var result MyData
	jsonErr := json.Unmarshal([]byte(payload), &result)
	if jsonErr != nil {
		// error occurred
	}

	age := result.Age

	fmt.Printf("%d %T\n", age, age)

	result.Age = 4

	transformedData, marshalErr := json.Marshal(result)
	if marshalErr != nil {
		// error occurred
	}

	fmt.Println("after", string(transformedData))
}
