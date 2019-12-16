/*
package main

import "encoding/json"
import "fmt"

type User struct {
	FullName string `json:"Name"`
	Asal string `json:"Asal"`
	Age      int
}

func main() {
	var jsonString = `[
    {"Name": "Khani","Asal":"Sulawesi", "Age": 40}

]`




	var data []User

	var json, err = json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		return
	}

	//fmt.Println("user 1:", data[0].FullName)
	//fmt.Println("user 2:", data[1].FullName)
}
*/

package main

import "encoding/json"
import "fmt"

type User struct {
	FullName string `json:"Name"`
	Asal     string `json:"Asal"`
	Age      int    `json:"Age"`
}

func main() {
	var object = []User{{"Khani","Sulawesi", 40},{"Agung","Jawa", 35}}
	var jsonData, err = json.Marshal(object)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var jsonString = string(jsonData)
	fmt.Println(jsonString)
}