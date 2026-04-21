package main

import (
	"fmt"
	"math"
)

type titik struct {
	x, y int
}

type lingkaran struct {
	pusat titik
	r     int
}

func jarak(a, b titik) float64 {
	return math.Sqrt(float64((a.x-b.x)*(a.x-b.x) + (a.y-b.y)*(a.y-b.y)))
}

func didalam(l lingkaran, p titik) bool {
	return jarak(l.pusat, p) <= float64(l.r)
}

func main() {
	var l1, l2 lingkaran
	var p titik

	// input lingkaran 1
	fmt.Scan(&l1.pusat.x, &l1.pusat.y, &l1.r)

	// input lingkaran 2
	fmt.Scan(&l2.pusat.x, &l2.pusat.y, &l2.r)

	// input titik
	fmt.Scan(&p.x, &p.y)

	d1 := didalam(l1, p)
	d2 := didalam(l2, p)

	if d1 && d2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if d1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if d2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
