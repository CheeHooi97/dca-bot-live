package repository

import (
	"dca-bot-live/app/model"
	"errors"
	"time"

	"gorm.io/gorm"
)

type TradeRepository interface {
	Create(trade *model.Trade) error
	GetById(id string) (*model.Trade, error)
	Update(trade *model.Trade) error
	Delete(id string) error
}

type tradeRepository struct {
	db *gorm.DB
}

func NewTradeRepository(db *gorm.DB) TradeRepository {
	return &tradeRepository{db: db}
}

func (r *tradeRepository) Create(trade *model.Trade) error {
	result := r.db.Create(trade)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *tradeRepository) GetById(id string) (*model.Trade, error) {
	var trade model.Trade
	result := r.db.First(&trade, id)
	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
	}
	return &trade, nil
}

func (r *tradeRepository) Update(trade *model.Trade) error {
	result := r.db.Save(trade)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *tradeRepository) Delete(id string) error {
	result := r.db.Model(&model.Trade{}).Where("id = ?", id).Updates(map[string]any{
		"status":          false,
		"updatedDateTime": time.Now().UTC(),
	})
	return result.Error
}
