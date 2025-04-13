package services

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/fiqrikm18/ISITransaction/internal/dto"
	"github.com/fiqrikm18/ISITransaction/internal/repositories"
)

type IAccountService interface {
	RegisterAccount(dto *dto.RegisterAccountRequest) error
	GetAccountByNik(nik string) (*dto.RegisterAccountResponse, error)
	GetAccountByAccountNumber(accountNumber string) (*dto.AccountDetailResponse, error)
	CheckBalance(accountNumber string) (float64, error)
	DepositBalance(accountNumber string, amount float64) error
	WithdrawBalance(accountNumber string, amount float64) error
}

type AccountService struct {
	accountRepository  repositories.IAccountRepository
	mutationRepository repositories.IMutationRepository
}

func NewAccountService(
	accountRepository repositories.IAccountRepository,
	mutationRepository repositories.IMutationRepository,
) *AccountService {
	return &AccountService{
		accountRepository:  accountRepository,
		mutationRepository: mutationRepository,
	}
}

func (accountService *AccountService) RegisterAccount(dto *dto.RegisterAccountRequest) error {
	// Validate the request
	if dto.Nama == "" {
		return fmt.Errorf("name is required")
	}

	if dto.Nik == "" {
		return fmt.Errorf("NIK is required")
	}

	if dto.PhoneNumber == "" {
		return fmt.Errorf("phone number is required")
	}

	// Check if the account already exists
	existingAccount, _ := accountService.accountRepository.FindAccountByNIK(dto.Nik)
	if existingAccount != nil {
		return fmt.Errorf("account with NIK already exists")
	}

	// Check if the phone number already exists
	existingAccount, _ = accountService.accountRepository.FindAccountByPhone(dto.PhoneNumber)
	if existingAccount != nil {
		return fmt.Errorf("account with phone number already exists")
	}

	// Create a new account
	accountNumber := accountService.generateAccountNumber()

	err := accountService.accountRepository.CreateAccount(accountNumber, dto.Nama, dto.Nik, dto.PhoneNumber)
	if err != nil {
		return err
	}

	return nil
}

func (accountService *AccountService) generateAccountNumber() string {
	rand.Seed(time.Now().UnixNano())
	accountNumber := ""

	for i := 0; i < 12; i++ {
		digit := rand.Intn(10) // 0-9
		accountNumber += strconv.Itoa(digit)
	}

	return accountNumber
}

func (accountService *AccountService) GetAccountByNik(nik string) (*dto.RegisterAccountResponse, error) {
	account, err := accountService.accountRepository.FindAccountByNIK(nik)
	if err != nil {
		return nil, err
	}

	if account == nil {
		return nil, fmt.Errorf("account not found")
	}

	return &dto.RegisterAccountResponse{
		NoRekening: account.AccountNumber,
	}, nil
}

func (accountService *AccountService) CheckBalance(accountNumber string) (float64, error) {
	account, err := accountService.accountRepository.FindAccountByAccountNumber(accountNumber)
	if err != nil {
		return 0, err
	}

	return account.Balance, nil
}

func (accountService *AccountService) DepositBalance(accountNumber string, amount float64) error {
	account, err := accountService.accountRepository.FindAccountByAccountNumber(accountNumber)
	if err != nil {
		return err
	}

	if account == nil {
		return fmt.Errorf("account not found")
	}

	// Update the balance
	account.Balance += amount

	err = accountService.accountRepository.UpdateAccountBalance(account)
	if err != nil {
		return err
	}

	// Create a mutation record
	err = accountService.mutationRepository.CreateMutation(account.ID, amount, "deposit")
	if err != nil {
		return err
	}

	return nil
}

func (accountService *AccountService) WithdrawBalance(accountNumber string, amount float64) error {
	account, err := accountService.accountRepository.FindAccountByAccountNumber(accountNumber)
	if err != nil {
		return err
	}

	if account == nil {
		return fmt.Errorf("account not found")
	}

	if account.Balance < amount {
		return fmt.Errorf("insufficient balance")
	}

	// Update the balance
	account.Balance -= amount

	err = accountService.accountRepository.UpdateAccountBalance(account)
	if err != nil {
		return err
	}

	// Create a mutation record
	err = accountService.mutationRepository.CreateMutation(account.ID, -amount, "withdraw")
	if err != nil {
		return err
	}

	return nil
}

func (accountService *AccountService) GetAccountByAccountNumber(accountNumber string) (*dto.AccountDetailResponse, error) {
	account, err := accountService.accountRepository.FindAccountByAccountNumber(accountNumber)
	if err != nil {
		return nil, err
	}

	if account == nil {
		return nil, fmt.Errorf("account not found")
	}

	return &dto.AccountDetailResponse{
		AccountNumber: account.AccountNumber,
		Name:          account.Name,
		Nik:           account.NIK,
		Phone:         account.Phone,
		Balance:       account.Balance,
	}, nil
}
