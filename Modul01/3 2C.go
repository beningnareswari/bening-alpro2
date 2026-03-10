package main

import "fmt"

func main() {

	var b int
	var jumlahFaktor int = 0

	fmt.Print("Bilatgat: ")
	fmt.Scanln(&b)

	fmt.Print("Faktor: ")

	for i := 1; i <= b; i++ {
		if b%i == 0 {
			fmt.Print(i, " ")
			jumlahFaktor++
		}
	}

	fmt.Println()

	if jumlahFaktor == 2 {
		fmt.Println("prima: true")
	} else {
		fmt.Println("prima: false")
	}

}
