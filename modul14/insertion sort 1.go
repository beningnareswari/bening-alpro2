package main

import "fmt"

func main() {

	var a [100]int
	var n, x int

	// Input data
	for {
		fmt.Scan(&x)

		if x < 0 {
			break
		}

		a[n] = x
		n++
	}

	// Insertion Sort
	for i := 1; i < n; i++ {

		temp := a[i]
		j := i - 1

		for j >= 0 && a[j] > temp {
			a[j+1] = a[j]
			j--
		}

		a[j+1] = temp
	}

	// Output array
	for i := 0; i < n; i++ {
		fmt.Print(a[i], " ")
	}

	fmt.Println()

	// Cek jarak
	jarak := a[1] - a[0]
	sama := true

	for i := 2; i < n; i++ {

		if a[i]-a[i-1] != jarak {
			sama = false
		}
	}

	if sama {
		fmt.Println("Data berjarak", jarak)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}
