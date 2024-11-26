// Sinta Sarwo - 2311102132

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Fungsi untuk melakukan selection sort pada slice
func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		// Tukar elemen
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin) // Membuat pembaca untuk membaca input dari pengguna

	// Baca jumlah daerah
	nInput, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(nInput)) // Mengubah input string menjadi integer

	// Slice untuk menyimpan hasil setiap daerah
	var hasil []string

	for i := 0; i < n; i++ {
		// Baca data per daerah
		line, _ := reader.ReadString('\n')
		data := strings.Fields(line)

		// Parse jumlah rumah kerabat
		m, _ := strconv.Atoi(data[0])
		if m < 1 {
			continue
		}

		// Parse nomor rumah
		rumah := make([]int, m)
		for j := 0; j < m; j++ {
			rumah[j], _ = strconv.Atoi(data[j+1])
		}

		// Urutkan nomor rumah menggunakan selection sort
		selectionSort(rumah)

		// Gabungkan hasil daerah ini menjadi string
		var daerahHasil []string
		for _, nomor := range rumah {
			daerahHasil = append(daerahHasil, strconv.Itoa(nomor))
		}
		hasil = append(hasil, strings.Join(daerahHasil, " "))
	}

	// Cetak semua hasil setelah semua input selesai
	fmt.Println()
	fmt.Println("Hasil nomor daerah: ")
	for _, daerah := range hasil {
		fmt.Println(daerah)
	}
}
