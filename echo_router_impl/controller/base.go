package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Base struct {
	//
}

func (controller *Base) Error(ctx echo.Context, err error, message string) error {
	errorObject := &Response{
		Status:       "ERROR",
		Message:      message,
		InnerMessage: err.Error(),
	}
	return ctx.JSON(http.StatusInternalServerError, errorObject)
}

func (controller *Base) NotFound(ctx echo.Context, err error, message string) error {
	errorObject := &Response{
		Status:       "ERROR",
		Message:      message,
		InnerMessage: err.Error(),
	}
	return ctx.JSON(http.StatusNotFound, errorObject)
}

func (controller *Base) Unauthorized(ctx echo.Context, err error, message string) error {
	errorObject := &Response{
		Status:       "ERROR",
		Message:      message,
		InnerMessage: err.Error(),
	}
	return ctx.JSON(http.StatusUnauthorized, errorObject)
}
