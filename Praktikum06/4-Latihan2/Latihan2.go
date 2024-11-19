// Sinta Sarwo - 2311102132

package main

import (
	"fmt"
)

// Fungsi untuk menghitung rata-rata berat wadah
func calculateAverage(beratWadah []float64, jumlahWadah int) float64 {
	total := 0.0
	for i := 0; i < jumlahWadah; i++ {
		total += beratWadah[i]
	}
	return total / float64(jumlahWadah)
}

func main() {
	var x, y int

	// Input jumlah ikan dan kapasitas wadah
	fmt.Print("Input jumlah ikan yang akan dijual (x) : ")
	fmt.Scan(&x)
	fmt.Print("Input jumlah ikan yang akan dimasukkan dalam wadah (y) : ")
	fmt.Scan(&y)

	// Validasi input
	if x <= 0 || x > 1000 || y <= 0 {
		fmt.Println("Jumlah ikan harus antara 1-1000, dan kapasitas wadah harus lebih dari 0!")
		return
	}

	// Membuat array kapasitas tetap untuk berat ikan
	var berat [1000]float64

	fmt.Printf("Masukkan berat %d ikan : ", x)
	for i := 0; i < x; i++ {
		_, err := fmt.Scan(&berat[i])
		if err != nil || berat[i] < 0 {
			fmt.Println("Input berat tidak valid! Masukkan bilangan positif.")
			return
		}
	}
	fmt.Println()

	// Menghitung jumlah wadah yang dibutuhkan
	jumlahWadah := (x + y - 1) / y // Pembulatan ke atas
	var beratWadah [1000]float64   // Array untuk berat total di setiap wadah

	// Distribusi berat ikan ke wadah
	for i := 0; i < x; i++ {
		wadahIndex := i / y
		beratWadah[wadahIndex] += berat[i]
	}

	// Menghitung rata-rata berat ikan di semua wadah
	averageBerat := calculateAverage(beratWadah[:], jumlahWadah)

	// Menampilkan total berat ikan di setiap wadah
	fmt.Println("Total berat ikan di setiap wadah :")
	for i := 0; i < jumlahWadah; i++ {
		fmt.Printf("Wadah %d: %.2f kg\n", i+1, beratWadah[i])
	}
	fmt.Println()

	// Menampilkan rata-rata berat ikan di semua wadah
	fmt.Print("Berat rata-rata ikan di setiap wadah : ")
	fmt.Printf("%.2f kg", averageBerat)
}
