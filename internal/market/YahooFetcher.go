package market

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type PriceQuote struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"regularMarketPrice"`
}

type YahooFetcher struct {
	client *http.Client
}

func NewYahooFetcher() *YahooFetcher {
	return &YahooFetcher{
		client: &http.Client{Timeout: 10 * time.Second},
	}
}

func (yf *YahooFetcher) Quotes(symbols []string) ([]*PriceQuote, error) {
	symbolsStr := joinSymbols(symbols)
	url := fmt.Sprintf("https://query1.finance.yahoo.com/v8/finance/quote?symbols=%s", symbolsStr)

	resp, err := yf.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch quotes: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var data struct {
		QuoteResponse struct {
			Result []*PriceQuote `json:"result"`
		} `json:"quoteResponse"`
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}

	return data.QuoteResponse.Result, nil
}

func (yf *YahooFetcher) Quote(symbol string) (*PriceQuote, error) {
	quotes, err := yf.Quotes([]string{symbol})
	if err != nil {
		return nil, err
	}
	if len(quotes) == 0 {
		return nil, fmt.Errorf("no quote found for symbol: %s", symbol)
	}
	return quotes[0], nil
}

func joinSymbols(symbols []string) string {
	symbolsStr := ""
	for i, symbol := range symbols {
		if i > 0 {
			symbolsStr += ","
		}
		symbolsStr += symbol
	}
	return symbolsStr
}
