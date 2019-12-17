package main

import (
	"oop/product"
	"oop/profile"
)

func main() {
	//var name = "Khani"
	profile := profile.User{}

	profile.GetProfile()
	profile.SetProfile("Khani", "Launggu", "Sulawesi")

	product := product.New("A", 200)
	product.GetProduct()

}
