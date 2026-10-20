# PROSES

## 1. Persona dan masalahnya
(2-3 kalimat dengan kata-katamu sendiri: siapa pengguna persona kamu, masalah apa yang diselesaikan fitur khasnya.)

## 2. Rancangan struktur data
Semua data disimpan di memori dalam satu struct `Store`:
- `sumberDaya []SumberDaya` dan `peminjaman []Peminjaman`. Slice dipilih karena data
  ditambah di akhir, jadi urutannya otomatis sama dengan urutan id naik dan daftar
  tidak perlu diurutkan lagi.
- `nextSDID` dan `nextPJID` sebagai counter id yang terpisah, masing-masing mulai dari 1.
  Counter dipakai, bukan len+1, supaya id tidak dipakai ulang kalau nanti ada data yang dihapus.
- Struct input memakai pointer supaya field yang tidak dikirim (nil) bisa dibedakan
  dari nilai kosong atau 0.
- `FilterPeminjaman` juga memakai pointer: nil berarti filter itu tidak dipakai.
- Mutex akan ditambahkan di B7.


## 3. Satu keputusan desain
`NewHandler()` membuat `Store` baru setiap dipanggil, dan store disimpan di struct `Server`,
bukan di variabel global. Alasannya, data tidak boleh bocor antar pemanggilan: tiap
`NewHandler()` harus mulai dari kosong dengan id kembali ke 1.

## 4. Bug atau kendala (minimal dua)
### Bug 1
- Gejala:
- Cara menemukan:
- Penyebab dan solusi:

### Bug 2
- Gejala:
- Cara menemukan:
- Penyebab dan solusi:

## 5. Change request
(Apa yang diminta, bagian kode mana yang kamu ubah, dan satu keputusan desain yang kamu ambil.)
