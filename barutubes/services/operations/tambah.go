package operations

import (
	"fmt"
	"invest_proj/models"
)

func TambahInvestasi() {
	var invest models.Investasi

	fmt.Println("=== Tambah Data Investasi ===")
	fmt.Print("Nama Aset       : ")
	fmt.Scan(&invest.NamaAset)

	fmt.Print("Jenis Aset      : ")
	fmt.Scan(&invest.JenisAset)

	fmt.Print("Harga Beli      : ")
	fmt.Scan(&invest.HargaBeli)

	fmt.Print("Harga Sekarang  : ")
	fmt.Scan(&invest.HargaSekarang)

	fmt.Print("Jumlah Unit     : ")
	fmt.Scan(&invest.JumlahUnit)

	if models.JumlahData < models.MaksData {
		models.Portofolio[models.JumlahData] = invest
		models.JumlahData++
		fmt.Println("Data berhasil diinputkan!\n")
	} else {
		fmt.Println("Data gagal ditambahkan: kapasitas maksimum tercapai.")
	}
}
