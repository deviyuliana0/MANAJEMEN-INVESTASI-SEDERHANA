package operations

import (
	"fmt"
	"proj_invest/models"
)

func riwayatTransaksi () {
	func riwayatTransaksi() {
	if len(models.DaftarTransaksi) == 0 {
		fmt.Println("Belum ada riwayat transaksi")
		return
	}

	fmt.Println("=== Riwayat Transaksi Investasi ===")
	for i := 0; i < len(models.DaftarTransaksi); i++ {
		models.riwayatTransaksi()
		fmt.Printf("%d. %s data: %s (%d)\n", i+1, models.DaftarTransaksi[i].NamaAset, models.DaftarTransaksi[i].JenisAset, models.DaftarTransaksi[i].JumlahUnit)
	}
}
}