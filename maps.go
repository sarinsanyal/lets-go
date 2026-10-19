package main

import (
	"fmt"
	"maps"
)

func main() {
	basic_map()
	operations()
}

func basic_map() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	print_keys(m)
}

func print_keys(m map[string]int) {
	v1 := m["k1"]
	v2 := m["k2"]

	fmt.Println("v1:", v1)
	fmt.Println("v2:", v2)

	fmt.Println("len:", len(m))
}

func operations() {
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
