package main

import "fmt"

func main() {

	var a [1000]int
	var n, x int

	for {
		fmt.Scan(&x)

		if x == -5313 {
			break
		}

		if x != 0 {
			a[n] = x
			n++
		} else {

			// selection sort
			for i := 0; i < n-1; i++ {
				for j := i + 1; j < n; j++ {

					if a[i] > a[j] {
						a[i], a[j] = a[j], a[i]
					}

				}
			}

			// cari mediannyaa
			if n%2 == 1 {
				fmt.Println(a[n/2])
			} else {
				tengah1 := a[n/2-1]
				tengah2 := a[n/2]

				fmt.Println((tengah1 + tengah2) / 2)
			}
		}
	}
}
