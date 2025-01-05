// Code generated by ent, DO NOT EDIT.

package openuem_ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/open-uem/openuem_ent/predicate"
	"github.com/open-uem/openuem_ent/sessions"
)

// SessionsDelete is the builder for deleting a Sessions entity.
type SessionsDelete struct {
	config
	hooks    []Hook
	mutation *SessionsMutation
}

// Where appends a list predicates to the SessionsDelete builder.
func (sd *SessionsDelete) Where(ps ...predicate.Sessions) *SessionsDelete {
	sd.mutation.Where(ps...)
	return sd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sd *SessionsDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, sd.sqlExec, sd.mutation, sd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (sd *SessionsDelete) ExecX(ctx context.Context) int {
	n, err := sd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sd *SessionsDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(sessions.Table, sqlgraph.NewFieldSpec(sessions.FieldID, field.TypeString))
	if ps := sd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, sd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	sd.mutation.done = true
	return affected, err
}

// SessionsDeleteOne is the builder for deleting a single Sessions entity.
type SessionsDeleteOne struct {
	sd *SessionsDelete
}

// Where appends a list predicates to the SessionsDelete builder.
func (sdo *SessionsDeleteOne) Where(ps ...predicate.Sessions) *SessionsDeleteOne {
	sdo.sd.mutation.Where(ps...)
	return sdo
}

// Exec executes the deletion query.
func (sdo *SessionsDeleteOne) Exec(ctx context.Context) error {
	n, err := sdo.sd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{sessions.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sdo *SessionsDeleteOne) ExecX(ctx context.Context) {
	if err := sdo.Exec(ctx); err != nil {
		panic(err)
	}
}
