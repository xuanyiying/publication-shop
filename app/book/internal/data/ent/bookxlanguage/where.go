// Code generated by ent, DO NOT EDIT.

package bookxlanguage

import (
	"entgo.io/ent/dialect/sql"
	"github.com/xuanyiying/publication-shop/app/book/internal/data/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldLTE(FieldID, id))
}

// LanguageID applies equality check predicate on the "language_id" field. It's identical to LanguageIDEQ.
func LanguageID(v int) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldEQ(FieldLanguageID, v))
}

// Isbn applies equality check predicate on the "isbn" field. It's identical to IsbnEQ.
func Isbn(v string) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldEQ(FieldIsbn, v))
}

// LanguageName applies equality check predicate on the "language_name" field. It's identical to LanguageNameEQ.
func LanguageName(v string) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldEQ(FieldLanguageName, v))
}

// BookID applies equality check predicate on the "book_id" field. It's identical to BookIDEQ.
func BookID(v int) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldEQ(FieldBookID, v))
}

// LanguageIDEQ applies the EQ predicate on the "language_id" field.
func LanguageIDEQ(v int) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldEQ(FieldLanguageID, v))
}

// LanguageIDNEQ applies the NEQ predicate on the "language_id" field.
func LanguageIDNEQ(v int) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldNEQ(FieldLanguageID, v))
}

// LanguageIDIn applies the In predicate on the "language_id" field.
func LanguageIDIn(vs ...int) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldIn(FieldLanguageID, vs...))
}

// LanguageIDNotIn applies the NotIn predicate on the "language_id" field.
func LanguageIDNotIn(vs ...int) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldNotIn(FieldLanguageID, vs...))
}

// LanguageIDGT applies the GT predicate on the "language_id" field.
func LanguageIDGT(v int) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldGT(FieldLanguageID, v))
}

// LanguageIDGTE applies the GTE predicate on the "language_id" field.
func LanguageIDGTE(v int) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldGTE(FieldLanguageID, v))
}

// LanguageIDLT applies the LT predicate on the "language_id" field.
func LanguageIDLT(v int) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldLT(FieldLanguageID, v))
}

// LanguageIDLTE applies the LTE predicate on the "language_id" field.
func LanguageIDLTE(v int) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldLTE(FieldLanguageID, v))
}

// IsbnEQ applies the EQ predicate on the "isbn" field.
func IsbnEQ(v string) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldEQ(FieldIsbn, v))
}

// IsbnNEQ applies the NEQ predicate on the "isbn" field.
func IsbnNEQ(v string) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldNEQ(FieldIsbn, v))
}

// IsbnIn applies the In predicate on the "isbn" field.
func IsbnIn(vs ...string) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldIn(FieldIsbn, vs...))
}

// IsbnNotIn applies the NotIn predicate on the "isbn" field.
func IsbnNotIn(vs ...string) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldNotIn(FieldIsbn, vs...))
}

// IsbnGT applies the GT predicate on the "isbn" field.
func IsbnGT(v string) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldGT(FieldIsbn, v))
}

// IsbnGTE applies the GTE predicate on the "isbn" field.
func IsbnGTE(v string) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldGTE(FieldIsbn, v))
}

// IsbnLT applies the LT predicate on the "isbn" field.
func IsbnLT(v string) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldLT(FieldIsbn, v))
}

// IsbnLTE applies the LTE predicate on the "isbn" field.
func IsbnLTE(v string) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldLTE(FieldIsbn, v))
}

// IsbnContains applies the Contains predicate on the "isbn" field.
func IsbnContains(v string) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldContains(FieldIsbn, v))
}

// IsbnHasPrefix applies the HasPrefix predicate on the "isbn" field.
func IsbnHasPrefix(v string) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldHasPrefix(FieldIsbn, v))
}

// IsbnHasSuffix applies the HasSuffix predicate on the "isbn" field.
func IsbnHasSuffix(v string) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldHasSuffix(FieldIsbn, v))
}

// IsbnEqualFold applies the EqualFold predicate on the "isbn" field.
func IsbnEqualFold(v string) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldEqualFold(FieldIsbn, v))
}

// IsbnContainsFold applies the ContainsFold predicate on the "isbn" field.
func IsbnContainsFold(v string) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldContainsFold(FieldIsbn, v))
}

// LanguageNameEQ applies the EQ predicate on the "language_name" field.
func LanguageNameEQ(v string) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldEQ(FieldLanguageName, v))
}

// LanguageNameNEQ applies the NEQ predicate on the "language_name" field.
func LanguageNameNEQ(v string) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldNEQ(FieldLanguageName, v))
}

// LanguageNameIn applies the In predicate on the "language_name" field.
func LanguageNameIn(vs ...string) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldIn(FieldLanguageName, vs...))
}

// LanguageNameNotIn applies the NotIn predicate on the "language_name" field.
func LanguageNameNotIn(vs ...string) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldNotIn(FieldLanguageName, vs...))
}

// LanguageNameGT applies the GT predicate on the "language_name" field.
func LanguageNameGT(v string) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldGT(FieldLanguageName, v))
}

// LanguageNameGTE applies the GTE predicate on the "language_name" field.
func LanguageNameGTE(v string) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldGTE(FieldLanguageName, v))
}

// LanguageNameLT applies the LT predicate on the "language_name" field.
func LanguageNameLT(v string) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldLT(FieldLanguageName, v))
}

// LanguageNameLTE applies the LTE predicate on the "language_name" field.
func LanguageNameLTE(v string) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldLTE(FieldLanguageName, v))
}

// LanguageNameContains applies the Contains predicate on the "language_name" field.
func LanguageNameContains(v string) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldContains(FieldLanguageName, v))
}

// LanguageNameHasPrefix applies the HasPrefix predicate on the "language_name" field.
func LanguageNameHasPrefix(v string) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldHasPrefix(FieldLanguageName, v))
}

// LanguageNameHasSuffix applies the HasSuffix predicate on the "language_name" field.
func LanguageNameHasSuffix(v string) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldHasSuffix(FieldLanguageName, v))
}

// LanguageNameIsNil applies the IsNil predicate on the "language_name" field.
func LanguageNameIsNil() predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldIsNull(FieldLanguageName))
}

// LanguageNameNotNil applies the NotNil predicate on the "language_name" field.
func LanguageNameNotNil() predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldNotNull(FieldLanguageName))
}

// LanguageNameEqualFold applies the EqualFold predicate on the "language_name" field.
func LanguageNameEqualFold(v string) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldEqualFold(FieldLanguageName, v))
}

// LanguageNameContainsFold applies the ContainsFold predicate on the "language_name" field.
func LanguageNameContainsFold(v string) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldContainsFold(FieldLanguageName, v))
}

// BookIDEQ applies the EQ predicate on the "book_id" field.
func BookIDEQ(v int) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldEQ(FieldBookID, v))
}

// BookIDNEQ applies the NEQ predicate on the "book_id" field.
func BookIDNEQ(v int) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldNEQ(FieldBookID, v))
}

// BookIDIn applies the In predicate on the "book_id" field.
func BookIDIn(vs ...int) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldIn(FieldBookID, vs...))
}

// BookIDNotIn applies the NotIn predicate on the "book_id" field.
func BookIDNotIn(vs ...int) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldNotIn(FieldBookID, vs...))
}

// BookIDGT applies the GT predicate on the "book_id" field.
func BookIDGT(v int) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldGT(FieldBookID, v))
}

// BookIDGTE applies the GTE predicate on the "book_id" field.
func BookIDGTE(v int) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldGTE(FieldBookID, v))
}

// BookIDLT applies the LT predicate on the "book_id" field.
func BookIDLT(v int) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldLT(FieldBookID, v))
}

// BookIDLTE applies the LTE predicate on the "book_id" field.
func BookIDLTE(v int) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldLTE(FieldBookID, v))
}

// BookIDIsNil applies the IsNil predicate on the "book_id" field.
func BookIDIsNil() predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldIsNull(FieldBookID))
}

// BookIDNotNil applies the NotNil predicate on the "book_id" field.
func BookIDNotNil() predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.FieldNotNull(FieldBookID))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.BookXLanguage) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.BookXLanguage) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.BookXLanguage) predicate.BookXLanguage {
	return predicate.BookXLanguage(sql.NotPredicates(p))
}