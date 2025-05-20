package data

func BinarySearch(p *[100]Investasi, n int, nama string) int {
	kiri := 0
	kanan := n - 1

	for kiri <= kanan {
		tengah := (kiri + kanan) / 2

		if p[tengah].NamaAset == nama {
			return tengah
		} else if p[tengah].NamaAset < nama {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	return -1
}
