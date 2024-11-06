// Sinta Sarwo - 2311102132

package main

import (
	"fmt"
	"math"
)

func indeksGanjil(arr []int) {
	fmt.Print("Elemen dengan nilai ganjil: ")
	for _, hasil := range arr {
		if hasil%2 != 0 {
			fmt.Print(hasil, " ")
		}
	}
	fmt.Println()
}

func indeksGenap(arr []int) {
	fmt.Print("Elemen dengan nilai genap: ")
	for _, hasil := range arr {
		if hasil%2 == 0 && hasil != 0 {
			fmt.Print(hasil, " ")
		}
	}
	fmt.Println()
}

func Kelipatan(arr []int, x int) {
	fmt.Printf("Elemen pada indeks kelipatan %d: ", x)
	for i := x - 1; i < len(arr); i += x {
		if arr[i] != 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()
}

func Delete(arr []int, delx int) []int {
	if delx >= 0 && delx < len(arr) {
		return append(arr[:delx], arr[delx+1:]...)
	}
	fmt.Println("Indeks tidak valid untuk penghapusan.")
	return arr
}

func mean(arr []int) int {
	sum := 0
	for _, hasil := range arr {
		sum += hasil
	}
	return sum / len(arr)
}

func StandarDeviasi(arr []int, rataRata float64) float64 {
	var variasiSum float64
	for _, hasil := range arr {
		variasiSum += math.Pow(float64(hasil)-rataRata, 2)
	}
	return math.Sqrt(variasiSum / float64(len(arr)))
}

func Frekuensi(arr []int, target int) int {
	frekuensi := 0
	for _, hasil := range arr {
		if hasil == target {
			frekuensi++
		}
	}
	return frekuensi
}

func main() {
	var N, x, delx, target, pilihan int

	fmt.Print("Input jumlah elemen array: ")
	fmt.Scan(&N)

	Array := make([]int, N)
	fmt.Println("Masukkan elemen array: ")
	for i := 0; i < N; i++ {
		fmt.Scan(&Array[i])
	}
	fmt.Println()

	for {
		fmt.Println("Menu:")
		fmt.Println("1. Menampilkan isi Array.")
		fmt.Println("2. Menampilkan elemen dengan nilai ganjil.")
		fmt.Println("3. Menampilkan elemen dengan nilai genap.")
		fmt.Println("4. Menampilkan elemen pada kelipatan x.")
		fmt.Println("5. Delete elemen dalam array.")
		fmt.Println("6. Menampilkan rata-rata dari array.")
		fmt.Println("7. Menampilkan standar deviasi dari array.")
		fmt.Println("8. Menampilkan frekuensi dari suatu bilangan.")
		fmt.Println("9. Exit.")
		fmt.Print("\nInput pilihan: ")
		fmt.Scan(&pilihan)
		fmt.Println()

		switch pilihan {
		case 1:
			fmt.Println("Isi array:", Array)
			fmt.Println()
		case 2:
			indeksGanjil(Array)
			fmt.Println()
		case 3:
			indeksGenap(Array)
			fmt.Println()
		case 4:
			fmt.Print("Input kelipatan: ")
			fmt.Scan(&x)
			Kelipatan(Array, x)
		case 5:
			fmt.Print("Input indeks elemen yang ingin dihapus: ")
			fmt.Scan(&delx)
			Array = Delete(Array, delx)
			fmt.Println("Array setelah dihapus:", Array)
			fmt.Println()
		case 6:
			fmt.Println("Rata-rata array:", mean(Array))
			fmt.Println()
		case 7:
			fmt.Println("Standar deviasi array:", StandarDeviasi(Array, float64(mean(Array))))
			fmt.Println()
		case 8:
			fmt.Print("Input bilangan target: ")
			fmt.Scan(&target)
			fmt.Printf("Frekuensi %d dalam array: %d\n", target, Frekuensi(Array, target))
			fmt.Println()
		case 9:
			fmt.Println("End Program ;)")
			return
		default:
			fmt.Println("Command Invalid!")
		}
	}
}
