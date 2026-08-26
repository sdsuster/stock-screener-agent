package news

import (
	"time"
)

// EmbeddingStatus represents the status of an embedding
type EmbeddingStatus string

const (
	EmbeddingStatusPending    EmbeddingStatus = "pending"
	EmbeddingStatusProcessing EmbeddingStatus = "processing"
	EmbeddingStatusCompleted  EmbeddingStatus = "completed"
	EmbeddingStatusFailed     EmbeddingStatus = "failed"
)

// News represents a news article
type News struct {
	ID          int64      `json:"id"`
	Title       string     `json:"title"`
	Content     string     `json:"content"`
	URL         string     `json:"url"`
	PublishedAt *time.Time `json:"published_at,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
}

// NewsChunk represents a chunk of news content for embedding
type NewsChunk struct {
	ID            int64           `json:"id"`
	NewsID        int64           `json:"news_id"`
	EmbeddingType string          `json:"embedding_type"` // "title" or "content"
	ChunkIndex    int             `json:"chunk_index"`
	Content       string          `json:"content"`
	Status        EmbeddingStatus `json:"status"`
	Embedding     []float32       `json:"embedding,omitempty"`
}

// CreateNewsParams represents parameters for creating news
type CreateNewsParams struct {
	Title       string     `json:"title"`
	Content     string     `json:"content"`
	URL         string     `json:"url"`
	PublishedAt *time.Time `json:"published_at,omitempty"`
}

// CreateNewsResponse represents the response for creating news
type CreateNewsResponse struct {
	News *News `json:"news"`
}

// GetNewsParams represents parameters for getting news
type GetNewsParams struct {
	ID int64 `json:"id"`
}

// GetNewsResponse represents the response for getting news
type GetNewsResponse struct {
	News *News `json:"news"`
}

// ListNewsParams represents parameters for listing news
type ListNewsParams struct {
	Limit  int        `json:"limit"`
	Offset int        `json:"offset"`
	Since  *time.Time `json:"since,omitempty"`
}

// ListNewsResponse represents the response for listing news
type ListNewsResponse struct {
	News  []News `json:"news"`
	Total int    `json:"total"`
}

// SearchNewsParams represents parameters for semantic search
type SearchNewsParams struct {
	Query      string  `json:"query"`
	Limit      int     `json:"limit"`
	Similarity float64 `json:"similarity"` // Minimum similarity threshold
}

// SearchNewsResponse represents the response for semantic search
type SearchNewsResponse struct {
	Results []SearchResult `json:"results"`
	Total   int            `json:"total"`
}

// SearchResult represents a single search result
type SearchResult struct {
	News       *News   `json:"news"`
	Chunk      *NewsChunk `json:"chunk,omitempty"`
	Similarity float64 `json:"similarity"`
}

// GetNewsChunksParams represents parameters for getting news chunks
type GetNewsChunksParams struct {
	NewsID int64  `json:"news_id"`
	Status *EmbeddingStatus `json:"status,omitempty"`
}

// GetNewsChunksResponse represents the response for getting news chunks
type GetNewsChunksResponse struct {
	Chunks []NewsChunk `json:"chunks"`
	Total  int         `json:"total"`
}
