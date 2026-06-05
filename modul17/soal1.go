package main

import "fmt"

func main() {

	var x float64
	var jumlah float64
	var rata float64
	var cacah int

	fmt.Scan(&x)

	for x != 9999 {

		jumlah = jumlah + x
		cacah++

		fmt.Scan(&x)
	}

	rata = jumlah / float64(cacah)

	fmt.Printf("%.2f\n", rata)
}
