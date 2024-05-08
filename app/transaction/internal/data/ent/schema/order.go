// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Order struct {
	ent.Schema
}

func (Order) Fields() []ent.Field {
	return []ent.Field{field.Int("id").Comment("订单ID，主键").StorageKey("order_id"), field.String("order_no").Optional().Comment("订单编号"), field.Int("tx_id").Comment("交易ID"), field.Enum("order_status").Comment("订单状态").Values("pending", "processing", "shipped", "delivered", "cancelled"), field.String("delivered_address").Optional().Comment("配送地址"), field.Float("shipping_cost").Optional().Comment("运费"), field.Float("total_amount").Optional().Comment("总金额"), field.Int("placed_user_id").Comment("下单用户ID"), field.Time("placed_at").Optional().Comment("下单时间"), field.String("shipped_address").Optional().Comment("发货地址"), field.Time("shipped_at").Optional().Comment("发货时间"), field.Int("payment_id").Optional().Comment("支付ID")}
}
func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("transaction", Transaction.Type),
		edge.To("txItems", TxItem.Type),
	}
}
func (Order) Annotations() []schema.Annotation {
	return nil
}
