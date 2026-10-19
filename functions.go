package main

import "fmt"

func main() {
	fmt.Println(add(2,3))
	ans := add(5,9)
	fmt.Println("add(5,9) gives:", ans)
}

func add (a int, b int) int{ // variable <type> in this order
	return a + b
}
