package main

import "fmt"

func pangkat(x int, y int) int {
	if y == 0 {
		return 1
	}

	return x * pangkat(x, y-1)
}

func main() {
	var angka int
	var pangkatKe int

	fmt.Print("angka: ")
	fmt.Scan(&angka)

	fmt.Print("pangkat: ")
	fmt.Scan(&pangkatKe)

	fmt.Println("Hasil:", pangkat(angka, pangkatKe))
}
