package main

import (
	"cfmt"

	_ "fmt"
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
	cfmt.Printf("ReD", "Hello World %d", leah.Age)

}
