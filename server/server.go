package server

import (
	"context"
	"net"
	"net/http"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"

	"github.com/Reskill-2022/hoarder/controllers"
	"github.com/Reskill-2022/hoarder/log"
	"github.com/Reskill-2022/hoarder/middlewares"
	"github.com/Reskill-2022/hoarder/repositories"
	"github.com/Reskill-2022/hoarder/services"
)

// Start starts the HTTP server, binding routes to the appropriate handler.
func Start(ctx context.Context, cts *controllers.Set, svs *services.Set, rcs *repositories.Set, port string) error {
	e := echo.New()

	e.GET("/health", controllers.HealthCheck)
	e.Use(
		echoMiddleware.Recover(),
		echoMiddleware.RequestID(),
	)
	bindRoutes(ctx, e, cts, svs, rcs)

	err := e.Start(addrFromPort(port))
	if err != http.ErrServerClosed {
		return err
	}
	return nil
}

func bindRoutes(ctx context.Context, e *echo.Echo, cts *controllers.Set, svs *services.Set, rcs *repositories.Set) {
	v1 := e.Group("/v1")
	v1.Use(middlewares.RequestLogger(log.FromContext(ctx)))

	// slack :- /v1/slack
	slack := v1.Group("/slack")
	{
		slack.POST("/events", cts.SlackController.Events(rcs.BigQuery))
	}

	// zendesk :- /v1/zendesk
	zendesk := v1.Group("/zendesk")
	{
		zendesk.POST("/tickets", cts.ZendeskController.CreateTicket(rcs.BigQuery, svs.SlackService))
	}

	calendly := v1.Group("/calendly")
	{
		calendly.POST("/events/:memberId", cts.CalendlyController.Events(svs.CalendlyService, rcs.BigQuery))
	}
}

func addrFromPort(port string) string {
	return net.JoinHostPort("", port)
}
