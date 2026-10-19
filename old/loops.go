package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println("value of", i, "th number is", i)
		i = i + 1
	}

    for j := 0; j < 3; j++ {
        fmt.Println(j)
    }

    fmt.Println("\nBackwards from 10")
    for k := 10; k > 0; k--{
        fmt.Println(k)
    }

    for i := range 3 {
        fmt.Println("range", i)
    }

    // while loops
    for {
        fmt.Println("1 iter of a while loop")
        break
    }
}
