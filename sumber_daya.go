package main

import (
	"encoding/json"
	"net/http"
)

// POST /sumber-daya
// Alur: decode body -> simpan -> 201.
// Validasi lengkap (trim, maks 50 karakter, jenis, stok >= 1) menyusul di B4.
// Pengecekan nil di bawah hanya supaya tidak panic.
func (sv *Server) buatSumberDaya(w http.ResponseWriter, r *http.Request) {
	var in SumberDayaInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "body bukan JSON yang valid"})
		return
	}
	if in.Nama == nil || in.Jenis == nil || in.Stok == nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "nama, jenis, dan stok wajib diisi"})
		return
	}
	sd := sv.store.tambahSumberDaya(SumberDaya{
		Nama:  *in.Nama,
		Jenis: *in.Jenis,
		Stok:  *in.Stok,
	})
	writeJSON(w, http.StatusCreated, sd)
}

// GET /sumber-daya : 200 + array urut id naik, kosong = [].
func (sv *Server) daftarSumberDaya(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, sv.store.daftarSumberDaya())
}

// GET /sumber-daya/{id} : bukan angka 400, tidak ada 404, sukses 200.
func (sv *Server) detailSumberDaya(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(r)
	if !ok {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "id harus berupa angka"})
		return
	}
	sd, ada := sv.store.ambilSumberDaya(id)
	if !ada {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "sumber daya tidak ditemukan"})
		return
	}
	writeJSON(w, http.StatusOK, sd)
}
