package gateway

import (
	"context"
	"testing"
)

func TestHealth(t *testing.T) {
	ctx := context.Background()
	
	resp, err := Health(ctx)
	
	if err != nil {
		t.Fatalf("Health check failed: %v", err)
	}
	
	if resp.Status != "healthy" {
		t.Errorf("Expected status 'healthy', got %s", resp.Status)
	}
	
	if len(resp.Services) == 0 {
		t.Error("Expected services map to be populated")
	}
}

func TestGetStockInfo(t *testing.T) {
	ctx := context.Background()
	
	resp, err := GetStockInfo(ctx, &GetStockInfoParams{
		StockCode:   "BBCA",
		IncludeNews: false,
	})
	
	if err != nil {
		t.Fatalf("GetStockInfo failed: %v", err)
	}
	
	if resp.Info == nil {
		t.Fatal("Expected info, got nil")
	}
	
	if resp.Info.Stock == nil {
		t.Error("Expected stock info")
	}
}

func TestSearch(t *testing.T) {
	ctx := context.Background()
	
	resp, err := Search(ctx, &SearchParams{
		Query:      "BBCA",
		SearchNews: false,
		Limit:      10,
	})
	
	if err != nil {
		t.Fatalf("Search failed: %v", err)
	}
	
	if resp == nil {
		t.Fatal("Expected response, got nil")
	}
}
