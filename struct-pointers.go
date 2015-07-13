package main

import "fmt"

type Vertes struct {
	X int
	Y int
}

func main() {
	v := Vertes{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}
