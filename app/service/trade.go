package service

import (
	"dca-bot-live/app/model"
	"dca-bot-live/app/repository"
)

type TradeService struct {
	tradeRepo repository.TradeRepository
}

func NewTradeService(tradeRepo repository.TradeRepository) *TradeService {
	return &TradeService{tradeRepo: tradeRepo}
}

func (s *TradeService) Create(trade *model.Trade) error {
	return s.tradeRepo.Create(trade)
}

func (s *TradeService) GetById(id string) (*model.Trade, error) {
	return s.tradeRepo.GetById(id)
}

func (s *TradeService) Update(trade *model.Trade) error {
	return s.tradeRepo.Update(trade)
}

func (s *TradeService) Delete(id string) error {
	return s.tradeRepo.Delete(id)
}
