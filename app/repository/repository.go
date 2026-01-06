package repository

import "gorm.io/gorm"

type Repositories struct {
	TradeRepo TradeRepository
}

func InitializeRepository(db *gorm.DB) *Repositories {
	return &Repositories{
		TradeRepo: NewTradeRepository(db),
	}
}
