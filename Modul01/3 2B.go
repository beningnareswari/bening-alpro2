package main

import "fmt"

func main() {
	var kiri, kanan float64
	var selisih float64

	for {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&kiri, &kanan)

		if kiri < 0 || kanan < 0 || kiri+kanan > 150 {
			break
		}

		if kiri > kanan {
			selisih = kiri - kanan
		} else {
			selisih = kanan - kiri
		}

		if selisih >= 9 {
			fmt.Println("Sepeda motor pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor pak Andi akan oleng: false")
		}
	}

	fmt.Println("Proses selesai.")
}
