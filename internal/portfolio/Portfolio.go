package portfolio

import (
	"errors"
)

type Position struct {
	AvgCost  float64 `json:"avgCost"`
	Quantity int     `json:"quantity"`
}

type Portfolio struct {
	Cash      float64                  `json:"cash"`
	Positions map[string]*Position `json:"positions"`
}

func NewPortfolio(cash float64) *Portfolio {
	return &Portfolio{
		Cash:      cash,
		Positions: make(map[string]*Position),
	}
}

func (p *Portfolio) Buy(symbol string, quantity int, price float64) error {
	if p.Cash < float64(quantity)*price {
		return errors.New("insufficient cash")
	}

	p.Cash -= float64(quantity) * price

	pos, exists := p.Positions[symbol]
	if !exists {
		pos = &Position{AvgCost: price, Quantity: quantity}
	} else {
		newTotalCost := pos.AvgCost*float64(pos.Quantity) + float64(quantity)*price
		pos.AvgCost = newTotalCost / float64(pos.Quantity+quantity)
		pos.Quantity += quantity
	}

	p.Positions[symbol] = pos
	return nil
}

func (p *Portfolio) Sell(symbol string, quantity int, price float64) error {
	pos, exists := p.Positions[symbol]
	if !exists || pos.Quantity < quantity {
		return errors.New("insufficient shares or no position")
	}

	p.Cash += float64(quantity) * price
	pos.Quantity -= quantity

	if pos.Quantity == 0 {
		delete(p.Positions, symbol)
	} else {
		p.Positions[symbol] = pos
	}

	return nil
}

func (p *Portfolio) UnrealizedPnL(currentPrices map[string]float64) float64 {
	var pnl float64

	for symbol, pos := range p.Positions {
		currentPrice, exists := currentPrices[symbol]
		if !exists {
			continue
		}
		pnl += (currentPrice - pos.AvgCost) * float64(pos.Quantity)
	}

	return pnl
}
