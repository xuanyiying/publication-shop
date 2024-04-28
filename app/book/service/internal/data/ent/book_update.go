// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/xuanyiying/publication-shop/app/book/service/internal/data/ent/book"
	"github.com/xuanyiying/publication-shop/app/book/service/internal/data/ent/predicate"
)

// BookUpdate is the builder for updating Book entities.
type BookUpdate struct {
	config
	hooks    []Hook
	mutation *BookMutation
}

// Where appends a list predicates to the BookUpdate builder.
func (bu *BookUpdate) Where(ps ...predicate.Book) *BookUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetIsbn sets the "isbn" field.
func (bu *BookUpdate) SetIsbn(s string) *BookUpdate {
	bu.mutation.SetIsbn(s)
	return bu
}

// SetNillableIsbn sets the "isbn" field if the given value is not nil.
func (bu *BookUpdate) SetNillableIsbn(s *string) *BookUpdate {
	if s != nil {
		bu.SetIsbn(*s)
	}
	return bu
}

// SetTitle sets the "title" field.
func (bu *BookUpdate) SetTitle(s string) *BookUpdate {
	bu.mutation.SetTitle(s)
	return bu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (bu *BookUpdate) SetNillableTitle(s *string) *BookUpdate {
	if s != nil {
		bu.SetTitle(*s)
	}
	return bu
}

// SetAuthor sets the "author" field.
func (bu *BookUpdate) SetAuthor(s string) *BookUpdate {
	bu.mutation.SetAuthor(s)
	return bu
}

// SetNillableAuthor sets the "author" field if the given value is not nil.
func (bu *BookUpdate) SetNillableAuthor(s *string) *BookUpdate {
	if s != nil {
		bu.SetAuthor(*s)
	}
	return bu
}

// SetTranslator sets the "translator" field.
func (bu *BookUpdate) SetTranslator(s string) *BookUpdate {
	bu.mutation.SetTranslator(s)
	return bu
}

// SetNillableTranslator sets the "translator" field if the given value is not nil.
func (bu *BookUpdate) SetNillableTranslator(s *string) *BookUpdate {
	if s != nil {
		bu.SetTranslator(*s)
	}
	return bu
}

// ClearTranslator clears the value of the "translator" field.
func (bu *BookUpdate) ClearTranslator() *BookUpdate {
	bu.mutation.ClearTranslator()
	return bu
}

// SetPublisherID sets the "publisher_id" field.
func (bu *BookUpdate) SetPublisherID(i int) *BookUpdate {
	bu.mutation.ResetPublisherID()
	bu.mutation.SetPublisherID(i)
	return bu
}

// SetNillablePublisherID sets the "publisher_id" field if the given value is not nil.
func (bu *BookUpdate) SetNillablePublisherID(i *int) *BookUpdate {
	if i != nil {
		bu.SetPublisherID(*i)
	}
	return bu
}

// AddPublisherID adds i to the "publisher_id" field.
func (bu *BookUpdate) AddPublisherID(i int) *BookUpdate {
	bu.mutation.AddPublisherID(i)
	return bu
}

// ClearPublisherID clears the value of the "publisher_id" field.
func (bu *BookUpdate) ClearPublisherID() *BookUpdate {
	bu.mutation.ClearPublisherID()
	return bu
}

// SetPublisher sets the "publisher" field.
func (bu *BookUpdate) SetPublisher(s string) *BookUpdate {
	bu.mutation.SetPublisher(s)
	return bu
}

// SetNillablePublisher sets the "publisher" field if the given value is not nil.
func (bu *BookUpdate) SetNillablePublisher(s *string) *BookUpdate {
	if s != nil {
		bu.SetPublisher(*s)
	}
	return bu
}

// ClearPublisher clears the value of the "publisher" field.
func (bu *BookUpdate) ClearPublisher() *BookUpdate {
	bu.mutation.ClearPublisher()
	return bu
}

// SetPublicationYear sets the "publication_year" field.
func (bu *BookUpdate) SetPublicationYear(i int16) *BookUpdate {
	bu.mutation.ResetPublicationYear()
	bu.mutation.SetPublicationYear(i)
	return bu
}

// SetNillablePublicationYear sets the "publication_year" field if the given value is not nil.
func (bu *BookUpdate) SetNillablePublicationYear(i *int16) *BookUpdate {
	if i != nil {
		bu.SetPublicationYear(*i)
	}
	return bu
}

// AddPublicationYear adds i to the "publication_year" field.
func (bu *BookUpdate) AddPublicationYear(i int16) *BookUpdate {
	bu.mutation.AddPublicationYear(i)
	return bu
}

// ClearPublicationYear clears the value of the "publication_year" field.
func (bu *BookUpdate) ClearPublicationYear() *BookUpdate {
	bu.mutation.ClearPublicationYear()
	return bu
}

// SetPublicationDate sets the "publication_date" field.
func (bu *BookUpdate) SetPublicationDate(t time.Time) *BookUpdate {
	bu.mutation.SetPublicationDate(t)
	return bu
}

// SetNillablePublicationDate sets the "publication_date" field if the given value is not nil.
func (bu *BookUpdate) SetNillablePublicationDate(t *time.Time) *BookUpdate {
	if t != nil {
		bu.SetPublicationDate(*t)
	}
	return bu
}

// SetEdition sets the "edition" field.
func (bu *BookUpdate) SetEdition(i int8) *BookUpdate {
	bu.mutation.ResetEdition()
	bu.mutation.SetEdition(i)
	return bu
}

// SetNillableEdition sets the "edition" field if the given value is not nil.
func (bu *BookUpdate) SetNillableEdition(i *int8) *BookUpdate {
	if i != nil {
		bu.SetEdition(*i)
	}
	return bu
}

// AddEdition adds i to the "edition" field.
func (bu *BookUpdate) AddEdition(i int8) *BookUpdate {
	bu.mutation.AddEdition(i)
	return bu
}

// ClearEdition clears the value of the "edition" field.
func (bu *BookUpdate) ClearEdition() *BookUpdate {
	bu.mutation.ClearEdition()
	return bu
}

// SetCategory sets the "category" field.
func (bu *BookUpdate) SetCategory(s string) *BookUpdate {
	bu.mutation.SetCategory(s)
	return bu
}

// SetNillableCategory sets the "category" field if the given value is not nil.
func (bu *BookUpdate) SetNillableCategory(s *string) *BookUpdate {
	if s != nil {
		bu.SetCategory(*s)
	}
	return bu
}

// ClearCategory clears the value of the "category" field.
func (bu *BookUpdate) ClearCategory() *BookUpdate {
	bu.mutation.ClearCategory()
	return bu
}

// SetPrice sets the "price" field.
func (bu *BookUpdate) SetPrice(f float64) *BookUpdate {
	bu.mutation.ResetPrice()
	bu.mutation.SetPrice(f)
	return bu
}

// SetNillablePrice sets the "price" field if the given value is not nil.
func (bu *BookUpdate) SetNillablePrice(f *float64) *BookUpdate {
	if f != nil {
		bu.SetPrice(*f)
	}
	return bu
}

// AddPrice adds f to the "price" field.
func (bu *BookUpdate) AddPrice(f float64) *BookUpdate {
	bu.mutation.AddPrice(f)
	return bu
}

// ClearPrice clears the value of the "price" field.
func (bu *BookUpdate) ClearPrice() *BookUpdate {
	bu.mutation.ClearPrice()
	return bu
}

// SetStockQuantity sets the "stock_quantity" field.
func (bu *BookUpdate) SetStockQuantity(i int32) *BookUpdate {
	bu.mutation.ResetStockQuantity()
	bu.mutation.SetStockQuantity(i)
	return bu
}

// SetNillableStockQuantity sets the "stock_quantity" field if the given value is not nil.
func (bu *BookUpdate) SetNillableStockQuantity(i *int32) *BookUpdate {
	if i != nil {
		bu.SetStockQuantity(*i)
	}
	return bu
}

// AddStockQuantity adds i to the "stock_quantity" field.
func (bu *BookUpdate) AddStockQuantity(i int32) *BookUpdate {
	bu.mutation.AddStockQuantity(i)
	return bu
}

// ClearStockQuantity clears the value of the "stock_quantity" field.
func (bu *BookUpdate) ClearStockQuantity() *BookUpdate {
	bu.mutation.ClearStockQuantity()
	return bu
}

// SetDescription sets the "description" field.
func (bu *BookUpdate) SetDescription(s string) *BookUpdate {
	bu.mutation.SetDescription(s)
	return bu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (bu *BookUpdate) SetNillableDescription(s *string) *BookUpdate {
	if s != nil {
		bu.SetDescription(*s)
	}
	return bu
}

// ClearDescription clears the value of the "description" field.
func (bu *BookUpdate) ClearDescription() *BookUpdate {
	bu.mutation.ClearDescription()
	return bu
}

// SetAddedOn sets the "added_on" field.
func (bu *BookUpdate) SetAddedOn(t time.Time) *BookUpdate {
	bu.mutation.SetAddedOn(t)
	return bu
}

// SetNillableAddedOn sets the "added_on" field if the given value is not nil.
func (bu *BookUpdate) SetNillableAddedOn(t *time.Time) *BookUpdate {
	if t != nil {
		bu.SetAddedOn(*t)
	}
	return bu
}

// ClearAddedOn clears the value of the "added_on" field.
func (bu *BookUpdate) ClearAddedOn() *BookUpdate {
	bu.mutation.ClearAddedOn()
	return bu
}

// SetCoverImage sets the "cover_image" field.
func (bu *BookUpdate) SetCoverImage(s string) *BookUpdate {
	bu.mutation.SetCoverImage(s)
	return bu
}

// SetNillableCoverImage sets the "cover_image" field if the given value is not nil.
func (bu *BookUpdate) SetNillableCoverImage(s *string) *BookUpdate {
	if s != nil {
		bu.SetCoverImage(*s)
	}
	return bu
}

// ClearCoverImage clears the value of the "cover_image" field.
func (bu *BookUpdate) ClearCoverImage() *BookUpdate {
	bu.mutation.ClearCoverImage()
	return bu
}

// SetPageCount sets the "page_count" field.
func (bu *BookUpdate) SetPageCount(i int32) *BookUpdate {
	bu.mutation.ResetPageCount()
	bu.mutation.SetPageCount(i)
	return bu
}

// SetNillablePageCount sets the "page_count" field if the given value is not nil.
func (bu *BookUpdate) SetNillablePageCount(i *int32) *BookUpdate {
	if i != nil {
		bu.SetPageCount(*i)
	}
	return bu
}

// AddPageCount adds i to the "page_count" field.
func (bu *BookUpdate) AddPageCount(i int32) *BookUpdate {
	bu.mutation.AddPageCount(i)
	return bu
}

// SetLanguageID sets the "language_id" field.
func (bu *BookUpdate) SetLanguageID(i int) *BookUpdate {
	bu.mutation.ResetLanguageID()
	bu.mutation.SetLanguageID(i)
	return bu
}

// SetNillableLanguageID sets the "language_id" field if the given value is not nil.
func (bu *BookUpdate) SetNillableLanguageID(i *int) *BookUpdate {
	if i != nil {
		bu.SetLanguageID(*i)
	}
	return bu
}

// AddLanguageID adds i to the "language_id" field.
func (bu *BookUpdate) AddLanguageID(i int) *BookUpdate {
	bu.mutation.AddLanguageID(i)
	return bu
}

// SetLanguage sets the "language" field.
func (bu *BookUpdate) SetLanguage(s string) *BookUpdate {
	bu.mutation.SetLanguage(s)
	return bu
}

// SetNillableLanguage sets the "language" field if the given value is not nil.
func (bu *BookUpdate) SetNillableLanguage(s *string) *BookUpdate {
	if s != nil {
		bu.SetLanguage(*s)
	}
	return bu
}

// SetAuthorID sets the "author_id" field.
func (bu *BookUpdate) SetAuthorID(i int) *BookUpdate {
	bu.mutation.ResetAuthorID()
	bu.mutation.SetAuthorID(i)
	return bu
}

// SetNillableAuthorID sets the "author_id" field if the given value is not nil.
func (bu *BookUpdate) SetNillableAuthorID(i *int) *BookUpdate {
	if i != nil {
		bu.SetAuthorID(*i)
	}
	return bu
}

// AddAuthorID adds i to the "author_id" field.
func (bu *BookUpdate) AddAuthorID(i int) *BookUpdate {
	bu.mutation.AddAuthorID(i)
	return bu
}

// SetCategoryID sets the "category_id" field.
func (bu *BookUpdate) SetCategoryID(i int) *BookUpdate {
	bu.mutation.ResetCategoryID()
	bu.mutation.SetCategoryID(i)
	return bu
}

// SetNillableCategoryID sets the "category_id" field if the given value is not nil.
func (bu *BookUpdate) SetNillableCategoryID(i *int) *BookUpdate {
	if i != nil {
		bu.SetCategoryID(*i)
	}
	return bu
}

// AddCategoryID adds i to the "category_id" field.
func (bu *BookUpdate) AddCategoryID(i int) *BookUpdate {
	bu.mutation.AddCategoryID(i)
	return bu
}

// Mutation returns the BookMutation object of the builder.
func (bu *BookUpdate) Mutation() *BookMutation {
	return bu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BookUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, bu.sqlSave, bu.mutation, bu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BookUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BookUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BookUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bu *BookUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(book.Table, book.Columns, sqlgraph.NewFieldSpec(book.FieldID, field.TypeInt))
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.Isbn(); ok {
		_spec.SetField(book.FieldIsbn, field.TypeString, value)
	}
	if value, ok := bu.mutation.Title(); ok {
		_spec.SetField(book.FieldTitle, field.TypeString, value)
	}
	if value, ok := bu.mutation.Author(); ok {
		_spec.SetField(book.FieldAuthor, field.TypeString, value)
	}
	if value, ok := bu.mutation.Translator(); ok {
		_spec.SetField(book.FieldTranslator, field.TypeString, value)
	}
	if bu.mutation.TranslatorCleared() {
		_spec.ClearField(book.FieldTranslator, field.TypeString)
	}
	if value, ok := bu.mutation.PublisherID(); ok {
		_spec.SetField(book.FieldPublisherID, field.TypeInt, value)
	}
	if value, ok := bu.mutation.AddedPublisherID(); ok {
		_spec.AddField(book.FieldPublisherID, field.TypeInt, value)
	}
	if bu.mutation.PublisherIDCleared() {
		_spec.ClearField(book.FieldPublisherID, field.TypeInt)
	}
	if value, ok := bu.mutation.Publisher(); ok {
		_spec.SetField(book.FieldPublisher, field.TypeString, value)
	}
	if bu.mutation.PublisherCleared() {
		_spec.ClearField(book.FieldPublisher, field.TypeString)
	}
	if value, ok := bu.mutation.PublicationYear(); ok {
		_spec.SetField(book.FieldPublicationYear, field.TypeInt16, value)
	}
	if value, ok := bu.mutation.AddedPublicationYear(); ok {
		_spec.AddField(book.FieldPublicationYear, field.TypeInt16, value)
	}
	if bu.mutation.PublicationYearCleared() {
		_spec.ClearField(book.FieldPublicationYear, field.TypeInt16)
	}
	if value, ok := bu.mutation.PublicationDate(); ok {
		_spec.SetField(book.FieldPublicationDate, field.TypeTime, value)
	}
	if value, ok := bu.mutation.Edition(); ok {
		_spec.SetField(book.FieldEdition, field.TypeInt8, value)
	}
	if value, ok := bu.mutation.AddedEdition(); ok {
		_spec.AddField(book.FieldEdition, field.TypeInt8, value)
	}
	if bu.mutation.EditionCleared() {
		_spec.ClearField(book.FieldEdition, field.TypeInt8)
	}
	if value, ok := bu.mutation.Category(); ok {
		_spec.SetField(book.FieldCategory, field.TypeString, value)
	}
	if bu.mutation.CategoryCleared() {
		_spec.ClearField(book.FieldCategory, field.TypeString)
	}
	if value, ok := bu.mutation.Price(); ok {
		_spec.SetField(book.FieldPrice, field.TypeFloat64, value)
	}
	if value, ok := bu.mutation.AddedPrice(); ok {
		_spec.AddField(book.FieldPrice, field.TypeFloat64, value)
	}
	if bu.mutation.PriceCleared() {
		_spec.ClearField(book.FieldPrice, field.TypeFloat64)
	}
	if value, ok := bu.mutation.StockQuantity(); ok {
		_spec.SetField(book.FieldStockQuantity, field.TypeInt32, value)
	}
	if value, ok := bu.mutation.AddedStockQuantity(); ok {
		_spec.AddField(book.FieldStockQuantity, field.TypeInt32, value)
	}
	if bu.mutation.StockQuantityCleared() {
		_spec.ClearField(book.FieldStockQuantity, field.TypeInt32)
	}
	if value, ok := bu.mutation.Description(); ok {
		_spec.SetField(book.FieldDescription, field.TypeString, value)
	}
	if bu.mutation.DescriptionCleared() {
		_spec.ClearField(book.FieldDescription, field.TypeString)
	}
	if value, ok := bu.mutation.AddedOn(); ok {
		_spec.SetField(book.FieldAddedOn, field.TypeTime, value)
	}
	if bu.mutation.AddedOnCleared() {
		_spec.ClearField(book.FieldAddedOn, field.TypeTime)
	}
	if value, ok := bu.mutation.CoverImage(); ok {
		_spec.SetField(book.FieldCoverImage, field.TypeString, value)
	}
	if bu.mutation.CoverImageCleared() {
		_spec.ClearField(book.FieldCoverImage, field.TypeString)
	}
	if value, ok := bu.mutation.PageCount(); ok {
		_spec.SetField(book.FieldPageCount, field.TypeInt32, value)
	}
	if value, ok := bu.mutation.AddedPageCount(); ok {
		_spec.AddField(book.FieldPageCount, field.TypeInt32, value)
	}
	if value, ok := bu.mutation.LanguageID(); ok {
		_spec.SetField(book.FieldLanguageID, field.TypeInt, value)
	}
	if value, ok := bu.mutation.AddedLanguageID(); ok {
		_spec.AddField(book.FieldLanguageID, field.TypeInt, value)
	}
	if value, ok := bu.mutation.Language(); ok {
		_spec.SetField(book.FieldLanguage, field.TypeString, value)
	}
	if value, ok := bu.mutation.AuthorID(); ok {
		_spec.SetField(book.FieldAuthorID, field.TypeInt, value)
	}
	if value, ok := bu.mutation.AddedAuthorID(); ok {
		_spec.AddField(book.FieldAuthorID, field.TypeInt, value)
	}
	if value, ok := bu.mutation.CategoryID(); ok {
		_spec.SetField(book.FieldCategoryID, field.TypeInt, value)
	}
	if value, ok := bu.mutation.AddedCategoryID(); ok {
		_spec.AddField(book.FieldCategoryID, field.TypeInt, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{book.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bu.mutation.done = true
	return n, nil
}

// BookUpdateOne is the builder for updating a single Book entity.
type BookUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BookMutation
}

// SetIsbn sets the "isbn" field.
func (buo *BookUpdateOne) SetIsbn(s string) *BookUpdateOne {
	buo.mutation.SetIsbn(s)
	return buo
}

// SetNillableIsbn sets the "isbn" field if the given value is not nil.
func (buo *BookUpdateOne) SetNillableIsbn(s *string) *BookUpdateOne {
	if s != nil {
		buo.SetIsbn(*s)
	}
	return buo
}

// SetTitle sets the "title" field.
func (buo *BookUpdateOne) SetTitle(s string) *BookUpdateOne {
	buo.mutation.SetTitle(s)
	return buo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (buo *BookUpdateOne) SetNillableTitle(s *string) *BookUpdateOne {
	if s != nil {
		buo.SetTitle(*s)
	}
	return buo
}

// SetAuthor sets the "author" field.
func (buo *BookUpdateOne) SetAuthor(s string) *BookUpdateOne {
	buo.mutation.SetAuthor(s)
	return buo
}

// SetNillableAuthor sets the "author" field if the given value is not nil.
func (buo *BookUpdateOne) SetNillableAuthor(s *string) *BookUpdateOne {
	if s != nil {
		buo.SetAuthor(*s)
	}
	return buo
}

// SetTranslator sets the "translator" field.
func (buo *BookUpdateOne) SetTranslator(s string) *BookUpdateOne {
	buo.mutation.SetTranslator(s)
	return buo
}

// SetNillableTranslator sets the "translator" field if the given value is not nil.
func (buo *BookUpdateOne) SetNillableTranslator(s *string) *BookUpdateOne {
	if s != nil {
		buo.SetTranslator(*s)
	}
	return buo
}

// ClearTranslator clears the value of the "translator" field.
func (buo *BookUpdateOne) ClearTranslator() *BookUpdateOne {
	buo.mutation.ClearTranslator()
	return buo
}

// SetPublisherID sets the "publisher_id" field.
func (buo *BookUpdateOne) SetPublisherID(i int) *BookUpdateOne {
	buo.mutation.ResetPublisherID()
	buo.mutation.SetPublisherID(i)
	return buo
}

// SetNillablePublisherID sets the "publisher_id" field if the given value is not nil.
func (buo *BookUpdateOne) SetNillablePublisherID(i *int) *BookUpdateOne {
	if i != nil {
		buo.SetPublisherID(*i)
	}
	return buo
}

// AddPublisherID adds i to the "publisher_id" field.
func (buo *BookUpdateOne) AddPublisherID(i int) *BookUpdateOne {
	buo.mutation.AddPublisherID(i)
	return buo
}

// ClearPublisherID clears the value of the "publisher_id" field.
func (buo *BookUpdateOne) ClearPublisherID() *BookUpdateOne {
	buo.mutation.ClearPublisherID()
	return buo
}

// SetPublisher sets the "publisher" field.
func (buo *BookUpdateOne) SetPublisher(s string) *BookUpdateOne {
	buo.mutation.SetPublisher(s)
	return buo
}

// SetNillablePublisher sets the "publisher" field if the given value is not nil.
func (buo *BookUpdateOne) SetNillablePublisher(s *string) *BookUpdateOne {
	if s != nil {
		buo.SetPublisher(*s)
	}
	return buo
}

// ClearPublisher clears the value of the "publisher" field.
func (buo *BookUpdateOne) ClearPublisher() *BookUpdateOne {
	buo.mutation.ClearPublisher()
	return buo
}

// SetPublicationYear sets the "publication_year" field.
func (buo *BookUpdateOne) SetPublicationYear(i int16) *BookUpdateOne {
	buo.mutation.ResetPublicationYear()
	buo.mutation.SetPublicationYear(i)
	return buo
}

// SetNillablePublicationYear sets the "publication_year" field if the given value is not nil.
func (buo *BookUpdateOne) SetNillablePublicationYear(i *int16) *BookUpdateOne {
	if i != nil {
		buo.SetPublicationYear(*i)
	}
	return buo
}

// AddPublicationYear adds i to the "publication_year" field.
func (buo *BookUpdateOne) AddPublicationYear(i int16) *BookUpdateOne {
	buo.mutation.AddPublicationYear(i)
	return buo
}

// ClearPublicationYear clears the value of the "publication_year" field.
func (buo *BookUpdateOne) ClearPublicationYear() *BookUpdateOne {
	buo.mutation.ClearPublicationYear()
	return buo
}

// SetPublicationDate sets the "publication_date" field.
func (buo *BookUpdateOne) SetPublicationDate(t time.Time) *BookUpdateOne {
	buo.mutation.SetPublicationDate(t)
	return buo
}

// SetNillablePublicationDate sets the "publication_date" field if the given value is not nil.
func (buo *BookUpdateOne) SetNillablePublicationDate(t *time.Time) *BookUpdateOne {
	if t != nil {
		buo.SetPublicationDate(*t)
	}
	return buo
}

// SetEdition sets the "edition" field.
func (buo *BookUpdateOne) SetEdition(i int8) *BookUpdateOne {
	buo.mutation.ResetEdition()
	buo.mutation.SetEdition(i)
	return buo
}

// SetNillableEdition sets the "edition" field if the given value is not nil.
func (buo *BookUpdateOne) SetNillableEdition(i *int8) *BookUpdateOne {
	if i != nil {
		buo.SetEdition(*i)
	}
	return buo
}

// AddEdition adds i to the "edition" field.
func (buo *BookUpdateOne) AddEdition(i int8) *BookUpdateOne {
	buo.mutation.AddEdition(i)
	return buo
}

// ClearEdition clears the value of the "edition" field.
func (buo *BookUpdateOne) ClearEdition() *BookUpdateOne {
	buo.mutation.ClearEdition()
	return buo
}

// SetCategory sets the "category" field.
func (buo *BookUpdateOne) SetCategory(s string) *BookUpdateOne {
	buo.mutation.SetCategory(s)
	return buo
}

// SetNillableCategory sets the "category" field if the given value is not nil.
func (buo *BookUpdateOne) SetNillableCategory(s *string) *BookUpdateOne {
	if s != nil {
		buo.SetCategory(*s)
	}
	return buo
}

// ClearCategory clears the value of the "category" field.
func (buo *BookUpdateOne) ClearCategory() *BookUpdateOne {
	buo.mutation.ClearCategory()
	return buo
}

// SetPrice sets the "price" field.
func (buo *BookUpdateOne) SetPrice(f float64) *BookUpdateOne {
	buo.mutation.ResetPrice()
	buo.mutation.SetPrice(f)
	return buo
}

// SetNillablePrice sets the "price" field if the given value is not nil.
func (buo *BookUpdateOne) SetNillablePrice(f *float64) *BookUpdateOne {
	if f != nil {
		buo.SetPrice(*f)
	}
	return buo
}

// AddPrice adds f to the "price" field.
func (buo *BookUpdateOne) AddPrice(f float64) *BookUpdateOne {
	buo.mutation.AddPrice(f)
	return buo
}

// ClearPrice clears the value of the "price" field.
func (buo *BookUpdateOne) ClearPrice() *BookUpdateOne {
	buo.mutation.ClearPrice()
	return buo
}

// SetStockQuantity sets the "stock_quantity" field.
func (buo *BookUpdateOne) SetStockQuantity(i int32) *BookUpdateOne {
	buo.mutation.ResetStockQuantity()
	buo.mutation.SetStockQuantity(i)
	return buo
}

// SetNillableStockQuantity sets the "stock_quantity" field if the given value is not nil.
func (buo *BookUpdateOne) SetNillableStockQuantity(i *int32) *BookUpdateOne {
	if i != nil {
		buo.SetStockQuantity(*i)
	}
	return buo
}

// AddStockQuantity adds i to the "stock_quantity" field.
func (buo *BookUpdateOne) AddStockQuantity(i int32) *BookUpdateOne {
	buo.mutation.AddStockQuantity(i)
	return buo
}

// ClearStockQuantity clears the value of the "stock_quantity" field.
func (buo *BookUpdateOne) ClearStockQuantity() *BookUpdateOne {
	buo.mutation.ClearStockQuantity()
	return buo
}

// SetDescription sets the "description" field.
func (buo *BookUpdateOne) SetDescription(s string) *BookUpdateOne {
	buo.mutation.SetDescription(s)
	return buo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (buo *BookUpdateOne) SetNillableDescription(s *string) *BookUpdateOne {
	if s != nil {
		buo.SetDescription(*s)
	}
	return buo
}

// ClearDescription clears the value of the "description" field.
func (buo *BookUpdateOne) ClearDescription() *BookUpdateOne {
	buo.mutation.ClearDescription()
	return buo
}

// SetAddedOn sets the "added_on" field.
func (buo *BookUpdateOne) SetAddedOn(t time.Time) *BookUpdateOne {
	buo.mutation.SetAddedOn(t)
	return buo
}

// SetNillableAddedOn sets the "added_on" field if the given value is not nil.
func (buo *BookUpdateOne) SetNillableAddedOn(t *time.Time) *BookUpdateOne {
	if t != nil {
		buo.SetAddedOn(*t)
	}
	return buo
}

// ClearAddedOn clears the value of the "added_on" field.
func (buo *BookUpdateOne) ClearAddedOn() *BookUpdateOne {
	buo.mutation.ClearAddedOn()
	return buo
}

// SetCoverImage sets the "cover_image" field.
func (buo *BookUpdateOne) SetCoverImage(s string) *BookUpdateOne {
	buo.mutation.SetCoverImage(s)
	return buo
}

// SetNillableCoverImage sets the "cover_image" field if the given value is not nil.
func (buo *BookUpdateOne) SetNillableCoverImage(s *string) *BookUpdateOne {
	if s != nil {
		buo.SetCoverImage(*s)
	}
	return buo
}

// ClearCoverImage clears the value of the "cover_image" field.
func (buo *BookUpdateOne) ClearCoverImage() *BookUpdateOne {
	buo.mutation.ClearCoverImage()
	return buo
}

// SetPageCount sets the "page_count" field.
func (buo *BookUpdateOne) SetPageCount(i int32) *BookUpdateOne {
	buo.mutation.ResetPageCount()
	buo.mutation.SetPageCount(i)
	return buo
}

// SetNillablePageCount sets the "page_count" field if the given value is not nil.
func (buo *BookUpdateOne) SetNillablePageCount(i *int32) *BookUpdateOne {
	if i != nil {
		buo.SetPageCount(*i)
	}
	return buo
}

// AddPageCount adds i to the "page_count" field.
func (buo *BookUpdateOne) AddPageCount(i int32) *BookUpdateOne {
	buo.mutation.AddPageCount(i)
	return buo
}

// SetLanguageID sets the "language_id" field.
func (buo *BookUpdateOne) SetLanguageID(i int) *BookUpdateOne {
	buo.mutation.ResetLanguageID()
	buo.mutation.SetLanguageID(i)
	return buo
}

// SetNillableLanguageID sets the "language_id" field if the given value is not nil.
func (buo *BookUpdateOne) SetNillableLanguageID(i *int) *BookUpdateOne {
	if i != nil {
		buo.SetLanguageID(*i)
	}
	return buo
}

// AddLanguageID adds i to the "language_id" field.
func (buo *BookUpdateOne) AddLanguageID(i int) *BookUpdateOne {
	buo.mutation.AddLanguageID(i)
	return buo
}

// SetLanguage sets the "language" field.
func (buo *BookUpdateOne) SetLanguage(s string) *BookUpdateOne {
	buo.mutation.SetLanguage(s)
	return buo
}

// SetNillableLanguage sets the "language" field if the given value is not nil.
func (buo *BookUpdateOne) SetNillableLanguage(s *string) *BookUpdateOne {
	if s != nil {
		buo.SetLanguage(*s)
	}
	return buo
}

// SetAuthorID sets the "author_id" field.
func (buo *BookUpdateOne) SetAuthorID(i int) *BookUpdateOne {
	buo.mutation.ResetAuthorID()
	buo.mutation.SetAuthorID(i)
	return buo
}

// SetNillableAuthorID sets the "author_id" field if the given value is not nil.
func (buo *BookUpdateOne) SetNillableAuthorID(i *int) *BookUpdateOne {
	if i != nil {
		buo.SetAuthorID(*i)
	}
	return buo
}

// AddAuthorID adds i to the "author_id" field.
func (buo *BookUpdateOne) AddAuthorID(i int) *BookUpdateOne {
	buo.mutation.AddAuthorID(i)
	return buo
}

// SetCategoryID sets the "category_id" field.
func (buo *BookUpdateOne) SetCategoryID(i int) *BookUpdateOne {
	buo.mutation.ResetCategoryID()
	buo.mutation.SetCategoryID(i)
	return buo
}

// SetNillableCategoryID sets the "category_id" field if the given value is not nil.
func (buo *BookUpdateOne) SetNillableCategoryID(i *int) *BookUpdateOne {
	if i != nil {
		buo.SetCategoryID(*i)
	}
	return buo
}

// AddCategoryID adds i to the "category_id" field.
func (buo *BookUpdateOne) AddCategoryID(i int) *BookUpdateOne {
	buo.mutation.AddCategoryID(i)
	return buo
}

// Mutation returns the BookMutation object of the builder.
func (buo *BookUpdateOne) Mutation() *BookMutation {
	return buo.mutation
}

// Where appends a list predicates to the BookUpdate builder.
func (buo *BookUpdateOne) Where(ps ...predicate.Book) *BookUpdateOne {
	buo.mutation.Where(ps...)
	return buo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BookUpdateOne) Select(field string, fields ...string) *BookUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Book entity.
func (buo *BookUpdateOne) Save(ctx context.Context) (*Book, error) {
	return withHooks(ctx, buo.sqlSave, buo.mutation, buo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BookUpdateOne) SaveX(ctx context.Context) *Book {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BookUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BookUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (buo *BookUpdateOne) sqlSave(ctx context.Context) (_node *Book, err error) {
	_spec := sqlgraph.NewUpdateSpec(book.Table, book.Columns, sqlgraph.NewFieldSpec(book.FieldID, field.TypeInt))
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Book.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, book.FieldID)
		for _, f := range fields {
			if !book.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != book.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.Isbn(); ok {
		_spec.SetField(book.FieldIsbn, field.TypeString, value)
	}
	if value, ok := buo.mutation.Title(); ok {
		_spec.SetField(book.FieldTitle, field.TypeString, value)
	}
	if value, ok := buo.mutation.Author(); ok {
		_spec.SetField(book.FieldAuthor, field.TypeString, value)
	}
	if value, ok := buo.mutation.Translator(); ok {
		_spec.SetField(book.FieldTranslator, field.TypeString, value)
	}
	if buo.mutation.TranslatorCleared() {
		_spec.ClearField(book.FieldTranslator, field.TypeString)
	}
	if value, ok := buo.mutation.PublisherID(); ok {
		_spec.SetField(book.FieldPublisherID, field.TypeInt, value)
	}
	if value, ok := buo.mutation.AddedPublisherID(); ok {
		_spec.AddField(book.FieldPublisherID, field.TypeInt, value)
	}
	if buo.mutation.PublisherIDCleared() {
		_spec.ClearField(book.FieldPublisherID, field.TypeInt)
	}
	if value, ok := buo.mutation.Publisher(); ok {
		_spec.SetField(book.FieldPublisher, field.TypeString, value)
	}
	if buo.mutation.PublisherCleared() {
		_spec.ClearField(book.FieldPublisher, field.TypeString)
	}
	if value, ok := buo.mutation.PublicationYear(); ok {
		_spec.SetField(book.FieldPublicationYear, field.TypeInt16, value)
	}
	if value, ok := buo.mutation.AddedPublicationYear(); ok {
		_spec.AddField(book.FieldPublicationYear, field.TypeInt16, value)
	}
	if buo.mutation.PublicationYearCleared() {
		_spec.ClearField(book.FieldPublicationYear, field.TypeInt16)
	}
	if value, ok := buo.mutation.PublicationDate(); ok {
		_spec.SetField(book.FieldPublicationDate, field.TypeTime, value)
	}
	if value, ok := buo.mutation.Edition(); ok {
		_spec.SetField(book.FieldEdition, field.TypeInt8, value)
	}
	if value, ok := buo.mutation.AddedEdition(); ok {
		_spec.AddField(book.FieldEdition, field.TypeInt8, value)
	}
	if buo.mutation.EditionCleared() {
		_spec.ClearField(book.FieldEdition, field.TypeInt8)
	}
	if value, ok := buo.mutation.Category(); ok {
		_spec.SetField(book.FieldCategory, field.TypeString, value)
	}
	if buo.mutation.CategoryCleared() {
		_spec.ClearField(book.FieldCategory, field.TypeString)
	}
	if value, ok := buo.mutation.Price(); ok {
		_spec.SetField(book.FieldPrice, field.TypeFloat64, value)
	}
	if value, ok := buo.mutation.AddedPrice(); ok {
		_spec.AddField(book.FieldPrice, field.TypeFloat64, value)
	}
	if buo.mutation.PriceCleared() {
		_spec.ClearField(book.FieldPrice, field.TypeFloat64)
	}
	if value, ok := buo.mutation.StockQuantity(); ok {
		_spec.SetField(book.FieldStockQuantity, field.TypeInt32, value)
	}
	if value, ok := buo.mutation.AddedStockQuantity(); ok {
		_spec.AddField(book.FieldStockQuantity, field.TypeInt32, value)
	}
	if buo.mutation.StockQuantityCleared() {
		_spec.ClearField(book.FieldStockQuantity, field.TypeInt32)
	}
	if value, ok := buo.mutation.Description(); ok {
		_spec.SetField(book.FieldDescription, field.TypeString, value)
	}
	if buo.mutation.DescriptionCleared() {
		_spec.ClearField(book.FieldDescription, field.TypeString)
	}
	if value, ok := buo.mutation.AddedOn(); ok {
		_spec.SetField(book.FieldAddedOn, field.TypeTime, value)
	}
	if buo.mutation.AddedOnCleared() {
		_spec.ClearField(book.FieldAddedOn, field.TypeTime)
	}
	if value, ok := buo.mutation.CoverImage(); ok {
		_spec.SetField(book.FieldCoverImage, field.TypeString, value)
	}
	if buo.mutation.CoverImageCleared() {
		_spec.ClearField(book.FieldCoverImage, field.TypeString)
	}
	if value, ok := buo.mutation.PageCount(); ok {
		_spec.SetField(book.FieldPageCount, field.TypeInt32, value)
	}
	if value, ok := buo.mutation.AddedPageCount(); ok {
		_spec.AddField(book.FieldPageCount, field.TypeInt32, value)
	}
	if value, ok := buo.mutation.LanguageID(); ok {
		_spec.SetField(book.FieldLanguageID, field.TypeInt, value)
	}
	if value, ok := buo.mutation.AddedLanguageID(); ok {
		_spec.AddField(book.FieldLanguageID, field.TypeInt, value)
	}
	if value, ok := buo.mutation.Language(); ok {
		_spec.SetField(book.FieldLanguage, field.TypeString, value)
	}
	if value, ok := buo.mutation.AuthorID(); ok {
		_spec.SetField(book.FieldAuthorID, field.TypeInt, value)
	}
	if value, ok := buo.mutation.AddedAuthorID(); ok {
		_spec.AddField(book.FieldAuthorID, field.TypeInt, value)
	}
	if value, ok := buo.mutation.CategoryID(); ok {
		_spec.SetField(book.FieldCategoryID, field.TypeInt, value)
	}
	if value, ok := buo.mutation.AddedCategoryID(); ok {
		_spec.AddField(book.FieldCategoryID, field.TypeInt, value)
	}
	_node = &Book{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{book.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	buo.mutation.done = true
	return _node, nil
}
