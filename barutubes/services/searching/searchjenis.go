package searching

import (
	"fmt"
	"invest_proj/models"
)

func CariBerdasarkanJenis() {
	if models.JumlahData == 0 {
		fmt.Println("Belum ada data untuk dicari.")
		return
	}

	var jenis string
	fmt.Print("Masukkan jenis aset yang dicari (saham, obligasi, emas, reksadana): ")
	fmt.Scan(&jenis)

	ditemukan := false
	fmt.Println("=== Hasil Pencarian ===")
	for i := 0; i < models.JumlahData; i++ {
		invest := models.Portofolio[i]
		if invest.JenisAset == jenis {
			fmt.Printf("%s (%s), Beli: %.2f, Sekarang: %.2f, Unit: %d\n",
				invest.NamaAset, invest.JenisAset, invest.HargaBeli, invest.HargaSekarang, invest.JumlahUnit)
			ditemukan = true
		}
	}

	if !ditemukan {
		fmt.Println("Data investasi tidak ditemukan dengan jenis tersebut.")
	}
}
