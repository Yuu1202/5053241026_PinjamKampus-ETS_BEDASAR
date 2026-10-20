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

Validasi saya pecah jadi fungsi kecil (`validasiTeks`, `validasiTanggal`, `validasiSlot`, `validasiPeminjaman`)
supaya aturan yang sama bisa dipakai ulang oleh POST dan PUT. Panjang teks dihitung dengan
`utf8.RuneCountInString`, bukan `len`, karena `len` menghitung byte sehingga huruf non-ASCII
bisa salah dihitung. Tanggal diperiksa dengan `time.Parse` ditambah cek panjang 10 karakter,
supaya format seperti `2026-1-5` ikut ditolak.

Bentrok dihitung dengan menjumlahkan `jumlah` pada tiap slot, bukan sekadar mengecek ada atau tidaknya
peminjaman lain. Dengan begitu satu aturan yang sama melayani ruang (stok 1) dan alat (stok lebih dari 1).
Slot diperiksa satu per satu dari mulai sampai selesai, jadi kasus tumpang tindih sebagian seperti
3-5 lawan 4-4 ikut terdeteksi, tidak hanya yang `slot_mulai`-nya sama. Aturan "status mana yang menahan slot"
saya taruh di satu fungsi (`menahanSlot`) supaya mudah diubah nanti.

## 4. Bug atau kendala (minimal dua)
### Kendala 1
- Gejala: Setup Go memakan banyak waktu dikarenakan baru pertamakali membuat project GO from scratch
- Cara menemukan:
- Penyebab dan solusi:

### Bug 2
- Gejala:
- Cara menemukan:
- Penyebab dan solusi:

## 5. Change request
(Apa yang diminta, bagian kode mana yang kamu ubah, dan satu keputusan desain yang kamu ambil.)
