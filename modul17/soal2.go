package main

import "fmt"

func main() {

	var x, kata string
	var n int

	// input string yang dicari
	fmt.Scanln(&x)

	// banyak data
	fmt.Scanln(&n)

	jumlah := 0
	posisi := -1

	for i := 1; i <= n; i++ {

		fmt.Scanln(&kata)

		if kata == x {

			jumlah++

			// simpan posisi pertama
			if posisi == -1 {
				posisi = i
			}
		}
	}

	// a. apakah ada
	if jumlah > 0 {
		fmt.Println("Ada")
	} else {
		fmt.Println("Tidak Ada")
	}

	// b. posisi pertama
	fmt.Println(posisi)

	// c. jumlah kemunculan
	fmt.Println(jumlah)

	// d. minimal 2 kali
	if jumlah >= 2 {
		fmt.Println("Ya")
	} else {
		fmt.Println("Tidak")
	}
}
