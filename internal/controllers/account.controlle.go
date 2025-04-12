package controllers

import (
	"github.com/fiqrikm18/ISITransaction/internal/dto"
	"github.com/fiqrikm18/ISITransaction/internal/services"
	"github.com/gofiber/fiber/v2"
)

type IAccountController interface {
	RegisterAccount(c *fiber.Ctx) error
	DepositBalance(c *fiber.Ctx) error
	WithdrawBalance(c *fiber.Ctx) error
	CheckBalance(c *fiber.Ctx) error
}

type AccountController struct {
	accountService services.IAccountService
}

func NewAccountController(
	accountService services.IAccountService,
) *AccountController {
	return &AccountController{
		accountService: accountService,
	}
}

// RegisterAccount create new account
//
// @Summary      Create new account
// @Description  Create new account
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param request body dto.RegisterAccountRequest true "request body"
// @Success      200  {object}  dto.RegisterAccountResponse
// @Success      400  {object}  dto.ErrorResponse
// @Router       /account/daftar [post]
func (controller *AccountController) RegisterAccount(c *fiber.Ctx) error {
	var request dto.RegisterAccountRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&dto.ErrorResponse{
			Remark: err,
		})
	}

	if err := controller.accountService.RegisterAccount(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&dto.ErrorResponse{
			Remark: err.Error(),
		})
	}

	account, err := controller.accountService.GetAccountByNik(request.Nik)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&dto.ErrorResponse{
			Remark: err.Error(),
		})
	}

	if account == nil {
		return c.Status(fiber.StatusBadRequest).JSON(&dto.ErrorResponse{
			Remark: "Account not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(&dto.RegisterAccountResponse{
		NoRekening: account.NoRekening,
	})
}

// DepositBalance add balance to account
//
// @Summary      Add balance to account
// @Description  Add balance to account
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param request body dto.TransactionRequest true "request body"
// @Success      200  {object}  dto.TransactionResponse
// @Success      400  {object}  dto.ErrorResponse
// @Router       /account/tabung [post]
func (controller *AccountController) DepositBalance(c *fiber.Ctx) error {
	var request dto.TransactionRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&dto.ErrorResponse{
			Remark: err,
		})
	}

	if err := controller.accountService.DepositBalance(request.NoRekening, request.Amount); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&dto.ErrorResponse{
			Remark: err.Error(),
		})
	}

	account, err := controller.accountService.GetAccountByAccountNumber(request.NoRekening)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&dto.ErrorResponse{
			Remark: err.Error(),
		})
	}

	if account == nil {
		return c.Status(fiber.StatusBadRequest).JSON(&dto.ErrorResponse{
			Remark: "Account not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(&dto.TransactionResponse{
		Saldo: account.Balance,
	})
}

// WithdrawBalance withdraw balance from account
//
// @Summary      Withdraw balance from account
// @Description  Withdraw balance from account
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param request body dto.TransactionRequest true "request body"
// @Success      200  {object}  dto.TransactionResponse
// @Success      400  {object}  dto.ErrorResponse
// @Router       /account/tarik [post]
func (controller *AccountController) WithdrawBalance(c *fiber.Ctx) error {
	var request dto.TransactionRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&dto.ErrorResponse{
			Remark: err,
		})
	}
	if err := controller.accountService.WithdrawBalance(request.NoRekening, request.Amount); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&dto.ErrorResponse{
			Remark: err.Error(),
		})
	}

	account, err := controller.accountService.GetAccountByAccountNumber(request.NoRekening)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&dto.ErrorResponse{
			Remark: err.Error(),
		})
	}

	if account == nil {
		return c.Status(fiber.StatusBadRequest).JSON(&dto.ErrorResponse{
			Remark: "Account not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(&dto.TransactionResponse{
		Saldo: account.Balance,
	})
}

// WithdrawBalance withdraw balance from account
//
// @Summary      Withdraw balance from account
// @Description  Withdraw balance from account
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param 		 no_rekening path string true "No Rekening"
// @Success      200  {object}  dto.TransactionResponse
// @Success      400  {object}  dto.ErrorResponse
// @Router       /account/saldo/{no_rekening} [get]
func (controller *AccountController) CheckBalance(c *fiber.Ctx) error {
	noRekening := c.Params("no_rekening")
	if noRekening == "" {
		return c.Status(fiber.StatusBadRequest).JSON(&dto.ErrorResponse{
			Remark: "No Rekening is required",
		})
	}

	account, err := controller.accountService.GetAccountByAccountNumber(noRekening)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&dto.ErrorResponse{
			Remark: err.Error(),
		})
	}

	if account == nil {
		return c.Status(fiber.StatusBadRequest).JSON(&dto.ErrorResponse{
			Remark: "Account not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(&dto.TransactionResponse{
		Saldo: account.Balance,
	})
}
