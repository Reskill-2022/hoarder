package server

import (
	"context"
	"net"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/Reskill-2022/hoarder/controllers"
)

// Start starts the HTTP server, binding routes to the appropriate handler.
func Start(ctx context.Context, cts *controllers.Set, port string) error {
	e := echo.New()
	bindRoutes(e, cts)

	err := e.Start(addrFromPort(port))
	if err != http.ErrServerClosed {
		return err
	}
	return nil
}

func bindRoutes(e *echo.Echo, cts *controllers.Set) {
	v1 := e.Group("/v1")

	v1.GET("/health", controllers.HealthCheck)

	// slack :- /v1/slack
	slack := v1.Group("/slack")
	{
		slack.POST("/events", cts.SlackController.Events())
	}
}

func addrFromPort(port string) string {
	return net.JoinHostPort("", port)
}
