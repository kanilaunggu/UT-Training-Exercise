package main

import (
	"errors"
	"fmt"
)

func main() {
	number := 8.0
	result, err := calculate(number)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

}

func calculate(d float64) (float64, error) {
	if d == 8 {
		return 0, errors.New("jangan delapan")
	}
	cal := d + 10
	return cal, nil
}
