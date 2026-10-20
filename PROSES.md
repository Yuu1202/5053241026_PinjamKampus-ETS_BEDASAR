# PROSES

## 1. Persona dan masalahnya
(2-3 kalimat dengan kata-katamu sendiri: siapa pengguna persona kamu, masalah apa yang diselesaikan fitur khasnya.)

## 2. Rancangan struktur data
(Tulis struct dan slice/map apa yang kamu simpan, dan kenapa dipilih begitu. Rancang ini sebelum menulis kode, lalu catat kalau berubah di tengah jalan.)

Semua data disimpan di memori dalam satu struct `Store`:
- `sumberDaya []SumberDaya`. Slice dipilih karena data ditambah di akhir, jadi urutannya
  otomatis sama dengan urutan id naik dan daftar tidak perlu diurutkan lagi.
- `nextSDID` sebagai counter id sumber daya, mulai dari 1. Counter dipakai, bukan len+1,
  supaya id tidak dipakai ulang kalau nanti ada data yang dihapus.
- Struct input memakai pointer supaya field yang tidak dikirim (nil) bisa dibedakan
  dari nilai kosong atau 0.
- Peminjaman, counter id peminjaman, dan mutex akan ditambahkan di B2 dan B7.


## 3. Satu keputusan desain
(Satu keputusan yang kamu ambil, pilihan lain yang kamu pertimbangkan, dan alasanmu memilih yang ini. Contoh: kenapa cek bentrok per slot, bukan per rentang.)

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
