package main

import (
	"context"

	"github.com/Reskill-2022/hoarder/controllers"
	"github.com/Reskill-2022/hoarder/errors"
	"github.com/Reskill-2022/hoarder/log"
	"github.com/Reskill-2022/hoarder/server"
)

func main() {
	ctx := context.Background()

	cts := controllers.NewSet()

	if err := server.Start(ctx, cts, "8001"); err != nil {
		log.FromContext(ctx).Named("main").Fatal("failed to start HTTP server", errors.ErrorLogFields(err)...)
	}
}
