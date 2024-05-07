// Code generated by ent, DO NOT EDIT.

package bookauthor

import (
	"entgo.io/ent/dialect/sql"
	"github.com/xuanyiying/publication-shop/app/book/internal/data/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldLTE(FieldID, id))
}

// AuthorID applies equality check predicate on the "author_id" field. It's identical to AuthorIDEQ.
func AuthorID(v int) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldEQ(FieldAuthorID, v))
}

// Isbn applies equality check predicate on the "isbn" field. It's identical to IsbnEQ.
func Isbn(v string) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldEQ(FieldIsbn, v))
}

// BookID applies equality check predicate on the "book_id" field. It's identical to BookIDEQ.
func BookID(v int) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldEQ(FieldBookID, v))
}

// Author applies equality check predicate on the "author" field. It's identical to AuthorEQ.
func Author(v string) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldEQ(FieldAuthor, v))
}

// AuthorAbout applies equality check predicate on the "author_about" field. It's identical to AuthorAboutEQ.
func AuthorAbout(v string) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldEQ(FieldAuthorAbout, v))
}

// AuthorIDEQ applies the EQ predicate on the "author_id" field.
func AuthorIDEQ(v int) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldEQ(FieldAuthorID, v))
}

// AuthorIDNEQ applies the NEQ predicate on the "author_id" field.
func AuthorIDNEQ(v int) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldNEQ(FieldAuthorID, v))
}

// AuthorIDIn applies the In predicate on the "author_id" field.
func AuthorIDIn(vs ...int) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldIn(FieldAuthorID, vs...))
}

// AuthorIDNotIn applies the NotIn predicate on the "author_id" field.
func AuthorIDNotIn(vs ...int) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldNotIn(FieldAuthorID, vs...))
}

// AuthorIDGT applies the GT predicate on the "author_id" field.
func AuthorIDGT(v int) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldGT(FieldAuthorID, v))
}

// AuthorIDGTE applies the GTE predicate on the "author_id" field.
func AuthorIDGTE(v int) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldGTE(FieldAuthorID, v))
}

// AuthorIDLT applies the LT predicate on the "author_id" field.
func AuthorIDLT(v int) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldLT(FieldAuthorID, v))
}

// AuthorIDLTE applies the LTE predicate on the "author_id" field.
func AuthorIDLTE(v int) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldLTE(FieldAuthorID, v))
}

// IsbnEQ applies the EQ predicate on the "isbn" field.
func IsbnEQ(v string) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldEQ(FieldIsbn, v))
}

// IsbnNEQ applies the NEQ predicate on the "isbn" field.
func IsbnNEQ(v string) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldNEQ(FieldIsbn, v))
}

// IsbnIn applies the In predicate on the "isbn" field.
func IsbnIn(vs ...string) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldIn(FieldIsbn, vs...))
}

// IsbnNotIn applies the NotIn predicate on the "isbn" field.
func IsbnNotIn(vs ...string) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldNotIn(FieldIsbn, vs...))
}

// IsbnGT applies the GT predicate on the "isbn" field.
func IsbnGT(v string) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldGT(FieldIsbn, v))
}

// IsbnGTE applies the GTE predicate on the "isbn" field.
func IsbnGTE(v string) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldGTE(FieldIsbn, v))
}

// IsbnLT applies the LT predicate on the "isbn" field.
func IsbnLT(v string) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldLT(FieldIsbn, v))
}

// IsbnLTE applies the LTE predicate on the "isbn" field.
func IsbnLTE(v string) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldLTE(FieldIsbn, v))
}

// IsbnContains applies the Contains predicate on the "isbn" field.
func IsbnContains(v string) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldContains(FieldIsbn, v))
}

// IsbnHasPrefix applies the HasPrefix predicate on the "isbn" field.
func IsbnHasPrefix(v string) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldHasPrefix(FieldIsbn, v))
}

// IsbnHasSuffix applies the HasSuffix predicate on the "isbn" field.
func IsbnHasSuffix(v string) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldHasSuffix(FieldIsbn, v))
}

// IsbnEqualFold applies the EqualFold predicate on the "isbn" field.
func IsbnEqualFold(v string) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldEqualFold(FieldIsbn, v))
}

// IsbnContainsFold applies the ContainsFold predicate on the "isbn" field.
func IsbnContainsFold(v string) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldContainsFold(FieldIsbn, v))
}

// BookIDEQ applies the EQ predicate on the "book_id" field.
func BookIDEQ(v int) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldEQ(FieldBookID, v))
}

// BookIDNEQ applies the NEQ predicate on the "book_id" field.
func BookIDNEQ(v int) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldNEQ(FieldBookID, v))
}

// BookIDIn applies the In predicate on the "book_id" field.
func BookIDIn(vs ...int) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldIn(FieldBookID, vs...))
}

// BookIDNotIn applies the NotIn predicate on the "book_id" field.
func BookIDNotIn(vs ...int) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldNotIn(FieldBookID, vs...))
}

// BookIDGT applies the GT predicate on the "book_id" field.
func BookIDGT(v int) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldGT(FieldBookID, v))
}

// BookIDGTE applies the GTE predicate on the "book_id" field.
func BookIDGTE(v int) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldGTE(FieldBookID, v))
}

// BookIDLT applies the LT predicate on the "book_id" field.
func BookIDLT(v int) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldLT(FieldBookID, v))
}

// BookIDLTE applies the LTE predicate on the "book_id" field.
func BookIDLTE(v int) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldLTE(FieldBookID, v))
}

// BookIDIsNil applies the IsNil predicate on the "book_id" field.
func BookIDIsNil() predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldIsNull(FieldBookID))
}

// BookIDNotNil applies the NotNil predicate on the "book_id" field.
func BookIDNotNil() predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldNotNull(FieldBookID))
}

// AuthorEQ applies the EQ predicate on the "author" field.
func AuthorEQ(v string) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldEQ(FieldAuthor, v))
}

// AuthorNEQ applies the NEQ predicate on the "author" field.
func AuthorNEQ(v string) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldNEQ(FieldAuthor, v))
}

// AuthorIn applies the In predicate on the "author" field.
func AuthorIn(vs ...string) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldIn(FieldAuthor, vs...))
}

// AuthorNotIn applies the NotIn predicate on the "author" field.
func AuthorNotIn(vs ...string) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldNotIn(FieldAuthor, vs...))
}

// AuthorGT applies the GT predicate on the "author" field.
func AuthorGT(v string) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldGT(FieldAuthor, v))
}

// AuthorGTE applies the GTE predicate on the "author" field.
func AuthorGTE(v string) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldGTE(FieldAuthor, v))
}

// AuthorLT applies the LT predicate on the "author" field.
func AuthorLT(v string) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldLT(FieldAuthor, v))
}

// AuthorLTE applies the LTE predicate on the "author" field.
func AuthorLTE(v string) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldLTE(FieldAuthor, v))
}

// AuthorContains applies the Contains predicate on the "author" field.
func AuthorContains(v string) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldContains(FieldAuthor, v))
}

// AuthorHasPrefix applies the HasPrefix predicate on the "author" field.
func AuthorHasPrefix(v string) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldHasPrefix(FieldAuthor, v))
}

// AuthorHasSuffix applies the HasSuffix predicate on the "author" field.
func AuthorHasSuffix(v string) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldHasSuffix(FieldAuthor, v))
}

// AuthorIsNil applies the IsNil predicate on the "author" field.
func AuthorIsNil() predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldIsNull(FieldAuthor))
}

// AuthorNotNil applies the NotNil predicate on the "author" field.
func AuthorNotNil() predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldNotNull(FieldAuthor))
}

// AuthorEqualFold applies the EqualFold predicate on the "author" field.
func AuthorEqualFold(v string) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldEqualFold(FieldAuthor, v))
}

// AuthorContainsFold applies the ContainsFold predicate on the "author" field.
func AuthorContainsFold(v string) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldContainsFold(FieldAuthor, v))
}

// AuthorAboutEQ applies the EQ predicate on the "author_about" field.
func AuthorAboutEQ(v string) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldEQ(FieldAuthorAbout, v))
}

// AuthorAboutNEQ applies the NEQ predicate on the "author_about" field.
func AuthorAboutNEQ(v string) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldNEQ(FieldAuthorAbout, v))
}

// AuthorAboutIn applies the In predicate on the "author_about" field.
func AuthorAboutIn(vs ...string) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldIn(FieldAuthorAbout, vs...))
}

// AuthorAboutNotIn applies the NotIn predicate on the "author_about" field.
func AuthorAboutNotIn(vs ...string) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldNotIn(FieldAuthorAbout, vs...))
}

// AuthorAboutGT applies the GT predicate on the "author_about" field.
func AuthorAboutGT(v string) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldGT(FieldAuthorAbout, v))
}

// AuthorAboutGTE applies the GTE predicate on the "author_about" field.
func AuthorAboutGTE(v string) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldGTE(FieldAuthorAbout, v))
}

// AuthorAboutLT applies the LT predicate on the "author_about" field.
func AuthorAboutLT(v string) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldLT(FieldAuthorAbout, v))
}

// AuthorAboutLTE applies the LTE predicate on the "author_about" field.
func AuthorAboutLTE(v string) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldLTE(FieldAuthorAbout, v))
}

// AuthorAboutContains applies the Contains predicate on the "author_about" field.
func AuthorAboutContains(v string) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldContains(FieldAuthorAbout, v))
}

// AuthorAboutHasPrefix applies the HasPrefix predicate on the "author_about" field.
func AuthorAboutHasPrefix(v string) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldHasPrefix(FieldAuthorAbout, v))
}

// AuthorAboutHasSuffix applies the HasSuffix predicate on the "author_about" field.
func AuthorAboutHasSuffix(v string) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldHasSuffix(FieldAuthorAbout, v))
}

// AuthorAboutIsNil applies the IsNil predicate on the "author_about" field.
func AuthorAboutIsNil() predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldIsNull(FieldAuthorAbout))
}

// AuthorAboutNotNil applies the NotNil predicate on the "author_about" field.
func AuthorAboutNotNil() predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldNotNull(FieldAuthorAbout))
}

// AuthorAboutEqualFold applies the EqualFold predicate on the "author_about" field.
func AuthorAboutEqualFold(v string) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldEqualFold(FieldAuthorAbout, v))
}

// AuthorAboutContainsFold applies the ContainsFold predicate on the "author_about" field.
func AuthorAboutContainsFold(v string) predicate.BookAuthor {
	return predicate.BookAuthor(sql.FieldContainsFold(FieldAuthorAbout, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.BookAuthor) predicate.BookAuthor {
	return predicate.BookAuthor(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.BookAuthor) predicate.BookAuthor {
	return predicate.BookAuthor(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.BookAuthor) predicate.BookAuthor {
	return predicate.BookAuthor(sql.NotPredicates(p))
}