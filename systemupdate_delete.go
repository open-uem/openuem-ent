// Code generated by ent, DO NOT EDIT.

package openuem_ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/open-uem/openuem_ent/predicate"
	"github.com/open-uem/openuem_ent/systemupdate"
)

// SystemUpdateDelete is the builder for deleting a SystemUpdate entity.
type SystemUpdateDelete struct {
	config
	hooks    []Hook
	mutation *SystemUpdateMutation
}

// Where appends a list predicates to the SystemUpdateDelete builder.
func (sud *SystemUpdateDelete) Where(ps ...predicate.SystemUpdate) *SystemUpdateDelete {
	sud.mutation.Where(ps...)
	return sud
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sud *SystemUpdateDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, sud.sqlExec, sud.mutation, sud.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (sud *SystemUpdateDelete) ExecX(ctx context.Context) int {
	n, err := sud.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sud *SystemUpdateDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(systemupdate.Table, sqlgraph.NewFieldSpec(systemupdate.FieldID, field.TypeInt))
	if ps := sud.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, sud.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	sud.mutation.done = true
	return affected, err
}

// SystemUpdateDeleteOne is the builder for deleting a single SystemUpdate entity.
type SystemUpdateDeleteOne struct {
	sud *SystemUpdateDelete
}

// Where appends a list predicates to the SystemUpdateDelete builder.
func (sudo *SystemUpdateDeleteOne) Where(ps ...predicate.SystemUpdate) *SystemUpdateDeleteOne {
	sudo.sud.mutation.Where(ps...)
	return sudo
}

// Exec executes the deletion query.
func (sudo *SystemUpdateDeleteOne) Exec(ctx context.Context) error {
	n, err := sudo.sud.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{systemupdate.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sudo *SystemUpdateDeleteOne) ExecX(ctx context.Context) {
	if err := sudo.Exec(ctx); err != nil {
		panic(err)
	}
}
