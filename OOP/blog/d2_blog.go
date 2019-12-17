package blog

import "fmt"

type BloInterface interface {
	GetBlog()
	SetBlog(string, string, string)
}

type Blog struct { //nama struktur Data-nya
	Title       string
	Descritpion string
	Authot      string
}

func New(title string, desc string, author string) Blog {
	return Blog{title, desc, author}
}

func (b Blog) GetBlog() {
	fmt.Println(b)
}

func (b Blog) SetBlog(title string, desc string, author string) {

}
