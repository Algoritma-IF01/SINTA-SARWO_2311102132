// Sinta Sarwo - 2311102132

package main

import (
	"fmt"
	"math"
)

type titik struct {
	x, y int
}

type lingkaran struct {
	pusat titik
	radius float64
}

func jarak(a, b titik) float64 {
	PenguranganX := math.Pow(float64(a.x-b.x), 2)
	PenguranganY := math.Pow(float64(a.y-b.y), 2)
	return math.Sqrt(PenguranganX + PenguranganY)
}

func didalam (c lingkaran, a titik ) bool {
	return jarak(a, c.pusat) < (c.radius)
}

func main() {
	var lingkaran1 lingkaran
	var lingkaran2 lingkaran
	var point titik

	fmt.Scan(&lingkaran1.pusat.x, &lingkaran1.pusat.y, &lingkaran1.radius)
	fmt.Scan(&lingkaran2.pusat.x, &lingkaran2.pusat.y, &lingkaran2.radius)
	fmt.Scan(&point.x, &point.y)

	didalam1 := didalam(lingkaran1, point)
	didalam2 := didalam(lingkaran2, point)

	if didalam1 && didalam2{
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	}else if didalam1 {
		fmt.Println("Titik di dalam lingkaran 1")
	}else if didalam2 {
		fmt.Println("Titik di dalam lingkaran 2")
	}else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}