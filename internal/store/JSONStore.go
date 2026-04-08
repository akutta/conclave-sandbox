package store

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/conclave-sandbox/internal/portfolio"
)

type JSONStore struct {
	Path string
}

func NewJSONStore(path string) *JSONStore {
	return &JSONStore{Path: path}
}

func (js *JSONStore) Load() (*portfolio.Portfolio, error) {
	file, err := os.Open(js.Path)
	if err != nil {
		if os.IsNotExist(err) {
			return portfolio.NewPortfolio(100000), nil
		}
		return nil, fmt.Errorf("failed to open portfolio file: %w", err)
	}
	defer file.Close()

	var portfolio portfolio.Portfolio
	err = json.NewDecoder(file).Decode(&portfolio)
	if err != nil {
		return nil, fmt.Errorf("failed to decode portfolio JSON: %w", err)
	}

	return &portfolio, nil
}

func (js *JSONStore) Save(portfolio *portfolio.Portfolio) error {
	file, err := os.Create(js.Path)
	if err != nil {
		return fmt.Errorf("failed to create portfolio file: %w", err)
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(portfolio)
	if err != nil {
		return fmt.Errorf("failed to encode portfolio JSON: %w", err)
	}

	return nil
}
