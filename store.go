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
