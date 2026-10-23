package main

import (

	"syscall"
	"fmt"
)


func main() {
	fmt.Println("hit ehre")
	buffer := make([]byte, 4096)
	_, err := syscall.Read(34, buffer)
	if err != nil {
		fmt.Println(err)
	}
	

}