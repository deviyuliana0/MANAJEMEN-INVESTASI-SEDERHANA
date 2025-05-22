package sorting

import (
	"fmt"
	"proj_invest/models"
)

func urutkanNama() { // Selection Sort
	n := len(models.Portofolio)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if models.Portofolio[j].NamaAset < models.Portofolio[minIndex].NamaAset {
				minIndex = j
			}
		}

		models.Portofolio[i], models.Portofolio[minIndex] = models.Portofolio[minIndex], models.Portofolio[i]
	}
	models.urutkanNama()
	fmt.Println("Data berhasil diurutkan berdasarkan nama investasi.")
}
