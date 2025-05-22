package operations

import (
	"fmt"
	"invest_proj/models"
)

func hitungKeuntungan() {
	if len(models.Portofolio) == 0 { // cek apakah portofolio kosong
		fmt.Println("Belum ada investasi.")
		return
	}

	fmt.Println("=== Hitung Keuntungan / Kerugian ===")
	for _, invest := range models.Portofolio {
		selisih := invest.HargaSekarang - invest.HargaBeli      // menghitung selisih harga
		totalKeuntungan := selisih * float64(invest.JumlahUnit) // menghitung total keuntungan

		if selisih > 0 {
			fmt.Printf("Yeay, Selamat Kamu Mendapatkan Keuntungan dari %s sebesar : Rp %.2f\n", invest.NamaAset, totalKeuntungan)
		} else if selisih < 0 {
			fmt.Printf("Yah, Jangan Bersedih Terdapat Kerugian dari %s : Rp %.2f\n", invest.NamaAset, -totalKeuntungan) // menampilkan pesan kerugian
		} else {
			fmt.Printf("Tidak ada keuntungan atau kerugian dari %s\n", invest.NamaAset) // menampilkan pesan tidak ada keuntungan atau kerugian
		}
		models.hitungKeuntungan()
	}
}
