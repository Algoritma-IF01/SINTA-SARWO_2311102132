// Sinta Sarwo - 2311102132

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Fungsi untuk melakukan selection sort pada slice secara menaik
func selectionSortAsc(arr []int) {
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

// Fungsi untuk melakukan selection sort pada slice secara menurun
func selectionSortDesc(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] > arr[maxIdx] {
				maxIdx = j
			}
		}
		// Tukar elemen
		arr[i], arr[maxIdx] = arr[maxIdx], arr[i]
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin) // Membuat pembaca untuk membaca input dari pengguna

	// Baca jumlah daerah
	nInput, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(nInput)) // Mengubah input string menjadi integer

	// Variabel untuk menyimpan hasil akhir untuk semua daerah
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

		// Pisahkan nomor rumah ganjil dan genap
		var ganjil, genap []int
		for _, nomor := range rumah {
			if nomor%2 == 0 {
				genap = append(genap, nomor)
			} else {
				ganjil = append(ganjil, nomor)
			}
		}

		// Urutkan nomor rumah ganjil secara menaik dan genap secara menurun
		selectionSortAsc(ganjil)
		selectionSortDesc(genap)

		// Gabungkan nomor rumah ganjil dan genap
		result := append(ganjil, genap...)

		// Simpan hasilnya untuk dicetak nanti
		var resultStr []string
		for j, nomor := range result {
			if j > 0 {
				resultStr = append(resultStr, fmt.Sprintf(" %d", nomor))
			} else {
				resultStr = append(resultStr, fmt.Sprintf("%d", nomor))
			}
		}

		// Gabungkan hasil per daerah menjadi satu string dan simpan
		hasil = append(hasil, strings.Join(resultStr, ""))
	}

	fmt.Println()
	fmt.Println("Hasil nomor daerah: ")
	// Cetak hasil setelah semua input diproses
	for _, result := range hasil {
		fmt.Println(result)
	}
}
