package portfolio

import (
	"testing"
)

func TestBuy(t *testing.T) {
	p := NewPortfolio(1000.0)
	err := p.Buy("AAPL", 1, 150.0)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if p.Cash != 850.0 {
		t.Errorf("expected cash to be 850.0, got %.2f", p.Cash)
	}
	pos := p.Positions["AAPL"]
	if pos == nil || pos.AvgCost != 150.0 || pos.Quantity != 1 {
		t.Errorf("unexpected position: %+v", pos)
	}

	err = p.Buy("AAPL", 2, 160.0)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if p.Cash != 530.0 {
		t.Errorf("expected cash to be 530.0, got %.2f", p.Cash)
	}
	pos = p.Positions["AAPL"]
	if pos == nil || pos.AvgCost != 156.67 || pos.Quantity != 3 {
		t.Errorf("unexpected position: %+v", pos)
	}

	err = p.Buy("MSFT", 1, 300.0)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if p.Cash != 230.0 {
		t.Errorf("expected cash to be 230.0, got %.2f", p.Cash)
	}
	pos = p.Positions["MSFT"]
	if pos == nil || pos.AvgCost != 300.0 || pos.Quantity != 1 {
		t.Errorf("unexpected position: %+v", pos)
	}

	err = p.Buy("GOOGL", 1, 2800.0)
	if err == nil {
		t.Errorf("expected error for insufficient cash")
	}
}

func TestSell(t *testing.T) {
	p := NewPortfolio(1000.0)
	p.Buy("AAPL", 3, 150.0)

	err := p.Sell("AAPL", 2, 160.0)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if p.Cash != 470.0 {
		t.Errorf("expected cash to be 470.0, got %.2f", p.Cash)
	}
	pos := p.Positions["AAPL"]
	if pos == nil || pos.AvgCost != 150.0 || pos.Quantity != 1 {
		t.Errorf("unexpected position: %+v", pos)
	}

	err = p.Sell("AAPL", 1, 160.0)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if p.Cash != 630.0 {
		t.Errorf("expected cash to be 630.0, got %.2f", p.Cash)
	}
	pos = p.Positions["AAPL"]
	if pos != nil {
		t.Errorf("unexpected position: %+v", pos)
	}

	err = p.Sell("MSFT", 1, 300.0)
	if err == nil {
		t.Errorf("expected error for non-existent position")
	}

	p.Buy("GOOGL", 2, 2800.0)
	err = p.Sell("GOOGL", 3, 2900.0)
	if err == nil {
		t.Errorf("expected error for over-selling")
	}
}

func TestUnrealizedPnL(t *testing.T) {
	p := NewPortfolio(1000.0)
	p.Buy("AAPL", 3, 150.0)
	p.Buy("MSFT", 2, 300.0)

	currentPrices := map[string]float64{
		"AAPL": 160.0,
		"MSFT": 290.0,
	}

	pnl := p.UnrealizedPnL(currentPrices)
	expectedPnL := (160-150)*3 + (290-300)*2
	if pnl != expectedPnL {
		t.Errorf("expected P&L to be %.2f, got %.2f", expectedPnL, pnl)
	}
}
