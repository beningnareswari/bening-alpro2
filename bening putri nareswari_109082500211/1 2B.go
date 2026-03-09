package main

import "fmt"

func main() {
	var w1, w2, w3, w4 string
	var hasil = true

	for i := 1; i <= 5; i++ {
		fmt.Scan(&w1, &w2, &w3, &w4)

		if w1 == "merah" && w2 == "kuning" && w3 == "hijau" && w4 == "ungu" {
			hasil = true
		} else {
			hasil = false
		}
	}

	fmt.Println(hasil)
}
