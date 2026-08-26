package marketdata

import (
	"context"
	"testing"
)

func TestGetStock(t *testing.T) {
	ctx := context.Background()
	
	resp, err := GetStock(ctx, &GetStockParams{
		StockCode: "BBCA",
	})
	
	if err != nil {
		t.Fatalf("GetStock failed: %v", err)
	}
	
	if resp.Stock == nil {
		t.Fatal("Expected stock, got nil")
	}
	
	if resp.Stock.StockCode != "BBCA" {
		t.Errorf("Expected stock code BBCA, got %s", resp.Stock.StockCode)
	}
}

func TestListStocks(t *testing.T) {
	ctx := context.Background()
	
	resp, err := ListStocks(ctx)
	
	if err != nil {
		t.Fatalf("ListStocks failed: %v", err)
	}
	
	if resp == nil {
		t.Fatal("Expected response, got nil")
	}
}

func TestCreateStock(t *testing.T) {
	ctx := context.Background()
	
	resp, err := CreateStock(ctx, &CreateStockParams{
		StockCode: "TEST",
		StockName: "Test Stock",
	})
	
	if err != nil {
		t.Fatalf("CreateStock failed: %v", err)
	}
	
	if resp.Stock == nil {
		t.Fatal("Expected stock, got nil")
	}
	
	if resp.Stock.StockCode != "TEST" {
		t.Errorf("Expected stock code TEST, got %s", resp.Stock.StockCode)
	}
}
