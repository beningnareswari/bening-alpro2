package main

import "fmt"

// SOAL 4A variabel Global
var teks string
var idx int

func start() { //START
	idx = 0
}

func maju() { //MAJU
	idx++
}

func eop() bool { //EOP

	if idx >= len(teks) {
		return true
	}

	return teks[idx] == '.'
}

func cc() byte { //CC
	return teks[idx]
}

func main() {

	fmt.Print("Masukkan teks (akhiri dengan titik): ")
	fmt.Scanln(&teks)

	fmt.Println("\nKarakter yang terbaca:")

	start()

	for !eop() {
		fmt.Printf("%c ", cc())
		maju()
	}

	totalKarakter := 0
	jumlahA := 0
	jumlahLE := 0

	start()

	for idx < len(teks)-1 && !eop() {

		totalKarakter++

		if cc() == 'A' {
			jumlahA++
		}

		if teks[idx] == 'L' &&
			teks[idx+1] == 'E' {

			jumlahLE++
		}

		maju()
	}

	frekuensiA := 0.0

	if totalKarakter > 0 {
		frekuensiA = float64(jumlahA) / float64(totalKarakter)
	}

	fmt.Println("\nJumlah karakter =", totalKarakter)
	fmt.Println("Jumlah huruf A =", jumlahA)
	fmt.Println("Frekuensi huruf A =", frekuensiA)
	fmt.Println("Jumlah kata LE =", jumlahLE)
}
