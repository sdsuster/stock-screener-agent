package gateway

import (
	"context"
	"time"

	"encore.app/marketdata"
	"encore.app/news"
	"encore.dev/rlog"
)

// HealthResponse represents the health check response
type HealthResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Services  map[string]string `json:"services"`
}

// Health checks the health of all services
//
//encore:api public method=GET path=/health
func Health(ctx context.Context) (*HealthResponse, error) {
	rlog.Info("Health check requested")
	
	return &HealthResponse{
		Status:    "healthy",
		Timestamp: time.Now(),
		Services: map[string]string{
			"gateway":    "healthy",
			"marketdata": "healthy",
			"news":       "healthy",
		},
	}, nil
}

// StockInfo represents combined stock information
type StockInfo struct {
	Stock         *marketdata.Stock         `json:"stock"`
	LatestSummary *marketdata.StockSummary  `json:"latest_summary,omitempty"`
	RelatedNews   []news.News               `json:"related_news,omitempty"`
}

// GetStockInfoParams represents parameters for getting comprehensive stock info
type GetStockInfoParams struct {
	StockCode   string `json:"stock_code"`
	IncludeNews bool   `json:"include_news"`
}

// GetStockInfoResponse represents the response for comprehensive stock info
type GetStockInfoResponse struct {
	Info *StockInfo `json:"info"`
}

// GetStockInfo retrieves comprehensive stock information including news
//
//encore:api public method=GET path=/api/stocks/:stock_code
func GetStockInfo(ctx context.Context, params *GetStockInfoParams) (*GetStockInfoResponse, error) {
	rlog.Info("Getting comprehensive stock info", "stock_code", params.StockCode)
	
	// Get stock data from marketdata service
	stockResp, err := marketdata.GetStock(ctx, &marketdata.GetStockParams{
		StockCode: params.StockCode,
	})
	if err != nil {
		return nil, err
	}
	
	info := &StockInfo{
		Stock: stockResp.Stock,
	}
	
	// Optionally fetch related news
	if params.IncludeNews {
		newsResp, err := news.ListNews(ctx, &news.ListNewsParams{
			Limit:  10,
			Offset: 0,
		})
		if err != nil {
			rlog.Error("Failed to fetch news", "error", err)
		} else {
			info.RelatedNews = newsResp.News
		}
	}
	
	return &GetStockInfoResponse{
		Info: info,
	}, nil
}

// SearchParams represents parameters for cross-service search
type SearchParams struct {
	Query      string `json:"query"`
	SearchNews bool   `json:"search_news"`
	Limit      int    `json:"limit"`
}

// SearchResponse represents the response for cross-service search
type SearchResponse struct {
	Stocks []marketdata.Stock `json:"stocks,omitempty"`
	News   []news.SearchResult `json:"news,omitempty"`
}

// Search performs cross-service search
//
//encore:api public method=POST path=/api/search
func Search(ctx context.Context, params *SearchParams) (*SearchResponse, error) {
	rlog.Info("Cross-service search", "query", params.Query)
	
	response := &SearchResponse{}
	
	// Search stocks
	stocksResp, err := marketdata.ListStocks(ctx)
	if err != nil {
		rlog.Error("Failed to search stocks", "error", err)
	} else {
		response.Stocks = stocksResp.Stocks
	}
	
	// Search news if requested
	if params.SearchNews {
		newsResp, err := news.SearchNews(ctx, &news.SearchNewsParams{
			Query:      params.Query,
			Limit:      params.Limit,
			Similarity: 0.7,
		})
		if err != nil {
			rlog.Error("Failed to search news", "error", err)
		} else {
			response.News = newsResp.Results
		}
	}
	
	return response, nil
}
