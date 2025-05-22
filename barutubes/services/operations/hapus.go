package operations

import (
	"fmt"
	"invest_proj/models"
)

func hapusInvestasi() {
	if len(models.Portofolio) == 0 { // cek apakah portofolio kosong
		fmt.Println("Belum ada data untuk dihapus")
		return
	}

	var index int                                                // variabel untuk menampung index investasi yang ingin dihapus
	fmt.Print("Pilih nomor data Investasi yang ingin dihapus: ") // menampilkan pesan
	fmt.Scan(&index)

	if index < 1 || index > len(models.Portofolio) { // cek apakah index valid
		fmt.Println("Index tidak valid")
		return
	}

	i := index - 1                                                                // mengubah index menjadi indeks array
	models.Portofolio = append(models.Portofolio[:i], models.Portofolio[i+1:]...) // menghapus data investasi
	models.hapusInvest()
	fmt.Print("Data berhasil dihapus!\n\n")
}
