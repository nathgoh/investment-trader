package main

import (
	"context"

	"github.com/nathgoh/investment-trader/alpaca/api/routes"
)

func main() {
	ctx := context.Background()

	router := routes.Handler(ctx)
	router.Run(":8080")
}
