package models

import "time"

type Users struct {
	Id                  int        `json:"id"`
	Nama                string     `json:"nama"`
	Username            string     `json:"username"`
	Password            string     `json:"password"`
	Jenis               int        `json:"jenis"`
	JenisUser           string     `json:"jenisUser"`
	IdCorporate         int        `json:"idCorporate"`
	NamaCorporate       string     `json:"namaCorporate"`
	Email               string     `json:"email"`
	EmailVerificationAt *time.Time `json:"emailVerificationAt"`
	RememberToken       *string    `json:"rememberToken"`
	Created_at          time.Time  `json:"createdAt"`
	Updated_at          time.Time  `json:"updatedAt"`
	Deleted_at          time.Time  `json:"deletedAt"`
	EncCredentials      string     `json:"encCredentials"`
	Role                string     `json:"role"`
}

type RequestUpdateUser struct {
	Id             int    `json:"id" validate:"required"`
	Nama           string `json:"nama" validate:"required"`
	Username       string `json:"username" validate:"required,min=8"`
	Jenis          int    `json:"jenis" validate:"required"`
	IdCorporate    int    `json:"idCorporate" validate:"required"`
	Email          string `json:"email" validate:"required,email"`
	EncCredentials string `json:"encCredentials"`
	IdRole         int    `json:"idRole"`
}

type RequestUpdatePassword struct {
	Id       int    `json:"id" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
}

type RequestUpdateEnc struct {
	Id  int    `json:"id" validate:"required"`
	Enc string `json:"enc"`
}

type RequestAddUser struct {
	Nama           string `json:"nama" validate:"required"`
	Username       string `json:"username" validate:"required,min=8"`
	Password       string `json:"password" validate:"required,min=8"`
	Jenis          int    `json:"jenis" validate:"required"`
	IdCorporate    int    `json:"idCorporate" validate:"required"`
	Email          string `json:"email" validate:"required,email"`
	EncCredentials string `json:"encCredentials"`
	IdRole         int    `json:"idRole"`
}

type RequestDeleteUser struct {
	Id int `json:"id" validate:"required"`
}

type ResponseUser struct {
	Id             int     `json:"id"`
	Nama           string  `json:"nama"`
	Username       string  `json:"username"`
	IdJenis        int     `json:"idJenis"`
	JenisUser      string  `json:"jenisUser"`
	IdCorporate    int     `json:"idCorporate"`
	NamaCorporate  string  `json:"namaCorporate"`
	Email          string  `json:"email"`
	EncCredentials string  `json:"encCredentials"`
	IdRole         *int    `json:"idRole"`
	Role           *string `json:"role"`
}

type ResponseJenisUser struct {
	Id        int    `json:"id"`
	JenisUser string `json:"jenisUser"`
}
