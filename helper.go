package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
)

// writeJSON: header dulu, baru status, baru body.
func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

// writeError: semua jalur error (400, 404, 409) lewat sini.
func writeError(w http.ResponseWriter, status int, pesan string) {
	writeJSON(w, status, map[string]string{"error": pesan})
}

// parseID membaca {id} dari path. ok=false bila bukan angka (-> 400).
func parseID(r *http.Request) (int, bool) {
	n, err := strconv.ParseUint(r.PathValue("id"), 10, 31)
	if err != nil {
		return 0, false
	}
	return int(n), true
}

// decodeBody mengubah body JSON ke struct input.
// Body kosong, JSON rusak, tipe field salah (teks / pecahan untuk int),
// atau ada sisa data setelah objek -> error (handler menjawab 400).
func decodeBody(r *http.Request, dst any) error {
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(dst); err != nil {
		return err
	}
	var sisa json.RawMessage
	if err := dec.Decode(&sisa); !errors.Is(err, io.EOF) {
		return errors.New("body harus berisi satu objek JSON")
	}
	return nil
}
