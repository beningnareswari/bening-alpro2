package main

import "fmt"

func main() {
	var a, b, c, d, e int
	var x, y, z rune

	fmt.Scan(&a, &b, &c, &d, &e)

	fmt.Printf("%c%c%c%c%c\n", a, b, c, d, e)

	fmt.Scanf("%c%c%c", &x, &y, &z)

	fmt.Printf("%d%d%d\n", x, y, z)
}
