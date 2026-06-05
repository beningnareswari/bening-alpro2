package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var n int

	// banyak titik hujan
	fmt.Scan(&n)

	var A, B, C, D int

	for i := 0; i < n; i++ {

		// menghasilkan angka acak 0 sampai 1
		x := rand.Float64()
		y := rand.Float64()

		// daerah A (kiri bawah)
		if x < 0.5 && y < 0.5 {

			A++

			// daerah B (kanan bawah)
		} else if x >= 0.5 && y < 0.5 {

			B++

			// daerah C (kanan atas)
		} else if x >= 0.5 && y >= 0.5 {

			C++

			// daerah D (kiri atas)
		} else {

			D++
		}
	}

	fmt.Println("Daerah A :", A)
	fmt.Println("Daerah B :", B)
	fmt.Println("Daerah C :", C)
	fmt.Println("Daerah D :", D)
}
