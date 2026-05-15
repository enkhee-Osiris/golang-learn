package main

import (
	"fmt"
)

func main() {

	i := 1
	for i <= 3 { // same as while loop
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j < 3; j++ { // same as classical for loop
		fmt.Println(j)
	}

	for m := range 3 {
		fmt.Println("range", m)
	}

	for {
		fmt.Println("loop")
		break // To break from loop
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}

		fmt.Println(n)
	}
}
