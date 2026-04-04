package main

import "fmt"

type point struct {
	X int
	Y int
}

type pair struct {
	V int
	f bool
}

type Vertex struct {
	Latitude, Longitued float64
}

func main() {
	var i, j int = 42, 112
	p1 := &i
	p2 := &j
	fmt.Println(*p1)
	fmt.Println(*p2)

	v := point{100, 100}
	v.X = 200
	fmt.Println(v.X)
	p := &v
	p.X = 1e9
	fmt.Println(v)
	fmt.Println(p)

	var a [10]int

	for i := 0; i < 10; i++ {
		a[i] = i + 1
	}

	fmt.Println(a)

	primes := [10]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}

	fmt.Println(primes)

	var slice []int = primes[0:7]
	fmt.Println(slice)

	names := [5]string{
		"Naruto",
		"Sasuke",
		"Obito",
		"Kakashi",
		"Shikamaru",
	}

	slice1 := names[0:3]
	slice2 := names[1:4]

	slice1[2] = "Rin"
	fmt.Println(slice1, slice2)

	fmt.Println(names)

	// slice literals
	q := []int{2, 4, 6, 8, 10}
	fmt.Println(q)

	z := []pair{{2, true}, {3, false}}
	z[0].f = false
	z[0].V += (z[0].V) / 2
	fmt.Println(z)

	newSlice := names[:]
	fmt.Println(cap(newSlice))
	fmt.Println(len(newSlice))

	var zd []int
	// if zd == nil {
	// 	fmt.Println("zd is nil!")
	// }
	fmt.Println(zd)

	za := make([]int, 5)
	fmt.Println(za, len(za), cap(za))

	zb := make([]int, 0, 5)
	fmt.Println(zb, len(zb), cap(zb))

	for i := 0; i < 5; i++ {
		zb = append(zb, i)
	}

	fmt.Println(zb)

	pow := [10]int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512}

	for i, v := range pow {
		fmt.Printf("2^%d = %d\n", i, v)
	}

	for _, v := range pow {
		fmt.Printf("%d ", v)
	}
	fmt.Println()

	var m map[string]Vertex
	m = make(map[string]Vertex)
	m["Bell labs"] = Vertex{40.111, -73.8867}
	m["Google"] = Vertex{37.422, -112.7993}
	element, ok := m["Google"]
	fmt.Println(element)
	delete(m, "Bell labs")
	fmt.Println(ok)
	fmt.Println(m)
}
