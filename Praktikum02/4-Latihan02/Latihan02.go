//Sinta Sarow - 2311102132

package main

import (
	"fmt"
)

func main(){
	var tahun int
	fmt.Print("Masukkan input tahun: ")
	fmt.Scanln(&tahun)

	if (tahun%400 == 0) || (tahun%4 == 0 && tahun%100 != 0){
		fmt.Println("Kabisat: True")
	}else {
		fmt.Println("Kabisat: False")
	}
}