package main

import "fmt"

func main() {
	// Masukkan nilai tagihan dari pengguna
	var tagihan float64
	fmt.Print("Masukan total nilai tagihan: ")
	fmt.Scan(&tagihan)

	// Hitung tip berdasarkan aturan
	tip := calculateTip(tagihan)

	// Hitung total nilai
	total := tagihan + tip

	// Tampilkan hasil di konsol
	fmt.Printf("Tagihannya %.2f, tipnya %.2f, dan total nilainya %.2f\n", tagihan, tip, total)
}

func calculateTip(tagihan float64) float64 {
	var tip float64

	// Hitung tip berdasarkan aturan
	if tagihan >= 50 && tagihan <= 300 {
		tip = tagihan * 0.15
	} else {
		tip = tagihan * 0.2
	}

	return tip
}
