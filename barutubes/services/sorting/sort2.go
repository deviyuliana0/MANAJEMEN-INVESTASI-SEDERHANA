package sorting

import (
	"fmt"
	"invest_proj/models"
)

// Fungsi bantu untuk menghitung return
func hitungReturn(invest models.Investasi) float64 {
	return (invest.HargaSekarang - invest.HargaBeli) * float64(invest.JumlahUnit)
}

// Insertion sort berdasarkan return tertinggi ke terendah
func UrutkanReturn() {
	for i := 1; i < models.JumlahData; i++ {
		key := models.Portofolio[i]
		j := i - 1

		for j >= 0 && hitungReturn(models.Portofolio[j]) < hitungReturn(key) {
			models.Portofolio[j+1] = models.Portofolio[j]
			j--
		}
		models.Portofolio[j+1] = key
	}
	fmt.Println("Data berhasil diurutkan berdasarkan return investasi (dari tertinggi ke terendah).")
}
