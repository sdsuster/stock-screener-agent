package marketdata

import (
	"context"
	"time"
)

// Stock represents a stock ticker
type Stock struct {
	StockCode string `json:"stock_code"`
	StockName string `json:"stock_name"`
}

// StockSummary represents daily stock market data
type StockSummary struct {
	IDStockSummary   int64     `json:"id_stock_summary"`
	Date             time.Time `json:"date"`
	StockCode        string    `json:"stock_code"`
	StockName        string    `json:"stock_name"`
	Remarks          string    `json:"remarks"`
	Previous         int64     `json:"previous"`
	OpenPrice        int64     `json:"open_price"`
	FirstTrade       int64     `json:"first_trade"`
	High             int64     `json:"high"`
	Low              int64     `json:"low"`
	Close            int64     `json:"close"`
	Change           int64     `json:"change"`
	Volume           int64     `json:"volume"`
	Value            int64     `json:"value"`
	Frequency        int64     `json:"frequency"`
	IndexIndividual  float64   `json:"index_individual"`
	Offer            int64     `json:"offer"`
	OfferVolume      int64     `json:"offer_volume"`
	Bid              int64     `json:"bid"`
	BidVolume        int64     `json:"bid_volume"`
	ListedShares     int64     `json:"listed_shares"`
	TradeableShares  int64     `json:"tradeable_shares"`
	WeightForIndex   int64     `json:"weight_for_index"`
	ForeignSell      int64     `json:"foreign_sell"`
	ForeignBuy       int64     `json:"foreign_buy"`
	DelistingDate    *time.Time `json:"delisting_date,omitempty"`
	NonRegularVolume int64     `json:"non_regular_volume"`
	NonRegularValue  int64     `json:"non_regular_value"`
	NonRegularFrequency int64  `json:"non_regular_frequency"`
	Percentage       *float64  `json:"percentage,omitempty"`
}

// GetStockParams represents parameters for getting a stock
type GetStockParams struct {
	StockCode string `json:"stock_code"`
}

// GetStockResponse represents the response for getting a stock
type GetStockResponse struct {
	Stock *Stock `json:"stock"`
}

// ListStocksResponse represents the response for listing stocks
type ListStocksResponse struct {
	Stocks []Stock `json:"stocks"`
	Total  int     `json:"total"`
}

// CreateStockParams represents parameters for creating a stock
type CreateStockParams struct {
	StockCode string `json:"stock_code"`
	StockName string `json:"stock_name"`
}

// GetStockSummaryParams represents parameters for getting stock summaries
type GetStockSummaryParams struct {
	StockCode string     `json:"stock_code"`
	StartDate *time.Time `json:"start_date,omitempty"`
	EndDate   *time.Time `json:"end_date,omitempty"`
	Limit     int        `json:"limit"`
}

// GetStockSummaryResponse represents the response for getting stock summaries
type GetStockSummaryResponse struct {
	Summaries []StockSummary `json:"summaries"`
	Total     int            `json:"total"`
}
