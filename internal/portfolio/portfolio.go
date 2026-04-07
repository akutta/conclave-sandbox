package portfolio

import (
	"errors"
	"time"
)

var errNotImplemented = errors.New("not implemented")

// Position represents an open equity holding.
type Position struct {
	Symbol   string
	Quantity float64
	AvgCost  float64
	OpenedAt time.Time
}

// Portfolio tracks cash and open positions.
type Portfolio struct {
	Cash      float64
	Positions map[string]*Position
}

// New creates a portfolio with the given starting cash balance.
func New(startingCash float64) *Portfolio {
	return &Portfolio{
		Cash:      startingCash,
		Positions: make(map[string]*Position),
	}
}

// UnrealizedPnL returns total unrealized profit/loss given current prices.
// currentPrices maps symbol → current market price.
func (p *Portfolio) UnrealizedPnL(currentPrices map[string]float64) float64 {
	// TODO: implement
	return 0
}

// Buy adds or increases a position. Returns error if insufficient cash.
func (p *Portfolio) Buy(symbol string, quantity, price float64) error {
	// TODO: implement
	return errNotImplemented
}

// Sell reduces or closes a position. Returns error if position does not exist.
func (p *Portfolio) Sell(symbol string, quantity, price float64) error {
	// TODO: implement
	return errNotImplemented
}
