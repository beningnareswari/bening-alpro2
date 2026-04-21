package main

import "fmt"

func main() {
	var n, i int
	var arr [100]int
	var x, hapus int

	// input jumlah elemen
	fmt.Print("Jumlah elemen: ")
	fmt.Scan(&n)

	// input isi array
	fmt.Println("Masukkan elemen array:")
	for i = 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	// a. tampilkan seluruh isi array
	fmt.Print("Isi array: ")
	for i = 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	// b. tampilkan elemen indeks ganjil
	fmt.Print("Indeks ganjil: ")
	for i = 0; i < n; i++ {
		if i%2 != 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	// c. tampilkan elemen indeks genap
	fmt.Print("Indeks genap: ")
	for i = 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	// d. tampilkan elemen indeks kelipatan x
	fmt.Print("Masukkan x: ")
	fmt.Scan(&x)

	fmt.Print("Indeks kelipatan ", x, ": ")
	for i = 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	// e. hapus elemen pada indeks tertentu
	fmt.Print("Masukkan indeks yang dihapus: ")
	fmt.Scan(&hapus)

	for i = hapus; i < n-1; i++ {
		arr[i] = arr[i+1]
	}
	n = n - 1

	fmt.Print("Array setelah dihapus: ")
	for i = 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}
