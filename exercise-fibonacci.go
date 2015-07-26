package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	fibp := 1
	fibn := 1
	return func() int {
		fibp, fibn = fibn, fibn+fibp
		return fibp
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
