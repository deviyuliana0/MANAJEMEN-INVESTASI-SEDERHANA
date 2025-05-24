package main

import (
	"fmt"

	"invest_proj/services/display"
	"invest_proj/services/operations"
	"invest_proj/services/searching"
	"invest_proj/services/sorting"
	"invest_proj/services/statistic"
)

func main() {
	for {
		fmt.Println("\n===================================")
		fmt.Println("=== Aplikasi Manajemen Investasi ===")
		fmt.Println("===================================")
		fmt.Println("1.  Tambah Investasi")
		fmt.Println("2.  Tampilkan Portofolio")
		fmt.Println("3.  Ubah Investasi")
		fmt.Println("4.  Hapus Investasi")
		fmt.Println("5.  Hitung Keuntungan / Kerugian")
		fmt.Println("6.  Cari Investasi berdasarkan Nama")
		fmt.Println("7.  Cari Investasi berdasarkan Jenis Aset")
		fmt.Println("8.  Urutkan berdasarkan Nama (A-Z)")
		fmt.Println("9.  Urutkan berdasarkan Return (tinggi → rendah)")
		fmt.Println("10. Tampilkan Update Laporan Portofolio")
		fmt.Println("11. Tampilkan Statistik Investasi")
		fmt.Println("12.  Keluar")
		fmt.Print("Pilih Menu Aplikasi: ")

		var menu int
		fmt.Scan(&menu)

		switch menu {
		case 1:
			operations.TambahInvestasi()
		case 2:
			display.TampilkanInvestasi()
		case 3:
			operations.UbahInvestasi()
		case 4:
			operations.HapusInvestasi()
		case 5:
			operations.HitungKeuntungan()
		case 6:
			searching.CariBerdasarkanNama()
		case 7:
			searching.CariBerdasarkanJenis()
		case 8:
			sorting.UrutkanNama()
		case 9:
			sorting.UrutkanReturn()
		case 10:
			operations.UpdateLaporanInvestasi()
		case 11:
			statistic.TampilkanStatistikInvestasi()
		case 0:
			fmt.Println("\nHallo, terima kasih telah menggunakan Aplikasi Manajemen Investasi Sederhana ! Sampai jumpa.\n")
			return
		default:
			fmt.Println("Menu tidak tersedia. Silakan pilih menu yang sesuai.")
		}
	}
}
