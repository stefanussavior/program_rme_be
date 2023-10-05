package models

import "github.com/google/uuid"

type Users struct {
	Id       uuid.UUID `json:"id"`
	Nama     string    `json:"nama"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Alamat   string    `json:"alamat"`
	NoTelp   int64     `json:"no_telp"`
}
