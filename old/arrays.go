package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp:", a)

	a[1] = 100
	fmt.Println("emp:", a)
	fmt.Println("First element is:", a[1])

	fmt.Println("Len of Array:", len(a))
	
	b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)

	b = [...]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)

	// position based values
	b = [...]int{100, 3: 400, 500}
    fmt.Println("idx:", b)

	// 2d
	var twoD [2][3]int
    for i := range 2 {
        for j := range 3 {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)

}
