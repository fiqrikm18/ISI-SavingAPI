package models

type Mutation struct {
	ID        int     `db:"id" json:"id"`
	AccountID int     `db:"account_id" json:"account_id"`
	Type      string  `db:"type" json:"type"`
	Amount    float64 `db:"amount" json:"amount"`
	CreatedAt string  `db:"created_at" json:"created_at"`
	UpdatedAt string  `db:"updated_at" json:"updated_at"`
	DeletedAt string  `db:"deleted_at" json:"deleted_at"`
}
