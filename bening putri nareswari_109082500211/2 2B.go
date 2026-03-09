package main

import "fmt"

func main() {
	var bunga string
	var pita string
	var i int = 1
	var jumlah int = 0

	for {
		fmt.Print("Bunga ", i, ": ")
		fmt.Scan(&bunga)

		if bunga == "SELESAI" {
			break
		}

		pita = pita + bunga + " - "
		jumlah++

		i++
	}

	fmt.Println("Pita:", pita)
	fmt.Println("Bunga:", jumlah)
}
