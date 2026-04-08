package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/conclave-sandbox/internal/market"
	"github.com/conclave-sandbox/internal/portfolio"
	"github.com/conclave-sandbox/internal/store"
)

func main() {
	watchlist := flag.String("watchlist", "AAPL,MSFT,GOOGL,AMZN,NVDA", "comma-separated list of symbols to watch")
	cash := flag.Float64("cash", 100000, "starting cash amount for the portfolio")
	flag.Parse()

	symbols := splitSymbols(*watchlist)
	yf := market.NewYahooFetcher()
	store := store.NewJSONStore("portfolio.json")

	portfolio, err := store.Load()
	if err != nil {
		log.Fatalf("failed to load portfolio: %v", err)
	}

	prices, err := yf.Quotes(symbols)
	if err != nil {
		log.Fatalf("failed to fetch quotes: %v", err)
	}

	currentPrices := make(map[string]float64)
	for _, price := range prices {
		currentPrices[price.Symbol] = price.Price
	}

	for _, symbol := range symbols {
		price, exists := currentPrices[symbol]
		if !exists || price <= 0 {
			continue
		}

		pos, exists := portfolio.Positions[symbol]
		if !exists {
			err := portfolio.Buy(symbol, 1, price)
			if err != nil {
				log.Printf("failed to buy %s: %v", symbol, err)
			} else {
				log.Printf("bought 1 share of %s at $%.2f", symbol, price)
			}
		} else {
			pnl := (price - pos.AvgCost) * float64(pos.Quantity)
			log.Printf("%s position: %d shares, avg cost $%.2f, unrealized P&L $%.2f", symbol, pos.Quantity, pos.AvgCost, pnl)
		}
	}

	err = store.Save(portfolio)
	if err != nil {
		log.Fatalf("failed to save portfolio: %v", err)
	}

	fmt.Println("Portfolio summary:")
	data, _ := json.MarshalIndent(portfolio, "", "  ")
	fmt.Println(string(data))
}

func splitSymbols(symbolsStr string) []string {
	var symbols []string
	for _, symbol := range splitString(symbolsStr, ",") {
		if symbol != "" {
			symbols = append(symbols, symbol)
		}
	}
	return symbols
}

func splitString(str, sep string) []string {
	var result []string
	start := 0
	end := 0

	for end < len(str) {
		if str[end] == sep[0] {
			result = append(result, str[start:end])
			start = end + 1
		}
		end++
	}

	result = append(result, str[start:])
	return result
}
