package Menyurat

import (
	"fmt"
	"testing"

	"github.com/AdeCandra12/surat/model"
	"github.com/AdeCandra12/surat/module"
)

func TestInsertSurat(t *testing.T) {
	no_surat := 96785
	status_surat := model.Status{
		Id_status:  82,
		Keterangan: "Terkirim",
	}
	perihal := "surat Pindah"
	id_pos := model.Kodepos{
		Kode_daerah: 22337,
		Nama_daerah: "Jawa Tengah",
	}
	pengirim_srt := model.Pengirim{
		Nama_pengirim: "Dito",
		Alamat:        "Jawa Tengah",
		Tgl_kirim:     "21 Januari 2023",
	}
	penerima_srt := model.Penerima{
		Nama_penerima: "Gani",
		Alamat:        "Kalimantan Timur",
		Tgl_terima:    "21 februari 2023",
	}

	hasil := module.InsertSurat(module.MongoConn, "surat", no_surat, status_surat, perihal, id_pos, pengirim_srt, penerima_srt)
	fmt.Println(hasil)
}

func TestInsertDisposisi(t *testing.T) {
	kode_disposisi := 10314
	tgl_disposisi := "20 Februari 2021"
	penerima_surat := model.Penerima{
		Nama_penerima: "Gani",
		Alamat:        "Kalimantan Timur",
		Tgl_terima:    "21 februari 2023",
	}
	stat_disposisi := model.Status{
		Id_status:  82,
		Keterangan: "Terkirim",
	}

	hasil := module.InsertDisposisi(module.MongoConn, "disposisi", kode_disposisi, tgl_disposisi, penerima_surat, stat_disposisi)
	fmt.Println(hasil)
}

func TestInsertStatus(t *testing.T) {
	id_status := 82
	keterangan := "Terkirim"

	hasil := module.InsertStatus(module.MongoConn, "status", id_status, keterangan)
	fmt.Println(hasil)
}

func TestInsertPengirim(t *testing.T) {
	nama_pengirim := "Dito"
	alamat := "Jawa Tengah"
	tgl_kirim := "21 Januari 2023"

	hasil := module.InsertPengirim(module.MongoConn, "pengirim", nama_pengirim, alamat, tgl_kirim)
	fmt.Println(hasil)
}

func TestInsertPenerima(t *testing.T) {
	nama_penerima := "Gani"
	alamat := "Kalimantan Timur"
	tgl_terima := "21 Februari 2023"

	hasil := module.InsertPenerima(module.MongoConn, "penerima", nama_penerima, alamat, tgl_terima)
	fmt.Println(hasil)
}

// test getFunction

func TestGetSuratFromNoSurat(t *testing.T) {
	NoSurat := 12143113
	surat := module.GetSuratFromNoSurat(NoSurat, module.MongoConn, "surat")
	fmt.Println(surat)
}

func TestGetDisposisiFromKodeDisposisi(t *testing.T) {
	kode_disposisi := 10244
	disposisi := module.GetDisposisiFromTglDisposisi(kode_disposisi, module.MongoConn, "disposisi")
	fmt.Println(disposisi)
}

func TestGetStatusFromIdStatus(t *testing.T) {
	id_status := 1326
	status := module.GetStatusFromIdStatus(id_status, module.MongoConn, "status")
	fmt.Println(status)
}
func TestGetPengirimFromNamaPengirim(t *testing.T) {
	nama_pengirim := "Ucok"
	pengirim := module.GetPengirimFromNamaPengirim(nama_pengirim, module.MongoConn, "nama_pengirim")
	fmt.Println(pengirim)
}

func TestGetPenerimaFromNamaPenerima(t *testing.T) {
	nama_penerima := "Dani"
	penerima := module.GetPenerimaFromNamaPenerima(nama_penerima, module.MongoConn, "nama_penerima")
	fmt.Println(penerima)
}

//test getFunctionAll

func TestGetAllSurat(t *testing.T) {
	data := module.GetAllSurat(module.MongoConn, "Surat")
	fmt.Println(data)
}
func TestGetAllDisposisi(t *testing.T) {
	data := module.GetAllDisposisi(module.MongoConn, "disposisi")
	fmt.Println(data)
}
func TestGetAllStatus(t *testing.T) {
	data := module.GetAllStatus(module.MongoConn, "Status")
	fmt.Println(data)
}
func TestGetAllPengirim(t *testing.T) {
	data := module.GetAllPengirim(module.MongoConn, "pengirim")
	fmt.Println(data)
}
func TestGetAllPenerima(t *testing.T) {
	data := module.GetAllPenerima(module.MongoConn, "penerima")
	fmt.Println(data)
}
