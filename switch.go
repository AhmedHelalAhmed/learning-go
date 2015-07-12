package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Os")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println(os)
	default:
		fmt.Printf("%s", os)
	}
}
