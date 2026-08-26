package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Stock holds the schema definition for the Stock entity.
type Stock struct {
	ent.Schema
}

// Fields of the Stock.
func (Stock) Fields() []ent.Field {
	return []ent.Field{
		field.String("stock_code").
			Unique().
			NotEmpty().
			MaxLen(20).
			Comment("Stock ticker code"),
		field.String("stock_name").
			NotEmpty().
			MaxLen(255).
			Comment("Stock full name"),
	}
}

// Edges of the Stock.
func (Stock) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("summaries", StockSummary.Type),
	}
}

// Indexes of the Stock.
func (Stock) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("stock_code").Unique(),
	}
}
