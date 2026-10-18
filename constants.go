package main

import (
    "fmt"
    "math"
)

const s string = "constant"

func main() {
    fmt.Println(s)

    const n = 500000000 // 5e8

    const d = 3e20 / n
    fmt.Println(d)

    fmt.Println(int64(d))

    fmt.Println(math.Sin(n))
	fmt.Println(math.Sin(math.Pi))
}

/*
Should give:

constant
6e11
600000000000
sin value
~ 0
*/