package routes

import (
	"github.com/fiqrikm18/ISITransaction/internal/controllers"
	"github.com/fiqrikm18/ISITransaction/internal/repositories"
	"github.com/fiqrikm18/ISITransaction/internal/services"
	"github.com/gofiber/fiber/v2"
)

var (
	accountRepository  repositories.IAccountRepository
	mutationRepository repositories.IMutationRepository
)

var (
	accountService services.IAccountService
)

var (
	pingController    controllers.IPingController
	accountController controllers.IAccountController
)

func init() {
	accountRepository = repositories.NewAccountRepository()
	mutationRepository = repositories.NewMutationRepository()

	accountService = services.NewAccountService(accountRepository, mutationRepository)

	pingController = controllers.NewPingController()
	accountController = controllers.NewAccountController(accountService)
}

func RegisterApiRouter(route *fiber.App) {
	api := route.Group("/api/v1")

	// Ping
	api.Get("/ping", pingController.Ping)

	// Account
	accountRoute := api.Group("/account")
	accountRoute.Post("/daftar", accountController.RegisterAccount)
	accountRoute.Post("/tabung", accountController.DepositBalance)
	accountRoute.Post("/tarik", accountController.WithdrawBalance)
	accountRoute.Get("/saldo/:no_rekening", accountController.CheckBalance)
}
