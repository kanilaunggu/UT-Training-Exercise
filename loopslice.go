package main

import "fmt"

func main() {
	var fruits = []string{"foo1", "Bar1", "Foo2", "Bar2"}

    for _,i := range fruits {
	    fmt.Println(i)
      }
}