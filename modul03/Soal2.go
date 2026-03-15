package main

import "fmt"

// f(x) = x^2
func f(x int) int {
	return x * x
}

// g(x) = x - 2
func g(x int) int {
	return x - 2
}

// h(x) = x + 1
func h(x int) int {
	return x + 1
}

func main() {
	var a, b, c int

	fmt.Scan(&a, &b, &c)

	// (fogoh)(a) = f(g(h(a)))
	hasil1 := f(g(h(a)))

	// (gohof)(b) = g(h(f(b)))
	hasil2 := g(h(f(b)))

	// (hofog)(c) = h(f(g(c)))
	hasil3 := h(f(g(c)))

	fmt.Println(hasil1)
	fmt.Println(hasil2)
	fmt.Println(hasil3)
}
