// Sinta Sarwo - 2311102132

package main

import (
	"fmt"
)

func cetakDeret(n int) {
	fmt.Print(n)

	for n != 1 {
		if n%2 == 0 {
			n = n/2
		}else {
			n = 3*n + 1
		}
		fmt.Print(" ", n)
	}

	fmt.Println()
}

func main() {
	const Maxn = 1000000
	var n int

	fmt.Print("Masukan nilai suku awal: ")
	fmt.Scan(&n)

	if n < Maxn {
		cetakDeret(n)
	} else {
		fmt.Println("Nilai suku awal lebih besar dari 1000000.")
	}
}