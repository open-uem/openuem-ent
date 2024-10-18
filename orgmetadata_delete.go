// Code generated by ent, DO NOT EDIT.

package openuem_ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/doncicuto/openuem_ent/orgmetadata"
	"github.com/doncicuto/openuem_ent/predicate"
)

// OrgMetadataDelete is the builder for deleting a OrgMetadata entity.
type OrgMetadataDelete struct {
	config
	hooks    []Hook
	mutation *OrgMetadataMutation
}

// Where appends a list predicates to the OrgMetadataDelete builder.
func (omd *OrgMetadataDelete) Where(ps ...predicate.OrgMetadata) *OrgMetadataDelete {
	omd.mutation.Where(ps...)
	return omd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (omd *OrgMetadataDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, omd.sqlExec, omd.mutation, omd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (omd *OrgMetadataDelete) ExecX(ctx context.Context) int {
	n, err := omd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (omd *OrgMetadataDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(orgmetadata.Table, sqlgraph.NewFieldSpec(orgmetadata.FieldID, field.TypeInt))
	if ps := omd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, omd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	omd.mutation.done = true
	return affected, err
}

// OrgMetadataDeleteOne is the builder for deleting a single OrgMetadata entity.
type OrgMetadataDeleteOne struct {
	omd *OrgMetadataDelete
}

// Where appends a list predicates to the OrgMetadataDelete builder.
func (omdo *OrgMetadataDeleteOne) Where(ps ...predicate.OrgMetadata) *OrgMetadataDeleteOne {
	omdo.omd.mutation.Where(ps...)
	return omdo
}

// Exec executes the deletion query.
func (omdo *OrgMetadataDeleteOne) Exec(ctx context.Context) error {
	n, err := omdo.omd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{orgmetadata.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (omdo *OrgMetadataDeleteOne) ExecX(ctx context.Context) {
	if err := omdo.Exec(ctx); err != nil {
		panic(err)
	}
}
