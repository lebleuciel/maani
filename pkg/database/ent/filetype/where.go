// Code generated by ent, DO NOT EDIT.

package filetype

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/lebleuciel/maani/pkg/database/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Filetype {
	return predicate.Filetype(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Filetype {
	return predicate.Filetype(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Filetype {
	return predicate.Filetype(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Filetype {
	return predicate.Filetype(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Filetype {
	return predicate.Filetype(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Filetype {
	return predicate.Filetype(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Filetype {
	return predicate.Filetype(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Filetype {
	return predicate.Filetype(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Filetype {
	return predicate.Filetype(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Filetype {
	return predicate.Filetype(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Filetype {
	return predicate.Filetype(sql.FieldContainsFold(FieldID, id))
}

// AllowedSize applies equality check predicate on the "allowed_size" field. It's identical to AllowedSizeEQ.
func AllowedSize(v int) predicate.Filetype {
	return predicate.Filetype(sql.FieldEQ(FieldAllowedSize, v))
}

// IsBanned applies equality check predicate on the "is_banned" field. It's identical to IsBannedEQ.
func IsBanned(v bool) predicate.Filetype {
	return predicate.Filetype(sql.FieldEQ(FieldIsBanned, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Filetype {
	return predicate.Filetype(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Filetype {
	return predicate.Filetype(sql.FieldEQ(FieldUpdatedAt, v))
}

// AllowedSizeEQ applies the EQ predicate on the "allowed_size" field.
func AllowedSizeEQ(v int) predicate.Filetype {
	return predicate.Filetype(sql.FieldEQ(FieldAllowedSize, v))
}

// AllowedSizeNEQ applies the NEQ predicate on the "allowed_size" field.
func AllowedSizeNEQ(v int) predicate.Filetype {
	return predicate.Filetype(sql.FieldNEQ(FieldAllowedSize, v))
}

// AllowedSizeIn applies the In predicate on the "allowed_size" field.
func AllowedSizeIn(vs ...int) predicate.Filetype {
	return predicate.Filetype(sql.FieldIn(FieldAllowedSize, vs...))
}

// AllowedSizeNotIn applies the NotIn predicate on the "allowed_size" field.
func AllowedSizeNotIn(vs ...int) predicate.Filetype {
	return predicate.Filetype(sql.FieldNotIn(FieldAllowedSize, vs...))
}

// AllowedSizeGT applies the GT predicate on the "allowed_size" field.
func AllowedSizeGT(v int) predicate.Filetype {
	return predicate.Filetype(sql.FieldGT(FieldAllowedSize, v))
}

// AllowedSizeGTE applies the GTE predicate on the "allowed_size" field.
func AllowedSizeGTE(v int) predicate.Filetype {
	return predicate.Filetype(sql.FieldGTE(FieldAllowedSize, v))
}

// AllowedSizeLT applies the LT predicate on the "allowed_size" field.
func AllowedSizeLT(v int) predicate.Filetype {
	return predicate.Filetype(sql.FieldLT(FieldAllowedSize, v))
}

// AllowedSizeLTE applies the LTE predicate on the "allowed_size" field.
func AllowedSizeLTE(v int) predicate.Filetype {
	return predicate.Filetype(sql.FieldLTE(FieldAllowedSize, v))
}

// IsBannedEQ applies the EQ predicate on the "is_banned" field.
func IsBannedEQ(v bool) predicate.Filetype {
	return predicate.Filetype(sql.FieldEQ(FieldIsBanned, v))
}

// IsBannedNEQ applies the NEQ predicate on the "is_banned" field.
func IsBannedNEQ(v bool) predicate.Filetype {
	return predicate.Filetype(sql.FieldNEQ(FieldIsBanned, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Filetype {
	return predicate.Filetype(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Filetype {
	return predicate.Filetype(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Filetype {
	return predicate.Filetype(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Filetype {
	return predicate.Filetype(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Filetype {
	return predicate.Filetype(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Filetype {
	return predicate.Filetype(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Filetype {
	return predicate.Filetype(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Filetype {
	return predicate.Filetype(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.Filetype {
	return predicate.Filetype(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.Filetype {
	return predicate.Filetype(sql.FieldNotNull(FieldCreatedAt))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Filetype {
	return predicate.Filetype(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Filetype {
	return predicate.Filetype(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Filetype {
	return predicate.Filetype(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Filetype {
	return predicate.Filetype(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Filetype {
	return predicate.Filetype(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Filetype {
	return predicate.Filetype(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Filetype {
	return predicate.Filetype(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Filetype {
	return predicate.Filetype(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasFiles applies the HasEdge predicate on the "files" edge.
func HasFiles() predicate.Filetype {
	return predicate.Filetype(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FilesTable, FilesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFilesWith applies the HasEdge predicate on the "files" edge with a given conditions (other predicates).
func HasFilesWith(preds ...predicate.File) predicate.Filetype {
	return predicate.Filetype(func(s *sql.Selector) {
		step := newFilesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Filetype) predicate.Filetype {
	return predicate.Filetype(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Filetype) predicate.Filetype {
	return predicate.Filetype(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Filetype) predicate.Filetype {
	return predicate.Filetype(sql.NotPredicates(p))
}
