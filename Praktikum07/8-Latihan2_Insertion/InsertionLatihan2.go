// Sinta Sarwo - 2311102132

package main

import (
	"fmt"
)

const nMax = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku [nMax]Buku

// Procedure DaftarkanBuku
func DaftarkanBuku(pustaka *DaftarBuku, n *int) {
	fmt.Println("Masukkan data buku (id, judul, penulis, penerbit, eksemplar, tahun, rating):")
	for i := 0; i < *n; i++ {
		fmt.Printf("Buku %d:\n", i+1)
		fmt.Scan(&pustaka[i].id, &pustaka[i].judul, &pustaka[i].penulis, &pustaka[i].penerbit,
			&pustaka[i].eksemplar, &pustaka[i].tahun, &pustaka[i].rating)
	}
}

// Procedure CetakTerfavorit
func CetakTerfavorit(pustaka DaftarBuku, n int) {
	if n == 0 {
		fmt.Println("Tidak ada buku di pustaka.")
		return
	}
	maxRating := pustaka[0].rating
	for i := 1; i < n; i++ {
		if pustaka[i].rating > maxRating {
			maxRating = pustaka[i].rating
		}
	}
	fmt.Println("Buku terfavorit: ")
	for i := 0; i < n; i++ {
		if pustaka[i].rating == maxRating {
			fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Rating: %d\n", pustaka[i].judul, pustaka[i].penulis, pustaka[i].penerbit, pustaka[i].tahun, pustaka[i].rating)
		}
	}
}

// Procedure UrutBuku
func UrutBuku(pustaka *DaftarBuku, n int) {
	for i := 1; i < n; i++ {
		key := (*pustaka)[i]
		j := i - 1
		for j >= 0 && (*pustaka)[j].rating < key.rating {
			(*pustaka)[j+1] = (*pustaka)[j]
			j--
		}
		(*pustaka)[j+1] = key
	}
}

// Procedure Cetak5Terbaru
func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	fmt.Println("5 Judul Buku dengan Rating Tertinggi:")
	for i := 0; i < n && i < 5; i++ {
		fmt.Println(pustaka[i].judul)
	}
}

// Procedure CariBuku (menggunakan pencarian biner)
func CariBuku(pustaka DaftarBuku, rating int) {
	found := false
	fmt.Printf("Buku dengan rating %d:\n", rating)
	for _, buku := range pustaka {
		if buku.rating == rating {
			fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Eksemplar: %d, Tahun: %d, Rating: %d\n",
				buku.judul, buku.penulis, buku.penerbit, buku.eksemplar, buku.tahun, buku.rating)
			found = true
		}
	}
	if !found {
		fmt.Println("Tidak ada buku dengan rating tersebut.")
	}
}

func main() {
	var pustaka DaftarBuku
	var n int
	var rating int

	// Masukkan jumlah buku
	fmt.Print("Masukkan jumlah buku: ")
	fmt.Scan(&n)

	// Daftarkan buku
	DaftarkanBuku(&pustaka, &n)

	// Cetak buku terfavorit
	fmt.Println()
	CetakTerfavorit(pustaka, n)

	// Urutkan buku
	UrutBuku(&pustaka, n)

	// Cetak 5 buku dengan rating tertinggi
	fmt.Println()
	Cetak5Terbaru(pustaka, n)

	// Cari buku dengan rating tertentu
	fmt.Println()
	fmt.Print("Masukkan rating buku yang ingin dicari: ")
	fmt.Scan(&rating)
	CariBuku(pustaka, rating)
}
