package http

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestHandlers(t *testing.T) {
	tests := []struct {
		description  string
		route        string
		method       string
		input        interface{}
		expectedCode int
	}{
		{
			description:  "get HTTP status 200 for balance route",
			method:       "GET",
			route:        "/balance",
			expectedCode: 200,
		},
		{
			description:  "get HTTP status 404, when route is not exists",
			method:       "GET",
			route:        "/not-found",
			expectedCode: 404,
		},
		{
			description:  "get HTTP status 200, when input is ok",
			method:       "GET",
			route:        "/transaction",
			input:        TransactionInput{Value: 1, Type: "c", Description: "testing"},
			expectedCode: 200,
		},
	}

	app := fiber.New()
	app.Get("/balance", BalanceHandler)
	app.Get("/transaction", TransactionHandler)

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			b, _ := json.Marshal(test.input)
			req := httptest.NewRequest(test.method, test.route, bytes.NewReader(b))
			req.Header.Set("Content-Type", "application/json")

			resp, _ := app.Test(req, 1)

			assert.Equal(t, test.expectedCode, resp.StatusCode)
		})
	}

}
