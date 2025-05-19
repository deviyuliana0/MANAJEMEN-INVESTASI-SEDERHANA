package services

import (
	"fmt"
	"invest_proj/data"
)

func MakeInvest(p *[100]data.Investasi, n *int) {
	if *n >= 100 {
		fmt.Println("Portofolio sudah penuh.")
		return
	}

	var nama, jenis string
	var beli, sekarang float64
	var unit int

	fmt.Print("Nama Aset: ")
	fmt.Scan(&nama)
	fmt.Print("Jenis Aset: ")
	fmt.Scan(&jenis)
	fmt.Print("Harga Beli: ")
	fmt.Scan(&beli)
	fmt.Print("Harga Sekarang: ")
	fmt.Scan(&sekarang)
	fmt.Print("Jumlah Unit: ")
	fmt.Scan(&unit)

	data.TambahInvestasi(p, n, nama, jenis, beli, sekarang, unit)
}
