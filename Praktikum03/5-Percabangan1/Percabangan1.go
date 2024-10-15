// Sinta Sarwo - 2311102132

package main

import (
	"fmt"
)

func main(){
	var beratParsel, beratKg, sisaGram int
	var biaya, biayaTambahan, total_biaya int

	const biayaperKg = 10000
	const biayaLebih = 5
	const biayaKurang = 15
	const batasBeratGratis = 10
	
	fmt.Print("Berat pasel (gram): ")
	fmt.Scan(&beratParsel)

	beratKg = beratParsel / 1000
	sisaGram = beratParsel % 1000

	biaya = beratKg * biayaperKg

	if sisaGram >= 500 {
		biayaTambahan = sisaGram * biayaLebih
	} else {
		biayaTambahan = sisaGram * biayaKurang
	}

	total_biaya = biaya + biayaTambahan
	
	if beratKg > 10 {
		total_biaya = biaya
		fmt.Printf("Detail berat: %d kg + %d gr\n", beratKg, sisaGram)
		fmt.Printf("Detail biaya: Rp.%d + %d\n", biaya, biayaTambahan)
		fmt.Printf("Total biaya: Rp. %d\n", total_biaya)	
	} else {
		fmt.Printf("Detail berat: %d kg + %d gr\n", beratKg, sisaGram)
		fmt.Printf("Detail biaya: Rp.%d + %d\n", biaya, biayaTambahan)
		fmt.Printf("Total biaya: Rp. %d\n", total_biaya)
	}
}