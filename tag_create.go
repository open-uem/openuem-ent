// Code generated by ent, DO NOT EDIT.

package openuem_ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/doncicuto/openuem_ent/agent"
	"github.com/doncicuto/openuem_ent/tag"
)

// TagCreate is the builder for creating a Tag entity.
type TagCreate struct {
	config
	mutation *TagMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetDescription sets the "description" field.
func (tc *TagCreate) SetDescription(s string) *TagCreate {
	tc.mutation.SetDescription(s)
	return tc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (tc *TagCreate) SetNillableDescription(s *string) *TagCreate {
	if s != nil {
		tc.SetDescription(*s)
	}
	return tc
}

// SetColor sets the "color" field.
func (tc *TagCreate) SetColor(s string) *TagCreate {
	tc.mutation.SetColor(s)
	return tc
}

// SetID sets the "id" field.
func (tc *TagCreate) SetID(s string) *TagCreate {
	tc.mutation.SetID(s)
	return tc
}

// AddOwnerIDs adds the "owner" edge to the Agent entity by IDs.
func (tc *TagCreate) AddOwnerIDs(ids ...string) *TagCreate {
	tc.mutation.AddOwnerIDs(ids...)
	return tc
}

// AddOwner adds the "owner" edges to the Agent entity.
func (tc *TagCreate) AddOwner(a ...*Agent) *TagCreate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return tc.AddOwnerIDs(ids...)
}

// SetParentID sets the "parent" edge to the Tag entity by ID.
func (tc *TagCreate) SetParentID(id string) *TagCreate {
	tc.mutation.SetParentID(id)
	return tc
}

// SetNillableParentID sets the "parent" edge to the Tag entity by ID if the given value is not nil.
func (tc *TagCreate) SetNillableParentID(id *string) *TagCreate {
	if id != nil {
		tc = tc.SetParentID(*id)
	}
	return tc
}

// SetParent sets the "parent" edge to the Tag entity.
func (tc *TagCreate) SetParent(t *Tag) *TagCreate {
	return tc.SetParentID(t.ID)
}

// AddChildIDs adds the "children" edge to the Tag entity by IDs.
func (tc *TagCreate) AddChildIDs(ids ...string) *TagCreate {
	tc.mutation.AddChildIDs(ids...)
	return tc
}

// AddChildren adds the "children" edges to the Tag entity.
func (tc *TagCreate) AddChildren(t ...*Tag) *TagCreate {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tc.AddChildIDs(ids...)
}

// Mutation returns the TagMutation object of the builder.
func (tc *TagCreate) Mutation() *TagMutation {
	return tc.mutation
}

// Save creates the Tag in the database.
func (tc *TagCreate) Save(ctx context.Context) (*Tag, error) {
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TagCreate) SaveX(ctx context.Context) *Tag {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TagCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TagCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TagCreate) check() error {
	if _, ok := tc.mutation.Color(); !ok {
		return &ValidationError{Name: "color", err: errors.New(`openuem_ent: missing required field "Tag.color"`)}
	}
	if v, ok := tc.mutation.ID(); ok {
		if err := tag.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`openuem_ent: validator failed for field "Tag.id": %w`, err)}
		}
	}
	return nil
}

func (tc *TagCreate) sqlSave(ctx context.Context) (*Tag, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Tag.ID type: %T", _spec.ID.Value)
		}
	}
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TagCreate) createSpec() (*Tag, *sqlgraph.CreateSpec) {
	var (
		_node = &Tag{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(tag.Table, sqlgraph.NewFieldSpec(tag.FieldID, field.TypeString))
	)
	_spec.OnConflict = tc.conflict
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tc.mutation.Description(); ok {
		_spec.SetField(tag.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := tc.mutation.Color(); ok {
		_spec.SetField(tag.FieldColor, field.TypeString, value)
		_node.Color = value
	}
	if nodes := tc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tag.OwnerTable,
			Columns: tag.OwnerPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(agent.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tag.ParentTable,
			Columns: []string{tag.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.tag_children = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tag.ChildrenTable,
			Columns: []string{tag.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeString),
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
//	client.Tag.Create().
//		SetDescription(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TagUpsert) {
//			SetDescription(v+v).
//		}).
//		Exec(ctx)
func (tc *TagCreate) OnConflict(opts ...sql.ConflictOption) *TagUpsertOne {
	tc.conflict = opts
	return &TagUpsertOne{
		create: tc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Tag.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tc *TagCreate) OnConflictColumns(columns ...string) *TagUpsertOne {
	tc.conflict = append(tc.conflict, sql.ConflictColumns(columns...))
	return &TagUpsertOne{
		create: tc,
	}
}

type (
	// TagUpsertOne is the builder for "upsert"-ing
	//  one Tag node.
	TagUpsertOne struct {
		create *TagCreate
	}

	// TagUpsert is the "OnConflict" setter.
	TagUpsert struct {
		*sql.UpdateSet
	}
)

// SetDescription sets the "description" field.
func (u *TagUpsert) SetDescription(v string) *TagUpsert {
	u.Set(tag.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *TagUpsert) UpdateDescription() *TagUpsert {
	u.SetExcluded(tag.FieldDescription)
	return u
}

// ClearDescription clears the value of the "description" field.
func (u *TagUpsert) ClearDescription() *TagUpsert {
	u.SetNull(tag.FieldDescription)
	return u
}

// SetColor sets the "color" field.
func (u *TagUpsert) SetColor(v string) *TagUpsert {
	u.Set(tag.FieldColor, v)
	return u
}

// UpdateColor sets the "color" field to the value that was provided on create.
func (u *TagUpsert) UpdateColor() *TagUpsert {
	u.SetExcluded(tag.FieldColor)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Tag.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(tag.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *TagUpsertOne) UpdateNewValues() *TagUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(tag.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Tag.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *TagUpsertOne) Ignore() *TagUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TagUpsertOne) DoNothing() *TagUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TagCreate.OnConflict
// documentation for more info.
func (u *TagUpsertOne) Update(set func(*TagUpsert)) *TagUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TagUpsert{UpdateSet: update})
	}))
	return u
}

// SetDescription sets the "description" field.
func (u *TagUpsertOne) SetDescription(v string) *TagUpsertOne {
	return u.Update(func(s *TagUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *TagUpsertOne) UpdateDescription() *TagUpsertOne {
	return u.Update(func(s *TagUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *TagUpsertOne) ClearDescription() *TagUpsertOne {
	return u.Update(func(s *TagUpsert) {
		s.ClearDescription()
	})
}

// SetColor sets the "color" field.
func (u *TagUpsertOne) SetColor(v string) *TagUpsertOne {
	return u.Update(func(s *TagUpsert) {
		s.SetColor(v)
	})
}

// UpdateColor sets the "color" field to the value that was provided on create.
func (u *TagUpsertOne) UpdateColor() *TagUpsertOne {
	return u.Update(func(s *TagUpsert) {
		s.UpdateColor()
	})
}

// Exec executes the query.
func (u *TagUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("openuem_ent: missing options for TagCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TagUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *TagUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("openuem_ent: TagUpsertOne.ID is not supported by MySQL driver. Use TagUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *TagUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// TagCreateBulk is the builder for creating many Tag entities in bulk.
type TagCreateBulk struct {
	config
	err      error
	builders []*TagCreate
	conflict []sql.ConflictOption
}

// Save creates the Tag entities in the database.
func (tcb *TagCreateBulk) Save(ctx context.Context) ([]*Tag, error) {
	if tcb.err != nil {
		return nil, tcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Tag, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TagMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = tcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TagCreateBulk) SaveX(ctx context.Context) []*Tag {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TagCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TagCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Tag.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TagUpsert) {
//			SetDescription(v+v).
//		}).
//		Exec(ctx)
func (tcb *TagCreateBulk) OnConflict(opts ...sql.ConflictOption) *TagUpsertBulk {
	tcb.conflict = opts
	return &TagUpsertBulk{
		create: tcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Tag.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tcb *TagCreateBulk) OnConflictColumns(columns ...string) *TagUpsertBulk {
	tcb.conflict = append(tcb.conflict, sql.ConflictColumns(columns...))
	return &TagUpsertBulk{
		create: tcb,
	}
}

// TagUpsertBulk is the builder for "upsert"-ing
// a bulk of Tag nodes.
type TagUpsertBulk struct {
	create *TagCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Tag.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(tag.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *TagUpsertBulk) UpdateNewValues() *TagUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(tag.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Tag.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *TagUpsertBulk) Ignore() *TagUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TagUpsertBulk) DoNothing() *TagUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TagCreateBulk.OnConflict
// documentation for more info.
func (u *TagUpsertBulk) Update(set func(*TagUpsert)) *TagUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TagUpsert{UpdateSet: update})
	}))
	return u
}

// SetDescription sets the "description" field.
func (u *TagUpsertBulk) SetDescription(v string) *TagUpsertBulk {
	return u.Update(func(s *TagUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *TagUpsertBulk) UpdateDescription() *TagUpsertBulk {
	return u.Update(func(s *TagUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *TagUpsertBulk) ClearDescription() *TagUpsertBulk {
	return u.Update(func(s *TagUpsert) {
		s.ClearDescription()
	})
}

// SetColor sets the "color" field.
func (u *TagUpsertBulk) SetColor(v string) *TagUpsertBulk {
	return u.Update(func(s *TagUpsert) {
		s.SetColor(v)
	})
}

// UpdateColor sets the "color" field to the value that was provided on create.
func (u *TagUpsertBulk) UpdateColor() *TagUpsertBulk {
	return u.Update(func(s *TagUpsert) {
		s.UpdateColor()
	})
}

// Exec executes the query.
func (u *TagUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("openuem_ent: OnConflict was set for builder %d. Set it on the TagCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("openuem_ent: missing options for TagCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TagUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
