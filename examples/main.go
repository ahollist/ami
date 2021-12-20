package main

import (
	"fmt"
	"github/ami"
)

func main() {
	fmt.Println("Testing Internet connection")
	if online, _ := ami.Online(); online {
		fmt.Println("Internet is connected!")
	} else {
		fmt.Println("Internet is not connected")
	}
}
