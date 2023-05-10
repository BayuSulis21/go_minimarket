package models

type ParamListBarang struct {
	Kategori    string `json:"kategori"`
	Nama_barang string `json:"nama"`
}

type ParamSignature struct {
	Secret string `json:"secretKey"`
	Client string `json:"clientKey"`
}

type ParamRegister struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Nama_lengkap string `json:"nama_lengkap"`
	Email        string `json:"email"`
	Telp         string `json:"telp"`
}

type ParamSendWA struct {
	Target  string `json:"target"`
	Message string `json:"message"`
}
