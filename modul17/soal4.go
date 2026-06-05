package main

import "fmt"

func main() {

	var n int
	var jumlah float64

	fmt.Scan(&n)

	for i := 1; i <= n; i++ {

		// hitg penyebut: 1,3,5,7,...
		penyebut := float64(2*i - 1)

		// nilai suku
		suku := 1.0 / penyebut

		// jika suku ke-i genap, ubah menjadi negatif
		if i%2 == 0 {
			suku = -suku
		}

		jumlah = jumlah + suku
	}

	pi := 4 * jumlah

	fmt.Printf("%.6f\n", pi)
}
