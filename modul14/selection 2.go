package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m)

		var a [100]int
		var ganjil [100]int
		var genap [100]int
		var g, e int

		// Input data
		for j := 0; j < m; j++ {
			fmt.Scan(&a[j])
			if a[j]%2 == 1 {
				ganjil[g] = a[j]
				g++
			} else {
				genap[e] = a[j]
				e++
			}
		}

		// Urut ganjil naik
		for j := 0; j < g-1; j++ {
			for k := j + 1; k < g; k++ {
				if ganjil[j] > ganjil[k] {
					ganjil[j], ganjil[k] = ganjil[k], ganjil[j]
				}
			}
		}

		// Urut genap turun
		for j := 0; j < e-1; j++ {
			for k := j + 1; k < e; k++ {
				if genap[j] < genap[k] {
					genap[j], genap[k] = genap[k], genap[j]
				}
			}
		}

		// Output ganjil
		for j := 0; j < g; j++ {
			fmt.Print(ganjil[j], " ")
		}

		// Output genap
		for j := 0; j < e; j++ {
			fmt.Print(genap[j], " ")
		}

		fmt.Println()
	}
}
