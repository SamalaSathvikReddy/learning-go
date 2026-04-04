package main

import "fmt"
import "math"

type Vertex struct {
	X int
	Y int
}

type Myfloat float64

func (f Myfloat) abs() Myfloat{
	if f < 0 {
		return (-f)
	}
	return f
}

func (V Vertex) norm1() float64 {
	return (math.Sqrt((float64)(V.X*V.X + V.Y*V.Y)))
}

func (V *Vertex) scale(s int) {
	V.X = V.X * s
	V.Y = V.Y * s
}

func main() {
	V := Vertex{3, 4}
	fmt.Println(V.norm1())
	var f Myfloat = -10.2
	fmt.Println(f.abs())
	V.scale(5)
	fmt.Println(V)
}
