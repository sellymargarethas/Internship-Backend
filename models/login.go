package models

import "github.com/dgrijalva/jwt-go"

type RequestLogin struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Jenis    int    `json:"jenis"`
}

type LoginResponse struct {
	Id        int           `json:"id"`
	Nama      string        `json:"nama"`
	Token     string        `json:"token"`
	HirarkiId string        `json:"hirarkiId"`
	Role      string        `json:"role"`
	Privilege []RoleDetails `json:"privilege"`
}

type JWTClaim struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}
