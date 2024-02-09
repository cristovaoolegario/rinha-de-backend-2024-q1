package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Post("/clientes/{id}/transacoes", TransactionHandler)
	app.Get("/clientes/{id}/extrato", BalanceHandler)

	app.Listen(":3000")
}

func TransactionHandler(ctx *fiber.Ctx) error {
	return ctx.SendString("Transaction Handler!")
}

func BalanceHandler(ctx *fiber.Ctx) error {
	return ctx.SendString("Balance Handler!")
}
