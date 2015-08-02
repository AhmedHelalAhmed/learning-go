package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}
func main() {
	a := Person{"Roshan Jossey", 45}
	r := Person{"Irunnidem Kulam", -34}
	fmt.Println(a, r)
}
