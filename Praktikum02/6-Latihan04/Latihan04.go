//Sinta Sarow - 2311102132

package main

import (
	"fmt"
)

func main(){
	var Celsius, Fahrenheit, Reamur, Kelvin float64
	
	fmt.Print("Input Temperature Celsius: ")
	fmt.Scanln(&Celsius)

	Reamur = Celsius * 4 / 5
	Fahrenheit = (Celsius * 9 / 5)+32
	Kelvin = Celsius +  273.15
	
	fmt.Println("Derajat Reamur: ",Reamur)
	fmt.Println("Derajat Fahrenheit: ", Fahrenheit)
	fmt.Println("Derajat Kelvin: ", Kelvin)
}