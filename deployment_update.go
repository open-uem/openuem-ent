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
	"github.com/open-uem/openuem_ent/predicate"
)

// DeploymentUpdate is the builder for updating Deployment entities.
type DeploymentUpdate struct {
	config
	hooks     []Hook
	mutation  *DeploymentMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the DeploymentUpdate builder.
func (du *DeploymentUpdate) Where(ps ...predicate.Deployment) *DeploymentUpdate {
	du.mutation.Where(ps...)
	return du
}

// SetPackageID sets the "package_id" field.
func (du *DeploymentUpdate) SetPackageID(s string) *DeploymentUpdate {
	du.mutation.SetPackageID(s)
	return du
}

// SetNillablePackageID sets the "package_id" field if the given value is not nil.
func (du *DeploymentUpdate) SetNillablePackageID(s *string) *DeploymentUpdate {
	if s != nil {
		du.SetPackageID(*s)
	}
	return du
}

// SetName sets the "name" field.
func (du *DeploymentUpdate) SetName(s string) *DeploymentUpdate {
	du.mutation.SetName(s)
	return du
}

// SetNillableName sets the "name" field if the given value is not nil.
func (du *DeploymentUpdate) SetNillableName(s *string) *DeploymentUpdate {
	if s != nil {
		du.SetName(*s)
	}
	return du
}

// SetVersion sets the "version" field.
func (du *DeploymentUpdate) SetVersion(s string) *DeploymentUpdate {
	du.mutation.SetVersion(s)
	return du
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (du *DeploymentUpdate) SetNillableVersion(s *string) *DeploymentUpdate {
	if s != nil {
		du.SetVersion(*s)
	}
	return du
}

// ClearVersion clears the value of the "version" field.
func (du *DeploymentUpdate) ClearVersion() *DeploymentUpdate {
	du.mutation.ClearVersion()
	return du
}

// SetInstalled sets the "installed" field.
func (du *DeploymentUpdate) SetInstalled(t time.Time) *DeploymentUpdate {
	du.mutation.SetInstalled(t)
	return du
}

// SetNillableInstalled sets the "installed" field if the given value is not nil.
func (du *DeploymentUpdate) SetNillableInstalled(t *time.Time) *DeploymentUpdate {
	if t != nil {
		du.SetInstalled(*t)
	}
	return du
}

// ClearInstalled clears the value of the "installed" field.
func (du *DeploymentUpdate) ClearInstalled() *DeploymentUpdate {
	du.mutation.ClearInstalled()
	return du
}

// SetUpdated sets the "updated" field.
func (du *DeploymentUpdate) SetUpdated(t time.Time) *DeploymentUpdate {
	du.mutation.SetUpdated(t)
	return du
}

// ClearUpdated clears the value of the "updated" field.
func (du *DeploymentUpdate) ClearUpdated() *DeploymentUpdate {
	du.mutation.ClearUpdated()
	return du
}

// SetOwnerID sets the "owner" edge to the Agent entity by ID.
func (du *DeploymentUpdate) SetOwnerID(id string) *DeploymentUpdate {
	du.mutation.SetOwnerID(id)
	return du
}

// SetOwner sets the "owner" edge to the Agent entity.
func (du *DeploymentUpdate) SetOwner(a *Agent) *DeploymentUpdate {
	return du.SetOwnerID(a.ID)
}

// Mutation returns the DeploymentMutation object of the builder.
func (du *DeploymentUpdate) Mutation() *DeploymentMutation {
	return du.mutation
}

// ClearOwner clears the "owner" edge to the Agent entity.
func (du *DeploymentUpdate) ClearOwner() *DeploymentUpdate {
	du.mutation.ClearOwner()
	return du
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DeploymentUpdate) Save(ctx context.Context) (int, error) {
	du.defaults()
	return withHooks(ctx, du.sqlSave, du.mutation, du.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (du *DeploymentUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DeploymentUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DeploymentUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (du *DeploymentUpdate) defaults() {
	if _, ok := du.mutation.Updated(); !ok && !du.mutation.UpdatedCleared() {
		v := deployment.UpdateDefaultUpdated()
		du.mutation.SetUpdated(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (du *DeploymentUpdate) check() error {
	if du.mutation.OwnerCleared() && len(du.mutation.OwnerIDs()) > 0 {
		return errors.New(`openuem_ent: clearing a required unique edge "Deployment.owner"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (du *DeploymentUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *DeploymentUpdate {
	du.modifiers = append(du.modifiers, modifiers...)
	return du
}

func (du *DeploymentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := du.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(deployment.Table, deployment.Columns, sqlgraph.NewFieldSpec(deployment.FieldID, field.TypeInt))
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.PackageID(); ok {
		_spec.SetField(deployment.FieldPackageID, field.TypeString, value)
	}
	if value, ok := du.mutation.Name(); ok {
		_spec.SetField(deployment.FieldName, field.TypeString, value)
	}
	if value, ok := du.mutation.Version(); ok {
		_spec.SetField(deployment.FieldVersion, field.TypeString, value)
	}
	if du.mutation.VersionCleared() {
		_spec.ClearField(deployment.FieldVersion, field.TypeString)
	}
	if value, ok := du.mutation.Installed(); ok {
		_spec.SetField(deployment.FieldInstalled, field.TypeTime, value)
	}
	if du.mutation.InstalledCleared() {
		_spec.ClearField(deployment.FieldInstalled, field.TypeTime)
	}
	if value, ok := du.mutation.Updated(); ok {
		_spec.SetField(deployment.FieldUpdated, field.TypeTime, value)
	}
	if du.mutation.UpdatedCleared() {
		_spec.ClearField(deployment.FieldUpdated, field.TypeTime)
	}
	if du.mutation.OwnerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.OwnerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(du.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{deployment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	du.mutation.done = true
	return n, nil
}

// DeploymentUpdateOne is the builder for updating a single Deployment entity.
type DeploymentUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *DeploymentMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetPackageID sets the "package_id" field.
func (duo *DeploymentUpdateOne) SetPackageID(s string) *DeploymentUpdateOne {
	duo.mutation.SetPackageID(s)
	return duo
}

// SetNillablePackageID sets the "package_id" field if the given value is not nil.
func (duo *DeploymentUpdateOne) SetNillablePackageID(s *string) *DeploymentUpdateOne {
	if s != nil {
		duo.SetPackageID(*s)
	}
	return duo
}

// SetName sets the "name" field.
func (duo *DeploymentUpdateOne) SetName(s string) *DeploymentUpdateOne {
	duo.mutation.SetName(s)
	return duo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (duo *DeploymentUpdateOne) SetNillableName(s *string) *DeploymentUpdateOne {
	if s != nil {
		duo.SetName(*s)
	}
	return duo
}

// SetVersion sets the "version" field.
func (duo *DeploymentUpdateOne) SetVersion(s string) *DeploymentUpdateOne {
	duo.mutation.SetVersion(s)
	return duo
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (duo *DeploymentUpdateOne) SetNillableVersion(s *string) *DeploymentUpdateOne {
	if s != nil {
		duo.SetVersion(*s)
	}
	return duo
}

// ClearVersion clears the value of the "version" field.
func (duo *DeploymentUpdateOne) ClearVersion() *DeploymentUpdateOne {
	duo.mutation.ClearVersion()
	return duo
}

// SetInstalled sets the "installed" field.
func (duo *DeploymentUpdateOne) SetInstalled(t time.Time) *DeploymentUpdateOne {
	duo.mutation.SetInstalled(t)
	return duo
}

// SetNillableInstalled sets the "installed" field if the given value is not nil.
func (duo *DeploymentUpdateOne) SetNillableInstalled(t *time.Time) *DeploymentUpdateOne {
	if t != nil {
		duo.SetInstalled(*t)
	}
	return duo
}

// ClearInstalled clears the value of the "installed" field.
func (duo *DeploymentUpdateOne) ClearInstalled() *DeploymentUpdateOne {
	duo.mutation.ClearInstalled()
	return duo
}

// SetUpdated sets the "updated" field.
func (duo *DeploymentUpdateOne) SetUpdated(t time.Time) *DeploymentUpdateOne {
	duo.mutation.SetUpdated(t)
	return duo
}

// ClearUpdated clears the value of the "updated" field.
func (duo *DeploymentUpdateOne) ClearUpdated() *DeploymentUpdateOne {
	duo.mutation.ClearUpdated()
	return duo
}

// SetOwnerID sets the "owner" edge to the Agent entity by ID.
func (duo *DeploymentUpdateOne) SetOwnerID(id string) *DeploymentUpdateOne {
	duo.mutation.SetOwnerID(id)
	return duo
}

// SetOwner sets the "owner" edge to the Agent entity.
func (duo *DeploymentUpdateOne) SetOwner(a *Agent) *DeploymentUpdateOne {
	return duo.SetOwnerID(a.ID)
}

// Mutation returns the DeploymentMutation object of the builder.
func (duo *DeploymentUpdateOne) Mutation() *DeploymentMutation {
	return duo.mutation
}

// ClearOwner clears the "owner" edge to the Agent entity.
func (duo *DeploymentUpdateOne) ClearOwner() *DeploymentUpdateOne {
	duo.mutation.ClearOwner()
	return duo
}

// Where appends a list predicates to the DeploymentUpdate builder.
func (duo *DeploymentUpdateOne) Where(ps ...predicate.Deployment) *DeploymentUpdateOne {
	duo.mutation.Where(ps...)
	return duo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (duo *DeploymentUpdateOne) Select(field string, fields ...string) *DeploymentUpdateOne {
	duo.fields = append([]string{field}, fields...)
	return duo
}

// Save executes the query and returns the updated Deployment entity.
func (duo *DeploymentUpdateOne) Save(ctx context.Context) (*Deployment, error) {
	duo.defaults()
	return withHooks(ctx, duo.sqlSave, duo.mutation, duo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DeploymentUpdateOne) SaveX(ctx context.Context) *Deployment {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DeploymentUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DeploymentUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (duo *DeploymentUpdateOne) defaults() {
	if _, ok := duo.mutation.Updated(); !ok && !duo.mutation.UpdatedCleared() {
		v := deployment.UpdateDefaultUpdated()
		duo.mutation.SetUpdated(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (duo *DeploymentUpdateOne) check() error {
	if duo.mutation.OwnerCleared() && len(duo.mutation.OwnerIDs()) > 0 {
		return errors.New(`openuem_ent: clearing a required unique edge "Deployment.owner"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (duo *DeploymentUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *DeploymentUpdateOne {
	duo.modifiers = append(duo.modifiers, modifiers...)
	return duo
}

func (duo *DeploymentUpdateOne) sqlSave(ctx context.Context) (_node *Deployment, err error) {
	if err := duo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(deployment.Table, deployment.Columns, sqlgraph.NewFieldSpec(deployment.FieldID, field.TypeInt))
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`openuem_ent: missing "Deployment.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := duo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, deployment.FieldID)
		for _, f := range fields {
			if !deployment.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("openuem_ent: invalid field %q for query", f)}
			}
			if f != deployment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := duo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := duo.mutation.PackageID(); ok {
		_spec.SetField(deployment.FieldPackageID, field.TypeString, value)
	}
	if value, ok := duo.mutation.Name(); ok {
		_spec.SetField(deployment.FieldName, field.TypeString, value)
	}
	if value, ok := duo.mutation.Version(); ok {
		_spec.SetField(deployment.FieldVersion, field.TypeString, value)
	}
	if duo.mutation.VersionCleared() {
		_spec.ClearField(deployment.FieldVersion, field.TypeString)
	}
	if value, ok := duo.mutation.Installed(); ok {
		_spec.SetField(deployment.FieldInstalled, field.TypeTime, value)
	}
	if duo.mutation.InstalledCleared() {
		_spec.ClearField(deployment.FieldInstalled, field.TypeTime)
	}
	if value, ok := duo.mutation.Updated(); ok {
		_spec.SetField(deployment.FieldUpdated, field.TypeTime, value)
	}
	if duo.mutation.UpdatedCleared() {
		_spec.ClearField(deployment.FieldUpdated, field.TypeTime)
	}
	if duo.mutation.OwnerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.OwnerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(duo.modifiers...)
	_node = &Deployment{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{deployment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	duo.mutation.done = true
	return _node, nil
}
