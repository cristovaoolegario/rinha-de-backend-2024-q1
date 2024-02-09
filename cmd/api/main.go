package main

import (
	"github.com/cristovaoolegario/rinha-de-backend-2024-q1/internal/infra/http"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Post("/clientes/{id}/transacoes", http.TransactionHandler)
	app.Get("/clientes/{id}/extrato", http.BalanceHandler)

	app.Listen(":3000")
}
