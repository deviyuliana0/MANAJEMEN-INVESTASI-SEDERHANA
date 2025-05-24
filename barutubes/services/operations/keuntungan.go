package operations

import (
	"fmt"
	"invest_proj/models"
)

func HitungKeuntungan() {
	if models.JumlahData == 0 {
		fmt.Println("Belum ada investasi.")
		return
	}

	fmt.Println("=== Hitung Keuntungan / Kerugian ===")
	for i := 0; i < models.JumlahData; i++ {
		invest := models.Portofolio[i]
		selisih := invest.HargaSekarang - invest.HargaBeli
		totalKeuntungan := selisih * float64(invest.JumlahUnit)

		if selisih > 0 {
			fmt.Printf("Keuntungan dari %s sebesar: Rp %.2f\n", invest.NamaAset, totalKeuntungan)
		} else if selisih < 0 {
			fmt.Printf("Kerugian dari %s sebesar: Rp %.2f\n", invest.NamaAset, -totalKeuntungan)
		} else {
			fmt.Printf("Tidak ada keuntungan atau kerugian dari %s\n", invest.NamaAset)
		}
	}
}
