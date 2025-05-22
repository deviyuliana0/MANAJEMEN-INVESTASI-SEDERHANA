package operations

import (
	"fmt"
	"invest_proj/models"
)

func TambahInvestasi() { // fungsi untuk menambahkan investasi ke portofolio
	var invest models.Investasi // variabel untuk menampung data investasi yang akan ditambahkan

	fmt.Println("=== Tambah Data Investasi ===")
	fmt.Print("Nama Aset	:")   // menampilkan input nama aset
	fmt.Scan(&invest.NamaAset) // menginput nama aset

	fmt.Print("Jenis Aset	:")   // jenis aset yang diinvestasikan
	fmt.Scan(&invest.JenisAset) // jenis aset yang diinvestasikan

	fmt.Print("Harga Beli	:")   // harga beli aset per unit
	fmt.Scan(&invest.HargaBeli) // input harga beli aset per unit

	fmt.Print("Harga Sekarang	:")   // harga aset per unit saat ini
	fmt.Scan(&invest.HargaSekarang) // input harga aset saat ini

	fmt.Print("Jumlah Unit	:")   // jumlah unit aset yang diinvestasikan
	fmt.Scan(&invest.JumlahUnit) // input jumlah unit aset yang diinvestasikan

	models.Portofolio = append(models.Portofolio, invest) // menambahkan data investasi ke array portofolio
	fmt.Print("Data berhasil diinputkan!\n\n")            // menampilkan pesan berhasil input data
}
