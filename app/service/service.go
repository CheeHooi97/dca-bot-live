package service

import "dca-bot-live/app/repository"

type Services struct {
	TradeService *TradeService
}

func InitializeService(repos *repository.Repositories) *Services {
	return &Services{
		TradeService: NewTradeService(repos.TradeRepo),
	}
}
