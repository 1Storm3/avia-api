package main

import (
	"context"

	"github.com/1Storm3/avia-api/internal/app"
)

func main() {

	a := app.New()

	ctx := context.Background()

	if err := a.Run(ctx); err != nil {
		panic(err)
	}
}
