package main

import (
	"fmt"
	"math/rand"
	"time"
)

// SOAL 1A
type Domino struct {
	sisi1 int
	sisi2 int
	nilai int
	balak bool
}

type Dominoes struct {
	kartu [28]Domino
	sisa  int
}

// Membuat 1 set domino
func buatDominoes() Dominoes {

	var d Dominoes
	idx := 0

	for i := 0; i <= 6; i++ {
		for j := i; j <= 6; j++ {

			d.kartu[idx] = Domino{
				sisi1: i,
				sisi2: j,
				nilai: i + j,
				balak: i == j,
			}

			idx++
		}
	}

	d.sisa = 28
	return d
}

// SOAL 1B Kocok Kartu

func kocokKartu(d *Dominoes) {

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < d.sisa; i++ {

		j := rand.Intn(d.sisa)

		d.kartu[i], d.kartu[j] =
			d.kartu[j], d.kartu[i]
	}
}

// SOAL 1C ambil Kartu
func ambilKartu(d *Dominoes) Domino {

	k := d.kartu[d.sisa-1]
	d.sisa--

	return k
}

// SOAL 1D Gambar Kartu
func gambarKartu(k Domino, suit int) int {

	if suit == 1 {
		return k.sisi1
	}

	return k.sisi2
}

// SOAL 1E nilai Kartu
func nilaiKartu(k Domino) int {
	return k.nilai
}

// SOAL 2A Gali Kartu
func galiKartu(d *Dominoes, acuan Domino) {

	for d.sisa > 0 {

		k := ambilKartu(d)

		if k.sisi1 == acuan.sisi1 ||
			k.sisi1 == acuan.sisi2 ||
			k.sisi2 == acuan.sisi1 ||
			k.sisi2 == acuan.sisi2 {

			fmt.Println("Kartu cocok ditemukan:", k)
			return
		}
	}

	fmt.Println("Tidak ada kartu yang cocok")
}

// SOAL 2B sepasang Kartu
func sepasangKartu(k1 Domino, k2 Domino) bool {
	return k1.nilai+k2.nilai == 12
}

// SOAL 3 Tampilkan Kartu
func tampilKartu(tangan [7]Domino) {

	fmt.Println("\nKartu Pemain")

	for i := 0; i < 7; i++ {

		fmt.Printf("%d. [%d|%d]\n",
			i+1,
			tangan[i].sisi1,
			tangan[i].sisi2)
	}
}

// SOAL 3 cek Bisa Pasang

func bisaPasang(k Domino, ujung int) bool {

	return k.sisi1 == ujung ||
		k.sisi2 == ujung
}

func main() {

	domino := buatDominoes()

	kocokKartu(&domino)

	// Membagikan 7 kartu

	var tangan [7]Domino

	for i := 0; i < 7; i++ {
		tangan[i] = ambilKartu(&domino)
	}

	// Membuka kartu pertama

	rantai := ambilKartu(&domino)

	kiri := rantai.sisi1
	kanan := rantai.sisi2

	skor := 0

	fmt.Printf("Kartu Awal : [%d|%d]\n",
		rantai.sisi1,
		rantai.sisi2)

	for {

		tampilKartu(tangan)

		fmt.Println("\nUjung kiri :", kiri)
		fmt.Println("Ujung kanan:", kanan)

		fmt.Println("\nPilihan")
		fmt.Println("-1 s/d -7 : pasang kiri")
		fmt.Println("1 s/d 7   : pasang kanan")
		fmt.Println("0         : selesai ronde")
		fmt.Println("9         : selesai permainan")

		var pilih int

		fmt.Print("Masukkan pilihan : ")
		fmt.Scan(&pilih)

		if pilih == 9 {
			fmt.Println("Permainan selesai")
			break
		}

		if pilih == 0 {
			fmt.Println("Ronde selesai")
			break
		}

		index := pilih

		if pilih < 0 {
			index = -pilih
		}

		index--

		if index < 0 || index >= 7 {
			fmt.Println("Pilihan tidak valid")
			continue
		}

		k := tangan[index]

		if pilih < 0 {

			if bisaPasang(k, kiri) {

				if k.sisi1 == kiri {
					kiri = k.sisi2
				} else {
					kiri = k.sisi1
				}

				fmt.Println("Kartu berhasil dipasang")
				skor++

			} else {
				fmt.Println("Kartu tidak cocok")
			}
		}

		if pilih > 0 {

			if bisaPasang(k, kanan) {

				if k.sisi1 == kanan {
					kanan = k.sisi2
				} else {
					kanan = k.sisi1
				}

				fmt.Println("Kartu berhasil dipasang")
				skor++

			} else {
				fmt.Println("Kartu tidak cocok")
			}
		}
	}

	fmt.Println("\nNilai ronde =", skor)
}
