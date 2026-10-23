package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)


func main() {
	fmt.Println(os.Getpid())
	var p [2]int
	err := syscall.Pipe(p[:])
	if err != nil {
		fmt.Println(err)
	}
	attr := &syscall.ProcAttr{
		Dir:   "",         // Current working directory (empty means inherit)
		Env:   os.Environ(), // Inherit environment variables
		Files: []uintptr{
			uintptr(syscall.Stdin),  // fd 0: Standard Input
			uintptr(syscall.Stdout), // fd 1: Standard Output
			uintptr(syscall.Stderr), // fd 2: Standard Error
		},
		Sys: nil, // Additional OS-specific process options if needed
	}
	pid, err := syscall.ForkExec("./main", []string{}, attr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pid)



	time.Sleep(10*time.Minute)
	

}