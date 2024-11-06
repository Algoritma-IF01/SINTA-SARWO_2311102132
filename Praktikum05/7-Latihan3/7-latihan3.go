// Sinta Sarwo - 2311102132

package main

import (
	"fmt"
)

// Struct untuk menyimpan variable
type MatchResult struct {
	MatchNumber int
	Winner      string
	IsDraw      bool
}

func main() {
	var clubA, clubB string
	var scoreA, scoreB int
	var results []MatchResult

	// Input nama club 
	fmt.Print("Nama Klub A: ")
	fmt.Scanln(&clubA)
	fmt.Print("Nama Klub B: ")
	fmt.Scanln(&clubB)

	// Proses menentukan hasil pertandingan dan meyimpan
	matchCount := 1
	for {
		fmt.Printf("Pertandingan %d: ", matchCount)
		fmt.Scan(&scoreA)
		fmt.Scan(&scoreB)

		// Program behenti jika score kurang dari 0
		if scoreA < 0 || scoreB < 0 {
			fmt.Println("Pertandingan selesai")
			break
		}

		// Penentuan Pemenang dari pertandingan
		if scoreA > scoreB {
			results = append(results, MatchResult{MatchNumber: matchCount, Winner: clubA, IsDraw: false})
		} else if scoreB > scoreA {
			results = append(results, MatchResult{MatchNumber: matchCount, Winner: clubB, IsDraw: false})
		} else {
			results = append(results, MatchResult{MatchNumber: matchCount, Winner: "Draw", IsDraw: true})
		}

		matchCount++
	}

	// Menampilkan hasil dari semua pertandingan setelah selesai
	for _, result := range results {
		if result.IsDraw {
			fmt.Printf("Hasil %d : Draw\n", result.MatchNumber)
		} else {
			fmt.Printf("Hasil %d : %s \n", result.MatchNumber, result.Winner)
		}
	}
	fmt.Print("Pertandingan Selesai.")

}
