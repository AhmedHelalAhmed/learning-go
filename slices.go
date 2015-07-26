package main

import "fmt"

func main() {
	s := []int{2, 3, 4, 4, 55}
	fmt.Println("s==", s)
	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] is %d\n", i, s[i])
	}
	fmt.Println(s[:3])
	fmt.Println(s[2:])

}
