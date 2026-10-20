# PROSES

## 1. Persona dan masalahnya
belum mengerjakan persona

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

Cek bentrok dan penyimpanan saya lakukan dalam satu `Lock()` di dalam method store
(`buatPeminjaman` dan `ubahPeminjaman`). Sebelumnya handler mengecek bentrok lalu menyimpan
sebagai dua langkah terpisah, sehingga dua request bisa sama-sama lolos cek dan sama-sama menyimpan
(dua 201 untuk satu slot). Fungsi `terpakai` dan `bentrok` tidak mengunci sendiri karena dipanggil
saat kunci sudah dipegang. `sync.Mutex` tidak reentrant, jadi mengunci dua kali di goroutine yang sama
akan deadlock. Method daftar mengembalikan salinan slice supaya pemanggil tidak memegang data internal
di luar kunci.

## 4. Bug atau kendala (minimal dua)
### Kendala 1
- Gejala: Setup Go memakan banyak waktu dikarenakan baru pertamakali membuat project GO from scratch
- Cara menemukan: waktu
- Penyebab dan solusi: waktu

### Bug 2
- Gejala:postman lag
- Cara menemukan: lag
- Penyebab dan solusi: menggunakan Curl

## 5. Change request
belum dikerjakan