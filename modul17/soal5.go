package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var n int
	var toppingPizza int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {

		x := rand.Float64()
		y := rand.Float64()

		// cekkk ap titik berada di dalam pizza
		if (x-0.5)*(x-0.5)+(y-0.5)*(y-0.5) <= 0.25 {

			toppingPizza++
		}
	}

	pi := 4.0 * float64(toppingPizza) / float64(n)

	fmt.Println("Topping pada Pizza:", toppingPizza)
	fmt.Printf("PI : %.10f\n", pi)
}
