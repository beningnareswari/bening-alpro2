package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(arr arrBalita, n int, min *float64, max *float64) {
	*min = arr[0]
	*max = arr[0]

	for i := 1; i < n; i++ {
		if arr[i] < *min {
			*min = arr[i]
		}
		if arr[i] > *max {
			*max = arr[i]
		}
	}
}

func rerata(arr arrBalita, n int) float64 {
	total := 0.0
	for i := 0; i < n; i++ {
		total += arr[i]
	}
	return total / float64(n)
}

func main() {
	var n int
	var data arrBalita

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}

	var min, max float64
	hitungMinMax(data, n, &min, &max)

	rata := rerata(data, n)

	fmt.Println("Berat balita minimum:", min, "kg")
	fmt.Println("Berat balita maksimum:", max, "kg")
	fmt.Println("Rerata berat balita:", rata, "kg")
}
