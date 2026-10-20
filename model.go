package main

//===== CP1 =====
//==== CP1 | B1 =====

// SumberDaya adalah ruang atau alat yang bisa dipinjam.
type SumberDaya struct {
	ID    int    `json:"id"`
	Nama  string `json:"nama"`
	Jenis string `json:"jenis"`
	Stok  int    `json:"stok"`
}

// SumberDayaInput untuk body POST /sumber-daya (tanpa id).
// Pointer: nil = field tidak dikirim, beda dari "" atau 0.
type SumberDayaInput struct {
	Nama  *string `json:"nama"`
	Jenis *string `json:"jenis"`
	Stok  *int    `json:"stok"`
}

//===== CP1 | B2 =====

const (
	SlotMin     = 1
	SlotMax     = 8
	StatusAktif = "aktif"
)

// Peminjaman: satu sumber daya, satu tanggal, satu rentang slot, satu jumlah.
type Peminjaman struct {
	ID           int    `json:"id"`
	SumberDayaID int    `json:"sumber_daya_id"`
	Peminjam     string `json:"peminjam"`
	Tanggal      string `json:"tanggal"` // YYYY-MM-DD
	SlotMulai    int    `json:"slot_mulai"`
	SlotSelesai  int    `json:"slot_selesai"`
	Jumlah       int    `json:"jumlah"`
	Status       string `json:"status"`
}

// PeminjamanInput untuk body POST /peminjaman (tanpa id dan status).
type PeminjamanInput struct {
	SumberDayaID *int    `json:"sumber_daya_id"`
	Peminjam     *string `json:"peminjam"`
	Tanggal      *string `json:"tanggal"`
	SlotMulai    *int    `json:"slot_mulai"`
	SlotSelesai  *int    `json:"slot_selesai"`
	Jumlah       *int    `json:"jumlah"`
}
