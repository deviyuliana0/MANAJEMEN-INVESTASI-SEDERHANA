package searching

import (
	"fmt"
	"invest_proj/models"
)

func cariBerdasarkanNama() {
	if len(models.Portofolio) == 0 {
		fmt.Println("Belum ada data untuk dicari")
		return
	}

	var nama string
	fmt.Print("Masukkan nama aset yang dicari : ")
	fmt.Scan(&nama)

	ditemukan := false
	for _, invest := range models.Portofolio {
		if invest.NamaAset == nama {
			fmt.Println("Ditemukan:")
			fmt.Printf("%s (%s), Beli : %.2f, Sekarang : %.2f, Unit : %d\n", invest.NamaAset, invest.JenisAset, invest.HargaBeli, invest.HargaSekarang, invest.JumlahUnit) // menampilkan data aset yang ditemukan
			ditemukan = true
			break
		}
	}

	if !ditemukan {
		models.cariBerdasarkanNama()
		fmt.Println("Data Tidak ditemukan.")
	}
}
