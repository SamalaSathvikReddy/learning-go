package main

import (
	"fmt"
	"math"
)

func add(x, y int) int {
	return x + y
}

func norm2(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println("Square root of 100", math.Sqrt(100))
	fmt.Println(math.Pi)
	fmt.Println(add(4, 3))
	fmt.Println(norm2(4.0, 3.0))
	var a, b string = swap("string 1", "string 2")
	fmt.Println(a)
	fmt.Println(b)

	var python, java bool = false, false
	fmt.Println(python, java)

	var i int = 77
	var j string = (string)(i)
	fmt.Println(j)

	fmt.Printf("j is of type %T\n", j)

	const Truth = true
	fmt.Println(Truth)
}
