package main

import (
	"fmt"
	"math"
	"runtime"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func flipCoin(s string) string {
	if s == "money" {
		return "Heads"
	}
	return "Tails"
}

func maxPow(x, n, lim float64) bool {
	if v := math.Pow(x, n); v <= lim {
		return true
	}
	return false
}

func binSqrt(x int) int {
	l := 0
	r := x

	for r > l+1 {
		m := l + (r-l)/2
		if m*m >= x {
			r = m
		} else {
			l = m
		}
	}
	return r
}

func main() {
	var sum int = 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	var cnt int = 0
	for sum > 0 {
		cnt += 1
		sum /= 2
	}
	fmt.Println(cnt)

	// for {
	// infinite loop
	// fmt.Print("1 ")
	// }

	fmt.Println(sqrt(2))
	fmt.Println(sqrt(-4))

	fmt.Println(flipCoin("money"))
	fmt.Println(flipCoin("cash"))
	fmt.Println(maxPow(2, 2, 4))
	fmt.Println(maxPow(3, 2, 8))

	fmt.Println(binSqrt(101))

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s \n", os)
	}

	defer fmt.Println("Gozibus")
	defer fmt.Println("Konnichiwa")
	fmt.Println("Arigato")

	// first Arigato then it's prev -> Konnichiwa then it's prev -> Gozibus

		

}
