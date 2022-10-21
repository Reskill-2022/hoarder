package server

import (
	"context"
	"net"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Start(ctx context.Context, port string) error {
	e := echo.New()

	err := e.Start(addrFromPort(port))
	if err != http.ErrServerClosed {
		return err
	}
	return nil
}

func bindRoutes(e echo.Echo) {}

func addrFromPort(port string) string {
	return net.JoinHostPort("", port)
}
