# AI-LOG

Catat setiap penggunaan AI yang menghasilkan kode. Jika tidak memakai AI sama sekali, tulis satu baris: **tidak memakai AI**.

Log yang jujur tidak mengurangi nilai, sebanyak apa pun AI dipakai. Yang mengurangi nilai adalah log yang tidak sesuai dengan kode, atau isi log yang tidak bisa dijelaskan.

## (Nama fitur atau fungsi, misalnya: Deteksi bentrok slot, B5)
- Prompt: (ringkas, apa yang kamu minta)
- Dipakai: (bagian kode yang kamu ambil, misalnya nama fungsi dan file)
- Diubah: (apa yang kamu ubah atau perbaiki dari hasil AI, dan kenapa)

## 1
- Prompt:  kerjakan B1 dengan todo sebagai berikut

```txt
todo:
1. model.go (bentuk datanya dulu)
   - buat konstanta: JenisRuang = "ruang", JenisAlat = "alat", MaxPanjangTeks = 50
   - buat struct SumberDaya: id, nama, jenis, stok
     (tag json persis: id, nama, jenis, stok)
   - buat struct SumberDayaInput: nama, jenis, stok
     (pakai pointer supaya "tidak dikirim" beda dengan 0 atau "")
     (tidak ada id, id milik server)

2. store.go (tempat simpan data di memori)
   - catat dulu rancangan struktur data di PROSES.md (struct apa, slice atau map, kenapa)
   - buat struct Store: daftar sumber daya, daftar peminjaman, counter id sumber daya, counter id peminjaman
     (counter dipisah, mulai dari 1, jangan pakai len()+1)
   - buat newStore() yang mengembalikan Store kosong
     (slice diinisialisasi, bukan nil, supaya JSON-nya [] bukan null)
   - buat method tambahSumberDaya: beri id dari counter, counter +1, simpan, kembalikan datanya
   - buat method ambilSumberDaya(id): kembalikan data dan true/false (ada atau tidak)
   - buat method daftarSumberDaya: kembalikan salinan, urut id naik

3. helper.go
   - buat writeJSON(w, status, data)
     (urutan: set header Content-Type application/json, lalu WriteHeader, lalu tulis body)
   - buat parseID(r): ambil r.PathValue("id"), ubah ke angka
     (kalau bukan angka kembalikan ok=false, nanti dijawab 400)

4. handler.go
   - buat struct Server yang menyimpan store
   - buat func NewHandler() http.Handler
     (tiap dipanggil: buat store BARU, jangan pakai variabel global)
   - di dalamnya buat http.NewServeMux dan daftarkan rute:
     POST /sumber-daya
     GET /sumber-daya
     GET /sumber-daya/{id}
   - return mux

5. sumber_daya.go (tiga handler)
   - buatSumberDaya (POST):
     decode body ke SumberDayaInput -> kalau JSON rusak jawab 400
     -> trim nama dan jenis, pastikan field tidak nil
     -> simpan lewat store -> jawab 201 + objek lengkap dengan id
   - daftarSumberDaya (GET): jawab 200 + array (kalau kosong tetap [])
   - detailSumberDaya (GET /{id}):
     id bukan angka -> 400
     data tidak ada -> 404
     sukses -> 200 + objek
   - catatan: validasi lengkap (nama maks 50 karakter, jenis, stok >= 1) menyusul di CP2/B4
```

- Dipakai: semua file
- Diubah: terlampir di todo



## (Entri berikutnya)
- Prompt: 
- Dipakai:
- Diubah:
