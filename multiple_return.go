package main

import "fmt"

func vals() (int, int) {
    return 3, 7
}

func main() {

    a, b := vals() // a = 3, b = 7
    fmt.Println(a) // 3 
    fmt.Println(b) // 7

    _, c := vals() // _ = 3, c = 7
    fmt.Println(c) // 7
}