package product

import "fmt"

type Product struct { //nama struktur Data-nya
	Productname string
	Qty         int32
}

func New(name string, qty int32) Product {
	p := Product{
		Productname: name,
		Qty:         qty,
	}
	return p
}

func (p Product) GetProduct() {
	fmt.Println(p)
}
