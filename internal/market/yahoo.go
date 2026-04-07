package market

import (
	"context"
	"errors"
)

var errNotImplemented = errors.New("not implemented")

// PriceQuote holds a market price snapshot for a single symbol.
type PriceQuote struct {
	Symbol string
	Price  float64
	// TODO: add bid/ask, volume, timestamp
}

// Fetcher retrieves real-time price quotes.
type Fetcher interface {
	Quote(ctx context.Context, symbol string) (*PriceQuote, error)
	Quotes(ctx context.Context, symbols []string) ([]*PriceQuote, error)
}

// YahooFetcher fetches real-time quotes from Yahoo Finance.
// Uses the unofficial v8/finance/quote endpoint (no API key required).
type YahooFetcher struct {
	// TODO: add HTTP client with timeout
}

func (f *YahooFetcher) Quote(ctx context.Context, symbol string) (*PriceQuote, error) {
	return nil, errNotImplemented
}

func (f *YahooFetcher) Quotes(ctx context.Context, symbols []string) ([]*PriceQuote, error) {
	return nil, errNotImplemented
}
