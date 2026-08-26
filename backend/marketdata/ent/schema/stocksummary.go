package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// StockSummary holds the schema definition for the StockSummary entity.
type StockSummary struct {
	ent.Schema
}

// Fields of the StockSummary.
func (StockSummary) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id_stock_summary").
			Positive().
			Comment("External ID from data source"),
		field.Time("date").
			Comment("Trading date"),
		field.String("stock_code").
			NotEmpty().
			MaxLen(20).
			Comment("Stock ticker code"),
		field.String("stock_name").
			NotEmpty().
			MaxLen(255).
			Comment("Stock name"),
		field.String("remarks").
			Default("").
			Comment("Trading remarks"),
		
		// Price fields
		field.Int64("previous").Default(0).Comment("Previous close price"),
		field.Int64("open_price").Default(0).Comment("Opening price"),
		field.Int64("first_trade").Default(0).Comment("First trade price"),
		field.Int64("high").Default(0).Comment("Highest price"),
		field.Int64("low").Default(0).Comment("Lowest price"),
		field.Int64("close").Default(0).Comment("Closing price"),
		field.Int64("change").Default(0).Comment("Price change"),
		
		// Volume fields
		field.Int64("volume").Default(0).Comment("Trading volume"),
		field.Int64("value").Default(0).Comment("Trading value"),
		field.Int64("frequency").Default(0).Comment("Trading frequency"),
		
		// Index and offer/bid fields
		field.Float("index_individual").Default(0).Comment("Individual stock index"),
		field.Int64("offer").Default(0).Comment("Offer price"),
		field.Int64("offer_volume").Default(0).Comment("Offer volume"),
		field.Int64("bid").Default(0).Comment("Bid price"),
		field.Int64("bid_volume").Default(0).Comment("Bid volume"),
		
		// Share fields
		field.Int64("listed_shares").Default(0).Comment("Listed shares"),
		field.Int64("tradeable_shares").Default(0).Comment("Tradeable shares"),
		field.Int64("weight_for_index").Default(0).Comment("Weight for index calculation"),
		
		// Foreign transaction fields
		field.Int64("foreign_sell").Default(0).Comment("Foreign sell volume"),
		field.Int64("foreign_buy").Default(0).Comment("Foreign buy volume"),
		
		// Non-regular market fields
		field.Int64("non_regular_volume").Default(0).Comment("Non-regular market volume"),
		field.Int64("non_regular_value").Default(0).Comment("Non-regular market value"),
		field.Int64("non_regular_frequency").Default(0).Comment("Non-regular market frequency"),
		
		// Optional fields
		field.Time("delisting_date").
			Optional().
			Nillable().
			Comment("Delisting date if applicable"),
		field.Float("percentage").
			Optional().
			Nillable().
			Comment("Price change percentage"),
	}
}

// Edges of the StockSummary.
func (StockSummary) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("stock", Stock.Type).
			Ref("summaries").
			Field("stock_code").
			Required().
			Unique(),
	}
}

// Indexes of the StockSummary.
func (StockSummary) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("stock_code", "date").Unique(),
		index.Fields("date"),
		index.Fields("stock_code"),
	}
}
