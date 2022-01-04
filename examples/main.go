package main

import (
	"fmt"
	"github/ami"
)

func main() {
	fmt.Println("Testing Internet connection")
	if online, _ := ami.Online(ami.WithAddress("8.8.4.4")); online {
		fmt.Println("Internet is connected!")
	} else {
		fmt.Println("Internet is not connected")
	}
}
