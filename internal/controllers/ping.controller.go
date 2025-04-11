package controllers

import "github.com/gofiber/fiber/v2"

type IPingController interface {
	Ping(c *fiber.Ctx) error
}

type PingController struct{}

type PingResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func NewPingController() *PingController {
	return &PingController{}
}

// Ping godoc
// @Summary      ping
// @Description  ping
// @Tags         Ping
// @Accept       json
// @Produce      json
// @Success      200  {object}  PingResponse
// @Router       /ping [get]
func (p *PingController) Ping(c *fiber.Ctx) error {
	response := PingResponse{
		Code:    fiber.StatusOK,
		Message: "pong",
		Data:    nil,
	}

	c.Set("Content-Type", "application/json")
	c.Set("Access-Control-Allow-Origin", "*")
	return c.Status(fiber.StatusOK).JSON(response)
}
