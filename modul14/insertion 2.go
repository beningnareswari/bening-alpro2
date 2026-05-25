package main

import "fmt"

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

func main() {

	var n int
	fmt.Scan(&n)

	var b [100]Buku

	// Input data buku
	for i := 0; i < n; i++ {
		fmt.Scan(
			&b[i].id,
			&b[i].judul,
			&b[i].penulis,
			&b[i].penerbit,
			&b[i].eksemplar,
			&b[i].tahun,
			&b[i].rating,
		)
	}

	// Insertion Sort Descending berdasarkan rating
	for i := 1; i < n; i++ {

		temp := b[i]
		j := i - 1

		for j >= 0 && b[j].rating < temp.rating {
			b[j+1] = b[j]
			j--
		}

		b[j+1] = temp
	}

	// Buku terfavorit
	fmt.Println("Buku Terfavorit:")
	fmt.Println(b[0].judul, b[0].rating)

	// 5 rating tertinggi
	fmt.Println("5 Buku Rating Tertinggi:")

	batas := 5
	if n < 5 {
		batas = n
	}

	for i := 0; i < batas; i++ {
		fmt.Println(b[i].judul)
	}

	// Cari rating
	var cari int
	fmt.Scan(&cari)

	kiri := 0
	kanan := n - 1
	ketemu := false

	for kiri <= kanan {

		tengah := (kiri + kanan) / 2

		if b[tengah].rating == cari {

			fmt.Println("Data Buku:")
			fmt.Println(b[tengah].judul)
			fmt.Println(b[tengah].penulis)
			fmt.Println(b[tengah].penerbit)
			fmt.Println(b[tengah].tahun)
			fmt.Println(b[tengah].eksemplar)
			fmt.Println(b[tengah].rating)

			ketemu = true
			break

		} else if cari < b[tengah].rating {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}

	if !ketemu {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}
