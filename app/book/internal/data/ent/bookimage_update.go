// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/xuanyiying/publication-shop/app/book/internal/data/ent/bookimage"
	"github.com/xuanyiying/publication-shop/app/book/internal/data/ent/predicate"
)

// BookImageUpdate is the builder for updating BookImage entities.
type BookImageUpdate struct {
	config
	hooks    []Hook
	mutation *BookImageMutation
}

// Where appends a list predicates to the BookImageUpdate builder.
func (biu *BookImageUpdate) Where(ps ...predicate.BookImage) *BookImageUpdate {
	biu.mutation.Where(ps...)
	return biu
}

// SetImgURL sets the "img_url" field.
func (biu *BookImageUpdate) SetImgURL(s string) *BookImageUpdate {
	biu.mutation.SetImgURL(s)
	return biu
}

// SetNillableImgURL sets the "img_url" field if the given value is not nil.
func (biu *BookImageUpdate) SetNillableImgURL(s *string) *BookImageUpdate {
	if s != nil {
		biu.SetImgURL(*s)
	}
	return biu
}

// ClearImgURL clears the value of the "img_url" field.
func (biu *BookImageUpdate) ClearImgURL() *BookImageUpdate {
	biu.mutation.ClearImgURL()
	return biu
}

// SetIsbn sets the "isbn" field.
func (biu *BookImageUpdate) SetIsbn(i int) *BookImageUpdate {
	biu.mutation.ResetIsbn()
	biu.mutation.SetIsbn(i)
	return biu
}

// SetNillableIsbn sets the "isbn" field if the given value is not nil.
func (biu *BookImageUpdate) SetNillableIsbn(i *int) *BookImageUpdate {
	if i != nil {
		biu.SetIsbn(*i)
	}
	return biu
}

// AddIsbn adds i to the "isbn" field.
func (biu *BookImageUpdate) AddIsbn(i int) *BookImageUpdate {
	biu.mutation.AddIsbn(i)
	return biu
}

// SetImgEncode sets the "img_encode" field.
func (biu *BookImageUpdate) SetImgEncode(s string) *BookImageUpdate {
	biu.mutation.SetImgEncode(s)
	return biu
}

// SetNillableImgEncode sets the "img_encode" field if the given value is not nil.
func (biu *BookImageUpdate) SetNillableImgEncode(s *string) *BookImageUpdate {
	if s != nil {
		biu.SetImgEncode(*s)
	}
	return biu
}

// ClearImgEncode clears the value of the "img_encode" field.
func (biu *BookImageUpdate) ClearImgEncode() *BookImageUpdate {
	biu.mutation.ClearImgEncode()
	return biu
}

// SetBookID sets the "book_id" field.
func (biu *BookImageUpdate) SetBookID(i int) *BookImageUpdate {
	biu.mutation.ResetBookID()
	biu.mutation.SetBookID(i)
	return biu
}

// SetNillableBookID sets the "book_id" field if the given value is not nil.
func (biu *BookImageUpdate) SetNillableBookID(i *int) *BookImageUpdate {
	if i != nil {
		biu.SetBookID(*i)
	}
	return biu
}

// AddBookID adds i to the "book_id" field.
func (biu *BookImageUpdate) AddBookID(i int) *BookImageUpdate {
	biu.mutation.AddBookID(i)
	return biu
}

// ClearBookID clears the value of the "book_id" field.
func (biu *BookImageUpdate) ClearBookID() *BookImageUpdate {
	biu.mutation.ClearBookID()
	return biu
}

// SetMainFlag sets the "mainFlag" field.
func (biu *BookImageUpdate) SetMainFlag(i int32) *BookImageUpdate {
	biu.mutation.ResetMainFlag()
	biu.mutation.SetMainFlag(i)
	return biu
}

// SetNillableMainFlag sets the "mainFlag" field if the given value is not nil.
func (biu *BookImageUpdate) SetNillableMainFlag(i *int32) *BookImageUpdate {
	if i != nil {
		biu.SetMainFlag(*i)
	}
	return biu
}

// AddMainFlag adds i to the "mainFlag" field.
func (biu *BookImageUpdate) AddMainFlag(i int32) *BookImageUpdate {
	biu.mutation.AddMainFlag(i)
	return biu
}

// ClearMainFlag clears the value of the "mainFlag" field.
func (biu *BookImageUpdate) ClearMainFlag() *BookImageUpdate {
	biu.mutation.ClearMainFlag()
	return biu
}

// Mutation returns the BookImageMutation object of the builder.
func (biu *BookImageUpdate) Mutation() *BookImageMutation {
	return biu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (biu *BookImageUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, biu.sqlSave, biu.mutation, biu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (biu *BookImageUpdate) SaveX(ctx context.Context) int {
	affected, err := biu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (biu *BookImageUpdate) Exec(ctx context.Context) error {
	_, err := biu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (biu *BookImageUpdate) ExecX(ctx context.Context) {
	if err := biu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (biu *BookImageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(bookimage.Table, bookimage.Columns, sqlgraph.NewFieldSpec(bookimage.FieldID, field.TypeInt))
	if ps := biu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := biu.mutation.ImgURL(); ok {
		_spec.SetField(bookimage.FieldImgURL, field.TypeString, value)
	}
	if biu.mutation.ImgURLCleared() {
		_spec.ClearField(bookimage.FieldImgURL, field.TypeString)
	}
	if value, ok := biu.mutation.Isbn(); ok {
		_spec.SetField(bookimage.FieldIsbn, field.TypeInt, value)
	}
	if value, ok := biu.mutation.AddedIsbn(); ok {
		_spec.AddField(bookimage.FieldIsbn, field.TypeInt, value)
	}
	if value, ok := biu.mutation.ImgEncode(); ok {
		_spec.SetField(bookimage.FieldImgEncode, field.TypeString, value)
	}
	if biu.mutation.ImgEncodeCleared() {
		_spec.ClearField(bookimage.FieldImgEncode, field.TypeString)
	}
	if value, ok := biu.mutation.BookID(); ok {
		_spec.SetField(bookimage.FieldBookID, field.TypeInt, value)
	}
	if value, ok := biu.mutation.AddedBookID(); ok {
		_spec.AddField(bookimage.FieldBookID, field.TypeInt, value)
	}
	if biu.mutation.BookIDCleared() {
		_spec.ClearField(bookimage.FieldBookID, field.TypeInt)
	}
	if value, ok := biu.mutation.MainFlag(); ok {
		_spec.SetField(bookimage.FieldMainFlag, field.TypeInt32, value)
	}
	if value, ok := biu.mutation.AddedMainFlag(); ok {
		_spec.AddField(bookimage.FieldMainFlag, field.TypeInt32, value)
	}
	if biu.mutation.MainFlagCleared() {
		_spec.ClearField(bookimage.FieldMainFlag, field.TypeInt32)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, biu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bookimage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	biu.mutation.done = true
	return n, nil
}

// BookImageUpdateOne is the builder for updating a single BookImage entity.
type BookImageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BookImageMutation
}

// SetImgURL sets the "img_url" field.
func (biuo *BookImageUpdateOne) SetImgURL(s string) *BookImageUpdateOne {
	biuo.mutation.SetImgURL(s)
	return biuo
}

// SetNillableImgURL sets the "img_url" field if the given value is not nil.
func (biuo *BookImageUpdateOne) SetNillableImgURL(s *string) *BookImageUpdateOne {
	if s != nil {
		biuo.SetImgURL(*s)
	}
	return biuo
}

// ClearImgURL clears the value of the "img_url" field.
func (biuo *BookImageUpdateOne) ClearImgURL() *BookImageUpdateOne {
	biuo.mutation.ClearImgURL()
	return biuo
}

// SetIsbn sets the "isbn" field.
func (biuo *BookImageUpdateOne) SetIsbn(i int) *BookImageUpdateOne {
	biuo.mutation.ResetIsbn()
	biuo.mutation.SetIsbn(i)
	return biuo
}

// SetNillableIsbn sets the "isbn" field if the given value is not nil.
func (biuo *BookImageUpdateOne) SetNillableIsbn(i *int) *BookImageUpdateOne {
	if i != nil {
		biuo.SetIsbn(*i)
	}
	return biuo
}

// AddIsbn adds i to the "isbn" field.
func (biuo *BookImageUpdateOne) AddIsbn(i int) *BookImageUpdateOne {
	biuo.mutation.AddIsbn(i)
	return biuo
}

// SetImgEncode sets the "img_encode" field.
func (biuo *BookImageUpdateOne) SetImgEncode(s string) *BookImageUpdateOne {
	biuo.mutation.SetImgEncode(s)
	return biuo
}

// SetNillableImgEncode sets the "img_encode" field if the given value is not nil.
func (biuo *BookImageUpdateOne) SetNillableImgEncode(s *string) *BookImageUpdateOne {
	if s != nil {
		biuo.SetImgEncode(*s)
	}
	return biuo
}

// ClearImgEncode clears the value of the "img_encode" field.
func (biuo *BookImageUpdateOne) ClearImgEncode() *BookImageUpdateOne {
	biuo.mutation.ClearImgEncode()
	return biuo
}

// SetBookID sets the "book_id" field.
func (biuo *BookImageUpdateOne) SetBookID(i int) *BookImageUpdateOne {
	biuo.mutation.ResetBookID()
	biuo.mutation.SetBookID(i)
	return biuo
}

// SetNillableBookID sets the "book_id" field if the given value is not nil.
func (biuo *BookImageUpdateOne) SetNillableBookID(i *int) *BookImageUpdateOne {
	if i != nil {
		biuo.SetBookID(*i)
	}
	return biuo
}

// AddBookID adds i to the "book_id" field.
func (biuo *BookImageUpdateOne) AddBookID(i int) *BookImageUpdateOne {
	biuo.mutation.AddBookID(i)
	return biuo
}

// ClearBookID clears the value of the "book_id" field.
func (biuo *BookImageUpdateOne) ClearBookID() *BookImageUpdateOne {
	biuo.mutation.ClearBookID()
	return biuo
}

// SetMainFlag sets the "mainFlag" field.
func (biuo *BookImageUpdateOne) SetMainFlag(i int32) *BookImageUpdateOne {
	biuo.mutation.ResetMainFlag()
	biuo.mutation.SetMainFlag(i)
	return biuo
}

// SetNillableMainFlag sets the "mainFlag" field if the given value is not nil.
func (biuo *BookImageUpdateOne) SetNillableMainFlag(i *int32) *BookImageUpdateOne {
	if i != nil {
		biuo.SetMainFlag(*i)
	}
	return biuo
}

// AddMainFlag adds i to the "mainFlag" field.
func (biuo *BookImageUpdateOne) AddMainFlag(i int32) *BookImageUpdateOne {
	biuo.mutation.AddMainFlag(i)
	return biuo
}

// ClearMainFlag clears the value of the "mainFlag" field.
func (biuo *BookImageUpdateOne) ClearMainFlag() *BookImageUpdateOne {
	biuo.mutation.ClearMainFlag()
	return biuo
}

// Mutation returns the BookImageMutation object of the builder.
func (biuo *BookImageUpdateOne) Mutation() *BookImageMutation {
	return biuo.mutation
}

// Where appends a list predicates to the BookImageUpdate builder.
func (biuo *BookImageUpdateOne) Where(ps ...predicate.BookImage) *BookImageUpdateOne {
	biuo.mutation.Where(ps...)
	return biuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (biuo *BookImageUpdateOne) Select(field string, fields ...string) *BookImageUpdateOne {
	biuo.fields = append([]string{field}, fields...)
	return biuo
}

// Save executes the query and returns the updated BookImage entity.
func (biuo *BookImageUpdateOne) Save(ctx context.Context) (*BookImage, error) {
	return withHooks(ctx, biuo.sqlSave, biuo.mutation, biuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (biuo *BookImageUpdateOne) SaveX(ctx context.Context) *BookImage {
	node, err := biuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (biuo *BookImageUpdateOne) Exec(ctx context.Context) error {
	_, err := biuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (biuo *BookImageUpdateOne) ExecX(ctx context.Context) {
	if err := biuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (biuo *BookImageUpdateOne) sqlSave(ctx context.Context) (_node *BookImage, err error) {
	_spec := sqlgraph.NewUpdateSpec(bookimage.Table, bookimage.Columns, sqlgraph.NewFieldSpec(bookimage.FieldID, field.TypeInt))
	id, ok := biuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "BookImage.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := biuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, bookimage.FieldID)
		for _, f := range fields {
			if !bookimage.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != bookimage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := biuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := biuo.mutation.ImgURL(); ok {
		_spec.SetField(bookimage.FieldImgURL, field.TypeString, value)
	}
	if biuo.mutation.ImgURLCleared() {
		_spec.ClearField(bookimage.FieldImgURL, field.TypeString)
	}
	if value, ok := biuo.mutation.Isbn(); ok {
		_spec.SetField(bookimage.FieldIsbn, field.TypeInt, value)
	}
	if value, ok := biuo.mutation.AddedIsbn(); ok {
		_spec.AddField(bookimage.FieldIsbn, field.TypeInt, value)
	}
	if value, ok := biuo.mutation.ImgEncode(); ok {
		_spec.SetField(bookimage.FieldImgEncode, field.TypeString, value)
	}
	if biuo.mutation.ImgEncodeCleared() {
		_spec.ClearField(bookimage.FieldImgEncode, field.TypeString)
	}
	if value, ok := biuo.mutation.BookID(); ok {
		_spec.SetField(bookimage.FieldBookID, field.TypeInt, value)
	}
	if value, ok := biuo.mutation.AddedBookID(); ok {
		_spec.AddField(bookimage.FieldBookID, field.TypeInt, value)
	}
	if biuo.mutation.BookIDCleared() {
		_spec.ClearField(bookimage.FieldBookID, field.TypeInt)
	}
	if value, ok := biuo.mutation.MainFlag(); ok {
		_spec.SetField(bookimage.FieldMainFlag, field.TypeInt32, value)
	}
	if value, ok := biuo.mutation.AddedMainFlag(); ok {
		_spec.AddField(bookimage.FieldMainFlag, field.TypeInt32, value)
	}
	if biuo.mutation.MainFlagCleared() {
		_spec.ClearField(bookimage.FieldMainFlag, field.TypeInt32)
	}
	_node = &BookImage{config: biuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, biuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bookimage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	biuo.mutation.done = true
	return _node, nil
}