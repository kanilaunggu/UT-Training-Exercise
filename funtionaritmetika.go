package main

import 
( 
	"fmt"
)

func main() {
	x := 8
	y := 3

	var s string
	ulang:
	fmt.Print("Masukkan huruf m,k,d,a : ")
	fmt.Scan(&s)
	//fmt.Scanln(&s)

	if s == "a" {
		result := Add(x, y)
		fmt.Print("Hasil tambah")
		fmt.Println(result)
        goto ulang  
	} else if s == "m" {

		minku := Min(x, y)
		fmt.Print("Hasil kurang ")
		fmt.Println(minku)
		goto ulang
	} else if s == "k" {

		result1 := Multiple(x, y)
		fmt.Print("Hasil kali ")
		fmt.Println(result1)
		goto ulang
	} else if s == "d" {
		result2 := Division(x, y)
		fmt.Print("Hasil Bagi ")
		fmt.Println(result2)
		goto ulang
	} else {
		fmt.Println("Keluar")
	}

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

func Division(x, y int) float32 {
	var value = float32(x) / float32(y)
	return value
}
