// Sinta Sarwo - 2311102132

package main

import (
	"fmt"
)

// Fungsi untuk melakukan insertion sort pada array
func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		// Geser elemen yang lebih besar dari `key` ke kanan
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		// Letakkan `key` pada posisi yang tepat
		arr[j+1] = key
	}
}

// Fungsi untuk memeriksa apakah array memiliki jarak tetap
func checkEqualSpacing(arr []int) (bool, int) {
	if len(arr) < 2 {
		return true, 0 // Jika hanya 1 elemen atau kosong, dianggap jarak tetap
	}

	// Hitung jarak antara dua elemen pertama
	distance := arr[1] - arr[0]

	// Periksa jarak setiap pasangan elemen
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != distance {
			return false, 0
		}
	}
	return true, distance
}

func main() {
	var numbers []int

	// Membaca input hingga bilangan negatif ditemukan
	fmt.Println("Masukkan bilangan bulat (akhiri dengan bilangan negatif):")
	for {
		var num int
		fmt.Scan(&num)

		if num < 0 {
			break // Akhiri input jika bilangan negatif
		}
		numbers = append(numbers, num) // Simpan hanya bilangan non-negatif
	}

	// Urutkan array menggunakan insertion sort
	insertionSort(numbers)

	// Periksa apakah jarak antar elemen tetap
	isEqualSpacing, distance := checkEqualSpacing(numbers)

	// Cetak array setelah pengurutan
	fmt.Println("Array terurut:", numbers)

	// Cetak status jarak elemen
	if isEqualSpacing {
		fmt.Printf("Data berjarak %d\n", distance)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}
