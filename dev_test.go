package Menyurat

import (
	"fmt"
	"testing"

	"github.com/AdeCandra12/surat/model"
	"github.com/AdeCandra12/surat/module"
)

func TestInsertSurat(t *testing.T) {
	no_surat := 10234
	status_surat := model.Status{
		Id_status:  10,
		Keterangan: "Tidak Terkirim",
	}
	perihal := "surat Dispensasi"
	id_pos := model.Kodepos{
		Kode_daerah: 10887,
		Nama_daerah: "Kalimantan Barat",
	}
	pengirim_srt := model.Pengirim{
		Nama_pengirim: "Gino",
		Alamat:        "Kalimantan Barat",
		Tgl_kirim:     "6 Januari 2023",
	}
	penerima_srt := model.Penerima{
		Nama_penerima: "Tian",
		Alamat:        "Kalimantan Timur",
		Tgl_terima:    "16 februari 2023",
	}

	hasil := module.InsertSurat(module.MongoConn, "surat", no_surat, status_surat, perihal, id_pos, pengirim_srt, penerima_srt)
	fmt.Println(hasil)
}

func TestInsertDisposisi(t *testing.T) {
	kode_disposisi := 10814
	tgl_disposisi := "15 Februari 2021"
	penerima_surat := model.Penerima{
		Nama_penerima: "Tian",
		Alamat:        "Kalimantan Timur",
		Tgl_terima:    "16 februari 2023",
	}
	stat_disposisi := model.Status{
		Id_status:  10,
		Keterangan: "Tidak Terkirim",
	}

	hasil := module.InsertDisposisi(module.MongoConn, "disposisi", kode_disposisi, tgl_disposisi, penerima_surat, stat_disposisi)
	fmt.Println(hasil)
}

func TestInsertStatus(t *testing.T) {
	id_status := 107
	keterangan := "Tidak Terkirim"

	hasil := module.InsertStatus(module.MongoConn, "status", id_status, keterangan)
	fmt.Println(hasil)
}

func TestInsertPengirim(t *testing.T) {
	nama_pengirim := "Gino"
	alamat := "Kalimantan Barat"
	tgl_kirim := "6 Januari 2023"

	hasil := module.InsertPengirim(module.MongoConn, "pengirim", nama_pengirim, alamat, tgl_kirim)
	fmt.Println(hasil)
}

func TestInsertPenerima(t *testing.T) {
	nama_penerima := "Tian"
	alamat := "Kalimantan Timur"
	tgl_terima := "16 Februari 2023"

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
	kode_disposisi := 10814
	disposisi := module.GetDisposisiFromTglDisposisi(kode_disposisi, module.MongoConn, "disposisi")
	fmt.Println(disposisi)
}

func TestGetStatusFromIdStatus(t *testing.T) {
	id_status := 82
	status := module.GetStatusFromIdStatus(id_status, module.MongoConn, "status")
	fmt.Println(status)
}
func TestGetPengirimFromNamaPengirim(t *testing.T) {
	nama_pengirim := "Gisa"
	pengirim := module.GetPengirimFromNamaPengirim(nama_pengirim, module.MongoConn, "pengirim")
	fmt.Println(pengirim)
}

func TestGetPenerimaFromNamaPenerima(t *testing.T) {
	nama_penerima := "Tian"
	penerima := module.GetPenerimaFromNamaPenerima(nama_penerima, module.MongoConn, "penerima")
	fmt.Println(penerima)
}

//test getFunctionAll

func TestGetAllSurat(t *testing.T) {
	data := module.GetAllSurat(module.MongoConn, "surat")
	fmt.Println(data)
}
func TestGetAllDisposisi(t *testing.T) {
	data := module.GetAllDisposisi(module.MongoConn, "disposisi")
	fmt.Println(data)
}
func TestGetAllStatus(t *testing.T) {
	data := module.GetAllStatus(module.MongoConn, "status")
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
