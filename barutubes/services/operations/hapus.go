package operations

import (
	"fmt"
	"invest_proj/models"
)

func HapusInvestasi() {
	if models.JumlahData == 0 {
		fmt.Println("Belum ada data untuk dihapus.")
		return
	}

	var index int
	fmt.Printf("Pilih nomor data Investasi yang ingin dihapus (1 - %d): ", models.JumlahData)
	fmt.Scan(&index)

	if index < 1 || index > models.JumlahData {
		fmt.Println("Index tidak valid.")
		return
	}

	i := index - 1

	// Geser elemen setelah i ke kiri
	for j := i; j < models.JumlahData-1; j++ {
		models.Portofolio[j] = models.Portofolio[j+1]
	}

	models.JumlahData-- // kurangi jumlah data aktif
	fmt.Println("Data berhasil dihapus!\n")
}
