package main

import (
	"context"
	"fmt"

	"github.com/llowkeysam/go/application"
)

func main() {
	app := application.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Printf("failed to start app:", err)
	}
}
