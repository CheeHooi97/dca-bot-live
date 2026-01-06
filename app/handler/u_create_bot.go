package handler

import (
	"dca-bot-live/app/bot"
	"dca-bot-live/app/utils"

	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateBot(c echo.Context) error {
	var i struct {
		Symbol           string  `json:"symbol" validate:"required"`
		TotalUSDT        float64 `json:"totalUSDT" validate:"required"`
		DropPercent      float64 `json:"dropPercent" validate:"required"`
		SellPercent      float64 `json:"sellPercent" validate:"required"`
		FallbackBuyHours int     `json:"fallbackBuyHours" validate:"required"`
	}

	if msg, err := utils.ValidateRequest(c, &i); err != nil {
		return responseValidationError(c, msg)
	}

	// actor, err := middleware.GetActor(c)
	// if err != nil {
	// 	return responseError(c, errcode.ActorNotFound)
	// }

	// user, err := h.User.GetById(actor.Id)
	// if err != nil {
	// 	return responseError(c, errcode.InternalServerError)
	// } else if user.Id == "" {
	// 	return responseError(c, errcode.UserNotFound)
	// }

	// trade := model.NewTrade()
	// trade.Symbol = i.Symbol
	// trade.UserId = "hooi97"
	// trade.Amount = i.Amount
	// trade.TakeProfitPercent = i.TakeProfitPercent
	// trade.StopLossPercent = i.StopLossPercent
	// trade.ProfitLoss = i.ProfitLoss

	// if err := h.Trade.Create(trade); err != nil {
	// 	return responseError(c, errcode.InternalServerError)
	// }

	oneBuyUSDT := i.TotalUSDT * 0.01
	go bot.RunDCABot(
		i.Symbol,
		i.TotalUSDT,
		oneBuyUSDT,
		i.DropPercent,
		i.SellPercent,
		i.FallbackBuyHours,
	)

	return responseJSON(c, echo.Map{
		"success": true,
	})
}
