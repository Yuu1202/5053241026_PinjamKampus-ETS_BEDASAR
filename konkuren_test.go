package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
)

func kirimK(h http.Handler, method, path, body string) int {
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest(method, path, strings.NewReader(body)))
	return rec.Code
}

// 50 request serentak ke slot yang sama pada sumber daya stok 1:
// harus TEPAT satu 201 dan sisanya 409. Jalankan: go test -race ./...
func TestKonkuren(t *testing.T) {
	const pinjam = `{"sumber_daya_id":1,"peminjam":"Budi","tanggal":"2026-11-01","slot_mulai":1,"slot_selesai":2,"jumlah":1}`
	for ulang := 0; ulang < 5; ulang++ {
		h := NewHandler()
		kirimK(h, "POST", "/sumber-daya", `{"nama":"Aula A","jenis":"ruang","stok":1}`)

		var wg sync.WaitGroup
		var mu sync.Mutex
		hitung := map[int]int{}
		for i := 0; i < 50; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				kode := kirimK(h, "POST", "/peminjaman", pinjam)
				mu.Lock()
				hitung[kode]++
				mu.Unlock()
			}()
		}
		wg.Wait()
		if hitung[201] != 1 || hitung[409] != 49 {
			t.Fatalf("ulang %d: harus tepat 1x201 dan 49x409, dapat %v", ulang, hitung)
		}
	}
}
