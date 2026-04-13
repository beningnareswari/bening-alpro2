package main

import "fmt"

func cariFaktor(n int, pembagi int) {
	if pembagi > n {
		return
	}
	if n%pembagi == 0 {
		fmt.Print(pembagi, " ")
	}
	cariFaktor(n, pembagi+1)
}

func main() {
	var angka int

	fmt.Scan(&angka)

	cariFaktor(angka, 1)
}
