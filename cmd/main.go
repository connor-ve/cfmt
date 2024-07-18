package main

import (
	"cfmt"
	_ "fmt"
)

type MockData struct {
	MockString string
	MockInt    int
}

func main() {
	cfmt.Print("->bl !u/e", "HelloWorld")
}
