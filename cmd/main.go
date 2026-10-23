package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)


func main() {
	fmt.Println(os.Getpid())
	fmt.Println("hit ehre")
	buffer := make([]byte, 4096)
	_, err := os.Open("file.txt")
	_, err = syscall.Read(34, buffer)
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(10*time.Minute)
	

}