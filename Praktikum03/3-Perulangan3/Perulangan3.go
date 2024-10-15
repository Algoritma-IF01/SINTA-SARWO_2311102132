// Sinta Sarwo - 2311102132

package main

import (
	"fmt"
)

func main(){
	var berat1, berat2 float64
	var statusSepeda bool

	for {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&berat1, &berat2)

		if berat1 < 0 || berat2 < 0 {
			break
		}

		if berat1 + berat2 > 150 {
			break
		}

		//Pemeriksaan apakah motor pak Andi akan oleng
		if berat1 > berat2 {
			if berat1 - berat2 >= 9{
				statusSepeda = true
			} else{
				statusSepeda = false

			}
		}else {
			if berat2-berat1 >= 9 {
				statusSepeda = true
			} else{
				statusSepeda = false

			}
		}
		fmt.Println("Sepeda motor pak Andi akan oleng: ", statusSepeda)
	}

	fmt.Println("Proses selesai.")

}