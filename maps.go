package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bo": {34, 35},
	"98": {3, 5},
}

func main() {
	fmt.Println(m)
}
