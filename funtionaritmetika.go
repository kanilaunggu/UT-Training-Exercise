package main

import "fmt"

func main() {
	x := 8
	y := 2
	result := Add(x, y)
	fmt.Println(result)

	minku := Min(x, y)
	fmt.Println(minku)

	result1 := Multiple(x, y)
	fmt.Println(result1)

	result2 := Division(x, y)
	fmt.Println(result2)

}

func Add(x, y int) int {
	var value = x + y
	return value
}

func Min(x, y int) int {
	var value = x - y
	return value
}

func Multiple(x, y int) int {
	var value = x * y
	return value
}

func Division(x, y int) int {
	var value = x / y
	return value
}
