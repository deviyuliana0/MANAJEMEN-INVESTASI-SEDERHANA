package display

import (
	"fmt"
	"invest_proj/models"
)

func TampilkanInvestasi() {
	fmt.Println("=== Daftar Investasi ===")
	for i := 0; i < models.JumlahData; i++ {
		fmt.Printf("[%d] %s (%s), Beli: %.2f, Sekarang: %.2f, Unit: %d\n",
			i+1,
			models.Portofolio[i].NamaAset,
			models.Portofolio[i].JenisAset,
			models.Portofolio[i].HargaBeli,
			models.Portofolio[i].HargaSekarang,
			models.Portofolio[i].JumlahUnit)
	}
}
