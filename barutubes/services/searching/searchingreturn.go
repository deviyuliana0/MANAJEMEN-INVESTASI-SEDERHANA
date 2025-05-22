package searching

import (
	"fmt"
	"proj_invest/models"
)

func urutkanReturn() { // insertion sort
	for i := 1; i < len(models.Portofolio); i++ {
		key := models.Portofolio[i]
		j := i - 1
		for j >= 0 && hitungReturn(models.Portofolio[j]) > hitungReturn(key) {
			models.Portofolio[j+1] = models.Portofolio[j]
			j--
		}
		models.Portofolio[j+1] = key
	}
	models.urutkanReturn()
	fmt.Println("Data berhasil diurutkan berdasarkan return investasi.") // tertinggi ke rendah
}
