package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice("a", a)
	b := make([]int, 0, 5)
	printSlice("b", b)
	c := b[:2]
	for i := 0; i < 20; i++ {
		c = append(c, i)
		printSlice("c", c)
	}
	d := c[2:5]
	printSlice("d", d)
	var e []int
	for i := 0; i < 33; i++ {
		e = append(e, i)
		printSlice("e", e)
	}
}
func printSlice(s string, x []int) {
	fmt.Printf("%s len = %d cap = %d %v\n", s, len(x), cap(x), x)
}
