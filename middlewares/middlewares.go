package middlewares

import (
	"fmt"
	"net/http/httputil"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	"github.com/Reskill-2022/hoarder/log"
)

func RequestLogger(l *log.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			req := c.Request()
			res := c.Response()

			start := time.Now()
			if err = next(c); err != nil {
				c.Error(err)
			}
			stop := time.Now()

			id := req.Header.Get(echo.HeaderXRequestID)
			if id == "" {
				id = res.Header().Get(echo.HeaderXRequestID)
			}

			fields := []zap.Field{
				zap.String("method", req.Method),
				zap.String("status", fmt.Sprint(res.Status)),
				zap.String("uri", req.RequestURI),
				zap.String("remote_ip", c.RealIP()),
				zap.String("time", fmt.Sprint(time.Now().UTC())),
				zap.String("id", id),
				zap.String("host", req.Host),
				zap.String("user_agent", req.UserAgent()),
				zap.String("latency", fmt.Sprint(stop.Sub(start))),
				zap.String("referer", req.Referer()),
			}

			code := res.Status
			switch {
			case code >= 500:
				l.Error("server error", fields...)
				return
			case code >= 400 && code < 500:
				l.Error("client error", fields...)
				return
			}

			l.Debug("", fields...)
			return
		}
	}
}

func DumpRequest(l *log.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			req := c.Request()

			reqBody, err := httputil.DumpRequest(req, true)
			if err != nil {
				panic(err)
			}

			l.Debug(spew.Sprint(string(reqBody)))
			return next(c)
		}
	}
}
