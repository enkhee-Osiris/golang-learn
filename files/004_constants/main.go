package main

import (
	"fmt"
	"math"
)

const us = "string constant" // untyped string
const s string = "string constant"

func main() {
	fmt.Println(s)

	const n = 50000000
	const d = 3e20 / n // untyped float
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
