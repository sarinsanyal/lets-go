package main

import "fmt"

func main(){
	if 7%2 == 0	{
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}


	if (9%2 == 0 || 7%2 == 0) {
        fmt.Println("either 9 or 7 are even")
    } else {
		fmt.Println("neither 9 or 7 are even")
	}

	
}