package main

// menahanSlot: pusat aturan status mana yang menahan slot.
// Bila persona / change request menambah status, ubah di sini saja.
func menahanSlot(p Peminjaman) bool {
	return p.Status == StatusAktif
}

// terpakai menjumlahkan "jumlah" (bukan banyaknya peminjaman) pada satu
// kombinasi (sumber daya, tanggal, slot).
// abaikanID > 0 -> peminjaman dengan id itu dilewati (dipakai PUT).
// POST memakai abaikanID 0 karena id mulai dari 1 (tidak ada yang dilewati).
func (s *Store) terpakai(sdID int, tanggal string, slot, abaikanID int) int {
	total := 0
	for _, p := range s.peminjaman {
		if p.ID == abaikanID || p.SumberDayaID != sdID || p.Tanggal != tanggal || !menahanSlot(p) {
			continue
		}
		if slot >= p.SlotMulai && slot <= p.SlotSelesai {
			total += p.Jumlah
		}
	}
	return total
}

// bentrok memeriksa SETIAP slot dari mulai sampai selesai.
// Bentrok bila pada salah satu slot: terpakai + jumlah baru > stok.
func (s *Store) bentrok(p Peminjaman, stok, abaikanID int) bool {
	for slot := p.SlotMulai; slot <= p.SlotSelesai; slot++ {
		if s.terpakai(p.SumberDayaID, p.Tanggal, slot, abaikanID)+p.Jumlah > stok {
			return true
		}
	}
	return false
}
