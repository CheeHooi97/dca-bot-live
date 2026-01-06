package model

import (
	"dca-bot-live/app/utils"
	"time"
)

type Trade struct {
	Id                string
	UserId            string
	Symbol            string
	Amount            float64
	TakeProfitPercent float64
	StopLossPercent   float64
	ProfitLoss        float64
	BaseModel
}

func NewTrade() *Trade {
	now := time.Now().UTC()

	m := new(Trade)
	m.Id = utils.UniqueID()
	m.CreatedDateTime = now
	m.UpdatedDateTime = now

	return m
}

func (m *Trade) UpdateDt() {
	m.UpdatedDateTime = time.Now().UTC()
}
