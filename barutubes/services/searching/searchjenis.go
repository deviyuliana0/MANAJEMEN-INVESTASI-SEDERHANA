package searching

import (
	"fmt"
	"invest_proj/models"
)

func cariBerdasarkanJenis() {
	if len(models.Portofolio) == 0 {
		fmt.Println("Belum ada data untuk dicari")
		return
	}

	var jenis string
	fmt.Print("Masukkan jenis aset yang dicari (saham, obligasi, emas, reksana: ") // menampilkan pesan
	fmt.Scan(&jenis)

	ditemukan := false
	fmt.Println("=== Hasil Pencarian ===")
	for _, invest := range models.Portofolio {
		if invest.JenisAset == jenis {
			fmt.Printf("%s (%s), Beli : %.2f, Sekarang : %.2f, Unit : %d\n", invest.NamaAset, invest.JenisAset, invest.HargaBeli, invest.HargaSekarang, invest.JumlahUnit)
			ditemukan = true
		}
	}

	if !ditemukan {
		models.cariBerdasarkanNama()
		fmt.Println("Data Investasi Tidak ditemukan dengan jenis tersebut.")
	}
}
