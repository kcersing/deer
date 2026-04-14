package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// PaymentRecord holds the schema definition for the PaymentRecord entity.
type PaymentRecord struct {
	ent.Schema
}

// Fields of the PaymentRecord.
func (PaymentRecord) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("order_id").
			Comment("Order ID"),
		field.Int("amount").
			Comment("金额/分"),
		field.String("method").
			Comment("支付方式"),

		field.Time("paid_at").
			Optional().
			Nillable().
			Comment("支付时间"),
	}
}

// Mixin of the PaymentRecord.
func (PaymentRecord) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Edges of the PaymentRecord.
func (PaymentRecord) Edges() []ent.Edge {
	return nil
}
