//Setiap program dimulai dengan "package main"
package main

//Impor paket yang dibutuhkan, "fmt" berisi proses I/O standar
import (
	"fmt"
)

//Kode program utama dalam "fungsi main"
func main() {
	fmt.Println("Hello, World!")
	var greetings string = "Selamat datang di dunia Go!"
	var a, b int

	fmt.Println(greetings)
	fmt.Scanln(&a, &b)
	fmt.Printf("%v + %v = %v\n", a, b, a+b)
}