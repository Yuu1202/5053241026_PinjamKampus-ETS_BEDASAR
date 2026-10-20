package main

import (
	"encoding/json"
	"net/http"
)

// POST /peminjaman (versi dasar): decode -> simpan -> 201.
// Status awal "aktif" dari server, bukan dari input.
// Validasi dan cek bentrok menyusul di B4 dan B5.
// Pengecekan nil di bawah hanya supaya tidak panic.
func (sv *Server) buatPeminjaman(w http.ResponseWriter, r *http.Request) {
	var in PeminjamanInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "body bukan JSON yang valid"})
		return
	}
	if in.SumberDayaID == nil || in.Peminjam == nil || in.Tanggal == nil ||
		in.SlotMulai == nil || in.SlotSelesai == nil || in.Jumlah == nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "semua field wajib diisi"})
		return
	}
	p := sv.store.tambahPeminjaman(Peminjaman{
		SumberDayaID: *in.SumberDayaID,
		Peminjam:     *in.Peminjam,
		Tanggal:      *in.Tanggal,
		SlotMulai:    *in.SlotMulai,
		SlotSelesai:  *in.SlotSelesai,
		Jumlah:       *in.Jumlah,
	})
	writeJSON(w, http.StatusCreated, p)
}
