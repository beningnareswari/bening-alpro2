package main

import "fmt"

func hitungSkor(jmlSoal *int, totalWaktu *int) {
	*jmlSoal = 0
	*totalWaktu = 0

	for i := 0; i < 8; i++ {
		var waktuSoal int
		fmt.Scan(&waktuSoal)

		if waktuSoal < 301 {
			*jmlSoal++
			*totalWaktu += waktuSoal
		}
	}
}

func main() {
	var namaPeserta string

	terbanyakSoal := -1
	waktuTercepat := 1000000
	namaPemenang := ""

	for {
		fmt.Scan(&namaPeserta)

		if namaPeserta == "Selesai" {
			break
		}

		var soalSelesai, totalNilai int
		hitungSkor(&soalSelesai, &totalNilai)

		if soalSelesai > terbanyakSoal || (soalSelesai == terbanyakSoal && totalNilai < waktuTercepat) {
			terbanyakSoal = soalSelesai
			waktuTercepat = totalNilai
			namaPemenang = namaPeserta
		}
	}

	fmt.Println(namaPemenang, terbanyakSoal, waktuTercepat)
}
