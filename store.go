package main

type Store struct {
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

func (s *Store) tambahSumberDaya(sd SumberDaya) SumberDaya {
	sd.ID = s.nextSDID
	s.nextSDID++
	s.sumberDaya = append(s.sumberDaya, sd)
	return sd
}

func (s *Store) ambilSumberDaya(id int) (SumberDaya, bool) {
	for _, sd := range s.sumberDaya {
		if sd.ID == id {
			return sd, true
		}
	}
	return SumberDaya{}, false
}

func (s *Store) daftarSumberDaya() []SumberDaya {
	out := make([]SumberDaya, len(s.sumberDaya))
	copy(out, s.sumberDaya)
	return out
}

func (s *Store) tambahPeminjaman(p Peminjaman) Peminjaman {
	p.ID = s.nextPJID
	s.nextPJID++
	p.Status = StatusAktif
	s.peminjaman = append(s.peminjaman, p)
	return p
}

func (s *Store) ambilPeminjaman(id int) (Peminjaman, bool) {
	for _, p := range s.peminjaman {
		if p.ID == id {
			return p, true
		}
	}
	return Peminjaman{}, false
}

func (s *Store) daftarPeminjaman(f FilterPeminjaman) []Peminjaman {
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

// ===== B6 =====

// gantiPeminjaman mengganti keenam field; id dan status TIDAK berubah.
// Dipanggil hanya setelah validasi dan cek bentrok lolos, jadi bila 409
// data lama tetap utuh.
func (s *Store) gantiPeminjaman(id int, p Peminjaman) (Peminjaman, bool) {
	for i := range s.peminjaman {
		if s.peminjaman[i].ID == id {
			p.ID = id
			p.Status = s.peminjaman[i].Status
			s.peminjaman[i] = p
			return p, true
		}
	}
	return Peminjaman{}, false
}

// hapusPeminjaman membuang data; urutan sisanya tetap id naik.
func (s *Store) hapusPeminjaman(id int) bool {
	for i := range s.peminjaman {
		if s.peminjaman[i].ID == id {
			s.peminjaman = append(s.peminjaman[:i], s.peminjaman[i+1:]...)
			return true
		}
	}
	return false
}
