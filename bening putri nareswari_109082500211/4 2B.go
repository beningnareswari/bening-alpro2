package main

import "fmt"

func main() {
	var k float64
	var akar float64

	for {
		fmt.Print("Nilai k = ")
		fmt.Scan(&k)

		if k < 0 {
			break
		}

		akar = (4*k + 2) / (4*k + 1)

		fmt.Printf("Nilai akar 2 = %.10f\n", akar)
	}
}
