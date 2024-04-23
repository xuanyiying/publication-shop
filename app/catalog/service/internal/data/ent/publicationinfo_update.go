// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/publication-shop/app/catalog/service/internal/data/ent/predicate"
	"github.com/go-kratos/publication-shop/app/catalog/service/internal/data/ent/publicationinfo"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PublicationInfoUpdate is the builder for updating PublicationInfo entities.
type PublicationInfoUpdate struct {
	config
	hooks    []Hook
	mutation *PublicationInfoMutation
}

// Where appends a list predicates to the PublicationInfoUpdate builder.
func (piu *PublicationInfoUpdate) Where(ps ...predicate.PublicationInfo) *PublicationInfoUpdate {
	piu.mutation.Where(ps...)
	return piu
}

// SetPublicationName sets the "publication_name" field.
func (piu *PublicationInfoUpdate) SetPublicationName(s string) *PublicationInfoUpdate {
	piu.mutation.SetPublicationName(s)
	return piu
}

// SetNillablePublicationName sets the "publication_name" field if the given value is not nil.
func (piu *PublicationInfoUpdate) SetNillablePublicationName(s *string) *PublicationInfoUpdate {
	if s != nil {
		piu.SetPublicationName(*s)
	}
	return piu
}

// ClearPublicationName clears the value of the "publication_name" field.
func (piu *PublicationInfoUpdate) ClearPublicationName() *PublicationInfoUpdate {
	piu.mutation.ClearPublicationName()
	return piu
}

// SetOrgID sets the "org_id" field.
func (piu *PublicationInfoUpdate) SetOrgID(s string) *PublicationInfoUpdate {
	piu.mutation.SetOrgID(s)
	return piu
}

// SetNillableOrgID sets the "org_id" field if the given value is not nil.
func (piu *PublicationInfoUpdate) SetNillableOrgID(s *string) *PublicationInfoUpdate {
	if s != nil {
		piu.SetOrgID(*s)
	}
	return piu
}

// ClearOrgID clears the value of the "org_id" field.
func (piu *PublicationInfoUpdate) ClearOrgID() *PublicationInfoUpdate {
	piu.mutation.ClearOrgID()
	return piu
}

// SetPublishedTimes sets the "published_times" field.
func (piu *PublicationInfoUpdate) SetPublishedTimes(s string) *PublicationInfoUpdate {
	piu.mutation.SetPublishedTimes(s)
	return piu
}

// SetNillablePublishedTimes sets the "published_times" field if the given value is not nil.
func (piu *PublicationInfoUpdate) SetNillablePublishedTimes(s *string) *PublicationInfoUpdate {
	if s != nil {
		piu.SetPublishedTimes(*s)
	}
	return piu
}

// ClearPublishedTimes clears the value of the "published_times" field.
func (piu *PublicationInfoUpdate) ClearPublishedTimes() *PublicationInfoUpdate {
	piu.mutation.ClearPublishedTimes()
	return piu
}

// SetPrintTimes sets the "print_times" field.
func (piu *PublicationInfoUpdate) SetPrintTimes(s string) *PublicationInfoUpdate {
	piu.mutation.SetPrintTimes(s)
	return piu
}

// SetNillablePrintTimes sets the "print_times" field if the given value is not nil.
func (piu *PublicationInfoUpdate) SetNillablePrintTimes(s *string) *PublicationInfoUpdate {
	if s != nil {
		piu.SetPrintTimes(*s)
	}
	return piu
}

// ClearPrintTimes clears the value of the "print_times" field.
func (piu *PublicationInfoUpdate) ClearPrintTimes() *PublicationInfoUpdate {
	piu.mutation.ClearPrintTimes()
	return piu
}

// SetPrice sets the "price" field.
func (piu *PublicationInfoUpdate) SetPrice(f float64) *PublicationInfoUpdate {
	piu.mutation.ResetPrice()
	piu.mutation.SetPrice(f)
	return piu
}

// SetNillablePrice sets the "price" field if the given value is not nil.
func (piu *PublicationInfoUpdate) SetNillablePrice(f *float64) *PublicationInfoUpdate {
	if f != nil {
		piu.SetPrice(*f)
	}
	return piu
}

// AddPrice adds f to the "price" field.
func (piu *PublicationInfoUpdate) AddPrice(f float64) *PublicationInfoUpdate {
	piu.mutation.AddPrice(f)
	return piu
}

// ClearPrice clears the value of the "price" field.
func (piu *PublicationInfoUpdate) ClearPrice() *PublicationInfoUpdate {
	piu.mutation.ClearPrice()
	return piu
}

// SetIntroduction sets the "introduction" field.
func (piu *PublicationInfoUpdate) SetIntroduction(s string) *PublicationInfoUpdate {
	piu.mutation.SetIntroduction(s)
	return piu
}

// SetNillableIntroduction sets the "introduction" field if the given value is not nil.
func (piu *PublicationInfoUpdate) SetNillableIntroduction(s *string) *PublicationInfoUpdate {
	if s != nil {
		piu.SetIntroduction(*s)
	}
	return piu
}

// ClearIntroduction clears the value of the "introduction" field.
func (piu *PublicationInfoUpdate) ClearIntroduction() *PublicationInfoUpdate {
	piu.mutation.ClearIntroduction()
	return piu
}

// SetQuantity sets the "quantity" field.
func (piu *PublicationInfoUpdate) SetQuantity(i int32) *PublicationInfoUpdate {
	piu.mutation.ResetQuantity()
	piu.mutation.SetQuantity(i)
	return piu
}

// SetNillableQuantity sets the "quantity" field if the given value is not nil.
func (piu *PublicationInfoUpdate) SetNillableQuantity(i *int32) *PublicationInfoUpdate {
	if i != nil {
		piu.SetQuantity(*i)
	}
	return piu
}

// AddQuantity adds i to the "quantity" field.
func (piu *PublicationInfoUpdate) AddQuantity(i int32) *PublicationInfoUpdate {
	piu.mutation.AddQuantity(i)
	return piu
}

// ClearQuantity clears the value of the "quantity" field.
func (piu *PublicationInfoUpdate) ClearQuantity() *PublicationInfoUpdate {
	piu.mutation.ClearQuantity()
	return piu
}

// SetWordCount sets the "word_count" field.
func (piu *PublicationInfoUpdate) SetWordCount(s string) *PublicationInfoUpdate {
	piu.mutation.SetWordCount(s)
	return piu
}

// SetNillableWordCount sets the "word_count" field if the given value is not nil.
func (piu *PublicationInfoUpdate) SetNillableWordCount(s *string) *PublicationInfoUpdate {
	if s != nil {
		piu.SetWordCount(*s)
	}
	return piu
}

// ClearWordCount clears the value of the "word_count" field.
func (piu *PublicationInfoUpdate) ClearWordCount() *PublicationInfoUpdate {
	piu.mutation.ClearWordCount()
	return piu
}

// SetIsbn sets the "isbn" field.
func (piu *PublicationInfoUpdate) SetIsbn(s string) *PublicationInfoUpdate {
	piu.mutation.SetIsbn(s)
	return piu
}

// SetNillableIsbn sets the "isbn" field if the given value is not nil.
func (piu *PublicationInfoUpdate) SetNillableIsbn(s *string) *PublicationInfoUpdate {
	if s != nil {
		piu.SetIsbn(*s)
	}
	return piu
}

// ClearIsbn clears the value of the "isbn" field.
func (piu *PublicationInfoUpdate) ClearIsbn() *PublicationInfoUpdate {
	piu.mutation.ClearIsbn()
	return piu
}

// SetStorageBy sets the "storage_by" field.
func (piu *PublicationInfoUpdate) SetStorageBy(s string) *PublicationInfoUpdate {
	piu.mutation.SetStorageBy(s)
	return piu
}

// SetNillableStorageBy sets the "storage_by" field if the given value is not nil.
func (piu *PublicationInfoUpdate) SetNillableStorageBy(s *string) *PublicationInfoUpdate {
	if s != nil {
		piu.SetStorageBy(*s)
	}
	return piu
}

// ClearStorageBy clears the value of the "storage_by" field.
func (piu *PublicationInfoUpdate) ClearStorageBy() *PublicationInfoUpdate {
	piu.mutation.ClearStorageBy()
	return piu
}

// SetStorageAt sets the "storage_at" field.
func (piu *PublicationInfoUpdate) SetStorageAt(t time.Time) *PublicationInfoUpdate {
	piu.mutation.SetStorageAt(t)
	return piu
}

// SetNillableStorageAt sets the "storage_at" field if the given value is not nil.
func (piu *PublicationInfoUpdate) SetNillableStorageAt(t *time.Time) *PublicationInfoUpdate {
	if t != nil {
		piu.SetStorageAt(*t)
	}
	return piu
}

// ClearStorageAt clears the value of the "storage_at" field.
func (piu *PublicationInfoUpdate) ClearStorageAt() *PublicationInfoUpdate {
	piu.mutation.ClearStorageAt()
	return piu
}

// SetModifiedBy sets the "modified_by" field.
func (piu *PublicationInfoUpdate) SetModifiedBy(s string) *PublicationInfoUpdate {
	piu.mutation.SetModifiedBy(s)
	return piu
}

// SetNillableModifiedBy sets the "modified_by" field if the given value is not nil.
func (piu *PublicationInfoUpdate) SetNillableModifiedBy(s *string) *PublicationInfoUpdate {
	if s != nil {
		piu.SetModifiedBy(*s)
	}
	return piu
}

// ClearModifiedBy clears the value of the "modified_by" field.
func (piu *PublicationInfoUpdate) ClearModifiedBy() *PublicationInfoUpdate {
	piu.mutation.ClearModifiedBy()
	return piu
}

// SetCreatedAt sets the "created_at" field.
func (piu *PublicationInfoUpdate) SetCreatedAt(t time.Time) *PublicationInfoUpdate {
	piu.mutation.SetCreatedAt(t)
	return piu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (piu *PublicationInfoUpdate) SetNillableCreatedAt(t *time.Time) *PublicationInfoUpdate {
	if t != nil {
		piu.SetCreatedAt(*t)
	}
	return piu
}

// ClearCreatedAt clears the value of the "created_at" field.
func (piu *PublicationInfoUpdate) ClearCreatedAt() *PublicationInfoUpdate {
	piu.mutation.ClearCreatedAt()
	return piu
}

// SetModifiedAt sets the "modified_at" field.
func (piu *PublicationInfoUpdate) SetModifiedAt(t time.Time) *PublicationInfoUpdate {
	piu.mutation.SetModifiedAt(t)
	return piu
}

// SetNillableModifiedAt sets the "modified_at" field if the given value is not nil.
func (piu *PublicationInfoUpdate) SetNillableModifiedAt(t *time.Time) *PublicationInfoUpdate {
	if t != nil {
		piu.SetModifiedAt(*t)
	}
	return piu
}

// ClearModifiedAt clears the value of the "modified_at" field.
func (piu *PublicationInfoUpdate) ClearModifiedAt() *PublicationInfoUpdate {
	piu.mutation.ClearModifiedAt()
	return piu
}

// Mutation returns the PublicationInfoMutation object of the builder.
func (piu *PublicationInfoUpdate) Mutation() *PublicationInfoMutation {
	return piu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (piu *PublicationInfoUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, piu.sqlSave, piu.mutation, piu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (piu *PublicationInfoUpdate) SaveX(ctx context.Context) int {
	affected, err := piu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (piu *PublicationInfoUpdate) Exec(ctx context.Context) error {
	_, err := piu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (piu *PublicationInfoUpdate) ExecX(ctx context.Context) {
	if err := piu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (piu *PublicationInfoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(publicationinfo.Table, publicationinfo.Columns, sqlgraph.NewFieldSpec(publicationinfo.FieldID, field.TypeInt))
	if ps := piu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := piu.mutation.PublicationName(); ok {
		_spec.SetField(publicationinfo.FieldPublicationName, field.TypeString, value)
	}
	if piu.mutation.PublicationNameCleared() {
		_spec.ClearField(publicationinfo.FieldPublicationName, field.TypeString)
	}
	if value, ok := piu.mutation.OrgID(); ok {
		_spec.SetField(publicationinfo.FieldOrgID, field.TypeString, value)
	}
	if piu.mutation.OrgIDCleared() {
		_spec.ClearField(publicationinfo.FieldOrgID, field.TypeString)
	}
	if value, ok := piu.mutation.PublishedTimes(); ok {
		_spec.SetField(publicationinfo.FieldPublishedTimes, field.TypeString, value)
	}
	if piu.mutation.PublishedTimesCleared() {
		_spec.ClearField(publicationinfo.FieldPublishedTimes, field.TypeString)
	}
	if value, ok := piu.mutation.PrintTimes(); ok {
		_spec.SetField(publicationinfo.FieldPrintTimes, field.TypeString, value)
	}
	if piu.mutation.PrintTimesCleared() {
		_spec.ClearField(publicationinfo.FieldPrintTimes, field.TypeString)
	}
	if value, ok := piu.mutation.Price(); ok {
		_spec.SetField(publicationinfo.FieldPrice, field.TypeFloat64, value)
	}
	if value, ok := piu.mutation.AddedPrice(); ok {
		_spec.AddField(publicationinfo.FieldPrice, field.TypeFloat64, value)
	}
	if piu.mutation.PriceCleared() {
		_spec.ClearField(publicationinfo.FieldPrice, field.TypeFloat64)
	}
	if value, ok := piu.mutation.Introduction(); ok {
		_spec.SetField(publicationinfo.FieldIntroduction, field.TypeString, value)
	}
	if piu.mutation.IntroductionCleared() {
		_spec.ClearField(publicationinfo.FieldIntroduction, field.TypeString)
	}
	if value, ok := piu.mutation.Quantity(); ok {
		_spec.SetField(publicationinfo.FieldQuantity, field.TypeInt32, value)
	}
	if value, ok := piu.mutation.AddedQuantity(); ok {
		_spec.AddField(publicationinfo.FieldQuantity, field.TypeInt32, value)
	}
	if piu.mutation.QuantityCleared() {
		_spec.ClearField(publicationinfo.FieldQuantity, field.TypeInt32)
	}
	if value, ok := piu.mutation.WordCount(); ok {
		_spec.SetField(publicationinfo.FieldWordCount, field.TypeString, value)
	}
	if piu.mutation.WordCountCleared() {
		_spec.ClearField(publicationinfo.FieldWordCount, field.TypeString)
	}
	if value, ok := piu.mutation.Isbn(); ok {
		_spec.SetField(publicationinfo.FieldIsbn, field.TypeString, value)
	}
	if piu.mutation.IsbnCleared() {
		_spec.ClearField(publicationinfo.FieldIsbn, field.TypeString)
	}
	if value, ok := piu.mutation.StorageBy(); ok {
		_spec.SetField(publicationinfo.FieldStorageBy, field.TypeString, value)
	}
	if piu.mutation.StorageByCleared() {
		_spec.ClearField(publicationinfo.FieldStorageBy, field.TypeString)
	}
	if value, ok := piu.mutation.StorageAt(); ok {
		_spec.SetField(publicationinfo.FieldStorageAt, field.TypeTime, value)
	}
	if piu.mutation.StorageAtCleared() {
		_spec.ClearField(publicationinfo.FieldStorageAt, field.TypeTime)
	}
	if value, ok := piu.mutation.ModifiedBy(); ok {
		_spec.SetField(publicationinfo.FieldModifiedBy, field.TypeString, value)
	}
	if piu.mutation.ModifiedByCleared() {
		_spec.ClearField(publicationinfo.FieldModifiedBy, field.TypeString)
	}
	if value, ok := piu.mutation.CreatedAt(); ok {
		_spec.SetField(publicationinfo.FieldCreatedAt, field.TypeTime, value)
	}
	if piu.mutation.CreatedAtCleared() {
		_spec.ClearField(publicationinfo.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := piu.mutation.ModifiedAt(); ok {
		_spec.SetField(publicationinfo.FieldModifiedAt, field.TypeTime, value)
	}
	if piu.mutation.ModifiedAtCleared() {
		_spec.ClearField(publicationinfo.FieldModifiedAt, field.TypeTime)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, piu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{publicationinfo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	piu.mutation.done = true
	return n, nil
}

// PublicationInfoUpdateOne is the builder for updating a single PublicationInfo entity.
type PublicationInfoUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PublicationInfoMutation
}

// SetPublicationName sets the "publication_name" field.
func (piuo *PublicationInfoUpdateOne) SetPublicationName(s string) *PublicationInfoUpdateOne {
	piuo.mutation.SetPublicationName(s)
	return piuo
}

// SetNillablePublicationName sets the "publication_name" field if the given value is not nil.
func (piuo *PublicationInfoUpdateOne) SetNillablePublicationName(s *string) *PublicationInfoUpdateOne {
	if s != nil {
		piuo.SetPublicationName(*s)
	}
	return piuo
}

// ClearPublicationName clears the value of the "publication_name" field.
func (piuo *PublicationInfoUpdateOne) ClearPublicationName() *PublicationInfoUpdateOne {
	piuo.mutation.ClearPublicationName()
	return piuo
}

// SetOrgID sets the "org_id" field.
func (piuo *PublicationInfoUpdateOne) SetOrgID(s string) *PublicationInfoUpdateOne {
	piuo.mutation.SetOrgID(s)
	return piuo
}

// SetNillableOrgID sets the "org_id" field if the given value is not nil.
func (piuo *PublicationInfoUpdateOne) SetNillableOrgID(s *string) *PublicationInfoUpdateOne {
	if s != nil {
		piuo.SetOrgID(*s)
	}
	return piuo
}

// ClearOrgID clears the value of the "org_id" field.
func (piuo *PublicationInfoUpdateOne) ClearOrgID() *PublicationInfoUpdateOne {
	piuo.mutation.ClearOrgID()
	return piuo
}

// SetPublishedTimes sets the "published_times" field.
func (piuo *PublicationInfoUpdateOne) SetPublishedTimes(s string) *PublicationInfoUpdateOne {
	piuo.mutation.SetPublishedTimes(s)
	return piuo
}

// SetNillablePublishedTimes sets the "published_times" field if the given value is not nil.
func (piuo *PublicationInfoUpdateOne) SetNillablePublishedTimes(s *string) *PublicationInfoUpdateOne {
	if s != nil {
		piuo.SetPublishedTimes(*s)
	}
	return piuo
}

// ClearPublishedTimes clears the value of the "published_times" field.
func (piuo *PublicationInfoUpdateOne) ClearPublishedTimes() *PublicationInfoUpdateOne {
	piuo.mutation.ClearPublishedTimes()
	return piuo
}

// SetPrintTimes sets the "print_times" field.
func (piuo *PublicationInfoUpdateOne) SetPrintTimes(s string) *PublicationInfoUpdateOne {
	piuo.mutation.SetPrintTimes(s)
	return piuo
}

// SetNillablePrintTimes sets the "print_times" field if the given value is not nil.
func (piuo *PublicationInfoUpdateOne) SetNillablePrintTimes(s *string) *PublicationInfoUpdateOne {
	if s != nil {
		piuo.SetPrintTimes(*s)
	}
	return piuo
}

// ClearPrintTimes clears the value of the "print_times" field.
func (piuo *PublicationInfoUpdateOne) ClearPrintTimes() *PublicationInfoUpdateOne {
	piuo.mutation.ClearPrintTimes()
	return piuo
}

// SetPrice sets the "price" field.
func (piuo *PublicationInfoUpdateOne) SetPrice(f float64) *PublicationInfoUpdateOne {
	piuo.mutation.ResetPrice()
	piuo.mutation.SetPrice(f)
	return piuo
}

// SetNillablePrice sets the "price" field if the given value is not nil.
func (piuo *PublicationInfoUpdateOne) SetNillablePrice(f *float64) *PublicationInfoUpdateOne {
	if f != nil {
		piuo.SetPrice(*f)
	}
	return piuo
}

// AddPrice adds f to the "price" field.
func (piuo *PublicationInfoUpdateOne) AddPrice(f float64) *PublicationInfoUpdateOne {
	piuo.mutation.AddPrice(f)
	return piuo
}

// ClearPrice clears the value of the "price" field.
func (piuo *PublicationInfoUpdateOne) ClearPrice() *PublicationInfoUpdateOne {
	piuo.mutation.ClearPrice()
	return piuo
}

// SetIntroduction sets the "introduction" field.
func (piuo *PublicationInfoUpdateOne) SetIntroduction(s string) *PublicationInfoUpdateOne {
	piuo.mutation.SetIntroduction(s)
	return piuo
}

// SetNillableIntroduction sets the "introduction" field if the given value is not nil.
func (piuo *PublicationInfoUpdateOne) SetNillableIntroduction(s *string) *PublicationInfoUpdateOne {
	if s != nil {
		piuo.SetIntroduction(*s)
	}
	return piuo
}

// ClearIntroduction clears the value of the "introduction" field.
func (piuo *PublicationInfoUpdateOne) ClearIntroduction() *PublicationInfoUpdateOne {
	piuo.mutation.ClearIntroduction()
	return piuo
}

// SetQuantity sets the "quantity" field.
func (piuo *PublicationInfoUpdateOne) SetQuantity(i int32) *PublicationInfoUpdateOne {
	piuo.mutation.ResetQuantity()
	piuo.mutation.SetQuantity(i)
	return piuo
}

// SetNillableQuantity sets the "quantity" field if the given value is not nil.
func (piuo *PublicationInfoUpdateOne) SetNillableQuantity(i *int32) *PublicationInfoUpdateOne {
	if i != nil {
		piuo.SetQuantity(*i)
	}
	return piuo
}

// AddQuantity adds i to the "quantity" field.
func (piuo *PublicationInfoUpdateOne) AddQuantity(i int32) *PublicationInfoUpdateOne {
	piuo.mutation.AddQuantity(i)
	return piuo
}

// ClearQuantity clears the value of the "quantity" field.
func (piuo *PublicationInfoUpdateOne) ClearQuantity() *PublicationInfoUpdateOne {
	piuo.mutation.ClearQuantity()
	return piuo
}

// SetWordCount sets the "word_count" field.
func (piuo *PublicationInfoUpdateOne) SetWordCount(s string) *PublicationInfoUpdateOne {
	piuo.mutation.SetWordCount(s)
	return piuo
}

// SetNillableWordCount sets the "word_count" field if the given value is not nil.
func (piuo *PublicationInfoUpdateOne) SetNillableWordCount(s *string) *PublicationInfoUpdateOne {
	if s != nil {
		piuo.SetWordCount(*s)
	}
	return piuo
}

// ClearWordCount clears the value of the "word_count" field.
func (piuo *PublicationInfoUpdateOne) ClearWordCount() *PublicationInfoUpdateOne {
	piuo.mutation.ClearWordCount()
	return piuo
}

// SetIsbn sets the "isbn" field.
func (piuo *PublicationInfoUpdateOne) SetIsbn(s string) *PublicationInfoUpdateOne {
	piuo.mutation.SetIsbn(s)
	return piuo
}

// SetNillableIsbn sets the "isbn" field if the given value is not nil.
func (piuo *PublicationInfoUpdateOne) SetNillableIsbn(s *string) *PublicationInfoUpdateOne {
	if s != nil {
		piuo.SetIsbn(*s)
	}
	return piuo
}

// ClearIsbn clears the value of the "isbn" field.
func (piuo *PublicationInfoUpdateOne) ClearIsbn() *PublicationInfoUpdateOne {
	piuo.mutation.ClearIsbn()
	return piuo
}

// SetStorageBy sets the "storage_by" field.
func (piuo *PublicationInfoUpdateOne) SetStorageBy(s string) *PublicationInfoUpdateOne {
	piuo.mutation.SetStorageBy(s)
	return piuo
}

// SetNillableStorageBy sets the "storage_by" field if the given value is not nil.
func (piuo *PublicationInfoUpdateOne) SetNillableStorageBy(s *string) *PublicationInfoUpdateOne {
	if s != nil {
		piuo.SetStorageBy(*s)
	}
	return piuo
}

// ClearStorageBy clears the value of the "storage_by" field.
func (piuo *PublicationInfoUpdateOne) ClearStorageBy() *PublicationInfoUpdateOne {
	piuo.mutation.ClearStorageBy()
	return piuo
}

// SetStorageAt sets the "storage_at" field.
func (piuo *PublicationInfoUpdateOne) SetStorageAt(t time.Time) *PublicationInfoUpdateOne {
	piuo.mutation.SetStorageAt(t)
	return piuo
}

// SetNillableStorageAt sets the "storage_at" field if the given value is not nil.
func (piuo *PublicationInfoUpdateOne) SetNillableStorageAt(t *time.Time) *PublicationInfoUpdateOne {
	if t != nil {
		piuo.SetStorageAt(*t)
	}
	return piuo
}

// ClearStorageAt clears the value of the "storage_at" field.
func (piuo *PublicationInfoUpdateOne) ClearStorageAt() *PublicationInfoUpdateOne {
	piuo.mutation.ClearStorageAt()
	return piuo
}

// SetModifiedBy sets the "modified_by" field.
func (piuo *PublicationInfoUpdateOne) SetModifiedBy(s string) *PublicationInfoUpdateOne {
	piuo.mutation.SetModifiedBy(s)
	return piuo
}

// SetNillableModifiedBy sets the "modified_by" field if the given value is not nil.
func (piuo *PublicationInfoUpdateOne) SetNillableModifiedBy(s *string) *PublicationInfoUpdateOne {
	if s != nil {
		piuo.SetModifiedBy(*s)
	}
	return piuo
}

// ClearModifiedBy clears the value of the "modified_by" field.
func (piuo *PublicationInfoUpdateOne) ClearModifiedBy() *PublicationInfoUpdateOne {
	piuo.mutation.ClearModifiedBy()
	return piuo
}

// SetCreatedAt sets the "created_at" field.
func (piuo *PublicationInfoUpdateOne) SetCreatedAt(t time.Time) *PublicationInfoUpdateOne {
	piuo.mutation.SetCreatedAt(t)
	return piuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (piuo *PublicationInfoUpdateOne) SetNillableCreatedAt(t *time.Time) *PublicationInfoUpdateOne {
	if t != nil {
		piuo.SetCreatedAt(*t)
	}
	return piuo
}

// ClearCreatedAt clears the value of the "created_at" field.
func (piuo *PublicationInfoUpdateOne) ClearCreatedAt() *PublicationInfoUpdateOne {
	piuo.mutation.ClearCreatedAt()
	return piuo
}

// SetModifiedAt sets the "modified_at" field.
func (piuo *PublicationInfoUpdateOne) SetModifiedAt(t time.Time) *PublicationInfoUpdateOne {
	piuo.mutation.SetModifiedAt(t)
	return piuo
}

// SetNillableModifiedAt sets the "modified_at" field if the given value is not nil.
func (piuo *PublicationInfoUpdateOne) SetNillableModifiedAt(t *time.Time) *PublicationInfoUpdateOne {
	if t != nil {
		piuo.SetModifiedAt(*t)
	}
	return piuo
}

// ClearModifiedAt clears the value of the "modified_at" field.
func (piuo *PublicationInfoUpdateOne) ClearModifiedAt() *PublicationInfoUpdateOne {
	piuo.mutation.ClearModifiedAt()
	return piuo
}

// Mutation returns the PublicationInfoMutation object of the builder.
func (piuo *PublicationInfoUpdateOne) Mutation() *PublicationInfoMutation {
	return piuo.mutation
}

// Where appends a list predicates to the PublicationInfoUpdate builder.
func (piuo *PublicationInfoUpdateOne) Where(ps ...predicate.PublicationInfo) *PublicationInfoUpdateOne {
	piuo.mutation.Where(ps...)
	return piuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (piuo *PublicationInfoUpdateOne) Select(field string, fields ...string) *PublicationInfoUpdateOne {
	piuo.fields = append([]string{field}, fields...)
	return piuo
}

// Save executes the query and returns the updated PublicationInfo entity.
func (piuo *PublicationInfoUpdateOne) Save(ctx context.Context) (*PublicationInfo, error) {
	return withHooks(ctx, piuo.sqlSave, piuo.mutation, piuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (piuo *PublicationInfoUpdateOne) SaveX(ctx context.Context) *PublicationInfo {
	node, err := piuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (piuo *PublicationInfoUpdateOne) Exec(ctx context.Context) error {
	_, err := piuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (piuo *PublicationInfoUpdateOne) ExecX(ctx context.Context) {
	if err := piuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (piuo *PublicationInfoUpdateOne) sqlSave(ctx context.Context) (_node *PublicationInfo, err error) {
	_spec := sqlgraph.NewUpdateSpec(publicationinfo.Table, publicationinfo.Columns, sqlgraph.NewFieldSpec(publicationinfo.FieldID, field.TypeInt))
	id, ok := piuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PublicationInfo.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := piuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, publicationinfo.FieldID)
		for _, f := range fields {
			if !publicationinfo.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != publicationinfo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := piuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := piuo.mutation.PublicationName(); ok {
		_spec.SetField(publicationinfo.FieldPublicationName, field.TypeString, value)
	}
	if piuo.mutation.PublicationNameCleared() {
		_spec.ClearField(publicationinfo.FieldPublicationName, field.TypeString)
	}
	if value, ok := piuo.mutation.OrgID(); ok {
		_spec.SetField(publicationinfo.FieldOrgID, field.TypeString, value)
	}
	if piuo.mutation.OrgIDCleared() {
		_spec.ClearField(publicationinfo.FieldOrgID, field.TypeString)
	}
	if value, ok := piuo.mutation.PublishedTimes(); ok {
		_spec.SetField(publicationinfo.FieldPublishedTimes, field.TypeString, value)
	}
	if piuo.mutation.PublishedTimesCleared() {
		_spec.ClearField(publicationinfo.FieldPublishedTimes, field.TypeString)
	}
	if value, ok := piuo.mutation.PrintTimes(); ok {
		_spec.SetField(publicationinfo.FieldPrintTimes, field.TypeString, value)
	}
	if piuo.mutation.PrintTimesCleared() {
		_spec.ClearField(publicationinfo.FieldPrintTimes, field.TypeString)
	}
	if value, ok := piuo.mutation.Price(); ok {
		_spec.SetField(publicationinfo.FieldPrice, field.TypeFloat64, value)
	}
	if value, ok := piuo.mutation.AddedPrice(); ok {
		_spec.AddField(publicationinfo.FieldPrice, field.TypeFloat64, value)
	}
	if piuo.mutation.PriceCleared() {
		_spec.ClearField(publicationinfo.FieldPrice, field.TypeFloat64)
	}
	if value, ok := piuo.mutation.Introduction(); ok {
		_spec.SetField(publicationinfo.FieldIntroduction, field.TypeString, value)
	}
	if piuo.mutation.IntroductionCleared() {
		_spec.ClearField(publicationinfo.FieldIntroduction, field.TypeString)
	}
	if value, ok := piuo.mutation.Quantity(); ok {
		_spec.SetField(publicationinfo.FieldQuantity, field.TypeInt32, value)
	}
	if value, ok := piuo.mutation.AddedQuantity(); ok {
		_spec.AddField(publicationinfo.FieldQuantity, field.TypeInt32, value)
	}
	if piuo.mutation.QuantityCleared() {
		_spec.ClearField(publicationinfo.FieldQuantity, field.TypeInt32)
	}
	if value, ok := piuo.mutation.WordCount(); ok {
		_spec.SetField(publicationinfo.FieldWordCount, field.TypeString, value)
	}
	if piuo.mutation.WordCountCleared() {
		_spec.ClearField(publicationinfo.FieldWordCount, field.TypeString)
	}
	if value, ok := piuo.mutation.Isbn(); ok {
		_spec.SetField(publicationinfo.FieldIsbn, field.TypeString, value)
	}
	if piuo.mutation.IsbnCleared() {
		_spec.ClearField(publicationinfo.FieldIsbn, field.TypeString)
	}
	if value, ok := piuo.mutation.StorageBy(); ok {
		_spec.SetField(publicationinfo.FieldStorageBy, field.TypeString, value)
	}
	if piuo.mutation.StorageByCleared() {
		_spec.ClearField(publicationinfo.FieldStorageBy, field.TypeString)
	}
	if value, ok := piuo.mutation.StorageAt(); ok {
		_spec.SetField(publicationinfo.FieldStorageAt, field.TypeTime, value)
	}
	if piuo.mutation.StorageAtCleared() {
		_spec.ClearField(publicationinfo.FieldStorageAt, field.TypeTime)
	}
	if value, ok := piuo.mutation.ModifiedBy(); ok {
		_spec.SetField(publicationinfo.FieldModifiedBy, field.TypeString, value)
	}
	if piuo.mutation.ModifiedByCleared() {
		_spec.ClearField(publicationinfo.FieldModifiedBy, field.TypeString)
	}
	if value, ok := piuo.mutation.CreatedAt(); ok {
		_spec.SetField(publicationinfo.FieldCreatedAt, field.TypeTime, value)
	}
	if piuo.mutation.CreatedAtCleared() {
		_spec.ClearField(publicationinfo.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := piuo.mutation.ModifiedAt(); ok {
		_spec.SetField(publicationinfo.FieldModifiedAt, field.TypeTime, value)
	}
	if piuo.mutation.ModifiedAtCleared() {
		_spec.ClearField(publicationinfo.FieldModifiedAt, field.TypeTime)
	}
	_node = &PublicationInfo{config: piuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, piuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{publicationinfo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	piuo.mutation.done = true
	return _node, nil
}
