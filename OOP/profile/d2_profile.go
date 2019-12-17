package profile

import "fmt"

type User struct { //nama struktur Data-nya
	Firstname string
	Lastname  string
	Address   string
}

func (u User) GetProfile() {
	u.Firstname = "Khani"
	u.Lastname = "Launggu"
	u.Address = "Sulawesi"

	fmt.Println(u)
}

func (u User) SetProfile(fname string, lname string, addr string) {
	u.Firstname = fname
	u.Lastname = lname
	u.Address = addr

	fmt.Println(u)
}
