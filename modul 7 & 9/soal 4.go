package main

import "fmt"

func main() {
	var kata string
	var reverse string = ""

	fmt.Print("Masukkan kata: ")
	fmt.Scan(&kata)

	for i := len(kata) - 1; i >= 0; i-- {
		reverse = reverse + string(kata[i])
	}

	fmt.Println("Reverse:", reverse)

	if kata == reverse {
		fmt.Println("Palindrom")
	} else {
		fmt.Println("Bukan palindrom")
	}
}
