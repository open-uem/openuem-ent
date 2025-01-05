// Code generated by ent, DO NOT EDIT.

package openuem_ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/open-uem/openuem_ent/agent"
	"github.com/open-uem/openuem_ent/computer"
	"github.com/open-uem/openuem_ent/predicate"
)

// ComputerUpdate is the builder for updating Computer entities.
type ComputerUpdate struct {
	config
	hooks     []Hook
	mutation  *ComputerMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the ComputerUpdate builder.
func (cu *ComputerUpdate) Where(ps ...predicate.Computer) *ComputerUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetManufacturer sets the "manufacturer" field.
func (cu *ComputerUpdate) SetManufacturer(s string) *ComputerUpdate {
	cu.mutation.SetManufacturer(s)
	return cu
}

// SetNillableManufacturer sets the "manufacturer" field if the given value is not nil.
func (cu *ComputerUpdate) SetNillableManufacturer(s *string) *ComputerUpdate {
	if s != nil {
		cu.SetManufacturer(*s)
	}
	return cu
}

// ClearManufacturer clears the value of the "manufacturer" field.
func (cu *ComputerUpdate) ClearManufacturer() *ComputerUpdate {
	cu.mutation.ClearManufacturer()
	return cu
}

// SetModel sets the "model" field.
func (cu *ComputerUpdate) SetModel(s string) *ComputerUpdate {
	cu.mutation.SetModel(s)
	return cu
}

// SetNillableModel sets the "model" field if the given value is not nil.
func (cu *ComputerUpdate) SetNillableModel(s *string) *ComputerUpdate {
	if s != nil {
		cu.SetModel(*s)
	}
	return cu
}

// ClearModel clears the value of the "model" field.
func (cu *ComputerUpdate) ClearModel() *ComputerUpdate {
	cu.mutation.ClearModel()
	return cu
}

// SetSerial sets the "serial" field.
func (cu *ComputerUpdate) SetSerial(s string) *ComputerUpdate {
	cu.mutation.SetSerial(s)
	return cu
}

// SetNillableSerial sets the "serial" field if the given value is not nil.
func (cu *ComputerUpdate) SetNillableSerial(s *string) *ComputerUpdate {
	if s != nil {
		cu.SetSerial(*s)
	}
	return cu
}

// ClearSerial clears the value of the "serial" field.
func (cu *ComputerUpdate) ClearSerial() *ComputerUpdate {
	cu.mutation.ClearSerial()
	return cu
}

// SetMemory sets the "memory" field.
func (cu *ComputerUpdate) SetMemory(u uint64) *ComputerUpdate {
	cu.mutation.ResetMemory()
	cu.mutation.SetMemory(u)
	return cu
}

// SetNillableMemory sets the "memory" field if the given value is not nil.
func (cu *ComputerUpdate) SetNillableMemory(u *uint64) *ComputerUpdate {
	if u != nil {
		cu.SetMemory(*u)
	}
	return cu
}

// AddMemory adds u to the "memory" field.
func (cu *ComputerUpdate) AddMemory(u int64) *ComputerUpdate {
	cu.mutation.AddMemory(u)
	return cu
}

// ClearMemory clears the value of the "memory" field.
func (cu *ComputerUpdate) ClearMemory() *ComputerUpdate {
	cu.mutation.ClearMemory()
	return cu
}

// SetProcessor sets the "processor" field.
func (cu *ComputerUpdate) SetProcessor(s string) *ComputerUpdate {
	cu.mutation.SetProcessor(s)
	return cu
}

// SetNillableProcessor sets the "processor" field if the given value is not nil.
func (cu *ComputerUpdate) SetNillableProcessor(s *string) *ComputerUpdate {
	if s != nil {
		cu.SetProcessor(*s)
	}
	return cu
}

// ClearProcessor clears the value of the "processor" field.
func (cu *ComputerUpdate) ClearProcessor() *ComputerUpdate {
	cu.mutation.ClearProcessor()
	return cu
}

// SetProcessorCores sets the "processor_cores" field.
func (cu *ComputerUpdate) SetProcessorCores(i int64) *ComputerUpdate {
	cu.mutation.ResetProcessorCores()
	cu.mutation.SetProcessorCores(i)
	return cu
}

// SetNillableProcessorCores sets the "processor_cores" field if the given value is not nil.
func (cu *ComputerUpdate) SetNillableProcessorCores(i *int64) *ComputerUpdate {
	if i != nil {
		cu.SetProcessorCores(*i)
	}
	return cu
}

// AddProcessorCores adds i to the "processor_cores" field.
func (cu *ComputerUpdate) AddProcessorCores(i int64) *ComputerUpdate {
	cu.mutation.AddProcessorCores(i)
	return cu
}

// ClearProcessorCores clears the value of the "processor_cores" field.
func (cu *ComputerUpdate) ClearProcessorCores() *ComputerUpdate {
	cu.mutation.ClearProcessorCores()
	return cu
}

// SetProcessorArch sets the "processor_arch" field.
func (cu *ComputerUpdate) SetProcessorArch(s string) *ComputerUpdate {
	cu.mutation.SetProcessorArch(s)
	return cu
}

// SetNillableProcessorArch sets the "processor_arch" field if the given value is not nil.
func (cu *ComputerUpdate) SetNillableProcessorArch(s *string) *ComputerUpdate {
	if s != nil {
		cu.SetProcessorArch(*s)
	}
	return cu
}

// ClearProcessorArch clears the value of the "processor_arch" field.
func (cu *ComputerUpdate) ClearProcessorArch() *ComputerUpdate {
	cu.mutation.ClearProcessorArch()
	return cu
}

// SetOwnerID sets the "owner" edge to the Agent entity by ID.
func (cu *ComputerUpdate) SetOwnerID(id string) *ComputerUpdate {
	cu.mutation.SetOwnerID(id)
	return cu
}

// SetOwner sets the "owner" edge to the Agent entity.
func (cu *ComputerUpdate) SetOwner(a *Agent) *ComputerUpdate {
	return cu.SetOwnerID(a.ID)
}

// Mutation returns the ComputerMutation object of the builder.
func (cu *ComputerUpdate) Mutation() *ComputerMutation {
	return cu.mutation
}

// ClearOwner clears the "owner" edge to the Agent entity.
func (cu *ComputerUpdate) ClearOwner() *ComputerUpdate {
	cu.mutation.ClearOwner()
	return cu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ComputerUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ComputerUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ComputerUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ComputerUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *ComputerUpdate) check() error {
	if cu.mutation.OwnerCleared() && len(cu.mutation.OwnerIDs()) > 0 {
		return errors.New(`openuem_ent: clearing a required unique edge "Computer.owner"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cu *ComputerUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ComputerUpdate {
	cu.modifiers = append(cu.modifiers, modifiers...)
	return cu
}

func (cu *ComputerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(computer.Table, computer.Columns, sqlgraph.NewFieldSpec(computer.FieldID, field.TypeInt))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Manufacturer(); ok {
		_spec.SetField(computer.FieldManufacturer, field.TypeString, value)
	}
	if cu.mutation.ManufacturerCleared() {
		_spec.ClearField(computer.FieldManufacturer, field.TypeString)
	}
	if value, ok := cu.mutation.Model(); ok {
		_spec.SetField(computer.FieldModel, field.TypeString, value)
	}
	if cu.mutation.ModelCleared() {
		_spec.ClearField(computer.FieldModel, field.TypeString)
	}
	if value, ok := cu.mutation.Serial(); ok {
		_spec.SetField(computer.FieldSerial, field.TypeString, value)
	}
	if cu.mutation.SerialCleared() {
		_spec.ClearField(computer.FieldSerial, field.TypeString)
	}
	if value, ok := cu.mutation.Memory(); ok {
		_spec.SetField(computer.FieldMemory, field.TypeUint64, value)
	}
	if value, ok := cu.mutation.AddedMemory(); ok {
		_spec.AddField(computer.FieldMemory, field.TypeUint64, value)
	}
	if cu.mutation.MemoryCleared() {
		_spec.ClearField(computer.FieldMemory, field.TypeUint64)
	}
	if value, ok := cu.mutation.Processor(); ok {
		_spec.SetField(computer.FieldProcessor, field.TypeString, value)
	}
	if cu.mutation.ProcessorCleared() {
		_spec.ClearField(computer.FieldProcessor, field.TypeString)
	}
	if value, ok := cu.mutation.ProcessorCores(); ok {
		_spec.SetField(computer.FieldProcessorCores, field.TypeInt64, value)
	}
	if value, ok := cu.mutation.AddedProcessorCores(); ok {
		_spec.AddField(computer.FieldProcessorCores, field.TypeInt64, value)
	}
	if cu.mutation.ProcessorCoresCleared() {
		_spec.ClearField(computer.FieldProcessorCores, field.TypeInt64)
	}
	if value, ok := cu.mutation.ProcessorArch(); ok {
		_spec.SetField(computer.FieldProcessorArch, field.TypeString, value)
	}
	if cu.mutation.ProcessorArchCleared() {
		_spec.ClearField(computer.FieldProcessorArch, field.TypeString)
	}
	if cu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   computer.OwnerTable,
			Columns: []string{computer.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(agent.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   computer.OwnerTable,
			Columns: []string{computer.OwnerColumn},
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
	_spec.AddModifiers(cu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{computer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// ComputerUpdateOne is the builder for updating a single Computer entity.
type ComputerUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ComputerMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetManufacturer sets the "manufacturer" field.
func (cuo *ComputerUpdateOne) SetManufacturer(s string) *ComputerUpdateOne {
	cuo.mutation.SetManufacturer(s)
	return cuo
}

// SetNillableManufacturer sets the "manufacturer" field if the given value is not nil.
func (cuo *ComputerUpdateOne) SetNillableManufacturer(s *string) *ComputerUpdateOne {
	if s != nil {
		cuo.SetManufacturer(*s)
	}
	return cuo
}

// ClearManufacturer clears the value of the "manufacturer" field.
func (cuo *ComputerUpdateOne) ClearManufacturer() *ComputerUpdateOne {
	cuo.mutation.ClearManufacturer()
	return cuo
}

// SetModel sets the "model" field.
func (cuo *ComputerUpdateOne) SetModel(s string) *ComputerUpdateOne {
	cuo.mutation.SetModel(s)
	return cuo
}

// SetNillableModel sets the "model" field if the given value is not nil.
func (cuo *ComputerUpdateOne) SetNillableModel(s *string) *ComputerUpdateOne {
	if s != nil {
		cuo.SetModel(*s)
	}
	return cuo
}

// ClearModel clears the value of the "model" field.
func (cuo *ComputerUpdateOne) ClearModel() *ComputerUpdateOne {
	cuo.mutation.ClearModel()
	return cuo
}

// SetSerial sets the "serial" field.
func (cuo *ComputerUpdateOne) SetSerial(s string) *ComputerUpdateOne {
	cuo.mutation.SetSerial(s)
	return cuo
}

// SetNillableSerial sets the "serial" field if the given value is not nil.
func (cuo *ComputerUpdateOne) SetNillableSerial(s *string) *ComputerUpdateOne {
	if s != nil {
		cuo.SetSerial(*s)
	}
	return cuo
}

// ClearSerial clears the value of the "serial" field.
func (cuo *ComputerUpdateOne) ClearSerial() *ComputerUpdateOne {
	cuo.mutation.ClearSerial()
	return cuo
}

// SetMemory sets the "memory" field.
func (cuo *ComputerUpdateOne) SetMemory(u uint64) *ComputerUpdateOne {
	cuo.mutation.ResetMemory()
	cuo.mutation.SetMemory(u)
	return cuo
}

// SetNillableMemory sets the "memory" field if the given value is not nil.
func (cuo *ComputerUpdateOne) SetNillableMemory(u *uint64) *ComputerUpdateOne {
	if u != nil {
		cuo.SetMemory(*u)
	}
	return cuo
}

// AddMemory adds u to the "memory" field.
func (cuo *ComputerUpdateOne) AddMemory(u int64) *ComputerUpdateOne {
	cuo.mutation.AddMemory(u)
	return cuo
}

// ClearMemory clears the value of the "memory" field.
func (cuo *ComputerUpdateOne) ClearMemory() *ComputerUpdateOne {
	cuo.mutation.ClearMemory()
	return cuo
}

// SetProcessor sets the "processor" field.
func (cuo *ComputerUpdateOne) SetProcessor(s string) *ComputerUpdateOne {
	cuo.mutation.SetProcessor(s)
	return cuo
}

// SetNillableProcessor sets the "processor" field if the given value is not nil.
func (cuo *ComputerUpdateOne) SetNillableProcessor(s *string) *ComputerUpdateOne {
	if s != nil {
		cuo.SetProcessor(*s)
	}
	return cuo
}

// ClearProcessor clears the value of the "processor" field.
func (cuo *ComputerUpdateOne) ClearProcessor() *ComputerUpdateOne {
	cuo.mutation.ClearProcessor()
	return cuo
}

// SetProcessorCores sets the "processor_cores" field.
func (cuo *ComputerUpdateOne) SetProcessorCores(i int64) *ComputerUpdateOne {
	cuo.mutation.ResetProcessorCores()
	cuo.mutation.SetProcessorCores(i)
	return cuo
}

// SetNillableProcessorCores sets the "processor_cores" field if the given value is not nil.
func (cuo *ComputerUpdateOne) SetNillableProcessorCores(i *int64) *ComputerUpdateOne {
	if i != nil {
		cuo.SetProcessorCores(*i)
	}
	return cuo
}

// AddProcessorCores adds i to the "processor_cores" field.
func (cuo *ComputerUpdateOne) AddProcessorCores(i int64) *ComputerUpdateOne {
	cuo.mutation.AddProcessorCores(i)
	return cuo
}

// ClearProcessorCores clears the value of the "processor_cores" field.
func (cuo *ComputerUpdateOne) ClearProcessorCores() *ComputerUpdateOne {
	cuo.mutation.ClearProcessorCores()
	return cuo
}

// SetProcessorArch sets the "processor_arch" field.
func (cuo *ComputerUpdateOne) SetProcessorArch(s string) *ComputerUpdateOne {
	cuo.mutation.SetProcessorArch(s)
	return cuo
}

// SetNillableProcessorArch sets the "processor_arch" field if the given value is not nil.
func (cuo *ComputerUpdateOne) SetNillableProcessorArch(s *string) *ComputerUpdateOne {
	if s != nil {
		cuo.SetProcessorArch(*s)
	}
	return cuo
}

// ClearProcessorArch clears the value of the "processor_arch" field.
func (cuo *ComputerUpdateOne) ClearProcessorArch() *ComputerUpdateOne {
	cuo.mutation.ClearProcessorArch()
	return cuo
}

// SetOwnerID sets the "owner" edge to the Agent entity by ID.
func (cuo *ComputerUpdateOne) SetOwnerID(id string) *ComputerUpdateOne {
	cuo.mutation.SetOwnerID(id)
	return cuo
}

// SetOwner sets the "owner" edge to the Agent entity.
func (cuo *ComputerUpdateOne) SetOwner(a *Agent) *ComputerUpdateOne {
	return cuo.SetOwnerID(a.ID)
}

// Mutation returns the ComputerMutation object of the builder.
func (cuo *ComputerUpdateOne) Mutation() *ComputerMutation {
	return cuo.mutation
}

// ClearOwner clears the "owner" edge to the Agent entity.
func (cuo *ComputerUpdateOne) ClearOwner() *ComputerUpdateOne {
	cuo.mutation.ClearOwner()
	return cuo
}

// Where appends a list predicates to the ComputerUpdate builder.
func (cuo *ComputerUpdateOne) Where(ps ...predicate.Computer) *ComputerUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ComputerUpdateOne) Select(field string, fields ...string) *ComputerUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Computer entity.
func (cuo *ComputerUpdateOne) Save(ctx context.Context) (*Computer, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ComputerUpdateOne) SaveX(ctx context.Context) *Computer {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ComputerUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ComputerUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *ComputerUpdateOne) check() error {
	if cuo.mutation.OwnerCleared() && len(cuo.mutation.OwnerIDs()) > 0 {
		return errors.New(`openuem_ent: clearing a required unique edge "Computer.owner"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cuo *ComputerUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ComputerUpdateOne {
	cuo.modifiers = append(cuo.modifiers, modifiers...)
	return cuo
}

func (cuo *ComputerUpdateOne) sqlSave(ctx context.Context) (_node *Computer, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(computer.Table, computer.Columns, sqlgraph.NewFieldSpec(computer.FieldID, field.TypeInt))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`openuem_ent: missing "Computer.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, computer.FieldID)
		for _, f := range fields {
			if !computer.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("openuem_ent: invalid field %q for query", f)}
			}
			if f != computer.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Manufacturer(); ok {
		_spec.SetField(computer.FieldManufacturer, field.TypeString, value)
	}
	if cuo.mutation.ManufacturerCleared() {
		_spec.ClearField(computer.FieldManufacturer, field.TypeString)
	}
	if value, ok := cuo.mutation.Model(); ok {
		_spec.SetField(computer.FieldModel, field.TypeString, value)
	}
	if cuo.mutation.ModelCleared() {
		_spec.ClearField(computer.FieldModel, field.TypeString)
	}
	if value, ok := cuo.mutation.Serial(); ok {
		_spec.SetField(computer.FieldSerial, field.TypeString, value)
	}
	if cuo.mutation.SerialCleared() {
		_spec.ClearField(computer.FieldSerial, field.TypeString)
	}
	if value, ok := cuo.mutation.Memory(); ok {
		_spec.SetField(computer.FieldMemory, field.TypeUint64, value)
	}
	if value, ok := cuo.mutation.AddedMemory(); ok {
		_spec.AddField(computer.FieldMemory, field.TypeUint64, value)
	}
	if cuo.mutation.MemoryCleared() {
		_spec.ClearField(computer.FieldMemory, field.TypeUint64)
	}
	if value, ok := cuo.mutation.Processor(); ok {
		_spec.SetField(computer.FieldProcessor, field.TypeString, value)
	}
	if cuo.mutation.ProcessorCleared() {
		_spec.ClearField(computer.FieldProcessor, field.TypeString)
	}
	if value, ok := cuo.mutation.ProcessorCores(); ok {
		_spec.SetField(computer.FieldProcessorCores, field.TypeInt64, value)
	}
	if value, ok := cuo.mutation.AddedProcessorCores(); ok {
		_spec.AddField(computer.FieldProcessorCores, field.TypeInt64, value)
	}
	if cuo.mutation.ProcessorCoresCleared() {
		_spec.ClearField(computer.FieldProcessorCores, field.TypeInt64)
	}
	if value, ok := cuo.mutation.ProcessorArch(); ok {
		_spec.SetField(computer.FieldProcessorArch, field.TypeString, value)
	}
	if cuo.mutation.ProcessorArchCleared() {
		_spec.ClearField(computer.FieldProcessorArch, field.TypeString)
	}
	if cuo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   computer.OwnerTable,
			Columns: []string{computer.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(agent.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   computer.OwnerTable,
			Columns: []string{computer.OwnerColumn},
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
	_spec.AddModifiers(cuo.modifiers...)
	_node = &Computer{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{computer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
