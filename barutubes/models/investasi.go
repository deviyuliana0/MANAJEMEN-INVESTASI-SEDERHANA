package models

type Investasi struct {
	NamaAset      string
	JenisAset     string
	HargaBeli     float64
	HargaSekarang float64
	JumlahUnit    int
}

const MaksData = 100

var Portofolio [MaksData]Investasi // ✅ huruf besar → public/exported
var JumlahData int                 // ✅ agar bisa dipakai di package lain
