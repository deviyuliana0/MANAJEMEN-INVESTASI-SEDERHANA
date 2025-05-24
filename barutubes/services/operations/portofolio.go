package operations

import (
	"fmt"
	"invest_proj/models"
)

// Fungsi bantu untuk hitung return
func hitungReturn(invest models.Investasi) float64 {
	return (invest.HargaSekarang - invest.HargaBeli) * float64(invest.JumlahUnit)
}

func UpdateLaporanInvestasi() {
	if models.JumlahData == 0 {
		fmt.Println("Data Investasi kosong")
		return
	}

	fmt.Println("=== Laporan Portofolio Investasi ===")
	var totalReturn float64 = 0

	for i := 0; i < models.JumlahData; i++ {
		invest := models.Portofolio[i]
		returnInvest := hitungReturn(invest)
		totalReturn += returnInvest

		status := "Tidak Mencapai Target, Jangan Menyerah"
		if returnInvest > 0 {
			status = "Mencapai Target, Selamat!"
		} else if returnInvest == 0 {
			status = "Yah, Rugi bang!"
		}

		fmt.Printf("%s (%s) | Beli: %.2f | Sekarang: %.2f | Unit: %d | Return: %.2f (%s)\n",
			invest.NamaAset, invest.JenisAset,
			invest.HargaBeli, invest.HargaSekarang,
			invest.JumlahUnit, returnInvest, status)
	}

	fmt.Printf("💸 Total Return Anda: %.2f\n", totalReturn)
}
