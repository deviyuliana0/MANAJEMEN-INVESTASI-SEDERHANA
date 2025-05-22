package operations

import (
	"fmt"
	"proj_invest/models"
)

func updateLaporanInvestasi() {
	if len(models.Portofolio) == 0 {
		fmt.Println("Data Investasi kosong")
		return
	}

	fmt.Println("=== Laporan Portofolio Investasi ===")
	var totalReturn float64 = 0

	for _, invest := range models.Portofolio {
		// Hitung return investasi
		returnInvest := hitungReturn(invest)
		totalReturn += returnInvest

		status := "Tidak Mencapai Target, Jangan Menyerah"
		if returnInvest > 0 {
			status = "Mencapai Target, Selamat !"
		} else if returnInvest == 0 {
			status = "Yah, Rugi bang!"
		}

		fmt.Printf("%s(%s) | Beli : %.2f | Sekarang : %.2f | Umit : %d | Return : %2.f (%s)\n", invest.NamaAset, invest.JenisAset, invest.HargaBeli, invest.HargaSekarang, invest.JumlahUnit, returnInvest, status) // Tampilkan data investasi
	}

	models.updateLaporanInvestasi()
	fmt.Printf("💸 Total Return Return Anda : %.2f\n", totalReturn)
}
