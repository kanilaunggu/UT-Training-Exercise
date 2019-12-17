package main

import (
	"oop/profile"
)

func main() {
	//var name = "Khani"
	profile := profile.User{}

	profile.GetProfile()
	profile.SetProfile("Khani", "Launggu", "Sulawesi")

}
