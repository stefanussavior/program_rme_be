package models

import "github.com/google/uuid"

type User struct {
	Id uuid.UUID `json:"id"`
	NamaLengkap string `json:"nama_lengkap"`
	Alamat string `json:"alamat"`
	Ttl string `json:"ttl"`
}