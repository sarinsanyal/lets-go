package main

import (
	"fmt"
	"time"
)

func main(){
	i := 2
	fmt.Println("Value of i is:", i)
	switch i {
	case 1:
		fmt.Println("This is 1")
	case 2:
		fmt.Println("This is 2")
	case 3:
		fmt.Println("This is 3")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("This is a weekend")
	default:
		fmt.Println("This is a Weekday - go to work")
	}
	fmt.Println("The time now is", time.Now())
}
