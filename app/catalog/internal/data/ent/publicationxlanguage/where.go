// Code generated by ent, DO NOT EDIT.

package publicationxlanguage

import (
	"entgo.io/ent/dialect/sql"
	"github.com/xuanyiying/publication-shop/app/catalog/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldLTE(FieldID, id))
}

// LanguageID applies equality check predicate on the "language_id" field. It's identical to LanguageIDEQ.
func LanguageID(v int) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldEQ(FieldLanguageID, v))
}

// Isbn applies equality check predicate on the "isbn" field. It's identical to IsbnEQ.
func Isbn(v string) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldEQ(FieldIsbn, v))
}

// LanguageName applies equality check predicate on the "language_name" field. It's identical to LanguageNameEQ.
func LanguageName(v string) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldEQ(FieldLanguageName, v))
}

// PublicationID applies equality check predicate on the "publication_id" field. It's identical to PublicationIDEQ.
func PublicationID(v int) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldEQ(FieldPublicationID, v))
}

// LanguageIDEQ applies the EQ predicate on the "language_id" field.
func LanguageIDEQ(v int) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldEQ(FieldLanguageID, v))
}

// LanguageIDNEQ applies the NEQ predicate on the "language_id" field.
func LanguageIDNEQ(v int) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldNEQ(FieldLanguageID, v))
}

// LanguageIDIn applies the In predicate on the "language_id" field.
func LanguageIDIn(vs ...int) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldIn(FieldLanguageID, vs...))
}

// LanguageIDNotIn applies the NotIn predicate on the "language_id" field.
func LanguageIDNotIn(vs ...int) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldNotIn(FieldLanguageID, vs...))
}

// LanguageIDGT applies the GT predicate on the "language_id" field.
func LanguageIDGT(v int) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldGT(FieldLanguageID, v))
}

// LanguageIDGTE applies the GTE predicate on the "language_id" field.
func LanguageIDGTE(v int) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldGTE(FieldLanguageID, v))
}

// LanguageIDLT applies the LT predicate on the "language_id" field.
func LanguageIDLT(v int) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldLT(FieldLanguageID, v))
}

// LanguageIDLTE applies the LTE predicate on the "language_id" field.
func LanguageIDLTE(v int) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldLTE(FieldLanguageID, v))
}

// IsbnEQ applies the EQ predicate on the "isbn" field.
func IsbnEQ(v string) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldEQ(FieldIsbn, v))
}

// IsbnNEQ applies the NEQ predicate on the "isbn" field.
func IsbnNEQ(v string) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldNEQ(FieldIsbn, v))
}

// IsbnIn applies the In predicate on the "isbn" field.
func IsbnIn(vs ...string) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldIn(FieldIsbn, vs...))
}

// IsbnNotIn applies the NotIn predicate on the "isbn" field.
func IsbnNotIn(vs ...string) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldNotIn(FieldIsbn, vs...))
}

// IsbnGT applies the GT predicate on the "isbn" field.
func IsbnGT(v string) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldGT(FieldIsbn, v))
}

// IsbnGTE applies the GTE predicate on the "isbn" field.
func IsbnGTE(v string) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldGTE(FieldIsbn, v))
}

// IsbnLT applies the LT predicate on the "isbn" field.
func IsbnLT(v string) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldLT(FieldIsbn, v))
}

// IsbnLTE applies the LTE predicate on the "isbn" field.
func IsbnLTE(v string) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldLTE(FieldIsbn, v))
}

// IsbnContains applies the Contains predicate on the "isbn" field.
func IsbnContains(v string) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldContains(FieldIsbn, v))
}

// IsbnHasPrefix applies the HasPrefix predicate on the "isbn" field.
func IsbnHasPrefix(v string) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldHasPrefix(FieldIsbn, v))
}

// IsbnHasSuffix applies the HasSuffix predicate on the "isbn" field.
func IsbnHasSuffix(v string) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldHasSuffix(FieldIsbn, v))
}

// IsbnEqualFold applies the EqualFold predicate on the "isbn" field.
func IsbnEqualFold(v string) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldEqualFold(FieldIsbn, v))
}

// IsbnContainsFold applies the ContainsFold predicate on the "isbn" field.
func IsbnContainsFold(v string) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldContainsFold(FieldIsbn, v))
}

// LanguageNameEQ applies the EQ predicate on the "language_name" field.
func LanguageNameEQ(v string) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldEQ(FieldLanguageName, v))
}

// LanguageNameNEQ applies the NEQ predicate on the "language_name" field.
func LanguageNameNEQ(v string) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldNEQ(FieldLanguageName, v))
}

// LanguageNameIn applies the In predicate on the "language_name" field.
func LanguageNameIn(vs ...string) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldIn(FieldLanguageName, vs...))
}

// LanguageNameNotIn applies the NotIn predicate on the "language_name" field.
func LanguageNameNotIn(vs ...string) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldNotIn(FieldLanguageName, vs...))
}

// LanguageNameGT applies the GT predicate on the "language_name" field.
func LanguageNameGT(v string) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldGT(FieldLanguageName, v))
}

// LanguageNameGTE applies the GTE predicate on the "language_name" field.
func LanguageNameGTE(v string) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldGTE(FieldLanguageName, v))
}

// LanguageNameLT applies the LT predicate on the "language_name" field.
func LanguageNameLT(v string) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldLT(FieldLanguageName, v))
}

// LanguageNameLTE applies the LTE predicate on the "language_name" field.
func LanguageNameLTE(v string) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldLTE(FieldLanguageName, v))
}

// LanguageNameContains applies the Contains predicate on the "language_name" field.
func LanguageNameContains(v string) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldContains(FieldLanguageName, v))
}

// LanguageNameHasPrefix applies the HasPrefix predicate on the "language_name" field.
func LanguageNameHasPrefix(v string) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldHasPrefix(FieldLanguageName, v))
}

// LanguageNameHasSuffix applies the HasSuffix predicate on the "language_name" field.
func LanguageNameHasSuffix(v string) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldHasSuffix(FieldLanguageName, v))
}

// LanguageNameIsNil applies the IsNil predicate on the "language_name" field.
func LanguageNameIsNil() predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldIsNull(FieldLanguageName))
}

// LanguageNameNotNil applies the NotNil predicate on the "language_name" field.
func LanguageNameNotNil() predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldNotNull(FieldLanguageName))
}

// LanguageNameEqualFold applies the EqualFold predicate on the "language_name" field.
func LanguageNameEqualFold(v string) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldEqualFold(FieldLanguageName, v))
}

// LanguageNameContainsFold applies the ContainsFold predicate on the "language_name" field.
func LanguageNameContainsFold(v string) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldContainsFold(FieldLanguageName, v))
}

// PublicationIDEQ applies the EQ predicate on the "publication_id" field.
func PublicationIDEQ(v int) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldEQ(FieldPublicationID, v))
}

// PublicationIDNEQ applies the NEQ predicate on the "publication_id" field.
func PublicationIDNEQ(v int) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldNEQ(FieldPublicationID, v))
}

// PublicationIDIn applies the In predicate on the "publication_id" field.
func PublicationIDIn(vs ...int) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldIn(FieldPublicationID, vs...))
}

// PublicationIDNotIn applies the NotIn predicate on the "publication_id" field.
func PublicationIDNotIn(vs ...int) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldNotIn(FieldPublicationID, vs...))
}

// PublicationIDGT applies the GT predicate on the "publication_id" field.
func PublicationIDGT(v int) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldGT(FieldPublicationID, v))
}

// PublicationIDGTE applies the GTE predicate on the "publication_id" field.
func PublicationIDGTE(v int) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldGTE(FieldPublicationID, v))
}

// PublicationIDLT applies the LT predicate on the "publication_id" field.
func PublicationIDLT(v int) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldLT(FieldPublicationID, v))
}

// PublicationIDLTE applies the LTE predicate on the "publication_id" field.
func PublicationIDLTE(v int) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldLTE(FieldPublicationID, v))
}

// PublicationIDIsNil applies the IsNil predicate on the "publication_id" field.
func PublicationIDIsNil() predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldIsNull(FieldPublicationID))
}

// PublicationIDNotNil applies the NotNil predicate on the "publication_id" field.
func PublicationIDNotNil() predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.FieldNotNull(FieldPublicationID))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PublicationXLanguage) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PublicationXLanguage) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.PublicationXLanguage) predicate.PublicationXLanguage {
	return predicate.PublicationXLanguage(sql.NotPredicates(p))
}