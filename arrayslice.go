package main

import "fmt"

func main() {

	var names [4]string
	names[0] = "trafalgar"
	names[1] = "d"
	names[2] = "water"
	names[3] = "law"

	fmt.Println(names[0], names[1], names[2], names[3])

	//Slice

	var fruits = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(fruits[0]) // "apple"

	//var fruitsA = []string{"apple", "grape"}   // slice
	//var fruitsB = [2]string{"banana", "melon"} // array

}
