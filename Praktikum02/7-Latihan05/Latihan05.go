//Sinta Sarow - 2311102132

package main

import "fmt"

func main() {
	var a, b, c, d, e int
	var char string

	fmt.Println("Masukkan 5 angka integer antara 32 sampai 127: ")
	fmt.Scan(&a, &b, &c, &d, &e)
	
	fmt.Println("Masukkan 3 karakter:")
	fmt.Scan(&char)

	fmt.Println("")
	fmt.Println("Output:")
	
	fmt.Printf("%c%c%c%c%c\n", a, b, c, d, e)
	
	if len(char) == 3 {
		fmt.Printf("%c%c%c\n", char[0]+1, char[1]+1, char[2]+1)
	} else {
		fmt.Println("Error.")
	}
}