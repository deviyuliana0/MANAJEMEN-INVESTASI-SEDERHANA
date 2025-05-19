package main

import "fmt"

type Investasi struct { // struct untuk menampung data investasi
	NamaAset      string  // nama aset yang diinvestasikan
	JenisAset     string  // jenis aset yang diinvestasikan
	HargaBeli     float64 // harga beli aset per unit
	HargaSekarang float64 // harga aset per unit saat ini
	JumlahUnit    int     // jumlah unit aset yang diinvestasikan
}

var portofolio []Investasi // array untuk menampung data portofolio
var daftarTransaksi []Investasi

func tambahInvestasi() { // fungsi untuk menambahkan investasi ke portofolio
	var invest Investasi // variabel untuk menampung data investasi yang akan ditambahkan

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

	portofolio = append(portofolio, invest)    // menambahkan data investasi ke array portofolio
	fmt.Println("Data berhasil diinputkan!\n") // menampilkan pesan berhasil input data
}

func tampilkanInvestasi() {
	fmt.Println("=== Daftar Investasi ===")
	for i := 0; i < len(portofolio); i++ {
		fmt.Printf("[%d] %s (%s), Beli: %.2f, Sekarang: %.2f, Unit: %d\n",
			i+1,
			portofolio[i].NamaAset,
			portofolio[i].JenisAset,
			portofolio[i].HargaBeli,
			portofolio[i].HargaSekarang,
			portofolio[i].JumlahUnit)
	}
}

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
			tambahInvestasi() // menampilkan menu tambah investasi
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
			fmt.Println("Menu tidak tersedia. Silakan pilih menu yang tersedia.\n") // menampilkan pesan menu tidak tersedia
		}
	}
}

func ubahInvestasi() {
	if len(portofolio) == 0 { // cek apakah portofolio kosong
		fmt.Println("Belum ada data untuk diubah")
		return
	}

	tampilkanInvestasi()
	var index int                                              
	fmt.Print("Pilih nomor data Investasi yang ingin diubah:") 
	fmt.Scan(&index)                                          

	if index < 1 || index > len(portofolio) { // cek apakah index valid
		fmt.Println("Index tidak valid")
		return
	}

	i := index - 1 // mengubah index menjadi indeks array
	fmt.Println("Input data baru:")

	fmt.Print("Nama Aset	:")
	fmt.Scan(&portofolio[i].NamaAset) // menginput nama aset

	fmt.Print("Jenis Aset	:")
	fmt.Scan(&portofolio[i].JenisAset) // menginput jenis aset

	fmt.Print("Harga Beli	:")          // menampilkan pesan harga beli
	fmt.Scan(&portofolio[i].HargaBeli) // menginput harga beli

	fmt.Print("Harga Sekarang	:")          // menampilkan pesan harga sekarang
	fmt.Scan(&portofolio[i].HargaSekarang) // menginput harga sekarang

	fmt.Print("Jumlah Unit	:")          // menampilkan pesan jumlah unit
	fmt.Scan(&portofolio[i].JumlahUnit) // menginput jumlah unit

	fmt.Println("Data berhasil diubah!\n")
}

func hapusInvestasi() {
	if len(portofolio) == 0 { // cek apakah portofolio kosong
		fmt.Println("Belum ada data untuk dihapus")
		return
	}

	tampilkanInvestasi()
	var index int                                                // variabel untuk menampung index investasi yang ingin dihapus
	fmt.Print("Pilih nomor data Investasi yang ingin dihapus: ") // menampilkan pesan
	fmt.Scan(&index)

	if index < 1 || index > len(portofolio) { // cek apakah index valid
		fmt.Println("Index tidak valid")
		return
	}

	i := index - 1                                           // mengubah index menjadi indeks array
	portofolio = append(portofolio[:i], portofolio[i+1:]...) // menghapus data investasi
	fmt.Println("Data berhasil dihapus!\n")
}

func hitungKeuntungan() {
	if len(portofolio) == 0 { // cek apakah portofolio kosong
		fmt.Println("Belum ada investasi.")
		return
	}

	fmt.Println("=== Hitung Keuntungan / Kerugian ===")
	for _, invest := range portofolio {
		selisih := invest.HargaSekarang - invest.HargaBeli      // menghitung selisih harga
		totalKeuntungan := selisih * float64(invest.JumlahUnit) // menghitung total keuntungan

		if selisih > 0 {
			fmt.Printf("Yeay, Selamat Kamu Mendapatkan Keuntungan dari %s sebesar : Rp %.2f\n", invest.NamaAset, totalKeuntungan)
		} else if selisih < 0 {
			fmt.Printf("Yah, Jangan Bersedih Terdapat Kerugian dari %s : Rp %.2f\n", invest.NamaAset, -totalKeuntungan) // menampilkan pesan kerugian
		} else {
			fmt.Printf("Tidak ada keuntungan atau kerugian dari %s\n", invest.NamaAset) // menampilkan pesan tidak ada keuntungan atau kerugian
		}
	}
}

func cariBerdasarkanNama() {
	if len(portofolio) == 0 {
		fmt.Println("Belum ada data untuk dicari")
		return
	}

	var nama string
	fmt.Print("Masukkan nama aset yang dicari : ")
	fmt.Scan(&nama)

	ditemukan := false
	for _, invest := range portofolio {
		if invest.NamaAset == nama { 
			fmt.Println("Ditemukan:")
			fmt.Printf("%s (%s), Beli : %.2f, Sekarang : %.2f, Unit : %d\n", invest.NamaAset, invest.JenisAset, invest.HargaBeli, invest.HargaSekarang, invest.JumlahUnit) // menampilkan data aset yang ditemukan
			ditemukan = true
			break
		}
	}

	if !ditemukan {
		fmt.Println("Data Tidak ditemukan.")
	}
}

func cariBerdasarkanJenis() {
	if len(portofolio) == 0 { 
		fmt.Println("Belum ada data untuk dicari")
		return
	}

	var jenis string
	fmt.Print("Masukkan jenis aset yang dicari (saham, obligasi, emas, reksana: ") // menampilkan pesan
	fmt.Scan(&jenis)

	ditemukan := false
	fmt.Println("=== Hasil Pencarian ===")
	for _, invest := range portofolio {
		if invest.JenisAset == jenis { 
			fmt.Printf("%s (%s), Beli : %.2f, Sekarang : %.2f, Unit : %d\n", invest.NamaAset, invest.JenisAset, invest.HargaBeli, invest.HargaSekarang, invest.JumlahUnit)
			ditemukan = true
		}
	}

	if !ditemukan {
		fmt.Println("Data Investasi Tidak ditemukan dengan jenis tersebut.")
	}
}

func hitungReturn(invest Investasi) float64 {
	return (invest.HargaSekarang - invest.HargaBeli) * float64(invest.JumlahUnit)
}

func cariReturnTertinggi() {
	if len(portofolio) == 0 { 
		fmt.Println("Belum ada data untuk dicari")
		return
	}

	maxIndex := 0
	maxReturn := hitungReturn(portofolio[0])

	for i := 1; i < len(portofolio); i++ {
		investReturn := hitungReturn(portofolio[i])
		if investReturn > maxReturn {
			maxReturn = investReturn
			maxIndex = i
		}
	}

	invest := portofolio[maxIndex]
	fmt.Println("=== Investasi dengan Return Tertinggi ===")
	fmt.Printf("%s (%s), Return : Rp %.2f, Beli : %.2f, Sekarang : %.2f, Unit %d\n", invest.NamaAset, invest.JenisAset, maxReturn, invest.HargaBeli, invest.HargaSekarang, invest.JumlahUnit)
}

func urutkanNama() { // Selection Sort
	n := len(portofolio)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if portofolio[j].NamaAset < portofolio[minIndex].NamaAset {
				minIndex = j
			}
		}
		
		portofolio[i], portofolio[minIndex] = portofolio[minIndex], portofolio[i] 
	}
	fmt.Println("Data berhasil diurutkan berdasarkan nama investasi.")
}

func urutkanReturn() { // insertion sort
	for i := 1; i < len(portofolio); i++ {
		key := portofolio[i]
		j := i - 1
		for j >= 0 && hitungReturn(portofolio[j]) > hitungReturn(key) {
			portofolio[j+1] = portofolio[j]
			j--
		}
		portofolio[j+1] = key
	}
	fmt.Println("Data berhasil diurutkan berdasarkan return investasi.") // tertinggi ke rendah
}

func updateLaporanInvestasi() {
	if len(portofolio) == 0 {
		fmt.Println("Data Investasi kosong")
		return
	}

	fmt.Println("=== Laporan Portofolio Investasi ===")
	var totalReturn float64 = 0

	for _, invest := range portofolio {
		// Hitung return investasi
		returnInvest := hitungReturn(invest)
		totalReturn += returnInvest

		status := "Tidak Mencapai Target, Jangan Menyerah"
		if returnInvest > 0 {
			status = "Mencapai Target, Selamat !"
		} else if returnInvest == 0 {
			status = "Yah, Rugi bang!"
		}
		
		fmt.Printf("%s(%s) | Beli : %.2f | Sekarang : %.2f | Umit : %d | Return : %2.f (%s)\n", invest.NamaAset, invest.JenisAset, invest.HargaBeli, invest.HargaSekarang, invest.JumlahUnit, returnInvest, status) // Tampilkan data investasi
	}

	fmt.Printf("💸 Total Return Return Anda : %.2f\n", totalReturn)
}

func riwayatTransaksi() {
	if len(daftarTransaksi) == 0 {
		fmt.Println("Belum ada riwayat transaksi")
		return
	}

	fmt.Println("=== Riwayat Transaksi Investasi ===")
	for i := 0; i < len(daftarTransaksi); i++ {
		fmt.Printf("%d. %s data: %s (%s)\n",
			i+1,
			daftarTransaksi[i].NamaAset,
			daftarTransaksi[i].JenisAset,
			daftarTransaksi[i].JumlahUnit)
	}
}
