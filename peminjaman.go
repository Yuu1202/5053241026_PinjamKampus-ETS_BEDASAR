package main

import (
	"net/http"
	"strconv"
	"strings"
)

// POST /peminjaman : decode -> validasi (400) -> simpan -> 201.
// Cek bentrok (409) menyusul di B5.
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
