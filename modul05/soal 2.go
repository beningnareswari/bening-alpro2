package main

import "fmt"

func cetakBintang(baris int) {
	if baris == 0 {
		return
	}

	cetakBintang(baris - 1)

	for i := 0; i < baris; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}

func main() {
	var jumlahBaris int

	fmt.Print("Masukkan jumlah baris: ")
	fmt.Scan(&jumlahBaris)

	cetakBintang(jumlahBaris)
}
