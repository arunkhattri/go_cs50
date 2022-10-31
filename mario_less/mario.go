package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Enter height of the pyramid:")

	var h int
	for true {
		_, err := fmt.Scan(&h)
		if err == nil {
			break
		}
		fmt.Println("Not a valid entry...")
		var dump string
		fmt.Scanln(&dump)
	}

	sym := "#"

	for i := 0; i < h; i++ {
		fmt.Printf("%*s\n", h, strings.Repeat(sym, i+1))
	}

}
