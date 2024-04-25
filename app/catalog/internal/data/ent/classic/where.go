// Code generated by ent, DO NOT EDIT.

package classic

import (
	"entgo.io/ent/dialect/sql"
	"github.com/xuanyiying/publication-shop/app/catalog/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Classic {
	return predicate.Classic(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Classic {
	return predicate.Classic(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Classic {
	return predicate.Classic(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Classic {
	return predicate.Classic(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Classic {
	return predicate.Classic(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Classic {
	return predicate.Classic(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Classic {
	return predicate.Classic(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Classic {
	return predicate.Classic(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Classic {
	return predicate.Classic(sql.FieldLTE(FieldID, id))
}

// ClassicName applies equality check predicate on the "classic_name" field. It's identical to ClassicNameEQ.
func ClassicName(v string) predicate.Classic {
	return predicate.Classic(sql.FieldEQ(FieldClassicName, v))
}

// ClassicNameEQ applies the EQ predicate on the "classic_name" field.
func ClassicNameEQ(v string) predicate.Classic {
	return predicate.Classic(sql.FieldEQ(FieldClassicName, v))
}

// ClassicNameNEQ applies the NEQ predicate on the "classic_name" field.
func ClassicNameNEQ(v string) predicate.Classic {
	return predicate.Classic(sql.FieldNEQ(FieldClassicName, v))
}

// ClassicNameIn applies the In predicate on the "classic_name" field.
func ClassicNameIn(vs ...string) predicate.Classic {
	return predicate.Classic(sql.FieldIn(FieldClassicName, vs...))
}

// ClassicNameNotIn applies the NotIn predicate on the "classic_name" field.
func ClassicNameNotIn(vs ...string) predicate.Classic {
	return predicate.Classic(sql.FieldNotIn(FieldClassicName, vs...))
}

// ClassicNameGT applies the GT predicate on the "classic_name" field.
func ClassicNameGT(v string) predicate.Classic {
	return predicate.Classic(sql.FieldGT(FieldClassicName, v))
}

// ClassicNameGTE applies the GTE predicate on the "classic_name" field.
func ClassicNameGTE(v string) predicate.Classic {
	return predicate.Classic(sql.FieldGTE(FieldClassicName, v))
}

// ClassicNameLT applies the LT predicate on the "classic_name" field.
func ClassicNameLT(v string) predicate.Classic {
	return predicate.Classic(sql.FieldLT(FieldClassicName, v))
}

// ClassicNameLTE applies the LTE predicate on the "classic_name" field.
func ClassicNameLTE(v string) predicate.Classic {
	return predicate.Classic(sql.FieldLTE(FieldClassicName, v))
}

// ClassicNameContains applies the Contains predicate on the "classic_name" field.
func ClassicNameContains(v string) predicate.Classic {
	return predicate.Classic(sql.FieldContains(FieldClassicName, v))
}

// ClassicNameHasPrefix applies the HasPrefix predicate on the "classic_name" field.
func ClassicNameHasPrefix(v string) predicate.Classic {
	return predicate.Classic(sql.FieldHasPrefix(FieldClassicName, v))
}

// ClassicNameHasSuffix applies the HasSuffix predicate on the "classic_name" field.
func ClassicNameHasSuffix(v string) predicate.Classic {
	return predicate.Classic(sql.FieldHasSuffix(FieldClassicName, v))
}

// ClassicNameIsNil applies the IsNil predicate on the "classic_name" field.
func ClassicNameIsNil() predicate.Classic {
	return predicate.Classic(sql.FieldIsNull(FieldClassicName))
}

// ClassicNameNotNil applies the NotNil predicate on the "classic_name" field.
func ClassicNameNotNil() predicate.Classic {
	return predicate.Classic(sql.FieldNotNull(FieldClassicName))
}

// ClassicNameEqualFold applies the EqualFold predicate on the "classic_name" field.
func ClassicNameEqualFold(v string) predicate.Classic {
	return predicate.Classic(sql.FieldEqualFold(FieldClassicName, v))
}

// ClassicNameContainsFold applies the ContainsFold predicate on the "classic_name" field.
func ClassicNameContainsFold(v string) predicate.Classic {
	return predicate.Classic(sql.FieldContainsFold(FieldClassicName, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Classic) predicate.Classic {
	return predicate.Classic(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Classic) predicate.Classic {
	return predicate.Classic(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Classic) predicate.Classic {
	return predicate.Classic(sql.NotPredicates(p))
}
