package main

import "fmt"

func main() {
	var c float64

	fmt.Print("Masukkan Celsius: ")
	fmt.Scan(&c)

	reamur := c * 4 / 5
	fahrenheit := (c * 9 / 5) + 32
	kelvin := c + 273

	fmt.Println("Derajat Reamur:", reamur)
	fmt.Println("Derajat Fahrenheit:", fahrenheit)
	fmt.Println("Derajat Kelvin:", kelvin)
}
