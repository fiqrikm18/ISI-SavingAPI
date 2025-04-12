package models

import "time"

type Account struct {
	ID            int        `db:"id" json:"id"`
	AccountNumber string     `db:"account_number" json:"account_number"`
	Name          string     `db:"name" json:"name"`
	NIK           string     `db:"nik" json:"nik"`
	Phone         string     `db:"phone" json:"phone"`
	Balance       float64    `db:"balance" json:"balance"`
	CreatedAt     time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt     *time.Time `db:"deleted_at" json:"deleted_at"`
}
