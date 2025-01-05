// Code generated by ent, DO NOT EDIT.

package openuem_ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/open-uem/openuem_ent/agent"
	"github.com/open-uem/openuem_ent/deployment"
)

// DeploymentCreate is the builder for creating a Deployment entity.
type DeploymentCreate struct {
	config
	mutation *DeploymentMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetPackageID sets the "package_id" field.
func (dc *DeploymentCreate) SetPackageID(s string) *DeploymentCreate {
	dc.mutation.SetPackageID(s)
	return dc
}

// SetName sets the "name" field.
func (dc *DeploymentCreate) SetName(s string) *DeploymentCreate {
	dc.mutation.SetName(s)
	return dc
}

// SetVersion sets the "version" field.
func (dc *DeploymentCreate) SetVersion(s string) *DeploymentCreate {
	dc.mutation.SetVersion(s)
	return dc
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (dc *DeploymentCreate) SetNillableVersion(s *string) *DeploymentCreate {
	if s != nil {
		dc.SetVersion(*s)
	}
	return dc
}

// SetInstalled sets the "installed" field.
func (dc *DeploymentCreate) SetInstalled(t time.Time) *DeploymentCreate {
	dc.mutation.SetInstalled(t)
	return dc
}

// SetNillableInstalled sets the "installed" field if the given value is not nil.
func (dc *DeploymentCreate) SetNillableInstalled(t *time.Time) *DeploymentCreate {
	if t != nil {
		dc.SetInstalled(*t)
	}
	return dc
}

// SetUpdated sets the "updated" field.
func (dc *DeploymentCreate) SetUpdated(t time.Time) *DeploymentCreate {
	dc.mutation.SetUpdated(t)
	return dc
}

// SetNillableUpdated sets the "updated" field if the given value is not nil.
func (dc *DeploymentCreate) SetNillableUpdated(t *time.Time) *DeploymentCreate {
	if t != nil {
		dc.SetUpdated(*t)
	}
	return dc
}

// SetOwnerID sets the "owner" edge to the Agent entity by ID.
func (dc *DeploymentCreate) SetOwnerID(id string) *DeploymentCreate {
	dc.mutation.SetOwnerID(id)
	return dc
}

// SetOwner sets the "owner" edge to the Agent entity.
func (dc *DeploymentCreate) SetOwner(a *Agent) *DeploymentCreate {
	return dc.SetOwnerID(a.ID)
}

// Mutation returns the DeploymentMutation object of the builder.
func (dc *DeploymentCreate) Mutation() *DeploymentMutation {
	return dc.mutation
}

// Save creates the Deployment in the database.
func (dc *DeploymentCreate) Save(ctx context.Context) (*Deployment, error) {
	dc.defaults()
	return withHooks(ctx, dc.sqlSave, dc.mutation, dc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DeploymentCreate) SaveX(ctx context.Context) *Deployment {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DeploymentCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DeploymentCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dc *DeploymentCreate) defaults() {
	if _, ok := dc.mutation.Installed(); !ok {
		v := deployment.DefaultInstalled()
		dc.mutation.SetInstalled(v)
	}
	if _, ok := dc.mutation.Updated(); !ok {
		v := deployment.DefaultUpdated()
		dc.mutation.SetUpdated(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DeploymentCreate) check() error {
	if _, ok := dc.mutation.PackageID(); !ok {
		return &ValidationError{Name: "package_id", err: errors.New(`openuem_ent: missing required field "Deployment.package_id"`)}
	}
	if _, ok := dc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`openuem_ent: missing required field "Deployment.name"`)}
	}
	if len(dc.mutation.OwnerIDs()) == 0 {
		return &ValidationError{Name: "owner", err: errors.New(`openuem_ent: missing required edge "Deployment.owner"`)}
	}
	return nil
}

func (dc *DeploymentCreate) sqlSave(ctx context.Context) (*Deployment, error) {
	if err := dc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	dc.mutation.id = &_node.ID
	dc.mutation.done = true
	return _node, nil
}

func (dc *DeploymentCreate) createSpec() (*Deployment, *sqlgraph.CreateSpec) {
	var (
		_node = &Deployment{config: dc.config}
		_spec = sqlgraph.NewCreateSpec(deployment.Table, sqlgraph.NewFieldSpec(deployment.FieldID, field.TypeInt))
	)
	_spec.OnConflict = dc.conflict
	if value, ok := dc.mutation.PackageID(); ok {
		_spec.SetField(deployment.FieldPackageID, field.TypeString, value)
		_node.PackageID = value
	}
	if value, ok := dc.mutation.Name(); ok {
		_spec.SetField(deployment.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := dc.mutation.Version(); ok {
		_spec.SetField(deployment.FieldVersion, field.TypeString, value)
		_node.Version = value
	}
	if value, ok := dc.mutation.Installed(); ok {
		_spec.SetField(deployment.FieldInstalled, field.TypeTime, value)
		_node.Installed = value
	}
	if value, ok := dc.mutation.Updated(); ok {
		_spec.SetField(deployment.FieldUpdated, field.TypeTime, value)
		_node.Updated = value
	}
	if nodes := dc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deployment.OwnerTable,
			Columns: []string{deployment.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(agent.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.agent_deployments = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Deployment.Create().
//		SetPackageID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DeploymentUpsert) {
//			SetPackageID(v+v).
//		}).
//		Exec(ctx)
func (dc *DeploymentCreate) OnConflict(opts ...sql.ConflictOption) *DeploymentUpsertOne {
	dc.conflict = opts
	return &DeploymentUpsertOne{
		create: dc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Deployment.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dc *DeploymentCreate) OnConflictColumns(columns ...string) *DeploymentUpsertOne {
	dc.conflict = append(dc.conflict, sql.ConflictColumns(columns...))
	return &DeploymentUpsertOne{
		create: dc,
	}
}

type (
	// DeploymentUpsertOne is the builder for "upsert"-ing
	//  one Deployment node.
	DeploymentUpsertOne struct {
		create *DeploymentCreate
	}

	// DeploymentUpsert is the "OnConflict" setter.
	DeploymentUpsert struct {
		*sql.UpdateSet
	}
)

// SetPackageID sets the "package_id" field.
func (u *DeploymentUpsert) SetPackageID(v string) *DeploymentUpsert {
	u.Set(deployment.FieldPackageID, v)
	return u
}

// UpdatePackageID sets the "package_id" field to the value that was provided on create.
func (u *DeploymentUpsert) UpdatePackageID() *DeploymentUpsert {
	u.SetExcluded(deployment.FieldPackageID)
	return u
}

// SetName sets the "name" field.
func (u *DeploymentUpsert) SetName(v string) *DeploymentUpsert {
	u.Set(deployment.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *DeploymentUpsert) UpdateName() *DeploymentUpsert {
	u.SetExcluded(deployment.FieldName)
	return u
}

// SetVersion sets the "version" field.
func (u *DeploymentUpsert) SetVersion(v string) *DeploymentUpsert {
	u.Set(deployment.FieldVersion, v)
	return u
}

// UpdateVersion sets the "version" field to the value that was provided on create.
func (u *DeploymentUpsert) UpdateVersion() *DeploymentUpsert {
	u.SetExcluded(deployment.FieldVersion)
	return u
}

// ClearVersion clears the value of the "version" field.
func (u *DeploymentUpsert) ClearVersion() *DeploymentUpsert {
	u.SetNull(deployment.FieldVersion)
	return u
}

// SetInstalled sets the "installed" field.
func (u *DeploymentUpsert) SetInstalled(v time.Time) *DeploymentUpsert {
	u.Set(deployment.FieldInstalled, v)
	return u
}

// UpdateInstalled sets the "installed" field to the value that was provided on create.
func (u *DeploymentUpsert) UpdateInstalled() *DeploymentUpsert {
	u.SetExcluded(deployment.FieldInstalled)
	return u
}

// ClearInstalled clears the value of the "installed" field.
func (u *DeploymentUpsert) ClearInstalled() *DeploymentUpsert {
	u.SetNull(deployment.FieldInstalled)
	return u
}

// SetUpdated sets the "updated" field.
func (u *DeploymentUpsert) SetUpdated(v time.Time) *DeploymentUpsert {
	u.Set(deployment.FieldUpdated, v)
	return u
}

// UpdateUpdated sets the "updated" field to the value that was provided on create.
func (u *DeploymentUpsert) UpdateUpdated() *DeploymentUpsert {
	u.SetExcluded(deployment.FieldUpdated)
	return u
}

// ClearUpdated clears the value of the "updated" field.
func (u *DeploymentUpsert) ClearUpdated() *DeploymentUpsert {
	u.SetNull(deployment.FieldUpdated)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Deployment.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *DeploymentUpsertOne) UpdateNewValues() *DeploymentUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Deployment.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *DeploymentUpsertOne) Ignore() *DeploymentUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DeploymentUpsertOne) DoNothing() *DeploymentUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DeploymentCreate.OnConflict
// documentation for more info.
func (u *DeploymentUpsertOne) Update(set func(*DeploymentUpsert)) *DeploymentUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DeploymentUpsert{UpdateSet: update})
	}))
	return u
}

// SetPackageID sets the "package_id" field.
func (u *DeploymentUpsertOne) SetPackageID(v string) *DeploymentUpsertOne {
	return u.Update(func(s *DeploymentUpsert) {
		s.SetPackageID(v)
	})
}

// UpdatePackageID sets the "package_id" field to the value that was provided on create.
func (u *DeploymentUpsertOne) UpdatePackageID() *DeploymentUpsertOne {
	return u.Update(func(s *DeploymentUpsert) {
		s.UpdatePackageID()
	})
}

// SetName sets the "name" field.
func (u *DeploymentUpsertOne) SetName(v string) *DeploymentUpsertOne {
	return u.Update(func(s *DeploymentUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *DeploymentUpsertOne) UpdateName() *DeploymentUpsertOne {
	return u.Update(func(s *DeploymentUpsert) {
		s.UpdateName()
	})
}

// SetVersion sets the "version" field.
func (u *DeploymentUpsertOne) SetVersion(v string) *DeploymentUpsertOne {
	return u.Update(func(s *DeploymentUpsert) {
		s.SetVersion(v)
	})
}

// UpdateVersion sets the "version" field to the value that was provided on create.
func (u *DeploymentUpsertOne) UpdateVersion() *DeploymentUpsertOne {
	return u.Update(func(s *DeploymentUpsert) {
		s.UpdateVersion()
	})
}

// ClearVersion clears the value of the "version" field.
func (u *DeploymentUpsertOne) ClearVersion() *DeploymentUpsertOne {
	return u.Update(func(s *DeploymentUpsert) {
		s.ClearVersion()
	})
}

// SetInstalled sets the "installed" field.
func (u *DeploymentUpsertOne) SetInstalled(v time.Time) *DeploymentUpsertOne {
	return u.Update(func(s *DeploymentUpsert) {
		s.SetInstalled(v)
	})
}

// UpdateInstalled sets the "installed" field to the value that was provided on create.
func (u *DeploymentUpsertOne) UpdateInstalled() *DeploymentUpsertOne {
	return u.Update(func(s *DeploymentUpsert) {
		s.UpdateInstalled()
	})
}

// ClearInstalled clears the value of the "installed" field.
func (u *DeploymentUpsertOne) ClearInstalled() *DeploymentUpsertOne {
	return u.Update(func(s *DeploymentUpsert) {
		s.ClearInstalled()
	})
}

// SetUpdated sets the "updated" field.
func (u *DeploymentUpsertOne) SetUpdated(v time.Time) *DeploymentUpsertOne {
	return u.Update(func(s *DeploymentUpsert) {
		s.SetUpdated(v)
	})
}

// UpdateUpdated sets the "updated" field to the value that was provided on create.
func (u *DeploymentUpsertOne) UpdateUpdated() *DeploymentUpsertOne {
	return u.Update(func(s *DeploymentUpsert) {
		s.UpdateUpdated()
	})
}

// ClearUpdated clears the value of the "updated" field.
func (u *DeploymentUpsertOne) ClearUpdated() *DeploymentUpsertOne {
	return u.Update(func(s *DeploymentUpsert) {
		s.ClearUpdated()
	})
}

// Exec executes the query.
func (u *DeploymentUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("openuem_ent: missing options for DeploymentCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DeploymentUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *DeploymentUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *DeploymentUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// DeploymentCreateBulk is the builder for creating many Deployment entities in bulk.
type DeploymentCreateBulk struct {
	config
	err      error
	builders []*DeploymentCreate
	conflict []sql.ConflictOption
}

// Save creates the Deployment entities in the database.
func (dcb *DeploymentCreateBulk) Save(ctx context.Context) ([]*Deployment, error) {
	if dcb.err != nil {
		return nil, dcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Deployment, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DeploymentMutation)
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
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = dcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DeploymentCreateBulk) SaveX(ctx context.Context) []*Deployment {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DeploymentCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DeploymentCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Deployment.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DeploymentUpsert) {
//			SetPackageID(v+v).
//		}).
//		Exec(ctx)
func (dcb *DeploymentCreateBulk) OnConflict(opts ...sql.ConflictOption) *DeploymentUpsertBulk {
	dcb.conflict = opts
	return &DeploymentUpsertBulk{
		create: dcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Deployment.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dcb *DeploymentCreateBulk) OnConflictColumns(columns ...string) *DeploymentUpsertBulk {
	dcb.conflict = append(dcb.conflict, sql.ConflictColumns(columns...))
	return &DeploymentUpsertBulk{
		create: dcb,
	}
}

// DeploymentUpsertBulk is the builder for "upsert"-ing
// a bulk of Deployment nodes.
type DeploymentUpsertBulk struct {
	create *DeploymentCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Deployment.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *DeploymentUpsertBulk) UpdateNewValues() *DeploymentUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Deployment.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *DeploymentUpsertBulk) Ignore() *DeploymentUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DeploymentUpsertBulk) DoNothing() *DeploymentUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DeploymentCreateBulk.OnConflict
// documentation for more info.
func (u *DeploymentUpsertBulk) Update(set func(*DeploymentUpsert)) *DeploymentUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DeploymentUpsert{UpdateSet: update})
	}))
	return u
}

// SetPackageID sets the "package_id" field.
func (u *DeploymentUpsertBulk) SetPackageID(v string) *DeploymentUpsertBulk {
	return u.Update(func(s *DeploymentUpsert) {
		s.SetPackageID(v)
	})
}

// UpdatePackageID sets the "package_id" field to the value that was provided on create.
func (u *DeploymentUpsertBulk) UpdatePackageID() *DeploymentUpsertBulk {
	return u.Update(func(s *DeploymentUpsert) {
		s.UpdatePackageID()
	})
}

// SetName sets the "name" field.
func (u *DeploymentUpsertBulk) SetName(v string) *DeploymentUpsertBulk {
	return u.Update(func(s *DeploymentUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *DeploymentUpsertBulk) UpdateName() *DeploymentUpsertBulk {
	return u.Update(func(s *DeploymentUpsert) {
		s.UpdateName()
	})
}

// SetVersion sets the "version" field.
func (u *DeploymentUpsertBulk) SetVersion(v string) *DeploymentUpsertBulk {
	return u.Update(func(s *DeploymentUpsert) {
		s.SetVersion(v)
	})
}

// UpdateVersion sets the "version" field to the value that was provided on create.
func (u *DeploymentUpsertBulk) UpdateVersion() *DeploymentUpsertBulk {
	return u.Update(func(s *DeploymentUpsert) {
		s.UpdateVersion()
	})
}

// ClearVersion clears the value of the "version" field.
func (u *DeploymentUpsertBulk) ClearVersion() *DeploymentUpsertBulk {
	return u.Update(func(s *DeploymentUpsert) {
		s.ClearVersion()
	})
}

// SetInstalled sets the "installed" field.
func (u *DeploymentUpsertBulk) SetInstalled(v time.Time) *DeploymentUpsertBulk {
	return u.Update(func(s *DeploymentUpsert) {
		s.SetInstalled(v)
	})
}

// UpdateInstalled sets the "installed" field to the value that was provided on create.
func (u *DeploymentUpsertBulk) UpdateInstalled() *DeploymentUpsertBulk {
	return u.Update(func(s *DeploymentUpsert) {
		s.UpdateInstalled()
	})
}

// ClearInstalled clears the value of the "installed" field.
func (u *DeploymentUpsertBulk) ClearInstalled() *DeploymentUpsertBulk {
	return u.Update(func(s *DeploymentUpsert) {
		s.ClearInstalled()
	})
}

// SetUpdated sets the "updated" field.
func (u *DeploymentUpsertBulk) SetUpdated(v time.Time) *DeploymentUpsertBulk {
	return u.Update(func(s *DeploymentUpsert) {
		s.SetUpdated(v)
	})
}

// UpdateUpdated sets the "updated" field to the value that was provided on create.
func (u *DeploymentUpsertBulk) UpdateUpdated() *DeploymentUpsertBulk {
	return u.Update(func(s *DeploymentUpsert) {
		s.UpdateUpdated()
	})
}

// ClearUpdated clears the value of the "updated" field.
func (u *DeploymentUpsertBulk) ClearUpdated() *DeploymentUpsertBulk {
	return u.Update(func(s *DeploymentUpsert) {
		s.ClearUpdated()
	})
}

// Exec executes the query.
func (u *DeploymentUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("openuem_ent: OnConflict was set for builder %d. Set it on the DeploymentCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("openuem_ent: missing options for DeploymentCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DeploymentUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
