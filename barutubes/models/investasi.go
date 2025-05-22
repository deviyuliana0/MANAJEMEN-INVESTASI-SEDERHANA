package models

type Investasi struct { // struct untuk menampung data investasi
	NamaAset      string  // nama aset yang diinvestasikan
	JenisAset     string  // jenis aset yang diinvestasikan
	HargaBeli     float64 // harga beli aset per unit
	HargaSekarang float64 // harga aset per unit saat ini
	JumlahUnit    int     // jumlah unit aset yang diinvestasikan
}

var Portofolio []Investasi // array untuk menampung data portofolio
var DaftarTransaksi []Investasi
