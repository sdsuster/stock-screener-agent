package news

import (
	"context"
	"testing"
)

func TestCreateNews(t *testing.T) {
	ctx := context.Background()
	
	resp, err := CreateNews(ctx, &CreateNewsParams{
		Title:   "Test News",
		Content: "Test content",
		URL:     "https://example.com/test",
	})
	
	if err != nil {
		t.Fatalf("CreateNews failed: %v", err)
	}
	
	if resp.News == nil {
		t.Fatal("Expected news, got nil")
	}
	
	if resp.News.Title != "Test News" {
		t.Errorf("Expected title 'Test News', got %s", resp.News.Title)
	}
}

func TestGetNews(t *testing.T) {
	ctx := context.Background()
	
	resp, err := GetNews(ctx, &GetNewsParams{
		ID: 1,
	})
	
	if err != nil {
		t.Fatalf("GetNews failed: %v", err)
	}
	
	if resp.News == nil {
		t.Fatal("Expected news, got nil")
	}
}

func TestListNews(t *testing.T) {
	ctx := context.Background()
	
	resp, err := ListNews(ctx, &ListNewsParams{
		Limit:  10,
		Offset: 0,
	})
	
	if err != nil {
		t.Fatalf("ListNews failed: %v", err)
	}
	
	if resp == nil {
		t.Fatal("Expected response, got nil")
	}
}

func TestSearchNews(t *testing.T) {
	ctx := context.Background()
	
	resp, err := SearchNews(ctx, &SearchNewsParams{
		Query:      "test",
		Limit:      10,
		Similarity: 0.7,
	})
	
	if err != nil {
		t.Fatalf("SearchNews failed: %v", err)
	}
	
	if resp == nil {
		t.Fatal("Expected response, got nil")
	}
}
