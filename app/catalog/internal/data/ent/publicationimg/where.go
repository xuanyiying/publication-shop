// Code generated by ent, DO NOT EDIT.

package publicationimg

import (
	"entgo.io/ent/dialect/sql"
	"github.com/xuanyiying/publication-shop/app/catalog/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldLTE(FieldID, id))
}

// ImgURL applies equality check predicate on the "img_url" field. It's identical to ImgURLEQ.
func ImgURL(v string) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldEQ(FieldImgURL, v))
}

// Isbn applies equality check predicate on the "isbn" field. It's identical to IsbnEQ.
func Isbn(v int) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldEQ(FieldIsbn, v))
}

// ImgEncode applies equality check predicate on the "img_encode" field. It's identical to ImgEncodeEQ.
func ImgEncode(v string) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldEQ(FieldImgEncode, v))
}

// PublicationID applies equality check predicate on the "publication_id" field. It's identical to PublicationIDEQ.
func PublicationID(v int) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldEQ(FieldPublicationID, v))
}

// MainFlag applies equality check predicate on the "mainFlag" field. It's identical to MainFlagEQ.
func MainFlag(v int32) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldEQ(FieldMainFlag, v))
}

// ImgURLEQ applies the EQ predicate on the "img_url" field.
func ImgURLEQ(v string) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldEQ(FieldImgURL, v))
}

// ImgURLNEQ applies the NEQ predicate on the "img_url" field.
func ImgURLNEQ(v string) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldNEQ(FieldImgURL, v))
}

// ImgURLIn applies the In predicate on the "img_url" field.
func ImgURLIn(vs ...string) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldIn(FieldImgURL, vs...))
}

// ImgURLNotIn applies the NotIn predicate on the "img_url" field.
func ImgURLNotIn(vs ...string) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldNotIn(FieldImgURL, vs...))
}

// ImgURLGT applies the GT predicate on the "img_url" field.
func ImgURLGT(v string) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldGT(FieldImgURL, v))
}

// ImgURLGTE applies the GTE predicate on the "img_url" field.
func ImgURLGTE(v string) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldGTE(FieldImgURL, v))
}

// ImgURLLT applies the LT predicate on the "img_url" field.
func ImgURLLT(v string) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldLT(FieldImgURL, v))
}

// ImgURLLTE applies the LTE predicate on the "img_url" field.
func ImgURLLTE(v string) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldLTE(FieldImgURL, v))
}

// ImgURLContains applies the Contains predicate on the "img_url" field.
func ImgURLContains(v string) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldContains(FieldImgURL, v))
}

// ImgURLHasPrefix applies the HasPrefix predicate on the "img_url" field.
func ImgURLHasPrefix(v string) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldHasPrefix(FieldImgURL, v))
}

// ImgURLHasSuffix applies the HasSuffix predicate on the "img_url" field.
func ImgURLHasSuffix(v string) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldHasSuffix(FieldImgURL, v))
}

// ImgURLIsNil applies the IsNil predicate on the "img_url" field.
func ImgURLIsNil() predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldIsNull(FieldImgURL))
}

// ImgURLNotNil applies the NotNil predicate on the "img_url" field.
func ImgURLNotNil() predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldNotNull(FieldImgURL))
}

// ImgURLEqualFold applies the EqualFold predicate on the "img_url" field.
func ImgURLEqualFold(v string) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldEqualFold(FieldImgURL, v))
}

// ImgURLContainsFold applies the ContainsFold predicate on the "img_url" field.
func ImgURLContainsFold(v string) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldContainsFold(FieldImgURL, v))
}

// IsbnEQ applies the EQ predicate on the "isbn" field.
func IsbnEQ(v int) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldEQ(FieldIsbn, v))
}

// IsbnNEQ applies the NEQ predicate on the "isbn" field.
func IsbnNEQ(v int) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldNEQ(FieldIsbn, v))
}

// IsbnIn applies the In predicate on the "isbn" field.
func IsbnIn(vs ...int) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldIn(FieldIsbn, vs...))
}

// IsbnNotIn applies the NotIn predicate on the "isbn" field.
func IsbnNotIn(vs ...int) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldNotIn(FieldIsbn, vs...))
}

// IsbnGT applies the GT predicate on the "isbn" field.
func IsbnGT(v int) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldGT(FieldIsbn, v))
}

// IsbnGTE applies the GTE predicate on the "isbn" field.
func IsbnGTE(v int) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldGTE(FieldIsbn, v))
}

// IsbnLT applies the LT predicate on the "isbn" field.
func IsbnLT(v int) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldLT(FieldIsbn, v))
}

// IsbnLTE applies the LTE predicate on the "isbn" field.
func IsbnLTE(v int) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldLTE(FieldIsbn, v))
}

// ImgEncodeEQ applies the EQ predicate on the "img_encode" field.
func ImgEncodeEQ(v string) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldEQ(FieldImgEncode, v))
}

// ImgEncodeNEQ applies the NEQ predicate on the "img_encode" field.
func ImgEncodeNEQ(v string) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldNEQ(FieldImgEncode, v))
}

// ImgEncodeIn applies the In predicate on the "img_encode" field.
func ImgEncodeIn(vs ...string) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldIn(FieldImgEncode, vs...))
}

// ImgEncodeNotIn applies the NotIn predicate on the "img_encode" field.
func ImgEncodeNotIn(vs ...string) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldNotIn(FieldImgEncode, vs...))
}

// ImgEncodeGT applies the GT predicate on the "img_encode" field.
func ImgEncodeGT(v string) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldGT(FieldImgEncode, v))
}

// ImgEncodeGTE applies the GTE predicate on the "img_encode" field.
func ImgEncodeGTE(v string) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldGTE(FieldImgEncode, v))
}

// ImgEncodeLT applies the LT predicate on the "img_encode" field.
func ImgEncodeLT(v string) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldLT(FieldImgEncode, v))
}

// ImgEncodeLTE applies the LTE predicate on the "img_encode" field.
func ImgEncodeLTE(v string) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldLTE(FieldImgEncode, v))
}

// ImgEncodeContains applies the Contains predicate on the "img_encode" field.
func ImgEncodeContains(v string) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldContains(FieldImgEncode, v))
}

// ImgEncodeHasPrefix applies the HasPrefix predicate on the "img_encode" field.
func ImgEncodeHasPrefix(v string) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldHasPrefix(FieldImgEncode, v))
}

// ImgEncodeHasSuffix applies the HasSuffix predicate on the "img_encode" field.
func ImgEncodeHasSuffix(v string) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldHasSuffix(FieldImgEncode, v))
}

// ImgEncodeIsNil applies the IsNil predicate on the "img_encode" field.
func ImgEncodeIsNil() predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldIsNull(FieldImgEncode))
}

// ImgEncodeNotNil applies the NotNil predicate on the "img_encode" field.
func ImgEncodeNotNil() predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldNotNull(FieldImgEncode))
}

// ImgEncodeEqualFold applies the EqualFold predicate on the "img_encode" field.
func ImgEncodeEqualFold(v string) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldEqualFold(FieldImgEncode, v))
}

// ImgEncodeContainsFold applies the ContainsFold predicate on the "img_encode" field.
func ImgEncodeContainsFold(v string) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldContainsFold(FieldImgEncode, v))
}

// PublicationIDEQ applies the EQ predicate on the "publication_id" field.
func PublicationIDEQ(v int) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldEQ(FieldPublicationID, v))
}

// PublicationIDNEQ applies the NEQ predicate on the "publication_id" field.
func PublicationIDNEQ(v int) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldNEQ(FieldPublicationID, v))
}

// PublicationIDIn applies the In predicate on the "publication_id" field.
func PublicationIDIn(vs ...int) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldIn(FieldPublicationID, vs...))
}

// PublicationIDNotIn applies the NotIn predicate on the "publication_id" field.
func PublicationIDNotIn(vs ...int) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldNotIn(FieldPublicationID, vs...))
}

// PublicationIDGT applies the GT predicate on the "publication_id" field.
func PublicationIDGT(v int) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldGT(FieldPublicationID, v))
}

// PublicationIDGTE applies the GTE predicate on the "publication_id" field.
func PublicationIDGTE(v int) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldGTE(FieldPublicationID, v))
}

// PublicationIDLT applies the LT predicate on the "publication_id" field.
func PublicationIDLT(v int) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldLT(FieldPublicationID, v))
}

// PublicationIDLTE applies the LTE predicate on the "publication_id" field.
func PublicationIDLTE(v int) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldLTE(FieldPublicationID, v))
}

// PublicationIDIsNil applies the IsNil predicate on the "publication_id" field.
func PublicationIDIsNil() predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldIsNull(FieldPublicationID))
}

// PublicationIDNotNil applies the NotNil predicate on the "publication_id" field.
func PublicationIDNotNil() predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldNotNull(FieldPublicationID))
}

// MainFlagEQ applies the EQ predicate on the "mainFlag" field.
func MainFlagEQ(v int32) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldEQ(FieldMainFlag, v))
}

// MainFlagNEQ applies the NEQ predicate on the "mainFlag" field.
func MainFlagNEQ(v int32) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldNEQ(FieldMainFlag, v))
}

// MainFlagIn applies the In predicate on the "mainFlag" field.
func MainFlagIn(vs ...int32) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldIn(FieldMainFlag, vs...))
}

// MainFlagNotIn applies the NotIn predicate on the "mainFlag" field.
func MainFlagNotIn(vs ...int32) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldNotIn(FieldMainFlag, vs...))
}

// MainFlagGT applies the GT predicate on the "mainFlag" field.
func MainFlagGT(v int32) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldGT(FieldMainFlag, v))
}

// MainFlagGTE applies the GTE predicate on the "mainFlag" field.
func MainFlagGTE(v int32) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldGTE(FieldMainFlag, v))
}

// MainFlagLT applies the LT predicate on the "mainFlag" field.
func MainFlagLT(v int32) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldLT(FieldMainFlag, v))
}

// MainFlagLTE applies the LTE predicate on the "mainFlag" field.
func MainFlagLTE(v int32) predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldLTE(FieldMainFlag, v))
}

// MainFlagIsNil applies the IsNil predicate on the "mainFlag" field.
func MainFlagIsNil() predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldIsNull(FieldMainFlag))
}

// MainFlagNotNil applies the NotNil predicate on the "mainFlag" field.
func MainFlagNotNil() predicate.PublicationImg {
	return predicate.PublicationImg(sql.FieldNotNull(FieldMainFlag))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PublicationImg) predicate.PublicationImg {
	return predicate.PublicationImg(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PublicationImg) predicate.PublicationImg {
	return predicate.PublicationImg(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.PublicationImg) predicate.PublicationImg {
	return predicate.PublicationImg(sql.NotPredicates(p))
}