package models

type Produk struct {
	Id                 int     `json:"id"`
	KodeProduk         string  `json:"kodeProduk"`
	Nama               string  `json:"nama"`
	HargaJual          int     `json:"hargaJual"`
	Gambar             *string `json:"gambar"`
	Jenis              *bool   `json:"jenis"`
	IdCorporate        int     `json:"idCorporate"`
	NamaCorporate      string  `json:"namaCorporate"`
	IdSatuanProduk     int     `json:"idSatuanProduk"`
	NamaSatuanProduk   string  `json:"namaSatuanProduk"`
	IdKategoriProduk   int     `json:"idKategoriProduk"`
	NamaKategoriProduk string  `json:"namaKategoriProduk"`
	HargaRombongan     *int    `json:"hargaRombongan"`
	MinRombongan       *int    `json:"minRombongan"`
	StatusStok         *bool   `json:"statusStok"`
	CurrentStok        *int    `json:"currentStok"`
	StatusPaket        *bool   `json:"statusPaket"`
}

type RequestAddProduk struct {
	KodeProduk       string  `json:"kodeProduk" validate:"required"`
	Nama             string  `json:"nama" validate:"required"`
	HargaJual        int     `json:"hargaJual" validate:"required"`
	Gambar           *string `json:"gambar"`
	Jenis            *bool   `json:"jenis"`
	IdCorporate      int     `json:"idCorporate" validate:"required"`
	IdSatuanProduk   int     `json:"idSatuanProduk" validate:"required"`
	IdKategoriProduk int     `json:"idKategoriProduk" validate:"required"`
	HargaRombongan   *int    `json:"hargaRombongan"`
	MinRombongan     *int    `json:"minRombongan"`
	StatusStok       *bool   `json:"statusStok"`
	CurrentStok      *int    `json:"currentStok"`
	StatusPaket      *bool   `json:"statusPaket"`
	CreatedAt        string  `json:"createdAt"`
}

type RequestUpdateProduk struct {
	Id               int     `json:"id" validate:"required"`
	KodeProduk       string  `json:"kodeProduk" validate:"required"`
	Nama             string  `json:"nama" validate:"required"`
	HargaJual        int     `json:"hargaJual" validate:"required"`
	Gambar           *string `json:"gambar"`
	Jenis            *bool   `json:"jenis"`
	IdCorporate      int     `json:"idCorporate" validate:"required"`
	IdSatuanProduk   int     `json:"idSatuanProduk" validate:"required"`
	IdKategoriProduk int     `json:"idKategoriProduk" validate:"required"`
	HargaRombongan   *int    `json:"hargaRombongan"`
	MinRombongan     *int    `json:"minRombongan"`
	StatusStok       *bool   `json:"statusStok"`
	CurrentStok      *int    `json:"currentStok"`
	StatusPaket      *bool   `json:"statusPaket"`
	UpdatedAt        string  `json:"updatedAt"`
}

type RequestDeleteProduk struct {
	Id        int    `json:"id" validate:"required"`
	DeletedAt string `json:"deletedAt"`
}

type ProdukArray struct {
	ProdukArray []Produk `json:"produk_array"`
}
