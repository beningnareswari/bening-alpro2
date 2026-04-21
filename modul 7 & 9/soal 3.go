package main

import "fmt"

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var hasil [100]string
	var i int

	fmt.Print("Klub A: ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B: ")
	fmt.Scan(&klubB)

	i = 0

	for {
		fmt.Print("Masukkan skor: ")
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			fmt.Println("Hasil:", klubA)
			hasil[i] = klubA
			i++
		} else if skorB > skorA {
			fmt.Println("Hasil:", klubB)
			hasil[i] = klubB
			i++
		} else {
			fmt.Println("Hasil: Draw")
		}
	}

	fmt.Println("Pertandingan selesai")
}
