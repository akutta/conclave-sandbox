package store

import (
	"context"

	"akutta/papermoney/internal/portfolio"
)

// Store persists and loads portfolio state.
type Store interface {
	Save(ctx context.Context, p *portfolio.Portfolio) error
	Load(ctx context.Context) (*portfolio.Portfolio, error)
}

// JSONStore persists portfolio state to a local JSON file.
type JSONStore struct {
	Path string
}

func (s *JSONStore) Save(ctx context.Context, p *portfolio.Portfolio) error {
	// TODO: implement
	return nil
}

func (s *JSONStore) Load(ctx context.Context) (*portfolio.Portfolio, error) {
	// TODO: implement
	return nil, nil
}
