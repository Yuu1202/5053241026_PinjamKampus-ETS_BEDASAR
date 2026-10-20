package main

import "net/http"

// POST /sumber-daya : decode -> validasi -> simpan -> 201.
func (sv *Server) buatSumberDaya(w http.ResponseWriter, r *http.Request) {
	var in SumberDayaInput
	if err := decodeBody(r, &in); err != nil {
		writeError(w, http.StatusBadRequest, "body bukan JSON yang valid")
		return
	}
	sd, pesan := validasiSumberDaya(in)
	if pesan != "" {
		writeError(w, http.StatusBadRequest, pesan)
		return
	}
	writeJSON(w, http.StatusCreated, sv.store.tambahSumberDaya(sd))
}

// GET /sumber-daya : 200 + array urut id naik, kosong = [].
func (sv *Server) daftarSumberDaya(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, sv.store.daftarSumberDaya())
}

// GET /sumber-daya/{id} : bukan angka 400, tidak ada 404, sukses 200.
func (sv *Server) detailSumberDaya(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(r)
	if !ok {
		writeError(w, http.StatusBadRequest, "id harus berupa angka")
		return
	}
	sd, ada := sv.store.ambilSumberDaya(id)
	if !ada {
		writeError(w, http.StatusNotFound, "sumber daya tidak ditemukan")
		return
	}
	writeJSON(w, http.StatusOK, sd)
}
