package main

import (
	"strings"
	"time"
	"unicode/utf8"
)

// Semua fungsi mengembalikan pesan error ("" bila valid).
// Semua pelanggaran di file ini = 400.

// validasiTeks: trim dulu, lalu periksa tidak kosong dan maks 50 KARAKTER (rune, bukan byte).
func validasiTeks(label, teks string) (string, string) {
	t := strings.TrimSpace(teks)
	if t == "" {
		return "", label + " wajib diisi"
	}
	if utf8.RuneCountInString(t) > MaxPanjangTeks {
		return "", label + " maksimal 50 karakter"
	}
	return t, ""
}

func validasiSumberDaya(in SumberDayaInput) (SumberDaya, string) {
	if in.Nama == nil || in.Jenis == nil || in.Stok == nil {
		return SumberDaya{}, "nama, jenis, dan stok wajib diisi"
	}
	nama, pesan := validasiTeks("nama", *in.Nama)
	if pesan != "" {
		return SumberDaya{}, pesan
	}
	jenis := strings.TrimSpace(*in.Jenis)
	if jenis != JenisRuang && jenis != JenisAlat {
		return SumberDaya{}, `jenis harus "ruang" atau "alat"`
	}
	if *in.Stok < 1 {
		return SumberDaya{}, "stok minimal 1"
	}
	return SumberDaya{Nama: nama, Jenis: jenis, Stok: *in.Stok}, ""
}

// validasiTanggal: format YYYY-MM-DD dan tanggalnya benar-benar ada.
func validasiTanggal(teks string) (string, string) {
	t := strings.TrimSpace(teks)
	if len(t) != 10 {
		return "", "tanggal harus berformat YYYY-MM-DD"
	}
	if _, err := time.Parse("2006-01-02", t); err != nil {
		return "", "tanggal tidak valid, gunakan YYYY-MM-DD dan tanggal yang ada"
	}
	return t, ""
}

// validasiSlot: 1 sampai 8 dan mulai <= selesai (mulai == selesai sah).
func validasiSlot(mulai, selesai int) string {
	if mulai < SlotMin || mulai > SlotMax || selesai < SlotMin || selesai > SlotMax {
		return "slot harus antara 1 dan 8"
	}
	if mulai > selesai {
		return "slot_mulai tidak boleh lebih besar dari slot_selesai"
	}
	return ""
}

// validasiPeminjaman dipakai POST dan PUT. cari = pencari sumber daya.
func validasiPeminjaman(in PeminjamanInput, cari func(int) (SumberDaya, bool)) (Peminjaman, string) {
	if in.SumberDayaID == nil || in.Peminjam == nil || in.Tanggal == nil ||
		in.SlotMulai == nil || in.SlotSelesai == nil || in.Jumlah == nil {
		return Peminjaman{}, "semua field wajib diisi"
	}
	peminjam, pesan := validasiTeks("peminjam", *in.Peminjam)
	if pesan != "" {
		return Peminjaman{}, pesan
	}
	tanggal, pesan := validasiTanggal(*in.Tanggal)
	if pesan != "" {
		return Peminjaman{}, pesan
	}
	if pesan := validasiSlot(*in.SlotMulai, *in.SlotSelesai); pesan != "" {
		return Peminjaman{}, pesan
	}
	sd, ada := cari(*in.SumberDayaID)
	if !ada {
		return Peminjaman{}, "sumber_daya_id tidak ditemukan" // 400, bukan 404
	}
	if *in.Jumlah < 1 || *in.Jumlah > sd.Stok {
		return Peminjaman{}, "jumlah harus antara 1 dan stok sumber daya"
	}
	return Peminjaman{
		SumberDayaID: sd.ID,
		Peminjam:     peminjam,
		Tanggal:      tanggal,
		SlotMulai:    *in.SlotMulai,
		SlotSelesai:  *in.SlotSelesai,
		Jumlah:       *in.Jumlah,
	}, ""
}
