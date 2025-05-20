package data

import "fmt"

type Investasi struct {
	NamaAset      string
	JenisAset     string
	HargaBeli     float64
	HargaSekarang float64
	JumlahUnit    int
}

func TambahInvestasi(p *[100]Investasi, n *int, nama, jenis string, beli, sekarang float64, unit int) {
	if *n >= 100 {
		fmt.Println("Data penuh.")
		return
	}

	p[*n] = Investasi{
		NamaAset:      nama,
		JenisAset:     jenis,
		HargaBeli:     beli,
		HargaSekarang: sekarang,
		JumlahUnit:    unit,
	}

	*n++
	fmt.Println("Data berhasil ditambahkan.")
}
