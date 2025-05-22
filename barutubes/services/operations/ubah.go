package operations

import (
	"fmt"
	"invest_proj/models"
)

func ubahInvestasi() {
	if len(models.Portofolio) == 0 { // cek apakah portofolio kosong
		fmt.Println("Belum ada data untuk diubah")
		return
	}

	var index int
	fmt.Print("Pilih nomor data Investasi yang ingin diubah:")
	fmt.Scan(&index)

	if index < 1 || index > len(models.Portofolio) { // cek apakah index valid
		fmt.Println("Index tidak valid")
		return
	}

	i := index - 1 // mengubah index menjadi indeks array
	fmt.Println("Input data baru:")

	fmt.Print("Nama Aset	:")
	fmt.Scan(&models.Portofolio[i].NamaAset) // menginput nama aset

	fmt.Print("Jenis Aset	:")
	fmt.Scan(&models.Portofolio[i].JenisAset) // menginput jenis aset

	fmt.Print("Harga Beli	:")                 // menampilkan pesan harga beli
	fmt.Scan(&models.Portofolio[i].HargaBeli) // menginput harga beli

	fmt.Print("Harga Sekarang	:")                 // menampilkan pesan harga sekarang
	fmt.Scan(&models.Portofolio[i].HargaSekarang) // menginput harga sekarang

	fmt.Print("Jumlah Unit	:")                 // menampilkan pesan jumlah unit
	fmt.Scan(&models.Portofolio[i].JumlahUnit) // menginput jumlah unit

	models.ubahInvestasi()
	fmt.Print("Data berhasil diubah!\n\n")
}
