package main

import (
	"fmt"
	"github/ami"
)

func main() {
	fmt.Println("Testing Internet connection")
	if online := ami.Online(ami.WithAddress("8.8.4.4")); online {
		fmt.Println("Internet is connected!")
	} else {
		if err := ami.Error(); err != nil {
			fmt.Println("Error in Online(): ")
		}
		fmt.Println("Internet is not connected")
	}
}
