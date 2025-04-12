package repositories

import (
	"time"

	"github.com/fiqrikm18/ISITransaction/internal/configs"
	"github.com/fiqrikm18/ISITransaction/internal/models"
)

type IAccountRepository interface {
	CreateAccount(accountNumber, name, nik, phone string) error

	FindAccountByAccountNumber(accountNumber string) (*models.Account, error)
	FindAccountByNIK(nik string) (*models.Account, error)
	FindAccountByPhone(phone string) (*models.Account, error)
	FindAccountByID(id int) (*models.Account, error)
	FindAccountByName(name string) (*models.Account, error)

	UpdateAccountBalance(account *models.Account) error
}

type AccountRepository struct {
	dbConf *configs.DatabaseConnection
}

func NewAccountRepository() *AccountRepository {
	dbConfig := configs.NewDatabaseConnection()

	return &AccountRepository{
		dbConf: dbConfig,
	}
}

func (repository *AccountRepository) CreateAccount(accountNumber, name, nik, phone string) error {
	qs := `INSERT INTO accounts (account_number, name, nik, phone, balance, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := repository.dbConf.DB.Exec(qs, accountNumber, name, nik, phone, 0, time.Now(), time.Now())
	if err != nil {
		return err
	}

	return nil
}

func (repository *AccountRepository) FindAccountByAccountNumber(accountNumber string) (*models.Account, error) {
	qs := `SELECT * FROM accounts WHERE account_number = $1`
	var account models.Account
	err := repository.dbConf.DB.Get(&account, qs, accountNumber)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (repository *AccountRepository) FindAccountByNIK(nik string) (*models.Account, error) {
	qs := `SELECT * FROM accounts WHERE nik = $1`
	var account models.Account
	err := repository.dbConf.DB.Get(&account, qs, nik)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (repository *AccountRepository) FindAccountByPhone(phone string) (*models.Account, error) {
	qs := `SELECT * FROM accounts WHERE phone = $1`
	var account models.Account
	err := repository.dbConf.DB.Get(&account, qs, phone)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (repository *AccountRepository) FindAccountByID(id int) (*models.Account, error) {
	qs := `SELECT * FROM accounts WHERE id = $1`
	var account models.Account
	err := repository.dbConf.DB.Get(&account, qs, id)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (repository *AccountRepository) FindAccountByName(name string) (*models.Account, error) {
	qs := `SELECT * FROM accounts WHERE name = $1`
	var account models.Account
	err := repository.dbConf.DB.Get(&account, qs, name)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (repository *AccountRepository) UpdateAccountBalance(account *models.Account) error {
	qs := `UPDATE accounts SET balance = $1, updated_at = $2 WHERE id = $3`
	_, err := repository.dbConf.DB.Exec(qs, account.Balance, time.Now(), account.ID)
	if err != nil {
		return err
	}

	return nil
}
