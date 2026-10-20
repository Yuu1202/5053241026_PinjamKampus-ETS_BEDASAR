package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
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

// GET /peminjaman?sumber_daya_id=&tanggal=&status=
// Semua filter opsional dan digabung (AND). Tanpa hasil = [] (bukan error).
func (sv *Server) daftarPeminjaman(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	var f FilterPeminjaman
	if q.Has("sumber_daya_id") {
		n, err := strconv.Atoi(strings.TrimSpace(q.Get("sumber_daya_id")))
		if err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "sumber_daya_id harus berupa angka"})
			return
		}
		f.SumberDayaID = &n
	}
	if q.Has("tanggal") {
		v := strings.TrimSpace(q.Get("tanggal"))
		f.Tanggal = &v
	}
	if q.Has("status") {
		v := strings.TrimSpace(q.Get("status"))
		f.Status = &v
	}
	writeJSON(w, http.StatusOK, sv.store.daftarPeminjaman(f))
}

// GET /peminjaman/{id} : bukan angka 400, tidak ada 404, sukses 200.
func (sv *Server) detailPeminjaman(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(r)
	if !ok {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "id harus berupa angka"})
		return
	}
	p, ada := sv.store.ambilPeminjaman(id)
	if !ada {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "peminjaman tidak ditemukan"})
		return
	}
	writeJSON(w, http.StatusOK, p)
}
