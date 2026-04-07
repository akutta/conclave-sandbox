# papermoney

A paper trading simulator for US equities using real market data from Yahoo Finance.
This repository is the primary test subject for [conclave.ai](https://conclave.andrewkutta.com).

## What it does (eventually)

- Fetch real-time stock quotes (Yahoo Finance unofficial API — no key required)
- Maintain a paper portfolio: buy/sell positions, track cash balance
- Calculate unrealized P&L across open positions
- Persist portfolio state to a local JSON file

## Project structure

```
cmd/trader/        — CLI entry point
internal/market/   — Yahoo Finance price fetcher
internal/portfolio/ — Portfolio, Position, buy/sell logic
internal/store/    — JSON file persistence
```

## What's not yet implemented

- [ ] `YahooFetcher.Quote` / `Quotes` — HTTP call to Yahoo Finance v8 API
- [ ] `Portfolio.Buy` — deduct cash, open/increase position
- [ ] `Portfolio.Sell` — close/reduce position, return cash
- [ ] `Portfolio.UnrealizedPnL` — sum (currentPrice - avgCost) * quantity across positions
- [ ] `JSONStore.Save` / `Load` — marshal/unmarshal portfolio to disk
- [ ] `cmd/trader` CLI — wire everything together with flags for buy/sell/status

## Running

```bash
go build ./cmd/trader
./trader   # prints "not implemented" until the above is done
```
