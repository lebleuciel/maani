// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lebleuciel/maani/pkg/database/ent/file"
	"github.com/lebleuciel/maani/pkg/database/ent/filetype"
)

// FiletypeCreate is the builder for creating a Filetype entity.
type FiletypeCreate struct {
	config
	mutation *FiletypeMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetAllowedSize sets the "allowed_size" field.
func (fc *FiletypeCreate) SetAllowedSize(i int) *FiletypeCreate {
	fc.mutation.SetAllowedSize(i)
	return fc
}

// SetNillableAllowedSize sets the "allowed_size" field if the given value is not nil.
func (fc *FiletypeCreate) SetNillableAllowedSize(i *int) *FiletypeCreate {
	if i != nil {
		fc.SetAllowedSize(*i)
	}
	return fc
}

// SetIsBanned sets the "is_banned" field.
func (fc *FiletypeCreate) SetIsBanned(b bool) *FiletypeCreate {
	fc.mutation.SetIsBanned(b)
	return fc
}

// SetNillableIsBanned sets the "is_banned" field if the given value is not nil.
func (fc *FiletypeCreate) SetNillableIsBanned(b *bool) *FiletypeCreate {
	if b != nil {
		fc.SetIsBanned(*b)
	}
	return fc
}

// SetCreatedAt sets the "created_at" field.
func (fc *FiletypeCreate) SetCreatedAt(t time.Time) *FiletypeCreate {
	fc.mutation.SetCreatedAt(t)
	return fc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fc *FiletypeCreate) SetNillableCreatedAt(t *time.Time) *FiletypeCreate {
	if t != nil {
		fc.SetCreatedAt(*t)
	}
	return fc
}

// SetUpdatedAt sets the "updated_at" field.
func (fc *FiletypeCreate) SetUpdatedAt(t time.Time) *FiletypeCreate {
	fc.mutation.SetUpdatedAt(t)
	return fc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (fc *FiletypeCreate) SetNillableUpdatedAt(t *time.Time) *FiletypeCreate {
	if t != nil {
		fc.SetUpdatedAt(*t)
	}
	return fc
}

// SetID sets the "id" field.
func (fc *FiletypeCreate) SetID(s string) *FiletypeCreate {
	fc.mutation.SetID(s)
	return fc
}

// AddFileIDs adds the "files" edge to the File entity by IDs.
func (fc *FiletypeCreate) AddFileIDs(ids ...int) *FiletypeCreate {
	fc.mutation.AddFileIDs(ids...)
	return fc
}

// AddFiles adds the "files" edges to the File entity.
func (fc *FiletypeCreate) AddFiles(f ...*File) *FiletypeCreate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fc.AddFileIDs(ids...)
}

// Mutation returns the FiletypeMutation object of the builder.
func (fc *FiletypeCreate) Mutation() *FiletypeMutation {
	return fc.mutation
}

// Save creates the Filetype in the database.
func (fc *FiletypeCreate) Save(ctx context.Context) (*Filetype, error) {
	fc.defaults()
	return withHooks(ctx, fc.sqlSave, fc.mutation, fc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FiletypeCreate) SaveX(ctx context.Context) *Filetype {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fc *FiletypeCreate) Exec(ctx context.Context) error {
	_, err := fc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fc *FiletypeCreate) ExecX(ctx context.Context) {
	if err := fc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fc *FiletypeCreate) defaults() {
	if _, ok := fc.mutation.AllowedSize(); !ok {
		v := filetype.DefaultAllowedSize
		fc.mutation.SetAllowedSize(v)
	}
	if _, ok := fc.mutation.IsBanned(); !ok {
		v := filetype.DefaultIsBanned
		fc.mutation.SetIsBanned(v)
	}
	if _, ok := fc.mutation.CreatedAt(); !ok {
		v := filetype.DefaultCreatedAt()
		fc.mutation.SetCreatedAt(v)
	}
	if _, ok := fc.mutation.UpdatedAt(); !ok {
		v := filetype.DefaultUpdatedAt()
		fc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fc *FiletypeCreate) check() error {
	if _, ok := fc.mutation.AllowedSize(); !ok {
		return &ValidationError{Name: "allowed_size", err: errors.New(`ent: missing required field "Filetype.allowed_size"`)}
	}
	if _, ok := fc.mutation.IsBanned(); !ok {
		return &ValidationError{Name: "is_banned", err: errors.New(`ent: missing required field "Filetype.is_banned"`)}
	}
	if _, ok := fc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Filetype.updated_at"`)}
	}
	if v, ok := fc.mutation.ID(); ok {
		if err := filetype.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Filetype.id": %w`, err)}
		}
	}
	return nil
}

func (fc *FiletypeCreate) sqlSave(ctx context.Context) (*Filetype, error) {
	if err := fc.check(); err != nil {
		return nil, err
	}
	_node, _spec := fc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Filetype.ID type: %T", _spec.ID.Value)
		}
	}
	fc.mutation.id = &_node.ID
	fc.mutation.done = true
	return _node, nil
}

func (fc *FiletypeCreate) createSpec() (*Filetype, *sqlgraph.CreateSpec) {
	var (
		_node = &Filetype{config: fc.config}
		_spec = sqlgraph.NewCreateSpec(filetype.Table, sqlgraph.NewFieldSpec(filetype.FieldID, field.TypeString))
	)
	_spec.OnConflict = fc.conflict
	if id, ok := fc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := fc.mutation.AllowedSize(); ok {
		_spec.SetField(filetype.FieldAllowedSize, field.TypeInt, value)
		_node.AllowedSize = value
	}
	if value, ok := fc.mutation.IsBanned(); ok {
		_spec.SetField(filetype.FieldIsBanned, field.TypeBool, value)
		_node.IsBanned = value
	}
	if value, ok := fc.mutation.CreatedAt(); ok {
		_spec.SetField(filetype.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = &value
	}
	if value, ok := fc.mutation.UpdatedAt(); ok {
		_spec.SetField(filetype.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := fc.mutation.FilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   filetype.FilesTable,
			Columns: []string{filetype.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Filetype.Create().
//		SetAllowedSize(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FiletypeUpsert) {
//			SetAllowedSize(v+v).
//		}).
//		Exec(ctx)
func (fc *FiletypeCreate) OnConflict(opts ...sql.ConflictOption) *FiletypeUpsertOne {
	fc.conflict = opts
	return &FiletypeUpsertOne{
		create: fc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Filetype.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (fc *FiletypeCreate) OnConflictColumns(columns ...string) *FiletypeUpsertOne {
	fc.conflict = append(fc.conflict, sql.ConflictColumns(columns...))
	return &FiletypeUpsertOne{
		create: fc,
	}
}

type (
	// FiletypeUpsertOne is the builder for "upsert"-ing
	//  one Filetype node.
	FiletypeUpsertOne struct {
		create *FiletypeCreate
	}

	// FiletypeUpsert is the "OnConflict" setter.
	FiletypeUpsert struct {
		*sql.UpdateSet
	}
)

// SetAllowedSize sets the "allowed_size" field.
func (u *FiletypeUpsert) SetAllowedSize(v int) *FiletypeUpsert {
	u.Set(filetype.FieldAllowedSize, v)
	return u
}

// UpdateAllowedSize sets the "allowed_size" field to the value that was provided on create.
func (u *FiletypeUpsert) UpdateAllowedSize() *FiletypeUpsert {
	u.SetExcluded(filetype.FieldAllowedSize)
	return u
}

// AddAllowedSize adds v to the "allowed_size" field.
func (u *FiletypeUpsert) AddAllowedSize(v int) *FiletypeUpsert {
	u.Add(filetype.FieldAllowedSize, v)
	return u
}

// SetIsBanned sets the "is_banned" field.
func (u *FiletypeUpsert) SetIsBanned(v bool) *FiletypeUpsert {
	u.Set(filetype.FieldIsBanned, v)
	return u
}

// UpdateIsBanned sets the "is_banned" field to the value that was provided on create.
func (u *FiletypeUpsert) UpdateIsBanned() *FiletypeUpsert {
	u.SetExcluded(filetype.FieldIsBanned)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *FiletypeUpsert) SetCreatedAt(v time.Time) *FiletypeUpsert {
	u.Set(filetype.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *FiletypeUpsert) UpdateCreatedAt() *FiletypeUpsert {
	u.SetExcluded(filetype.FieldCreatedAt)
	return u
}

// ClearCreatedAt clears the value of the "created_at" field.
func (u *FiletypeUpsert) ClearCreatedAt() *FiletypeUpsert {
	u.SetNull(filetype.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *FiletypeUpsert) SetUpdatedAt(v time.Time) *FiletypeUpsert {
	u.Set(filetype.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FiletypeUpsert) UpdateUpdatedAt() *FiletypeUpsert {
	u.SetExcluded(filetype.FieldUpdatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Filetype.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(filetype.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *FiletypeUpsertOne) UpdateNewValues() *FiletypeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(filetype.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Filetype.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *FiletypeUpsertOne) Ignore() *FiletypeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FiletypeUpsertOne) DoNothing() *FiletypeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FiletypeCreate.OnConflict
// documentation for more info.
func (u *FiletypeUpsertOne) Update(set func(*FiletypeUpsert)) *FiletypeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FiletypeUpsert{UpdateSet: update})
	}))
	return u
}

// SetAllowedSize sets the "allowed_size" field.
func (u *FiletypeUpsertOne) SetAllowedSize(v int) *FiletypeUpsertOne {
	return u.Update(func(s *FiletypeUpsert) {
		s.SetAllowedSize(v)
	})
}

// AddAllowedSize adds v to the "allowed_size" field.
func (u *FiletypeUpsertOne) AddAllowedSize(v int) *FiletypeUpsertOne {
	return u.Update(func(s *FiletypeUpsert) {
		s.AddAllowedSize(v)
	})
}

// UpdateAllowedSize sets the "allowed_size" field to the value that was provided on create.
func (u *FiletypeUpsertOne) UpdateAllowedSize() *FiletypeUpsertOne {
	return u.Update(func(s *FiletypeUpsert) {
		s.UpdateAllowedSize()
	})
}

// SetIsBanned sets the "is_banned" field.
func (u *FiletypeUpsertOne) SetIsBanned(v bool) *FiletypeUpsertOne {
	return u.Update(func(s *FiletypeUpsert) {
		s.SetIsBanned(v)
	})
}

// UpdateIsBanned sets the "is_banned" field to the value that was provided on create.
func (u *FiletypeUpsertOne) UpdateIsBanned() *FiletypeUpsertOne {
	return u.Update(func(s *FiletypeUpsert) {
		s.UpdateIsBanned()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *FiletypeUpsertOne) SetCreatedAt(v time.Time) *FiletypeUpsertOne {
	return u.Update(func(s *FiletypeUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *FiletypeUpsertOne) UpdateCreatedAt() *FiletypeUpsertOne {
	return u.Update(func(s *FiletypeUpsert) {
		s.UpdateCreatedAt()
	})
}

// ClearCreatedAt clears the value of the "created_at" field.
func (u *FiletypeUpsertOne) ClearCreatedAt() *FiletypeUpsertOne {
	return u.Update(func(s *FiletypeUpsert) {
		s.ClearCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *FiletypeUpsertOne) SetUpdatedAt(v time.Time) *FiletypeUpsertOne {
	return u.Update(func(s *FiletypeUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FiletypeUpsertOne) UpdateUpdatedAt() *FiletypeUpsertOne {
	return u.Update(func(s *FiletypeUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *FiletypeUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for FiletypeCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FiletypeUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *FiletypeUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: FiletypeUpsertOne.ID is not supported by MySQL driver. Use FiletypeUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *FiletypeUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// FiletypeCreateBulk is the builder for creating many Filetype entities in bulk.
type FiletypeCreateBulk struct {
	config
	err      error
	builders []*FiletypeCreate
	conflict []sql.ConflictOption
}

// Save creates the Filetype entities in the database.
func (fcb *FiletypeCreateBulk) Save(ctx context.Context) ([]*Filetype, error) {
	if fcb.err != nil {
		return nil, fcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(fcb.builders))
	nodes := make([]*Filetype, len(fcb.builders))
	mutators := make([]Mutator, len(fcb.builders))
	for i := range fcb.builders {
		func(i int, root context.Context) {
			builder := fcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FiletypeMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, fcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = fcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, fcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fcb *FiletypeCreateBulk) SaveX(ctx context.Context) []*Filetype {
	v, err := fcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fcb *FiletypeCreateBulk) Exec(ctx context.Context) error {
	_, err := fcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fcb *FiletypeCreateBulk) ExecX(ctx context.Context) {
	if err := fcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Filetype.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FiletypeUpsert) {
//			SetAllowedSize(v+v).
//		}).
//		Exec(ctx)
func (fcb *FiletypeCreateBulk) OnConflict(opts ...sql.ConflictOption) *FiletypeUpsertBulk {
	fcb.conflict = opts
	return &FiletypeUpsertBulk{
		create: fcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Filetype.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (fcb *FiletypeCreateBulk) OnConflictColumns(columns ...string) *FiletypeUpsertBulk {
	fcb.conflict = append(fcb.conflict, sql.ConflictColumns(columns...))
	return &FiletypeUpsertBulk{
		create: fcb,
	}
}

// FiletypeUpsertBulk is the builder for "upsert"-ing
// a bulk of Filetype nodes.
type FiletypeUpsertBulk struct {
	create *FiletypeCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Filetype.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(filetype.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *FiletypeUpsertBulk) UpdateNewValues() *FiletypeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(filetype.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Filetype.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *FiletypeUpsertBulk) Ignore() *FiletypeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FiletypeUpsertBulk) DoNothing() *FiletypeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FiletypeCreateBulk.OnConflict
// documentation for more info.
func (u *FiletypeUpsertBulk) Update(set func(*FiletypeUpsert)) *FiletypeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FiletypeUpsert{UpdateSet: update})
	}))
	return u
}

// SetAllowedSize sets the "allowed_size" field.
func (u *FiletypeUpsertBulk) SetAllowedSize(v int) *FiletypeUpsertBulk {
	return u.Update(func(s *FiletypeUpsert) {
		s.SetAllowedSize(v)
	})
}

// AddAllowedSize adds v to the "allowed_size" field.
func (u *FiletypeUpsertBulk) AddAllowedSize(v int) *FiletypeUpsertBulk {
	return u.Update(func(s *FiletypeUpsert) {
		s.AddAllowedSize(v)
	})
}

// UpdateAllowedSize sets the "allowed_size" field to the value that was provided on create.
func (u *FiletypeUpsertBulk) UpdateAllowedSize() *FiletypeUpsertBulk {
	return u.Update(func(s *FiletypeUpsert) {
		s.UpdateAllowedSize()
	})
}

// SetIsBanned sets the "is_banned" field.
func (u *FiletypeUpsertBulk) SetIsBanned(v bool) *FiletypeUpsertBulk {
	return u.Update(func(s *FiletypeUpsert) {
		s.SetIsBanned(v)
	})
}

// UpdateIsBanned sets the "is_banned" field to the value that was provided on create.
func (u *FiletypeUpsertBulk) UpdateIsBanned() *FiletypeUpsertBulk {
	return u.Update(func(s *FiletypeUpsert) {
		s.UpdateIsBanned()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *FiletypeUpsertBulk) SetCreatedAt(v time.Time) *FiletypeUpsertBulk {
	return u.Update(func(s *FiletypeUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *FiletypeUpsertBulk) UpdateCreatedAt() *FiletypeUpsertBulk {
	return u.Update(func(s *FiletypeUpsert) {
		s.UpdateCreatedAt()
	})
}

// ClearCreatedAt clears the value of the "created_at" field.
func (u *FiletypeUpsertBulk) ClearCreatedAt() *FiletypeUpsertBulk {
	return u.Update(func(s *FiletypeUpsert) {
		s.ClearCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *FiletypeUpsertBulk) SetUpdatedAt(v time.Time) *FiletypeUpsertBulk {
	return u.Update(func(s *FiletypeUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FiletypeUpsertBulk) UpdateUpdatedAt() *FiletypeUpsertBulk {
	return u.Update(func(s *FiletypeUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *FiletypeUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the FiletypeCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for FiletypeCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FiletypeUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
