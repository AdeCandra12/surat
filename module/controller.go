package module

import (
	"context"
	"fmt"
	"os"

	"github.com/AdeCandra12/surat/model"
	"github.com/aiteung/atdb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var MongoString string = os.Getenv("MONGOSTRING")

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "Surat_db",
}

var MongoConn = atdb.MongoConnect(MongoInfo)

func InsertOneDoc(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

// functionInsert

func InsertSurat(db *mongo.Database, col string, no_surat int, status_surat model.Status, Perihal string, id_pos model.Kodepos, pengirim_srt model.Pengirim, penerima_srt model.Penerima) (InsertedID interface{}) {
	var surat model.Surat
	surat.No_surat = no_surat
	surat.Status_surat = status_surat
	surat.Perihal = Perihal
	surat.Id_pos = id_pos
	surat.Pengirim_srt = pengirim_srt
	surat.Penerima_srt = penerima_srt
	return InsertOneDoc(db, col, surat)
}

func InsertDisposisi(db *mongo.Database, col string, kode_disposisi int, tgl_disposisi string, penerima_surat model.Penerima, stat_disposisi model.Status) (InsertedID interface{}) {
	var disposisi model.Disposisi
	disposisi.Kode_disposisi = kode_disposisi
	disposisi.Tgl_disposisi = tgl_disposisi
	disposisi.Penerima_surat = penerima_surat
	disposisi.Stat_disposisi = stat_disposisi
	return InsertOneDoc(db, col, disposisi)
}

func InsertKodePos(db *mongo.Database, col string, kode_daerah int, nama_daerah string) (InsertedID interface{}) {
	var kodepos model.Kodepos
	kodepos.Kode_daerah = kode_daerah
	kodepos.Nama_daerah = nama_daerah
	return InsertOneDoc(db, col, kodepos)
}

func InsertStatus(db *mongo.Database, col string, id_status int, keterangan string) (InsertedID interface{}) {
	var status model.Status
	status.Id_status = id_status
	status.Keterangan = keterangan
	return InsertOneDoc(db, col, status)
}

func InsertPengirim(db *mongo.Database, col string, nama_pengirim string, alamat string, tgl_kirim string) (InsertedID interface{}) {
	var pengirim model.Pengirim
	pengirim.Nama_pengirim = nama_pengirim
	pengirim.Alamat = alamat
	pengirim.Tgl_kirim = tgl_kirim
	return InsertOneDoc(db, col, pengirim)
}

func InsertPenerima(db *mongo.Database, col string, nama_penerima string, alamat string, tgl_terima string) (InsertedID interface{}) {
	var penerima model.Penerima
	penerima.Nama_penerima = nama_penerima
	penerima.Alamat = alamat
	penerima.Tgl_terima = tgl_terima
	return InsertOneDoc(db, col, penerima)
}

func GetSuratFromNoSurat(No_surat int, db *mongo.Database, col string) (sm model.Surat) {
	surat := db.Collection(col)
	filter := bson.M{"no_surat": No_surat}
	err := surat.FindOne(context.TODO(), filter).Decode(&sm)
	if err != nil {
		fmt.Printf("getSuratFromNoSurat: %v\n", err)
	}
	return sm
}

func GetDisposisiFromTglDisposisi(tgl_disposisi int, db *mongo.Database, col string) (sm model.Disposisi) {
	disposisi := db.Collection(col)
	filter := bson.M{"tgl_disposisi": tgl_disposisi}
	err := disposisi.FindOne(context.TODO(), filter).Decode(&sm)
	if err != nil {
		fmt.Printf("getDisposisiFromTglSurat: %v\n", err)
	}
	return sm
}

func GetStatusFromIdStatus(id_status int, db *mongo.Database, col string) (sm model.Disposisi) {
	status := db.Collection(col)
	filter := bson.M{"id_status": id_status}
	err := status.FindOne(context.TODO(), filter).Decode(&sm)
	if err != nil {
		fmt.Printf("getStatusFromIdStatus: %v\n", err)
	}
	return sm
}

func GetPengirimFromNamaPengirim(nama_pengirim string, db *mongo.Database, col string) (sm model.Pengirim) {
	pengirim := db.Collection(col)
	filter := bson.M{"nama_penerima": nama_pengirim}
	err := pengirim.FindOne(context.TODO(), filter).Decode(&sm)
	if err != nil {
		fmt.Printf("getPenerimaFromNamaPengirim: %v\n", err)
	}
	return sm
}

func GetPenerimaFromNamaPenerima(nama_penerima string, db *mongo.Database, col string) (sm model.Penerima) {
	penerima := db.Collection(col)
	filter := bson.M{"nama_penerima": nama_penerima}
	err := penerima.FindOne(context.TODO(), filter).Decode(&sm)
	if err != nil {
		fmt.Printf("getPenerimaFromNamaPenerima: %v\n", err)
	}
	return sm
}

//GetFunctionAll

func GetAllSurat(db *mongo.Database, col string) (surat []model.Surat) {
	sm_surat := db.Collection(col)
	filter := bson.M{}
	cursor, err := sm_surat.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &surat)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetAllDisposisi(db *mongo.Database, col string) (disposisi []model.Disposisi) {
	sm_disposisi := db.Collection(col)
	filter := bson.M{}
	cursor, err := sm_disposisi.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &disposisi)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetAllStatus(db *mongo.Database, col string) (status []model.Status) {
	sm_status := db.Collection(col)
	filter := bson.M{}
	cursor, err := sm_status.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &status)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetAllPengirim(db *mongo.Database, col string) (pengirim []model.Pengirim) {
	sm_pengirim := db.Collection(col)
	filter := bson.M{}
	cursor, err := sm_pengirim.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &pengirim)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetAllPenerima(db *mongo.Database, col string) (penerima []model.Penerima) {
	sm_penerima := db.Collection(col)
	filter := bson.M{}
	cursor, err := sm_penerima.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &penerima)
	if err != nil {
		fmt.Println(err)
	}
	return
}
