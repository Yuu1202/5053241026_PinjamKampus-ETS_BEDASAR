package main

import "net/http"

// Server membungkus store agar handler bisa mengaksesnya.
type Server struct {
	store *Store
}

// NewHandler membuat store BARU yang kosong pada setiap pemanggilan
// (tidak ada variabel global).
func NewHandler() http.Handler {
	srv := &Server{store: newStore()}
	mux := http.NewServeMux()

	mux.HandleFunc("POST /sumber-daya", srv.buatSumberDaya)
	mux.HandleFunc("GET /sumber-daya", srv.daftarSumberDaya)
	mux.HandleFunc("GET /sumber-daya/{id}", srv.detailSumberDaya)

	return mux
}
