package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Surat struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	No_surat     int                `bson:"no_surat,omitempty" json:"no_surat,omitempty"`
	Status_surat Status             `bson:"status_surat,omitempty" json:"status_surat,omitempty"`
	Perihal      string             `bson:"perihal,omitempty" json:"perihal,omitempty"`
	Id_pos       Kodepos            `bson:"id_pos,omitempty" json:"id_pos,omitempty"`
	Pengirim_srt Pengirim           `bson:"pengirim_srt,omitempty" json:"pengirim_srt,omitempty"`
	Penerima_srt Penerima           `bson:"penerima_srt,omitempty" json:"penerima_srt,omitempty"`
}

type Disposisi struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Kode_disposisi int                `bson:"kode_disposisi,omitempty" json:"kode_disposisi,omitempty"`
	Tgl_disposisi  string             `bson:"tgl_disposisi,omitempty" json:"tgl_disposisi,omitempty"`
	Penerima_surat Penerima           `bson:"penerima_surat,omitempty" json:"penerima_surat,omitempty"`
	Stat_disposisi Status             `bson:"status_disposisi,omitempty" json:"status_disposisi,omitempty"`
}
type Kodepos struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Kode_daerah int                `bson:"kode_daerah,omitempty" json:"kode_daerah,omitempty"`
	Nama_daerah string             `bson:"nama_daerah,omitempty" json:"nama_dareah,omitempty"`
}

type Status struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Id_status  int                `bson:"id_status,omitempty" json:"id_status,omitempty"`
	Keterangan string             `bson:"keterangan,omitempty" json:"keterangan,omitempty"`
}

type Pengirim struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_pengirim string             `bson:"nama_pengirim,omitempty" json:"nama_pengirim,omitempty"`
	Alamat        string             `bson:"alamat,omitempty" json:"alamat,omitempty"`
	Tgl_kirim     string             `bson:"tgl_kirim,omitempty" json:"tgl_kirim,omitempty"`
}

type Penerima struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_penerima string             `bson:"nama_penerima,omitempty" json:"nama_penerima,omitempty"`
	Alamat        string             `bson:"alamat,omitempty" json:"alamat,omitempty"`
	Tgl_terima    string             `bson:"tgl_terima,omitempty" json:"tgl_terima,omitempty"`
}
