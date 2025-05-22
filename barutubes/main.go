package main

import (
	"fmt"
	"invest_proj/services"
)

func main() {
	for {
		fmt.Println("=== Menu ===") // menampilkan menu
		fmt.Println("1. Tambah Investasi")
		fmt.Println("2. Tampilkan Portofolio")
		fmt.Println("3. Ubah Investasi")
		fmt.Println("4. Hapus Investasi")
		fmt.Println("5. Hitung Keuntungan / Kerugian")
		fmt.Println("6. Cari Investasi berdasarkan Nama")
		fmt.Println("7. Cari Investasi berdasarkan Jenis Aset")
		fmt.Println("8. Cari Investasi dengan Return Tertinggi")
		fmt.Println("9. Urutkan berdasarkan Nama")            // selection sort
		fmt.Println("10. Urutkan berdasarkan Return")         // insertion sort
		fmt.Println("11. Tampilkan Update Laporan Investasi") // menampilkan portofolio lengkap
		fmt.Println("12. Tampilkan Riwayat Transaksi")
		fmt.Println("13. Keluar") // menampilkan input menu
		fmt.Print("Pilih Menu Aplikasi:")

		var menu int    // variabel untuk menampung input menu aplikasi
		fmt.Scan(&menu) // menginput menu aplikasi

		switch menu { // switch case untuk menampilkan menu aplikasi
		case 1:
			services.TambahInvestasi() // menampilkan menu tambah investasi
		case 2:
			tampilkanInvestasi() // menampilkan menu tampilkan portofolio
		case 3:
			ubahInvestasi() // menampilkan menu ubah investasi
		case 4:
			hapusInvestasi() // menampilkan menu hapus investasi
		case 5:
			hitungKeuntungan() // menampilkan menu hitung keuntungan
		case 6:
			cariBerdasarkanNama() // menampilkan menu cari berdasarkan nama
		case 7:
			cariBerdasarkanJenis() // menampilkan menu cari berdasrkan Jenis
		case 8:
			cariReturnTertinggi() // menampilkan menu cari berdasarkan return tertinggi
		case 9:
			urutkanNama() // menampilkan menu urutkan berdasarkan return
		case 10:
			urutkanReturn() // menampilkan menu urutkan berdasarkan return
		case 11:
			updateLaporanInvestasi() // menampilkan menu tampilkan portofolio lengkap
		case 12:
			riwayatTransaksi()
		case 0:
			fmt.Println("Hallo, terima kasih telah menggunakan Aplikasi Manajemen Investasi Sederhana !")
			return // mengakhiri aplikasi
		default: // default case untuk menu yang tidak ada
			fmt.Print("Menu tidak tersedia. Silakan pilih menu yang tersedia.\n\n") // menampilkan pesan menu tidak tersedia
		}
	}
}
