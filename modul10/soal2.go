package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	var berat [1000]float64

	for i := 0; i < x; i++ {
		fmt.Scan(&berat[i])
	}

	var totalPerWadah []float64
	index := 0

	for index < x {
		total := 0.0

		for i := 0; i < y && index < x; i++ {
			total += berat[index]
			index++
		}

		totalPerWadah = append(totalPerWadah, total)
	}

	totalSemua := 0.0
	for i := 0; i < len(totalPerWadah); i++ {
		fmt.Print(totalPerWadah[i], " ")
		totalSemua += totalPerWadah[i]
	}

	fmt.Println()

	rata := totalSemua / float64(len(totalPerWadah))
	fmt.Println(rata)
}
