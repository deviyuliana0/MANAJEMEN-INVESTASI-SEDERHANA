package sorting

import (
	"fmt"
	"invest_proj/models"
)

func UrutkanNama() { // Selection Sort ascending A-Z
	n := models.JumlahData
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if models.Portofolio[j].NamaAset < models.Portofolio[minIndex].NamaAset {
				minIndex = j
			}
		}
		models.Portofolio[i], models.Portofolio[minIndex] = models.Portofolio[minIndex], models.Portofolio[i]
	}
	fmt.Println("Data berhasil diurutkan berdasarkan nama aset.")
}
