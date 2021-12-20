package main

import (
	"fmt"
	"github/ami"
)

func main() {
	fmt.Println("Testing Internet connection")
	if ami.Online() {
		fmt.Println("Internet is connected!")
	} else {
		fmt.Println("Internet is not connected")
	}
}
