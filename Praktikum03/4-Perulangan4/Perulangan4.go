// Sinta Sarwo - 2311102132

package main

import (
	"fmt"
	"math"
)

func sqrt2(k int) float64 {
	hasil := 1.0
	var num, deno float64

	for i := 0; i <= k; i++ {
		// Menggunakan math.Pow untuk menghitung kuadrat
		num = math.Pow(4*float64(i)+2, 2)  // (4i + 2)^2
		deno = (4*float64(i) + 1) * (4*float64(i) + 3)  // (4i + 1)(4i + 3)
		hasil *= num / deno
	}

	return hasil
}

func main() {
	var K int

	fmt.Print("Nilai K: ")
	fmt.Scan(&K)

	hasil := sqrt2(K)

	fmt.Printf("Nilai akar 2 = %.10f\n", hasil)
}