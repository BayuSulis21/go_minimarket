package models

type Barang struct {
	Kode             string `json:"kode"  gorm:"column:kode"`
	Nama_barang      string `json:"nama_barang"  gorm:"column:nama"`
	Kategori         string `json:"kategori"  gorm:"column:kategori"`
	Satuan           string `json:"satuan"  gorm:"column:satuan"`
	Isi              string `json:"isi"  gorm:"column:isi"`
	Hpp              string `json:"hpp"  gorm:"column:hpp"`
	Harga_toko       string `json:"harga_toko"  gorm:"column:harga_toko"`
	Diskon           string `json:"diskon"  gorm:"column:diskon"`
	Supplier         string `json:"supplier"  gorm:"column:supplier"`
	Pajak            string `json:"pajak"  gorm:"column:pajak"`
	Kode_barcode     string `json:"kode_barcode"  gorm:"column:kode_barcode"`
	Gambar           string `json:"gambar"  gorm:"column:gambar"`
	TanggalExpired   string `json:"tgl_expired"  gorm:"column:expired"`
	AdaExpiredDate   string `json:"ada_expired_date"  gorm:"column:ada_expired_date"`
	Paket            string `json:"paket"  gorm:"column:paket"`
	Lokasi           string `json:"lokasi"  gorm:"column:lokasi"`
	Merk             string `json:"merk"  gorm:"column:merk"`
	Ket_tambahan     string `json:"ket_tambahan"  gorm:"column:ket_tambahan"`
	Item_description string `json:"item_description"  gorm:"column:item_description"`
}

type ClientData struct {
	ClientKey string `json:"clientKey"  gorm:"column:clientKey"`
	SecretKey string `json:"secretKey"  gorm:"column:secretKey"`
}

type AccountData struct {
	Username     string `json:"username" gorm:"username"`
	Password     string `json:"password" gorm:"password"`
	Nama_lengkap string `json:"nama_lengkap" gorm:"nama_lengkap"`
	Email        string `json:"email" gorm:"email"`
	Telp         string `json:"telp" gorm:"telp"`
}
