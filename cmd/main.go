package main

import (
	"cfmt"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	leah := Person{
		Name: "Leah Peck",
		Age:  24,
	}
	cfmt.Print("red!", leah)

	x := cfmt.Sprint("red", "Hwllo")
	y := cfmt.Sprint("red!", "Hwllo")
	fmt.Println(y + x + y)

}
