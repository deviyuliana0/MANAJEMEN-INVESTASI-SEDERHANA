package statistic

import (
	"fmt"
	"invest_proj/models"
)

func TampilkanStatistikInvestasi() {
	if models.JumlahData == 0 {
		fmt.Println("Belum ada data investasi untuk dihitung statistiknya.")
		return
	}

	var totalUnit int
	var totalBeli float64
	var totalSekarang float64
	var totalReturn float64

	for i := 0; i < models.JumlahData; i++ {
		invest := models.Portofolio[i]
		totalUnit += invest.JumlahUnit
		totalBeli += invest.HargaBeli * float64(invest.JumlahUnit)
		totalSekarang += invest.HargaSekarang * float64(invest.JumlahUnit)
		totalReturn += (invest.HargaSekarang - invest.HargaBeli) * float64(invest.JumlahUnit)
	}

	avgReturn := totalReturn / float64(models.JumlahData)

	// Tabel ringkasan
	fmt.Println("===============================================")
	fmt.Println("|       Statistik Portofolio Investasi        |")
	fmt.Println("===============================================")
	fmt.Printf("| Total Aset               | %17d |\n", models.JumlahData)
	fmt.Printf("| Total Unit Keseluruhan   | %17d |\n", totalUnit)
	fmt.Printf("| Total Nilai Beli         | Rp %14.2f |\n", totalBeli)
	fmt.Printf("| Total Nilai Sekarang     | Rp %14.2f |\n", totalSekarang)
	fmt.Printf("| Rata-rata Return/Aset    | Rp %14.2f |\n", avgReturn)
	fmt.Println("===============================================")
}
