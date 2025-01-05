// Code generated by ent, DO NOT EDIT.

package openuem_ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/open-uem/openuem_ent/networkadapter"
	"github.com/open-uem/openuem_ent/predicate"
)

// NetworkAdapterDelete is the builder for deleting a NetworkAdapter entity.
type NetworkAdapterDelete struct {
	config
	hooks    []Hook
	mutation *NetworkAdapterMutation
}

// Where appends a list predicates to the NetworkAdapterDelete builder.
func (nad *NetworkAdapterDelete) Where(ps ...predicate.NetworkAdapter) *NetworkAdapterDelete {
	nad.mutation.Where(ps...)
	return nad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (nad *NetworkAdapterDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, nad.sqlExec, nad.mutation, nad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (nad *NetworkAdapterDelete) ExecX(ctx context.Context) int {
	n, err := nad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (nad *NetworkAdapterDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(networkadapter.Table, sqlgraph.NewFieldSpec(networkadapter.FieldID, field.TypeInt))
	if ps := nad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, nad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	nad.mutation.done = true
	return affected, err
}

// NetworkAdapterDeleteOne is the builder for deleting a single NetworkAdapter entity.
type NetworkAdapterDeleteOne struct {
	nad *NetworkAdapterDelete
}

// Where appends a list predicates to the NetworkAdapterDelete builder.
func (nado *NetworkAdapterDeleteOne) Where(ps ...predicate.NetworkAdapter) *NetworkAdapterDeleteOne {
	nado.nad.mutation.Where(ps...)
	return nado
}

// Exec executes the deletion query.
func (nado *NetworkAdapterDeleteOne) Exec(ctx context.Context) error {
	n, err := nado.nad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{networkadapter.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (nado *NetworkAdapterDeleteOne) ExecX(ctx context.Context) {
	if err := nado.Exec(ctx); err != nil {
		panic(err)
	}
}
