package searching

import (
	"fmt"
	"invest_proj/models"
)

func CariBerdasarkanNama() {
	if models.JumlahData == 0 {
		fmt.Println("Belum ada data untuk dicari.")
		return
	}

	var nama string
	fmt.Print("Masukkan nama aset yang dicari: ")
	fmt.Scan(&nama)

	ditemukan := false
	for i := 0; i < models.JumlahData; i++ {
		invest := models.Portofolio[i]
		if invest.NamaAset == nama {
			fmt.Println("Ditemukan:")
			fmt.Printf("%s (%s), Beli: %.2f, Sekarang: %.2f, Unit: %d\n",
				invest.NamaAset, invest.JenisAset, invest.HargaBeli, invest.HargaSekarang, invest.JumlahUnit)
			ditemukan = true
			break
		}
	}

	if !ditemukan {
		fmt.Println("Data tidak ditemukan.")
	}
}
