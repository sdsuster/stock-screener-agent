package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// NewsChunk holds the schema definition for the NewsChunk entity.
type NewsChunk struct {
	ent.Schema
}

// Fields of the NewsChunk.
func (NewsChunk) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("news_id").
			Positive().
			Comment("Foreign key to news table"),
		field.String("embedding_type").
			NotEmpty().
			Comment("Type of embedding: 'title' or 'content'"),
		field.Int("chunk_index").
			NonNegative().
			Comment("Index of the chunk within the article"),
		field.Text("content").
			NotEmpty().
			Comment("Text content of this chunk"),
		field.Enum("status").
			Values("pending", "processing", "completed", "failed").
			Default("pending").
			Comment("Embedding processing status"),
		field.JSON("embedding", []float32{}).
			Optional().
			Comment("Vector embedding (1024 dimensions for pgvector)"),
	}
}

// Edges of the NewsChunk.
func (NewsChunk) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("news", News.Type).
			Ref("chunks").
			Field("news_id").
			Required().
			Unique(),
	}
}

// Indexes of the NewsChunk.
func (NewsChunk) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("news_id", "chunk_index").Unique(),
		index.Fields("news_id"),
		index.Fields("status"),
		index.Fields("embedding_type"),
	}
}
