package http

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

func TransactionHandler(ctx *fiber.Ctx) error {
	var t TransactionInput
	if err := ctx.BodyParser(&t); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// Execute Use case
	limit, balance := 100, 10000
	response := &TransactionResponse{
		Balance: balance,
		Limit:   limit,
	}
	b, err := json.Marshal(response)

	if err != nil {
		return err
	}
	return ctx.Send(b)
}

func BalanceHandler(ctx *fiber.Ctx) error {
	return ctx.SendString("Balance Handler!")
}
