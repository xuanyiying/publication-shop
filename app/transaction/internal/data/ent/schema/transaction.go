// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Transaction struct {
	ent.Schema
}

func (Transaction) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Comment("交易ID，主键").StorageKey("tx_id"),
		field.String("tx_no").Comment("交易编号"),
		field.Enum("tx_type").Comment("交易类型（买或卖）").Values("buy", "sell"),
		field.Int64("user_id").Comment("用户ID"),
		field.Int("quantity").Comment("交易数量"),
		field.Enum("tx_status").Comment("交易状态").Values("pending", "completed", "cancelled"),
		field.Time("tx_date").Optional().Comment("交易日期"),
		field.Float("tx_amount").Optional().Comment("交易金额"),
		field.Int64("payment_id").Optional().Comment("支付ID")}
}
func (Transaction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("txItems", TxItem.Type),
		edge.To("txOrder", Order.Type),
	}
}
func (Transaction) Annotations() []schema.Annotation {
	return nil
}
