package operations

import (
	"fmt"
	"invest_proj/models"
)

func UbahInvestasi() {
	if models.JumlahData == 0 {
		fmt.Println("Belum ada data untuk diubah.")
		return
	}

	fmt.Printf("Pilih nomor data Investasi yang ingin diubah (1 - %d): ", models.JumlahData)
	var index int
	fmt.Scan(&index)

	if index < 1 || index > models.JumlahData {
		fmt.Println("Index tidak valid.")
		return
	}

	i := index - 1
	fmt.Println("=== Masukkan Data Baru ===")

	fmt.Print("Nama Aset       : ")
	fmt.Scan(&models.Portofolio[i].NamaAset)

	fmt.Print("Jenis Aset      : ")
	fmt.Scan(&models.Portofolio[i].JenisAset)

	fmt.Print("Harga Beli      : ")
	fmt.Scan(&models.Portofolio[i].HargaBeli)

	fmt.Print("Harga Sekarang  : ")
	fmt.Scan(&models.Portofolio[i].HargaSekarang)

	fmt.Print("Jumlah Unit     : ")
	fmt.Scan(&models.Portofolio[i].JumlahUnit)

	fmt.Println("Data berhasil diubah!\n")
}
