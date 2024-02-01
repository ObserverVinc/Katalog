package models

type Katalog struct {
	Id_aplikasi   uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Id_kategori   string `json:"id_katalog"`
	Nama_aplikasi string `json:"nama_aplikasi"`
	Deskripsi     string `json:"deskripsi"`
	Link          string `json:"link"`
	Gambar        string `json:"gambar"`
}

type Kategori struct {
	Id_kategori string `json:"id_katalog"`
	Id_aplikasi uint   `json:"id"`
	Deskripsi_K string `json:"deskripsiK"`
}

//halo vincent saya maman, kodingan anda saya hek
