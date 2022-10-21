package controllers

import (
	"net/http"

	"github.com/jcobhams/echoresponse"
	"github.com/labstack/echo/v4"

	"github.com/Reskill-2022/hoarder/errors"
	"github.com/Reskill-2022/hoarder/log"
)

type Set struct {
	SlackController *SlackController
}

func NewSet() *Set {
	return &Set{
		SlackController: NewSlackController(),
	}
}

func ErrorHandler(c echo.Context, err error) error {
	var code int
	var msg interface{}

	switch e := err.(type) {
	case errors.Error:
		code = e.Code
		msg = e.Msg
	case *echo.HTTPError:
		code = e.Code
		msg = e.Message
	}

	if code < 100 {
		code = 500
	}

	if code > 0 && code < 500 {
		return echoresponse.Format(c, msg, nil, code)
	}

	log.FromContext(c.Request().Context()).Named("ErrorHandler").Error(err.Error(), errors.ErrorLogFields(err)...)

	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead {
			return c.NoContent(code)
		}
	}
	return echoresponse.Format(c, "Internal Service Error - Try Later", nil, http.StatusInternalServerError)
}
