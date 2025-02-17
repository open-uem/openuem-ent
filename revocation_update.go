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
	"github.com/open-uem/ent/predicate"
	"github.com/open-uem/ent/revocation"
)

// RevocationUpdate is the builder for updating Revocation entities.
type RevocationUpdate struct {
	config
	hooks     []Hook
	mutation  *RevocationMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the RevocationUpdate builder.
func (ru *RevocationUpdate) Where(ps ...predicate.Revocation) *RevocationUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetReason sets the "reason" field.
func (ru *RevocationUpdate) SetReason(i int) *RevocationUpdate {
	ru.mutation.ResetReason()
	ru.mutation.SetReason(i)
	return ru
}

// SetNillableReason sets the "reason" field if the given value is not nil.
func (ru *RevocationUpdate) SetNillableReason(i *int) *RevocationUpdate {
	if i != nil {
		ru.SetReason(*i)
	}
	return ru
}

// AddReason adds i to the "reason" field.
func (ru *RevocationUpdate) AddReason(i int) *RevocationUpdate {
	ru.mutation.AddReason(i)
	return ru
}

// ClearReason clears the value of the "reason" field.
func (ru *RevocationUpdate) ClearReason() *RevocationUpdate {
	ru.mutation.ClearReason()
	return ru
}

// SetInfo sets the "info" field.
func (ru *RevocationUpdate) SetInfo(s string) *RevocationUpdate {
	ru.mutation.SetInfo(s)
	return ru
}

// SetNillableInfo sets the "info" field if the given value is not nil.
func (ru *RevocationUpdate) SetNillableInfo(s *string) *RevocationUpdate {
	if s != nil {
		ru.SetInfo(*s)
	}
	return ru
}

// ClearInfo clears the value of the "info" field.
func (ru *RevocationUpdate) ClearInfo() *RevocationUpdate {
	ru.mutation.ClearInfo()
	return ru
}

// SetExpiry sets the "expiry" field.
func (ru *RevocationUpdate) SetExpiry(t time.Time) *RevocationUpdate {
	ru.mutation.SetExpiry(t)
	return ru
}

// SetNillableExpiry sets the "expiry" field if the given value is not nil.
func (ru *RevocationUpdate) SetNillableExpiry(t *time.Time) *RevocationUpdate {
	if t != nil {
		ru.SetExpiry(*t)
	}
	return ru
}

// ClearExpiry clears the value of the "expiry" field.
func (ru *RevocationUpdate) ClearExpiry() *RevocationUpdate {
	ru.mutation.ClearExpiry()
	return ru
}

// SetRevoked sets the "revoked" field.
func (ru *RevocationUpdate) SetRevoked(t time.Time) *RevocationUpdate {
	ru.mutation.SetRevoked(t)
	return ru
}

// SetNillableRevoked sets the "revoked" field if the given value is not nil.
func (ru *RevocationUpdate) SetNillableRevoked(t *time.Time) *RevocationUpdate {
	if t != nil {
		ru.SetRevoked(*t)
	}
	return ru
}

// Mutation returns the RevocationMutation object of the builder.
func (ru *RevocationUpdate) Mutation() *RevocationMutation {
	return ru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RevocationUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ru.sqlSave, ru.mutation, ru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RevocationUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RevocationUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RevocationUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ru *RevocationUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *RevocationUpdate {
	ru.modifiers = append(ru.modifiers, modifiers...)
	return ru
}

func (ru *RevocationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(revocation.Table, revocation.Columns, sqlgraph.NewFieldSpec(revocation.FieldID, field.TypeInt64))
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.Reason(); ok {
		_spec.SetField(revocation.FieldReason, field.TypeInt, value)
	}
	if value, ok := ru.mutation.AddedReason(); ok {
		_spec.AddField(revocation.FieldReason, field.TypeInt, value)
	}
	if ru.mutation.ReasonCleared() {
		_spec.ClearField(revocation.FieldReason, field.TypeInt)
	}
	if value, ok := ru.mutation.Info(); ok {
		_spec.SetField(revocation.FieldInfo, field.TypeString, value)
	}
	if ru.mutation.InfoCleared() {
		_spec.ClearField(revocation.FieldInfo, field.TypeString)
	}
	if value, ok := ru.mutation.Expiry(); ok {
		_spec.SetField(revocation.FieldExpiry, field.TypeTime, value)
	}
	if ru.mutation.ExpiryCleared() {
		_spec.ClearField(revocation.FieldExpiry, field.TypeTime)
	}
	if value, ok := ru.mutation.Revoked(); ok {
		_spec.SetField(revocation.FieldRevoked, field.TypeTime, value)
	}
	_spec.AddModifiers(ru.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{revocation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ru.mutation.done = true
	return n, nil
}

// RevocationUpdateOne is the builder for updating a single Revocation entity.
type RevocationUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *RevocationMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetReason sets the "reason" field.
func (ruo *RevocationUpdateOne) SetReason(i int) *RevocationUpdateOne {
	ruo.mutation.ResetReason()
	ruo.mutation.SetReason(i)
	return ruo
}

// SetNillableReason sets the "reason" field if the given value is not nil.
func (ruo *RevocationUpdateOne) SetNillableReason(i *int) *RevocationUpdateOne {
	if i != nil {
		ruo.SetReason(*i)
	}
	return ruo
}

// AddReason adds i to the "reason" field.
func (ruo *RevocationUpdateOne) AddReason(i int) *RevocationUpdateOne {
	ruo.mutation.AddReason(i)
	return ruo
}

// ClearReason clears the value of the "reason" field.
func (ruo *RevocationUpdateOne) ClearReason() *RevocationUpdateOne {
	ruo.mutation.ClearReason()
	return ruo
}

// SetInfo sets the "info" field.
func (ruo *RevocationUpdateOne) SetInfo(s string) *RevocationUpdateOne {
	ruo.mutation.SetInfo(s)
	return ruo
}

// SetNillableInfo sets the "info" field if the given value is not nil.
func (ruo *RevocationUpdateOne) SetNillableInfo(s *string) *RevocationUpdateOne {
	if s != nil {
		ruo.SetInfo(*s)
	}
	return ruo
}

// ClearInfo clears the value of the "info" field.
func (ruo *RevocationUpdateOne) ClearInfo() *RevocationUpdateOne {
	ruo.mutation.ClearInfo()
	return ruo
}

// SetExpiry sets the "expiry" field.
func (ruo *RevocationUpdateOne) SetExpiry(t time.Time) *RevocationUpdateOne {
	ruo.mutation.SetExpiry(t)
	return ruo
}

// SetNillableExpiry sets the "expiry" field if the given value is not nil.
func (ruo *RevocationUpdateOne) SetNillableExpiry(t *time.Time) *RevocationUpdateOne {
	if t != nil {
		ruo.SetExpiry(*t)
	}
	return ruo
}

// ClearExpiry clears the value of the "expiry" field.
func (ruo *RevocationUpdateOne) ClearExpiry() *RevocationUpdateOne {
	ruo.mutation.ClearExpiry()
	return ruo
}

// SetRevoked sets the "revoked" field.
func (ruo *RevocationUpdateOne) SetRevoked(t time.Time) *RevocationUpdateOne {
	ruo.mutation.SetRevoked(t)
	return ruo
}

// SetNillableRevoked sets the "revoked" field if the given value is not nil.
func (ruo *RevocationUpdateOne) SetNillableRevoked(t *time.Time) *RevocationUpdateOne {
	if t != nil {
		ruo.SetRevoked(*t)
	}
	return ruo
}

// Mutation returns the RevocationMutation object of the builder.
func (ruo *RevocationUpdateOne) Mutation() *RevocationMutation {
	return ruo.mutation
}

// Where appends a list predicates to the RevocationUpdate builder.
func (ruo *RevocationUpdateOne) Where(ps ...predicate.Revocation) *RevocationUpdateOne {
	ruo.mutation.Where(ps...)
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RevocationUpdateOne) Select(field string, fields ...string) *RevocationUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Revocation entity.
func (ruo *RevocationUpdateOne) Save(ctx context.Context) (*Revocation, error) {
	return withHooks(ctx, ruo.sqlSave, ruo.mutation, ruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RevocationUpdateOne) SaveX(ctx context.Context) *Revocation {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RevocationUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RevocationUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ruo *RevocationUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *RevocationUpdateOne {
	ruo.modifiers = append(ruo.modifiers, modifiers...)
	return ruo
}

func (ruo *RevocationUpdateOne) sqlSave(ctx context.Context) (_node *Revocation, err error) {
	_spec := sqlgraph.NewUpdateSpec(revocation.Table, revocation.Columns, sqlgraph.NewFieldSpec(revocation.FieldID, field.TypeInt64))
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Revocation.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, revocation.FieldID)
		for _, f := range fields {
			if !revocation.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != revocation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.Reason(); ok {
		_spec.SetField(revocation.FieldReason, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.AddedReason(); ok {
		_spec.AddField(revocation.FieldReason, field.TypeInt, value)
	}
	if ruo.mutation.ReasonCleared() {
		_spec.ClearField(revocation.FieldReason, field.TypeInt)
	}
	if value, ok := ruo.mutation.Info(); ok {
		_spec.SetField(revocation.FieldInfo, field.TypeString, value)
	}
	if ruo.mutation.InfoCleared() {
		_spec.ClearField(revocation.FieldInfo, field.TypeString)
	}
	if value, ok := ruo.mutation.Expiry(); ok {
		_spec.SetField(revocation.FieldExpiry, field.TypeTime, value)
	}
	if ruo.mutation.ExpiryCleared() {
		_spec.ClearField(revocation.FieldExpiry, field.TypeTime)
	}
	if value, ok := ruo.mutation.Revoked(); ok {
		_spec.SetField(revocation.FieldRevoked, field.TypeTime, value)
	}
	_spec.AddModifiers(ruo.modifiers...)
	_node = &Revocation{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{revocation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruo.mutation.done = true
	return _node, nil
}
