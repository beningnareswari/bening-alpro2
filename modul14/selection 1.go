package main

import "fmt"

func main() {
	var n, m int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {

		fmt.Scan(&m)

		var rumah [1000]int

		for j := 0; j < m; j++ {
			fmt.Scan(&rumah[j])
		}

		// Selection Sort Ascending
		for j := 0; j < m-1; j++ {

			min := j

			for k := j + 1; k < m; k++ {
				if rumah[k] < rumah[min] {
					min = k
				}
			}

			rumah[j], rumah[min] = rumah[min], rumah[j]
		}

		for j := 0; j < m; j++ {
			fmt.Print(rumah[j], " ")
		}

		fmt.Println()
	}
}
