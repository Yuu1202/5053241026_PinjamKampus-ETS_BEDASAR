package main

import (
	"net/http"
	"sync"
)

// Store menyimpan semua data di memori. Semua akses ke field di bawah
// WAJIB lewat mu, termasuk saat membaca.
type Store struct {
	mu         sync.Mutex
	sumberDaya []SumberDaya
	nextSDID   int
	peminjaman []Peminjaman
	nextPJID   int
}

func newStore() *Store {
	return &Store{
		sumberDaya: []SumberDaya{},
		nextSDID:   1,
		peminjaman: []Peminjaman{},
		nextPJID:   1,
	}
}

// ---------- sumber daya ----------

func (s *Store) tambahSumberDaya(sd SumberDaya) SumberDaya {
	s.mu.Lock()
	defer s.mu.Unlock()
	sd.ID = s.nextSDID
	s.nextSDID++
	s.sumberDaya = append(s.sumberDaya, sd)
	return sd
}

func (s *Store) ambilSumberDaya(id int) (SumberDaya, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, sd := range s.sumberDaya {
		if sd.ID == id {
			return sd, true
		}
	}
	return SumberDaya{}, false
}

// daftarSumberDaya mengembalikan SALINAN (tidak pernah nil -> JSON []).
func (s *Store) daftarSumberDaya() []SumberDaya {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make([]SumberDaya, len(s.sumberDaya))
	copy(out, s.sumberDaya)
	return out
}

// ---------- peminjaman ----------

func (s *Store) ambilPeminjaman(id int) (Peminjaman, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, p := range s.peminjaman {
		if p.ID == id {
			return p, true
		}
	}
	return Peminjaman{}, false
}

// daftarPeminjaman: filter digabung dengan AND, urut id naik, tidak pernah nil.
func (s *Store) daftarPeminjaman(f FilterPeminjaman) []Peminjaman {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := []Peminjaman{}
	for _, p := range s.peminjaman {
		if f.SumberDayaID != nil && p.SumberDayaID != *f.SumberDayaID {
			continue
		}
		if f.Tanggal != nil && p.Tanggal != *f.Tanggal {
			continue
		}
		if f.Status != nil && p.Status != *f.Status {
			continue
		}
		out = append(out, p)
	}
	return out
}

// buatPeminjaman: cek bentrok DAN simpan dalam SATU kunci.
// Dua request yang berebut slot sama: satu lolos (201), satu ditolak (409).
// Mengembalikan ok=false bila bentrok.
func (s *Store) buatPeminjaman(p Peminjaman, stok int) (Peminjaman, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.bentrok(p, stok, 0) { // dipanggil saat lock sudah dipegang
		return Peminjaman{}, false
	}
	p.ID = s.nextPJID
	s.nextPJID++
	p.Status = StatusAktif
	s.peminjaman = append(s.peminjaman, p)
	return p, true
}

// ubahPeminjaman: cek bentrok (mengabaikan diri sendiri) DAN ganti dalam satu kunci.
// Kode: 0 sukses, 404 tidak ada, 409 bentrok. Bila bukan 0, data lama tetap utuh.
func (s *Store) ubahPeminjaman(id int, p Peminjaman, stok int) (Peminjaman, int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i := range s.peminjaman {
		if s.peminjaman[i].ID != id {
			continue
		}
		if s.bentrok(p, stok, id) {
			return Peminjaman{}, http.StatusConflict
		}
		p.ID = id
		p.Status = s.peminjaman[i].Status
		s.peminjaman[i] = p
		return p, 0
	}
	return Peminjaman{}, http.StatusNotFound
}

func (s *Store) hapusPeminjaman(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i := range s.peminjaman {
		if s.peminjaman[i].ID == id {
			s.peminjaman = append(s.peminjaman[:i], s.peminjaman[i+1:]...)
			return true
		}
	}
	return false
}
