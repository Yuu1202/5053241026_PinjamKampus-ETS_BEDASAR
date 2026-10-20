package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// writeJSON: header dulu, baru status, baru body.
func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

// parseID membaca {id} dari path. ok=false bila bukan angka (-> 400).
func parseID(r *http.Request) (int, bool) {
	n, err := strconv.ParseUint(r.PathValue("id"), 10, 31)
	if err != nil {
		return 0, false
	}
	return int(n), true
}
