package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)


func main() {
	f , err := os.Open("file.txt")

	if err != nil {
		panic("oh my gosh")
	}

	dupFd, _ := syscall.Dup(int(f.Fd()))

	fmt.Println("original FD:", int(f.Fd()))
	fmt.Println("duplicate FD:", dupFd)
	
	// read from original fd
	buffer := make([]byte, 3)
	n1, _ := syscall.Read(int(f.Fd()), buffer)

	fmt.Printf("Original read: %q (%d bytes)\n", buffer[:n1], n1)

	buffer2 := make([]byte, 3)
	n2, _ := syscall.Read(int(f.Fd()), buffer)

	fmt.Printf("dup read: %q (%d bytes)\n", buffer2[:n2], n2)
	time.Sleep(10*time.Minute)
	

}