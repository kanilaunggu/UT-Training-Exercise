package main

import (
	"fmt"
	"oop/blog"
	"oop/product"
	"oop/profile"
	"oop/store"
)

func main() {
	//var name = "Khani"
	profile := profile.User{}

	profile.SetProfile("Khani", "Launggu", "Sulawesi")
	profile.GetProfile()
	//profile.SetProfile("Khani", "Launggu", "Sulawesi")

	product := product.New("A", 200)
	product.GetProduct()

	blog := blog.New("A", "B", "C")
	blog.GetBlog()

	brancstore := store.BranchStore{StoreName: "Store: Toko apa saja,", OwnerBranch: "Owner: Saya dan Dia"}
	store := store.Store{BranchStore: brancstore, StoreName: "Toko apa saja 1", Owner: "Saya dan Dia juga"}

	store.GetBranchStore()

	fmt.Println(store)

	//cab := BranchStore{StoreName: "Cab1,", OwnerBranch: "Own1"}
	//fmt.Println(cab)

}
