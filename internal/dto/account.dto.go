package dto

type RegisterAccountRequest struct {
	Nama        string `json:"nama" validate:"required"`
	Nik         string `json:"nik" validate:"required,unique"`
	PhoneNumber string `json:"phone_number" validate:"required,unique"`
}

type RegisterAccountResponse struct {
	NoRekening string `json:"no_rekening"`
}

type AccountDetailResponse struct {
	AccountNumber string  `json:"no_rekening" validate:"required"`
	Name          string  `json:"nama" validate:"required"`
	Nik           string  `json:"nik" validate:"required"`
	Phone         string  `json:"phone" validate:"required"`
	Balance       float64 `json:"saldo" validate:"required"`
}

type TransactionRequest struct {
	NoRekening string  `json:"no_rekening" validate:"required"`
	Amount     float64 `json:"amount" validate:"required"`
}

type TransactionResponse struct {
	Saldo float64 `json:"saldo" validate:"required"`
}
