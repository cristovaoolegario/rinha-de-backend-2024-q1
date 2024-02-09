package main

import (
	"context"
	"fmt"
	"github.com/cristovaoolegario/rinha-de-backend-2024-q1/internal/infra/http"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"

	"os"
)

func main() {

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer conn.Close(context.Background())

	app := fiber.New()

	app.Post("/clientes/:id/transacoes", http.TransactionHandler)
	app.Get("/clientes/:id/extrato", http.BalanceHandler)

	app.Listen(":3000")
}
