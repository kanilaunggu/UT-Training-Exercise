package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

func main() {
	var JsonString = `{"name":"Khani", "age":"49"}`
	var JsonData = []byte(JsonString)

	var data User

	var err = json.Unmarshal(JsonData, &data)

	if err != nil {git 
		fmt.Println("Error Unmarshalling Json " + err.Error())
		return
	}

	fmt.Println(data)

	fmt.Println("Nama : " + data.Name)
	fmt.Println("Umur : " + data.Age)
}
