package main

import (
	"fmt"
)

/*
Pointer adalah reference atau alamat memori. Variabel pointer berarti variabel yang berisi alamat memori suatu nilai.
Sebagai contoh sebuah variabel bertipe integer memiliki nilai 4,
maka yang dimaksud pointer adalah alamat memori dimana nilai 4 disimpan, bukan nilai 4 itu sendiri.
Variabel-variabel yang memiliki reference atau alamat pointer yang sama,
saling berhubungan satu sama lain dan nilainya pasti sama. Ketika ada perubahan nilai,
maka akan memberikan efek kepada variabel lain (yang referensi-nya sama) yaitu nilainya ikut berubah.
*/

func main() {
	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("numberA (value)   :", numberA)  // 4
	fmt.Println("numberA (address) :", &numberA) // 0xc20800a220

	fmt.Println("numberB (value)   :", *numberB) // 4
	fmt.Println("numberB (address) :", numberB)  // 0xc20800a220

}
