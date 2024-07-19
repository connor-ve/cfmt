package main

import (
	"fmt"

	cfmt "github.com/connor-ve/cfmt"
)

/*
This file is purely here for local testing
Feel free to call any and all functions to test your outputs and use case
*/
type MockData struct {
	MockString string
	MockInt    int
}

func main() {
	cfmt.Print("", "HelloWorld\n")
	cfmt.Println("orange", "HelloWorld")
	cfmt.Printf("", "HelloWorld")
	str := cfmt.Sprint("!#121212", "First Half")
	char := cfmt.Sprintf("/plum", "Second Half")
	fmt.Println(str + " " + char)
}
