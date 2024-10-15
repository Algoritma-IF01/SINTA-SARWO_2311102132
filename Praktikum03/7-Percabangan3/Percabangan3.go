// Sinta Sarwo - 2311102132

package main

import (
	"fmt"
)

func cariFaktor(n int) []int {
	var faktor []int
	for i:= 1; i <= n; i++{
		if n%i == 0 {
			faktor = append(faktor, i)
		}
	}
	return faktor
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0{
			return false
		}
	}
	return true
}

func main(){
	var b int

	fmt.Print("Bilangan: ")
	fmt.Scan(&b)

	if b <= 0 {
		fmt.Println("Bilangan tidak boleh 0 atau negatif.")
	}

	faktor := cariFaktor(b)
	fmt.Print("Faktor: ")
	for _, f := range faktor {
		fmt.Printf("%d ", f)
	}
	fmt.Println()

	prima := isPrime(b)
	fmt.Println("Prima: ",prima )
}