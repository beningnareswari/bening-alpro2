package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var berat float64

	// data pertama
	fmt.Scan(&berat)
	min := berat
	max := berat

	//sisa data
	for i := 2; i <= n; i++ {
		fmt.Scan(&berat)

		if berat < min {
			min = berat
		}
		if berat > max {
			max = berat
		}
	}

	fmt.Println(min, max)
}