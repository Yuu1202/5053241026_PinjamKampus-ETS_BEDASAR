# AI-LOG

Catat setiap penggunaan AI yang menghasilkan kode. Jika tidak memakai AI sama sekali, tulis satu baris: **tidak memakai AI**.

Log yang jujur tidak mengurangi nilai, sebanyak apa pun AI dipakai. Yang mengurangi nilai adalah log yang tidak sesuai dengan kode, atau isi log yang tidak bisa dijelaskan.

## 1 (B1: Kelola sumber daya)
- Prompt: kerjakan B1 dengan todo sebagai berikut

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

- Dipakai: model.go, store.go, helper.go, handler.go, sumber_daya.go
- Diubah: terlampir di todo

## 2 (B2: Buat peminjaman)
- Prompt: kerjakan B2 dengan todo sebagai berikut

```txt
todo:
1. model.go (tambah bentuk data peminjaman)
   - tambah konstanta: SlotMin = 1, SlotMax = 8, StatusAktif = "aktif"
   - buat struct Peminjaman dengan tag json PERSIS:
     id, sumber_daya_id, peminjam, tanggal, slot_mulai, slot_selesai, jumlah, status
     (jangan diubah jadi camelCase)
   - buat struct PeminjamanInput berisi 6 field:
     sumber_daya_id, peminjam, tanggal, slot_mulai, slot_selesai, jumlah
     (pakai pointer, tanpa id dan tanpa status karena itu milik server)

2. store.go
   - buat method tambahPeminjaman:
     beri id dari counter peminjaman (mulai 1, +1 tiap data baru yang tersimpan)
     set status = StatusAktif (status bukan dari input user)
     simpan ke daftar peminjaman, kembalikan datanya
   - id peminjaman terpisah dari id sumber daya

3. peminjaman.go
   - buat handler buatPeminjaman (POST):
     decode body ke PeminjamanInput -> kalau JSON rusak jawab 400
     -> ubah input jadi Peminjaman (ambil nilai dari pointer)
     -> simpan lewat store -> jawab 201 + objek (status "aktif")
   - versi CP1 ini cukup: decode -> simpan -> 201
     (validasi menyusul di B4, cek bentrok di B5)

4. handler.go
   - daftarkan rute: POST /peminjaman
```

- Dipakai: model.go, store.go, peminjaman.go, handler.go
- Diubah: terlampir di todo

## 3 (B3: Lihat peminjaman)
- Prompt: kerjakan B3 dengan todo sebagai berikut

```txt
todo:
1. model.go
   - buat struct FilterPeminjaman: sumber_daya_id, tanggal, status
     (pakai pointer: nil artinya filter tidak dipakai)

2. store.go
   - buat method ambilPeminjaman(id): kembalikan data dan true/false
   - buat method daftarPeminjaman(filter):
     mulai dari slice kosong (bukan nil, supaya hasilnya [])
     loop semua peminjaman, lewati yang tidak cocok dengan filter yang diisi
     (semua filter digabung AND)
     hasil urut id naik

3. peminjaman.go
   - buat handler daftarPeminjaman (GET /peminjaman):
     baca query ?sumber_daya_id= ?tanggal= ?status= (semuanya opsional)
     kalau sumber_daya_id ada tapi bukan angka -> 400
     isi FilterPeminjaman sesuai query yang dikirim
     jawab 200 + array
     (filter tanpa hasil -> [], bukan error dan bukan null)
   - buat handler detailPeminjaman (GET /peminjaman/{id}):
     id bukan angka -> 400
     data tidak ada -> 404
     sukses -> 200 + objek

4. handler.go
   - daftarkan rute:
     GET /peminjaman
     GET /peminjaman/{id}
```

- Dipakai: model.go, store.go, peminjaman.go, handler.go
- Diubah: terlampir di todo

## 4 (B4: Validasi dan error)
- Prompt: kerjakan B4 dengan todo sebagai berikut

```txt
todo:
1. helper.go (rapikan jalur error dulu)
   - buat writeError(w, status, pesan)
     body persis {"error": "pesan"}, lewat writeJSON supaya Content-Type tetap application/json
   - buat decodeBody(r, dst) untuk baca body JSON ke struct input
     body kosong, JSON rusak, tipe field salah (slot_mulai berisi teks atau 1.5) -> kembalikan error
     body yang punya sisa data setelah satu objek JSON juga dianggap error
   - ganti semua jalur error yang sudah ada di CP1 supaya memakai writeError

2. validasi.go (pecah jadi fungsi kecil, bukan satu fungsi raksasa)
   - buat fungsi bantu validasiTeks(label, teks):
     trim dulu, tidak boleh kosong setelah trim
     maks 50 KARAKTER (pakai utf8.RuneCountInString, bukan len())
     kembalikan teks yang sudah di-trim atau pesan error
   - validasiSumberDaya(input):
     nama, jenis, stok semuanya wajib ada (cek pointer tidak nil)
     nama lewat validasiTeks
     jenis hanya "ruang" atau "alat" (persis, huruf kecil)
     stok minimal 1
   - validasiTanggal(teks):
     format YYYY-MM-DD dan tanggalnya benar-benar ada (pakai time.Parse)
     tolak "2026-02-30", "2026-02-29", "2026-1-5", "05-11-2026", ""
     terima "2028-02-29"
   - validasiSlot(mulai, selesai):
     masing-masing 1 sampai 8
     mulai <= selesai (mulai == selesai sah, mulai > selesai ditolak)
   - validasiPeminjaman(input, pencari sumber daya):
     enam field wajib ada
     peminjam lewat validasiTeks
     tanggal lewat validasiTanggal, slot lewat validasiSlot
     sumber_daya_id harus ada, kalau tidak ada hasilnya 400 (BUKAN 404)
     jumlah antara 1 dan stok sumber daya tersebut, jumlah > stok -> 400 (bukan 409)
     tulis sekali, nanti dipakai POST dan PUT

3. sambungkan ke handler
   - sumber_daya.go: POST panggil validasiSumberDaya, semua pelanggaran -> 400
   - peminjaman.go: POST panggil validasiPeminjaman, gagal -> 400
   - simpan nilai yang sudah di-trim, bukan nilai mentah dari input
```

- Dipakai: helper.go, validasi.go, sumber_daya.go, peminjaman.go
- Diubah: terlampir di todo

## 5 (B5: Deteksi bentrok)
- Prompt: kerjakan B5 dengan todo sebagai berikut

```txt
todo:
1. fungsi aturan "siapa yang menahan slot" (bentrok.go)
   - buat menahanSlot(p Peminjaman) bool
     hanya status "aktif" yang menahan slot
   - taruh di satu fungsi saja supaya kalau nanti ada status baru (persona / CR) cukup ubah di sini

2. fungsi hitung unit yang sudah terpakai pada satu slot
   - buat terpakaiLocked(sumberDayaID, tanggal, slot, abaikanID)
   - loop semua peminjaman, lewati yang:
     id-nya sama dengan abaikanID (untuk PUT nanti)
     sumber daya atau tanggalnya beda
     tidak menahan slot
   - kalau slot yang dicek berada di antara slot_mulai dan slot_selesai peminjaman itu,
     tambahkan "jumlah"-nya (jumlahkan jumlah, bukan banyaknya peminjaman)
   - kembalikan totalnya

3. fungsi cek bentrok untuk calon peminjaman
   - buat bentrokLocked(calon, stok, abaikanID)
   - loop SETIAP slot dari slot_mulai sampai slot_selesai
   - kalau di salah satu slot (terpakai + jumlah calon) > stok -> bentrok (true)
   - kalau semua slot aman -> false
   - jangan hanya membandingkan slot_mulai

4. sambungkan ke POST /peminjaman
   - urutan: validasi dulu (400), baru cek bentrok (409)
   - kalau bentrok jawab 409 pakai writeError
   - kalau lolos baru simpan dan jawab 201
   - nama fungsi berakhiran Locked karena nanti dipanggil saat lock dipegang (B7)
```

- Dipakai: bentrok.go, store.go, peminjaman.go
- Diubah: terlampir di todo

## 6 (B6: Ubah dan hapus)
- Prompt: kerjakan B6 dengan todo sebagai berikut

```txt
todo:
1. store.go: method ubahPeminjaman(id, input) (PUT)
   - cari indeks peminjaman berdasarkan id, kalau tidak ada -> 404
   - validasi input pakai validasiPeminjaman yang sama dengan POST -> kalau gagal 400
   - cek bentrok pakai bentrokLocked dengan abaikanID = id
     (supaya peminjaman yang diubah tidak bentrok dengan dirinya sendiri)
   - kalau bentrok -> 409 dan JANGAN ubah data lama
     (hitung dulu pakai data calon, ganti data hanya kalau lolos)
   - kalau lolos: ganti KEENAM field sekaligus, id dan status tetap yang lama
   - kembalikan data baru

2. store.go: method hapusPeminjaman(id) (DELETE)
   - cari indeks, kalau tidak ada kembalikan false
   - buang dari slice (urutan sisanya tetap id naik)
   - counter id tidak diubah, id yang sudah dihapus tidak dipakai ulang

3. peminjaman.go: handler ubahPeminjaman (PUT /peminjaman/{id})
   - urutan cek:
     a. id bukan angka -> 400
     b. peminjaman tidak ada -> 404 (cek ini DULU sebelum baca body)
     c. body rusak / validasi gagal -> 400
     d. bentrok -> 409
     e. sukses -> 200 + objek baru

4. peminjaman.go: handler hapusPeminjaman (DELETE /peminjaman/{id})
   - id bukan angka -> 400
   - tidak ada -> 404
   - sukses -> WriteHeader(204) TANPA body (jangan pakai writeJSON)

5. handler.go
   - daftarkan rute:
     PUT /peminjaman/{id}
     DELETE /peminjaman/{id}
```

- Dipakai: store.go, peminjaman.go, handler.go
- Diubah: terlampir di todo

## 7 (B7: Permintaan bersamaan)
- Prompt: kerjakan B7 dengan todo sebagai berikut

```txt
todo:
1. store.go: tambah sync.Mutex di struct Store

2. pisahkan versi "tanpa lock" dan "dengan lock"
   - fungsi internal berakhiran Locked (cariSDLocked, indeksPJLocked, bentrokLocked, terpakaiLocked)
     tidak boleh memanggil Lock() sendiri
   - method yang dipakai handler memegang lock di awal
   - method yang sudah memegang lock tidak boleh memanggil method lain yang mengunci lagi (deadlock)

3. kunci semua method yang menyentuh data
   - pakai s.mu.Lock() lalu langsung defer s.mu.Unlock() di baris berikutnya
   - termasuk method yang hanya membaca (ambil satu, daftar, detail)
   - daftar mengembalikan salinan slice, jangan kembalikan slice internal ke luar lock

4. gabungkan cek + simpan dalam SATU kunci
   - buatPeminjaman di store: Lock -> validasi (butuh cari sumber daya) -> cek bentrok -> beri id -> simpan -> Unlock
   - ubahPeminjaman di store: Lock -> cari -> validasi -> cek bentrok -> ganti -> Unlock
   - handler cukup memanggil satu method store ini, tidak memanggil cek bentrok dan simpan terpisah

5. rapikan handler
   - handler POST dan PUT hanya decode body, panggil method store, lalu terjemahkan hasilnya
     (kode 400 / 404 / 409 / sukses) ke respons
```

- Dipakai: store.go, bentrok.go, validasi.go, peminjaman.go
- Diubah: terlampir di todo