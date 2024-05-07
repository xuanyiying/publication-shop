// Code generated by ent, DO NOT EDIT.

package author

import (
	"entgo.io/ent/dialect/sql"
	"github.com/xuanyiying/publication-shop/app/book/internal/data/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int32) predicate.Author {
	return predicate.Author(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int32) predicate.Author {
	return predicate.Author(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int32) predicate.Author {
	return predicate.Author(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int32) predicate.Author {
	return predicate.Author(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int32) predicate.Author {
	return predicate.Author(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int32) predicate.Author {
	return predicate.Author(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int32) predicate.Author {
	return predicate.Author(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int32) predicate.Author {
	return predicate.Author(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int32) predicate.Author {
	return predicate.Author(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Author {
	return predicate.Author(sql.FieldEQ(FieldName, v))
}

// Biography applies equality check predicate on the "biography" field. It's identical to BiographyEQ.
func Biography(v string) predicate.Author {
	return predicate.Author(sql.FieldEQ(FieldBiography, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Author {
	return predicate.Author(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Author {
	return predicate.Author(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Author {
	return predicate.Author(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Author {
	return predicate.Author(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Author {
	return predicate.Author(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Author {
	return predicate.Author(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Author {
	return predicate.Author(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Author {
	return predicate.Author(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Author {
	return predicate.Author(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Author {
	return predicate.Author(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Author {
	return predicate.Author(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Author {
	return predicate.Author(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Author {
	return predicate.Author(sql.FieldContainsFold(FieldName, v))
}

// BiographyEQ applies the EQ predicate on the "biography" field.
func BiographyEQ(v string) predicate.Author {
	return predicate.Author(sql.FieldEQ(FieldBiography, v))
}

// BiographyNEQ applies the NEQ predicate on the "biography" field.
func BiographyNEQ(v string) predicate.Author {
	return predicate.Author(sql.FieldNEQ(FieldBiography, v))
}

// BiographyIn applies the In predicate on the "biography" field.
func BiographyIn(vs ...string) predicate.Author {
	return predicate.Author(sql.FieldIn(FieldBiography, vs...))
}

// BiographyNotIn applies the NotIn predicate on the "biography" field.
func BiographyNotIn(vs ...string) predicate.Author {
	return predicate.Author(sql.FieldNotIn(FieldBiography, vs...))
}

// BiographyGT applies the GT predicate on the "biography" field.
func BiographyGT(v string) predicate.Author {
	return predicate.Author(sql.FieldGT(FieldBiography, v))
}

// BiographyGTE applies the GTE predicate on the "biography" field.
func BiographyGTE(v string) predicate.Author {
	return predicate.Author(sql.FieldGTE(FieldBiography, v))
}

// BiographyLT applies the LT predicate on the "biography" field.
func BiographyLT(v string) predicate.Author {
	return predicate.Author(sql.FieldLT(FieldBiography, v))
}

// BiographyLTE applies the LTE predicate on the "biography" field.
func BiographyLTE(v string) predicate.Author {
	return predicate.Author(sql.FieldLTE(FieldBiography, v))
}

// BiographyContains applies the Contains predicate on the "biography" field.
func BiographyContains(v string) predicate.Author {
	return predicate.Author(sql.FieldContains(FieldBiography, v))
}

// BiographyHasPrefix applies the HasPrefix predicate on the "biography" field.
func BiographyHasPrefix(v string) predicate.Author {
	return predicate.Author(sql.FieldHasPrefix(FieldBiography, v))
}

// BiographyHasSuffix applies the HasSuffix predicate on the "biography" field.
func BiographyHasSuffix(v string) predicate.Author {
	return predicate.Author(sql.FieldHasSuffix(FieldBiography, v))
}

// BiographyIsNil applies the IsNil predicate on the "biography" field.
func BiographyIsNil() predicate.Author {
	return predicate.Author(sql.FieldIsNull(FieldBiography))
}

// BiographyNotNil applies the NotNil predicate on the "biography" field.
func BiographyNotNil() predicate.Author {
	return predicate.Author(sql.FieldNotNull(FieldBiography))
}

// BiographyEqualFold applies the EqualFold predicate on the "biography" field.
func BiographyEqualFold(v string) predicate.Author {
	return predicate.Author(sql.FieldEqualFold(FieldBiography, v))
}

// BiographyContainsFold applies the ContainsFold predicate on the "biography" field.
func BiographyContainsFold(v string) predicate.Author {
	return predicate.Author(sql.FieldContainsFold(FieldBiography, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Author) predicate.Author {
	return predicate.Author(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Author) predicate.Author {
	return predicate.Author(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Author) predicate.Author {
	return predicate.Author(sql.NotPredicates(p))
}