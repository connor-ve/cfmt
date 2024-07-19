package main

import (
	_ "fmt"

	cfmt "github.com/connor-ve/cfmt"
)

type MockData struct {
	MockString string
	MockInt    int
}

func main() {
	cfmt.Print("", "HelloWorld")
}
