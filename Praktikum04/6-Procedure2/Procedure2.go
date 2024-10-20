// Sinta Sarwo - 2311102132

package main

import (
	"fmt"
	"strings"
)

// Prosedur hitungSkor untuk menghitung soal yang diselesaikan dan total waktu
func hitungSkor(waktu []int, totalSoal *int, skor *int) {
	*totalSoal = 0
	*skor = 0

	for _, w := range waktu {
		if w < 301 { // Batas waktu 301 menit
			*totalSoal++ // Menambah soal yang diselesaikan
			*skor += w   // Menambah total waktu
		}
	}
}

func main() {
	var pemenang string
	maxSoal := -1
	minWaktu := 99999 // Nilai awal untuk perbandingan waktu

	fmt.Println("Masukan nama peserta:")

	for {
		var nama string
		fmt.Scan(&nama)

		if strings.ToLower(nama) == "selesai" {
			break
		}

		// Membaca waktu penyelesaian untuk 8 soal
		waktu := make([]int, 8)
		for j := 0; j < 8; j++ {
			fmt.Scan(&waktu[j])
		}

		var totalSoal, totalWaktu int
		hitungSkor(waktu, &totalSoal, &totalWaktu)

		// Update pemenang jika peserta ini menyelesaikan lebih banyak soal,
		// atau jika soal sama tetapi waktu lebih cepat
		if totalSoal > maxSoal || (totalSoal == maxSoal && totalWaktu < minWaktu) {
			pemenang = nama
			maxSoal = totalSoal
			minWaktu = totalWaktu
		}
	}

	// Menampilkan hasil pemenang
	if maxSoal != -1 {
		fmt.Printf("%s %d %d\n", pemenang, maxSoal, minWaktu)
	} else {
		fmt.Println("Tidak ada peserta.")
	}
}
