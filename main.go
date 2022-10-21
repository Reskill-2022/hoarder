package main

import (
	"context"

	"github.com/Reskill-2022/hoarder/errors"
	"github.com/Reskill-2022/hoarder/log"
	"github.com/Reskill-2022/hoarder/server"
)

func main() {
	ctx := context.Background()

	if err := server.Start(ctx, "8000"); err != nil {
		log.FromContext(ctx).Named("main").Fatal("failed to start HTTP server", errors.ErrorLogFields(err)...)
	}
}
