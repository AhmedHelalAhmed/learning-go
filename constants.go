package main

import (
    "fmt"
    "math"
)

const ro string = "heoll"

func main(){
    fmt.Println(ro)
    const n = 5000
    const b = 4e10 / n
    fmt.Println(b,n)
    fmt.Println(int64(b))
    fmt.Println(math.Sin(n))
}
