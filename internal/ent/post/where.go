// Code generated by ent, DO NOT EDIT.

package post

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/Southclaws/storyden/internal/ent/predicate"
	"github.com/rs/xid"
)

// ID filters vertices based on their ID field.
func ID(id xid.ID) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id xid.ID) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id xid.ID) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...xid.ID) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...xid.ID) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id xid.ID) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id xid.ID) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id xid.ID) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id xid.ID) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldDeletedAt, v))
}

// First applies equality check predicate on the "first" field. It's identical to FirstEQ.
func First(v bool) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldFirst, v))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldTitle, v))
}

// Slug applies equality check predicate on the "slug" field. It's identical to SlugEQ.
func Slug(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldSlug, v))
}

// Pinned applies equality check predicate on the "pinned" field. It's identical to PinnedEQ.
func Pinned(v bool) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldPinned, v))
}

// RootPostID applies equality check predicate on the "root_post_id" field. It's identical to RootPostIDEQ.
func RootPostID(v xid.ID) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldRootPostID, v))
}

// ReplyToPostID applies equality check predicate on the "reply_to_post_id" field. It's identical to ReplyToPostIDEQ.
func ReplyToPostID(v xid.ID) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldReplyToPostID, v))
}

// Body applies equality check predicate on the "body" field. It's identical to BodyEQ.
func Body(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldBody, v))
}

// Short applies equality check predicate on the "short" field. It's identical to ShortEQ.
func Short(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldShort, v))
}

// CategoryID applies equality check predicate on the "category_id" field. It's identical to CategoryIDEQ.
func CategoryID(v xid.ID) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldCategoryID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Post {
	return predicate.Post(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Post {
	return predicate.Post(sql.FieldNotNull(FieldDeletedAt))
}

// FirstEQ applies the EQ predicate on the "first" field.
func FirstEQ(v bool) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldFirst, v))
}

// FirstNEQ applies the NEQ predicate on the "first" field.
func FirstNEQ(v bool) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldFirst, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Post {
	return predicate.Post(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleIsNil applies the IsNil predicate on the "title" field.
func TitleIsNil() predicate.Post {
	return predicate.Post(sql.FieldIsNull(FieldTitle))
}

// TitleNotNil applies the NotNil predicate on the "title" field.
func TitleNotNil() predicate.Post {
	return predicate.Post(sql.FieldNotNull(FieldTitle))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Post {
	return predicate.Post(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Post {
	return predicate.Post(sql.FieldContainsFold(FieldTitle, v))
}

// SlugEQ applies the EQ predicate on the "slug" field.
func SlugEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldSlug, v))
}

// SlugNEQ applies the NEQ predicate on the "slug" field.
func SlugNEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldSlug, v))
}

// SlugIn applies the In predicate on the "slug" field.
func SlugIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldSlug, vs...))
}

// SlugNotIn applies the NotIn predicate on the "slug" field.
func SlugNotIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldSlug, vs...))
}

// SlugGT applies the GT predicate on the "slug" field.
func SlugGT(v string) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldSlug, v))
}

// SlugGTE applies the GTE predicate on the "slug" field.
func SlugGTE(v string) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldSlug, v))
}

// SlugLT applies the LT predicate on the "slug" field.
func SlugLT(v string) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldSlug, v))
}

// SlugLTE applies the LTE predicate on the "slug" field.
func SlugLTE(v string) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldSlug, v))
}

// SlugContains applies the Contains predicate on the "slug" field.
func SlugContains(v string) predicate.Post {
	return predicate.Post(sql.FieldContains(FieldSlug, v))
}

// SlugHasPrefix applies the HasPrefix predicate on the "slug" field.
func SlugHasPrefix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasPrefix(FieldSlug, v))
}

// SlugHasSuffix applies the HasSuffix predicate on the "slug" field.
func SlugHasSuffix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasSuffix(FieldSlug, v))
}

// SlugIsNil applies the IsNil predicate on the "slug" field.
func SlugIsNil() predicate.Post {
	return predicate.Post(sql.FieldIsNull(FieldSlug))
}

// SlugNotNil applies the NotNil predicate on the "slug" field.
func SlugNotNil() predicate.Post {
	return predicate.Post(sql.FieldNotNull(FieldSlug))
}

// SlugEqualFold applies the EqualFold predicate on the "slug" field.
func SlugEqualFold(v string) predicate.Post {
	return predicate.Post(sql.FieldEqualFold(FieldSlug, v))
}

// SlugContainsFold applies the ContainsFold predicate on the "slug" field.
func SlugContainsFold(v string) predicate.Post {
	return predicate.Post(sql.FieldContainsFold(FieldSlug, v))
}

// PinnedEQ applies the EQ predicate on the "pinned" field.
func PinnedEQ(v bool) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldPinned, v))
}

// PinnedNEQ applies the NEQ predicate on the "pinned" field.
func PinnedNEQ(v bool) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldPinned, v))
}

// RootPostIDEQ applies the EQ predicate on the "root_post_id" field.
func RootPostIDEQ(v xid.ID) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldRootPostID, v))
}

// RootPostIDNEQ applies the NEQ predicate on the "root_post_id" field.
func RootPostIDNEQ(v xid.ID) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldRootPostID, v))
}

// RootPostIDIn applies the In predicate on the "root_post_id" field.
func RootPostIDIn(vs ...xid.ID) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldRootPostID, vs...))
}

// RootPostIDNotIn applies the NotIn predicate on the "root_post_id" field.
func RootPostIDNotIn(vs ...xid.ID) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldRootPostID, vs...))
}

// RootPostIDGT applies the GT predicate on the "root_post_id" field.
func RootPostIDGT(v xid.ID) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldRootPostID, v))
}

// RootPostIDGTE applies the GTE predicate on the "root_post_id" field.
func RootPostIDGTE(v xid.ID) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldRootPostID, v))
}

// RootPostIDLT applies the LT predicate on the "root_post_id" field.
func RootPostIDLT(v xid.ID) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldRootPostID, v))
}

// RootPostIDLTE applies the LTE predicate on the "root_post_id" field.
func RootPostIDLTE(v xid.ID) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldRootPostID, v))
}

// RootPostIDContains applies the Contains predicate on the "root_post_id" field.
func RootPostIDContains(v xid.ID) predicate.Post {
	vc := v.String()
	return predicate.Post(sql.FieldContains(FieldRootPostID, vc))
}

// RootPostIDHasPrefix applies the HasPrefix predicate on the "root_post_id" field.
func RootPostIDHasPrefix(v xid.ID) predicate.Post {
	vc := v.String()
	return predicate.Post(sql.FieldHasPrefix(FieldRootPostID, vc))
}

// RootPostIDHasSuffix applies the HasSuffix predicate on the "root_post_id" field.
func RootPostIDHasSuffix(v xid.ID) predicate.Post {
	vc := v.String()
	return predicate.Post(sql.FieldHasSuffix(FieldRootPostID, vc))
}

// RootPostIDIsNil applies the IsNil predicate on the "root_post_id" field.
func RootPostIDIsNil() predicate.Post {
	return predicate.Post(sql.FieldIsNull(FieldRootPostID))
}

// RootPostIDNotNil applies the NotNil predicate on the "root_post_id" field.
func RootPostIDNotNil() predicate.Post {
	return predicate.Post(sql.FieldNotNull(FieldRootPostID))
}

// RootPostIDEqualFold applies the EqualFold predicate on the "root_post_id" field.
func RootPostIDEqualFold(v xid.ID) predicate.Post {
	vc := v.String()
	return predicate.Post(sql.FieldEqualFold(FieldRootPostID, vc))
}

// RootPostIDContainsFold applies the ContainsFold predicate on the "root_post_id" field.
func RootPostIDContainsFold(v xid.ID) predicate.Post {
	vc := v.String()
	return predicate.Post(sql.FieldContainsFold(FieldRootPostID, vc))
}

// ReplyToPostIDEQ applies the EQ predicate on the "reply_to_post_id" field.
func ReplyToPostIDEQ(v xid.ID) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldReplyToPostID, v))
}

// ReplyToPostIDNEQ applies the NEQ predicate on the "reply_to_post_id" field.
func ReplyToPostIDNEQ(v xid.ID) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldReplyToPostID, v))
}

// ReplyToPostIDIn applies the In predicate on the "reply_to_post_id" field.
func ReplyToPostIDIn(vs ...xid.ID) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldReplyToPostID, vs...))
}

// ReplyToPostIDNotIn applies the NotIn predicate on the "reply_to_post_id" field.
func ReplyToPostIDNotIn(vs ...xid.ID) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldReplyToPostID, vs...))
}

// ReplyToPostIDGT applies the GT predicate on the "reply_to_post_id" field.
func ReplyToPostIDGT(v xid.ID) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldReplyToPostID, v))
}

// ReplyToPostIDGTE applies the GTE predicate on the "reply_to_post_id" field.
func ReplyToPostIDGTE(v xid.ID) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldReplyToPostID, v))
}

// ReplyToPostIDLT applies the LT predicate on the "reply_to_post_id" field.
func ReplyToPostIDLT(v xid.ID) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldReplyToPostID, v))
}

// ReplyToPostIDLTE applies the LTE predicate on the "reply_to_post_id" field.
func ReplyToPostIDLTE(v xid.ID) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldReplyToPostID, v))
}

// ReplyToPostIDContains applies the Contains predicate on the "reply_to_post_id" field.
func ReplyToPostIDContains(v xid.ID) predicate.Post {
	vc := v.String()
	return predicate.Post(sql.FieldContains(FieldReplyToPostID, vc))
}

// ReplyToPostIDHasPrefix applies the HasPrefix predicate on the "reply_to_post_id" field.
func ReplyToPostIDHasPrefix(v xid.ID) predicate.Post {
	vc := v.String()
	return predicate.Post(sql.FieldHasPrefix(FieldReplyToPostID, vc))
}

// ReplyToPostIDHasSuffix applies the HasSuffix predicate on the "reply_to_post_id" field.
func ReplyToPostIDHasSuffix(v xid.ID) predicate.Post {
	vc := v.String()
	return predicate.Post(sql.FieldHasSuffix(FieldReplyToPostID, vc))
}

// ReplyToPostIDIsNil applies the IsNil predicate on the "reply_to_post_id" field.
func ReplyToPostIDIsNil() predicate.Post {
	return predicate.Post(sql.FieldIsNull(FieldReplyToPostID))
}

// ReplyToPostIDNotNil applies the NotNil predicate on the "reply_to_post_id" field.
func ReplyToPostIDNotNil() predicate.Post {
	return predicate.Post(sql.FieldNotNull(FieldReplyToPostID))
}

// ReplyToPostIDEqualFold applies the EqualFold predicate on the "reply_to_post_id" field.
func ReplyToPostIDEqualFold(v xid.ID) predicate.Post {
	vc := v.String()
	return predicate.Post(sql.FieldEqualFold(FieldReplyToPostID, vc))
}

// ReplyToPostIDContainsFold applies the ContainsFold predicate on the "reply_to_post_id" field.
func ReplyToPostIDContainsFold(v xid.ID) predicate.Post {
	vc := v.String()
	return predicate.Post(sql.FieldContainsFold(FieldReplyToPostID, vc))
}

// BodyEQ applies the EQ predicate on the "body" field.
func BodyEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldBody, v))
}

// BodyNEQ applies the NEQ predicate on the "body" field.
func BodyNEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldBody, v))
}

// BodyIn applies the In predicate on the "body" field.
func BodyIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldBody, vs...))
}

// BodyNotIn applies the NotIn predicate on the "body" field.
func BodyNotIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldBody, vs...))
}

// BodyGT applies the GT predicate on the "body" field.
func BodyGT(v string) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldBody, v))
}

// BodyGTE applies the GTE predicate on the "body" field.
func BodyGTE(v string) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldBody, v))
}

// BodyLT applies the LT predicate on the "body" field.
func BodyLT(v string) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldBody, v))
}

// BodyLTE applies the LTE predicate on the "body" field.
func BodyLTE(v string) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldBody, v))
}

// BodyContains applies the Contains predicate on the "body" field.
func BodyContains(v string) predicate.Post {
	return predicate.Post(sql.FieldContains(FieldBody, v))
}

// BodyHasPrefix applies the HasPrefix predicate on the "body" field.
func BodyHasPrefix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasPrefix(FieldBody, v))
}

// BodyHasSuffix applies the HasSuffix predicate on the "body" field.
func BodyHasSuffix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasSuffix(FieldBody, v))
}

// BodyEqualFold applies the EqualFold predicate on the "body" field.
func BodyEqualFold(v string) predicate.Post {
	return predicate.Post(sql.FieldEqualFold(FieldBody, v))
}

// BodyContainsFold applies the ContainsFold predicate on the "body" field.
func BodyContainsFold(v string) predicate.Post {
	return predicate.Post(sql.FieldContainsFold(FieldBody, v))
}

// ShortEQ applies the EQ predicate on the "short" field.
func ShortEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldShort, v))
}

// ShortNEQ applies the NEQ predicate on the "short" field.
func ShortNEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldShort, v))
}

// ShortIn applies the In predicate on the "short" field.
func ShortIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldShort, vs...))
}

// ShortNotIn applies the NotIn predicate on the "short" field.
func ShortNotIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldShort, vs...))
}

// ShortGT applies the GT predicate on the "short" field.
func ShortGT(v string) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldShort, v))
}

// ShortGTE applies the GTE predicate on the "short" field.
func ShortGTE(v string) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldShort, v))
}

// ShortLT applies the LT predicate on the "short" field.
func ShortLT(v string) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldShort, v))
}

// ShortLTE applies the LTE predicate on the "short" field.
func ShortLTE(v string) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldShort, v))
}

// ShortContains applies the Contains predicate on the "short" field.
func ShortContains(v string) predicate.Post {
	return predicate.Post(sql.FieldContains(FieldShort, v))
}

// ShortHasPrefix applies the HasPrefix predicate on the "short" field.
func ShortHasPrefix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasPrefix(FieldShort, v))
}

// ShortHasSuffix applies the HasSuffix predicate on the "short" field.
func ShortHasSuffix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasSuffix(FieldShort, v))
}

// ShortEqualFold applies the EqualFold predicate on the "short" field.
func ShortEqualFold(v string) predicate.Post {
	return predicate.Post(sql.FieldEqualFold(FieldShort, v))
}

// ShortContainsFold applies the ContainsFold predicate on the "short" field.
func ShortContainsFold(v string) predicate.Post {
	return predicate.Post(sql.FieldContainsFold(FieldShort, v))
}

// MetadataIsNil applies the IsNil predicate on the "metadata" field.
func MetadataIsNil() predicate.Post {
	return predicate.Post(sql.FieldIsNull(FieldMetadata))
}

// MetadataNotNil applies the NotNil predicate on the "metadata" field.
func MetadataNotNil() predicate.Post {
	return predicate.Post(sql.FieldNotNull(FieldMetadata))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldStatus, vs...))
}

// CategoryIDEQ applies the EQ predicate on the "category_id" field.
func CategoryIDEQ(v xid.ID) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldCategoryID, v))
}

// CategoryIDNEQ applies the NEQ predicate on the "category_id" field.
func CategoryIDNEQ(v xid.ID) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldCategoryID, v))
}

// CategoryIDIn applies the In predicate on the "category_id" field.
func CategoryIDIn(vs ...xid.ID) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldCategoryID, vs...))
}

// CategoryIDNotIn applies the NotIn predicate on the "category_id" field.
func CategoryIDNotIn(vs ...xid.ID) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldCategoryID, vs...))
}

// CategoryIDGT applies the GT predicate on the "category_id" field.
func CategoryIDGT(v xid.ID) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldCategoryID, v))
}

// CategoryIDGTE applies the GTE predicate on the "category_id" field.
func CategoryIDGTE(v xid.ID) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldCategoryID, v))
}

// CategoryIDLT applies the LT predicate on the "category_id" field.
func CategoryIDLT(v xid.ID) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldCategoryID, v))
}

// CategoryIDLTE applies the LTE predicate on the "category_id" field.
func CategoryIDLTE(v xid.ID) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldCategoryID, v))
}

// CategoryIDContains applies the Contains predicate on the "category_id" field.
func CategoryIDContains(v xid.ID) predicate.Post {
	vc := v.String()
	return predicate.Post(sql.FieldContains(FieldCategoryID, vc))
}

// CategoryIDHasPrefix applies the HasPrefix predicate on the "category_id" field.
func CategoryIDHasPrefix(v xid.ID) predicate.Post {
	vc := v.String()
	return predicate.Post(sql.FieldHasPrefix(FieldCategoryID, vc))
}

// CategoryIDHasSuffix applies the HasSuffix predicate on the "category_id" field.
func CategoryIDHasSuffix(v xid.ID) predicate.Post {
	vc := v.String()
	return predicate.Post(sql.FieldHasSuffix(FieldCategoryID, vc))
}

// CategoryIDIsNil applies the IsNil predicate on the "category_id" field.
func CategoryIDIsNil() predicate.Post {
	return predicate.Post(sql.FieldIsNull(FieldCategoryID))
}

// CategoryIDNotNil applies the NotNil predicate on the "category_id" field.
func CategoryIDNotNil() predicate.Post {
	return predicate.Post(sql.FieldNotNull(FieldCategoryID))
}

// CategoryIDEqualFold applies the EqualFold predicate on the "category_id" field.
func CategoryIDEqualFold(v xid.ID) predicate.Post {
	vc := v.String()
	return predicate.Post(sql.FieldEqualFold(FieldCategoryID, vc))
}

// CategoryIDContainsFold applies the ContainsFold predicate on the "category_id" field.
func CategoryIDContainsFold(v xid.ID) predicate.Post {
	vc := v.String()
	return predicate.Post(sql.FieldContainsFold(FieldCategoryID, vc))
}

// HasAuthor applies the HasEdge predicate on the "author" edge.
func HasAuthor() predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AuthorTable, AuthorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAuthorWith applies the HasEdge predicate on the "author" edge with a given conditions (other predicates).
func HasAuthorWith(preds ...predicate.Account) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AuthorInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AuthorTable, AuthorColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCategory applies the HasEdge predicate on the "category" edge.
func HasCategory() predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CategoryTable, CategoryColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCategoryWith applies the HasEdge predicate on the "category" edge with a given conditions (other predicates).
func HasCategoryWith(preds ...predicate.Category) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CategoryInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CategoryTable, CategoryColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTags applies the HasEdge predicate on the "tags" edge.
func HasTags() predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, TagsTable, TagsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTagsWith applies the HasEdge predicate on the "tags" edge with a given conditions (other predicates).
func HasTagsWith(preds ...predicate.Tag) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TagsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, TagsTable, TagsPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRoot applies the HasEdge predicate on the "root" edge.
func HasRoot() predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RootTable, RootColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRootWith applies the HasEdge predicate on the "root" edge with a given conditions (other predicates).
func HasRootWith(preds ...predicate.Post) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RootTable, RootColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPosts applies the HasEdge predicate on the "posts" edge.
func HasPosts() predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PostsTable, PostsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPostsWith applies the HasEdge predicate on the "posts" edge with a given conditions (other predicates).
func HasPostsWith(preds ...predicate.Post) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PostsTable, PostsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasReplyTo applies the HasEdge predicate on the "replyTo" edge.
func HasReplyTo() predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ReplyToTable, ReplyToColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasReplyToWith applies the HasEdge predicate on the "replyTo" edge with a given conditions (other predicates).
func HasReplyToWith(preds ...predicate.Post) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ReplyToTable, ReplyToColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasReplies applies the HasEdge predicate on the "replies" edge.
func HasReplies() predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RepliesTable, RepliesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRepliesWith applies the HasEdge predicate on the "replies" edge with a given conditions (other predicates).
func HasRepliesWith(preds ...predicate.Post) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RepliesTable, RepliesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasReacts applies the HasEdge predicate on the "reacts" edge.
func HasReacts() predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ReactsTable, ReactsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasReactsWith applies the HasEdge predicate on the "reacts" edge with a given conditions (other predicates).
func HasReactsWith(preds ...predicate.React) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ReactsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ReactsTable, ReactsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAssets applies the HasEdge predicate on the "assets" edge.
func HasAssets() predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, AssetsTable, AssetsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAssetsWith applies the HasEdge predicate on the "assets" edge with a given conditions (other predicates).
func HasAssetsWith(preds ...predicate.Asset) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AssetsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, AssetsTable, AssetsPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCollections applies the HasEdge predicate on the "collections" edge.
func HasCollections() predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, CollectionsTable, CollectionsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCollectionsWith applies the HasEdge predicate on the "collections" edge with a given conditions (other predicates).
func HasCollectionsWith(preds ...predicate.Collection) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CollectionsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, CollectionsTable, CollectionsPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Post) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Post) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Post) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		p(s.Not())
	})
}
