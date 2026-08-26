package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// News holds the schema definition for the News entity.
type News struct {
	ent.Schema
}

// Fields of the News.
func (News) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
			NotEmpty().
			Comment("News article title"),
		field.Text("content").
			NotEmpty().
			Comment("Full news article content"),
		field.String("url").
			Unique().
			MaxLen(2048).
			Comment("News article URL"),
		field.Time("published_at").
			Optional().
			Nillable().
			Comment("Publication date"),
		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Comment("Creation timestamp"),
	}
}

// Edges of the News.
func (News) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("chunks", NewsChunk.Type).
			Comment("News content chunks for embeddings"),
	}
}

// Indexes of the News.
func (News) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("url").Unique(),
		index.Fields("created_at"),
		index.Fields("published_at"),
	}
}
