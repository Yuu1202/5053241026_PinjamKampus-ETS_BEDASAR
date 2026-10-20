package main

import (
	"net/http"
	"strconv"
	"strings"
)

// POST /peminjaman : decode -> validasi (400) -> cek bentrok (409) -> simpan -> 201.
func (sv *Server) buatPeminjaman(w http.ResponseWriter, r *http.Request) {
	var in PeminjamanInput
	if err := decodeBody(r, &in); err != nil {
		writeError(w, http.StatusBadRequest, "body bukan JSON yang valid")
		return
	}
	p, pesan := validasiPeminjaman(in, sv.store.ambilSumberDaya)
	if pesan != "" {
		writeError(w, http.StatusBadRequest, pesan)
		return
	}
	sd, _ := sv.store.ambilSumberDaya(p.SumberDayaID) // pasti ada, sudah divalidasi
	if sv.store.bentrok(p, sd.Stok, 0) {
		writeError(w, http.StatusConflict, "slot bentrok: kapasitas sumber daya tidak cukup")
		return
	}
	writeJSON(w, http.StatusCreated, sv.store.tambahPeminjaman(p))
}

// GET /peminjaman?sumber_daya_id=&tanggal=&status=
// Semua filter opsional dan digabung (AND). Tanpa hasil = [] (bukan error).
func (sv *Server) daftarPeminjaman(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	var f FilterPeminjaman
	if q.Has("sumber_daya_id") {
		n, err := strconv.Atoi(strings.TrimSpace(q.Get("sumber_daya_id")))
		if err != nil {
			writeError(w, http.StatusBadRequest, "sumber_daya_id harus berupa angka")
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
		writeError(w, http.StatusBadRequest, "id harus berupa angka")
		return
	}
	p, ada := sv.store.ambilPeminjaman(id)
	if !ada {
		writeError(w, http.StatusNotFound, "peminjaman tidak ditemukan")
		return
	}
	writeJSON(w, http.StatusOK, p)
}

// ===== B6 =====

// PUT /peminjaman/{id}
// Urutan: id bukan angka 400 -> tidak ada 404 -> body/validasi 400 -> bentrok 409 -> 200.
func (sv *Server) ubahPeminjaman(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(r)
	if !ok {
		writeError(w, http.StatusBadRequest, "id harus berupa angka")
		return
	}
	if _, ada := sv.store.ambilPeminjaman(id); !ada {
		writeError(w, http.StatusNotFound, "peminjaman tidak ditemukan")
		return
	}
	var in PeminjamanInput
	if err := decodeBody(r, &in); err != nil {
		writeError(w, http.StatusBadRequest, "body bukan JSON yang valid")
		return
	}
	p, pesan := validasiPeminjaman(in, sv.store.ambilSumberDaya)
	if pesan != "" {
		writeError(w, http.StatusBadRequest, pesan)
		return
	}
	sd, _ := sv.store.ambilSumberDaya(p.SumberDayaID)
	if sv.store.bentrok(p, sd.Stok, id) { // abaikan dirinya sendiri
		writeError(w, http.StatusConflict, "slot bentrok: kapasitas sumber daya tidak cukup")
		return
	}
	hasil, _ := sv.store.gantiPeminjaman(id, p)
	writeJSON(w, http.StatusOK, hasil)
}

// DELETE /peminjaman/{id} : bukan angka 400, tidak ada 404, sukses 204 TANPA body.
func (sv *Server) hapusPeminjaman(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(r)
	if !ok {
		writeError(w, http.StatusBadRequest, "id harus berupa angka")
		return
	}
	if !sv.store.hapusPeminjaman(id) {
		writeError(w, http.StatusNotFound, "peminjaman tidak ditemukan")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
