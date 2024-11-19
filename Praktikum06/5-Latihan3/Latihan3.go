// Sinta Sarwo - 2311102132

package main

import "fmt"

type arrBalita [100]float64

// Subprogram untuk menghitung berat minimum dan maksimum balita
func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	// Iterasi untuk menemukan nilai minimum dan maksimum
	for i := 1; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

// Fungsi untuk menghitung rata-rata berat balita
func rerata(arrBerat arrBalita, n int) float64 {
	total := 0.0
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}
	return total / float64(n)
}

func main() {
	var n int
	var beratBalita arrBalita
	var bMin, bMax float64

	fmt.Print("Masukkan jumlah data berat balita (maks 100): ")
	fmt.Scan(&n)

	// Validasi jumlah balita
	if n <= 0 || n > 100 {
		fmt.Println("Jumlah balita harus antara 1 hingga 100.")
		return
	}

	// Input berat balita
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan Berat balita ke-%d: ", i+1)
		fmt.Scan(&beratBalita[i])
		if beratBalita[i] < 0 {
			fmt.Println("Berat balita tidak boleh negatif.")
			return
		}
	}

	hitungMinMax(beratBalita, n, &bMin, &bMax)
	rataRata := rerata(beratBalita, n)

	// Menampilkan hasil
	fmt.Printf("\nBerat balita minimum: %.2f kg\n", bMin)
	fmt.Printf("Berat balita makasimum: %.2f kg\n", bMax)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rataRata)
}
