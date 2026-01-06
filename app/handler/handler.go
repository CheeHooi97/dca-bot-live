package handler

import (
	"dca-bot-live/app/errcode"
	"dca-bot-live/app/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	Trade *service.TradeService
}

func NewHandler(services *service.Services) *Handler {
	h := &Handler{
		Trade: services.TradeService,
	}

	return h
}

func responseError(c echo.Context, message errcode.ErrorCode) error {
	return c.JSON(http.StatusOK, map[string]any{
		"result": nil,
		"errmsg": message.Message,
		"error":  true,
		"status": false,
	})
}

func responseJSON(c echo.Context, result any) error {
	return c.JSON(http.StatusOK, map[string]any{
		"result": result,
		"errmsg": "",
		"error":  false,
		"status": true,
	})
}

func responseListJSON(c echo.Context, result any) error {
	return c.JSON(http.StatusOK, map[string]any{
		"result": result,
		"errmsg": "",
		"error":  false,
		"status": true,
	})
}

func responseValidationError(c echo.Context, message string) error {
	return c.JSON(http.StatusOK, map[string]any{
		"result": nil,
		"errmsg": message,
		"error":  true,
		"status": false,
	})
}
