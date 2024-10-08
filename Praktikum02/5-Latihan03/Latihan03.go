//Sinta Sarow - 2311102132

package main

import (
	"fmt"
)

func main(){
	var r, Volume_Bola, Luas_Bola float64
	var pi float64 = 3.14
	
	fmt.Print("Masukkan input nilai jari-jari: ")
	fmt.Scanln(&r)

	Luas_Bola = 4 * pi * r * r
	Volume_Bola = (4/3) * pi * r * r * r
	
	fmt.Println("Bola dengan nilai jari-jari", r, "memiliki volume", Volume_Bola, "dan luas kulit", Luas_Bola)
}