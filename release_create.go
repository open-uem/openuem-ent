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
	"github.com/open-uem/openuem_ent/release"
)

// ReleaseCreate is the builder for creating a Release entity.
type ReleaseCreate struct {
	config
	mutation *ReleaseMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetReleaseType sets the "release_type" field.
func (rc *ReleaseCreate) SetReleaseType(rt release.ReleaseType) *ReleaseCreate {
	rc.mutation.SetReleaseType(rt)
	return rc
}

// SetNillableReleaseType sets the "release_type" field if the given value is not nil.
func (rc *ReleaseCreate) SetNillableReleaseType(rt *release.ReleaseType) *ReleaseCreate {
	if rt != nil {
		rc.SetReleaseType(*rt)
	}
	return rc
}

// SetVersion sets the "version" field.
func (rc *ReleaseCreate) SetVersion(s string) *ReleaseCreate {
	rc.mutation.SetVersion(s)
	return rc
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (rc *ReleaseCreate) SetNillableVersion(s *string) *ReleaseCreate {
	if s != nil {
		rc.SetVersion(*s)
	}
	return rc
}

// SetChannel sets the "channel" field.
func (rc *ReleaseCreate) SetChannel(s string) *ReleaseCreate {
	rc.mutation.SetChannel(s)
	return rc
}

// SetNillableChannel sets the "channel" field if the given value is not nil.
func (rc *ReleaseCreate) SetNillableChannel(s *string) *ReleaseCreate {
	if s != nil {
		rc.SetChannel(*s)
	}
	return rc
}

// SetSummary sets the "summary" field.
func (rc *ReleaseCreate) SetSummary(s string) *ReleaseCreate {
	rc.mutation.SetSummary(s)
	return rc
}

// SetNillableSummary sets the "summary" field if the given value is not nil.
func (rc *ReleaseCreate) SetNillableSummary(s *string) *ReleaseCreate {
	if s != nil {
		rc.SetSummary(*s)
	}
	return rc
}

// SetReleaseNotes sets the "release_notes" field.
func (rc *ReleaseCreate) SetReleaseNotes(s string) *ReleaseCreate {
	rc.mutation.SetReleaseNotes(s)
	return rc
}

// SetNillableReleaseNotes sets the "release_notes" field if the given value is not nil.
func (rc *ReleaseCreate) SetNillableReleaseNotes(s *string) *ReleaseCreate {
	if s != nil {
		rc.SetReleaseNotes(*s)
	}
	return rc
}

// SetFileURL sets the "file_url" field.
func (rc *ReleaseCreate) SetFileURL(s string) *ReleaseCreate {
	rc.mutation.SetFileURL(s)
	return rc
}

// SetNillableFileURL sets the "file_url" field if the given value is not nil.
func (rc *ReleaseCreate) SetNillableFileURL(s *string) *ReleaseCreate {
	if s != nil {
		rc.SetFileURL(*s)
	}
	return rc
}

// SetChecksum sets the "checksum" field.
func (rc *ReleaseCreate) SetChecksum(s string) *ReleaseCreate {
	rc.mutation.SetChecksum(s)
	return rc
}

// SetNillableChecksum sets the "checksum" field if the given value is not nil.
func (rc *ReleaseCreate) SetNillableChecksum(s *string) *ReleaseCreate {
	if s != nil {
		rc.SetChecksum(*s)
	}
	return rc
}

// SetIsCritical sets the "is_critical" field.
func (rc *ReleaseCreate) SetIsCritical(b bool) *ReleaseCreate {
	rc.mutation.SetIsCritical(b)
	return rc
}

// SetNillableIsCritical sets the "is_critical" field if the given value is not nil.
func (rc *ReleaseCreate) SetNillableIsCritical(b *bool) *ReleaseCreate {
	if b != nil {
		rc.SetIsCritical(*b)
	}
	return rc
}

// SetReleaseDate sets the "release_date" field.
func (rc *ReleaseCreate) SetReleaseDate(t time.Time) *ReleaseCreate {
	rc.mutation.SetReleaseDate(t)
	return rc
}

// SetNillableReleaseDate sets the "release_date" field if the given value is not nil.
func (rc *ReleaseCreate) SetNillableReleaseDate(t *time.Time) *ReleaseCreate {
	if t != nil {
		rc.SetReleaseDate(*t)
	}
	return rc
}

// SetOs sets the "os" field.
func (rc *ReleaseCreate) SetOs(s string) *ReleaseCreate {
	rc.mutation.SetOs(s)
	return rc
}

// SetNillableOs sets the "os" field if the given value is not nil.
func (rc *ReleaseCreate) SetNillableOs(s *string) *ReleaseCreate {
	if s != nil {
		rc.SetOs(*s)
	}
	return rc
}

// SetArch sets the "arch" field.
func (rc *ReleaseCreate) SetArch(s string) *ReleaseCreate {
	rc.mutation.SetArch(s)
	return rc
}

// SetNillableArch sets the "arch" field if the given value is not nil.
func (rc *ReleaseCreate) SetNillableArch(s *string) *ReleaseCreate {
	if s != nil {
		rc.SetArch(*s)
	}
	return rc
}

// AddAgentIDs adds the "agents" edge to the Agent entity by IDs.
func (rc *ReleaseCreate) AddAgentIDs(ids ...string) *ReleaseCreate {
	rc.mutation.AddAgentIDs(ids...)
	return rc
}

// AddAgents adds the "agents" edges to the Agent entity.
func (rc *ReleaseCreate) AddAgents(a ...*Agent) *ReleaseCreate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return rc.AddAgentIDs(ids...)
}

// Mutation returns the ReleaseMutation object of the builder.
func (rc *ReleaseCreate) Mutation() *ReleaseMutation {
	return rc.mutation
}

// Save creates the Release in the database.
func (rc *ReleaseCreate) Save(ctx context.Context) (*Release, error) {
	return withHooks(ctx, rc.sqlSave, rc.mutation, rc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rc *ReleaseCreate) SaveX(ctx context.Context) *Release {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *ReleaseCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *ReleaseCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *ReleaseCreate) check() error {
	if v, ok := rc.mutation.ReleaseType(); ok {
		if err := release.ReleaseTypeValidator(v); err != nil {
			return &ValidationError{Name: "release_type", err: fmt.Errorf(`openuem_ent: validator failed for field "Release.release_type": %w`, err)}
		}
	}
	return nil
}

func (rc *ReleaseCreate) sqlSave(ctx context.Context) (*Release, error) {
	if err := rc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	rc.mutation.id = &_node.ID
	rc.mutation.done = true
	return _node, nil
}

func (rc *ReleaseCreate) createSpec() (*Release, *sqlgraph.CreateSpec) {
	var (
		_node = &Release{config: rc.config}
		_spec = sqlgraph.NewCreateSpec(release.Table, sqlgraph.NewFieldSpec(release.FieldID, field.TypeInt))
	)
	_spec.OnConflict = rc.conflict
	if value, ok := rc.mutation.ReleaseType(); ok {
		_spec.SetField(release.FieldReleaseType, field.TypeEnum, value)
		_node.ReleaseType = value
	}
	if value, ok := rc.mutation.Version(); ok {
		_spec.SetField(release.FieldVersion, field.TypeString, value)
		_node.Version = value
	}
	if value, ok := rc.mutation.Channel(); ok {
		_spec.SetField(release.FieldChannel, field.TypeString, value)
		_node.Channel = value
	}
	if value, ok := rc.mutation.Summary(); ok {
		_spec.SetField(release.FieldSummary, field.TypeString, value)
		_node.Summary = value
	}
	if value, ok := rc.mutation.ReleaseNotes(); ok {
		_spec.SetField(release.FieldReleaseNotes, field.TypeString, value)
		_node.ReleaseNotes = value
	}
	if value, ok := rc.mutation.FileURL(); ok {
		_spec.SetField(release.FieldFileURL, field.TypeString, value)
		_node.FileURL = value
	}
	if value, ok := rc.mutation.Checksum(); ok {
		_spec.SetField(release.FieldChecksum, field.TypeString, value)
		_node.Checksum = value
	}
	if value, ok := rc.mutation.IsCritical(); ok {
		_spec.SetField(release.FieldIsCritical, field.TypeBool, value)
		_node.IsCritical = value
	}
	if value, ok := rc.mutation.ReleaseDate(); ok {
		_spec.SetField(release.FieldReleaseDate, field.TypeTime, value)
		_node.ReleaseDate = value
	}
	if value, ok := rc.mutation.Os(); ok {
		_spec.SetField(release.FieldOs, field.TypeString, value)
		_node.Os = value
	}
	if value, ok := rc.mutation.Arch(); ok {
		_spec.SetField(release.FieldArch, field.TypeString, value)
		_node.Arch = value
	}
	if nodes := rc.mutation.AgentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   release.AgentsTable,
			Columns: []string{release.AgentsColumn},
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
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Release.Create().
//		SetReleaseType(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ReleaseUpsert) {
//			SetReleaseType(v+v).
//		}).
//		Exec(ctx)
func (rc *ReleaseCreate) OnConflict(opts ...sql.ConflictOption) *ReleaseUpsertOne {
	rc.conflict = opts
	return &ReleaseUpsertOne{
		create: rc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Release.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (rc *ReleaseCreate) OnConflictColumns(columns ...string) *ReleaseUpsertOne {
	rc.conflict = append(rc.conflict, sql.ConflictColumns(columns...))
	return &ReleaseUpsertOne{
		create: rc,
	}
}

type (
	// ReleaseUpsertOne is the builder for "upsert"-ing
	//  one Release node.
	ReleaseUpsertOne struct {
		create *ReleaseCreate
	}

	// ReleaseUpsert is the "OnConflict" setter.
	ReleaseUpsert struct {
		*sql.UpdateSet
	}
)

// SetReleaseType sets the "release_type" field.
func (u *ReleaseUpsert) SetReleaseType(v release.ReleaseType) *ReleaseUpsert {
	u.Set(release.FieldReleaseType, v)
	return u
}

// UpdateReleaseType sets the "release_type" field to the value that was provided on create.
func (u *ReleaseUpsert) UpdateReleaseType() *ReleaseUpsert {
	u.SetExcluded(release.FieldReleaseType)
	return u
}

// ClearReleaseType clears the value of the "release_type" field.
func (u *ReleaseUpsert) ClearReleaseType() *ReleaseUpsert {
	u.SetNull(release.FieldReleaseType)
	return u
}

// SetVersion sets the "version" field.
func (u *ReleaseUpsert) SetVersion(v string) *ReleaseUpsert {
	u.Set(release.FieldVersion, v)
	return u
}

// UpdateVersion sets the "version" field to the value that was provided on create.
func (u *ReleaseUpsert) UpdateVersion() *ReleaseUpsert {
	u.SetExcluded(release.FieldVersion)
	return u
}

// ClearVersion clears the value of the "version" field.
func (u *ReleaseUpsert) ClearVersion() *ReleaseUpsert {
	u.SetNull(release.FieldVersion)
	return u
}

// SetChannel sets the "channel" field.
func (u *ReleaseUpsert) SetChannel(v string) *ReleaseUpsert {
	u.Set(release.FieldChannel, v)
	return u
}

// UpdateChannel sets the "channel" field to the value that was provided on create.
func (u *ReleaseUpsert) UpdateChannel() *ReleaseUpsert {
	u.SetExcluded(release.FieldChannel)
	return u
}

// ClearChannel clears the value of the "channel" field.
func (u *ReleaseUpsert) ClearChannel() *ReleaseUpsert {
	u.SetNull(release.FieldChannel)
	return u
}

// SetSummary sets the "summary" field.
func (u *ReleaseUpsert) SetSummary(v string) *ReleaseUpsert {
	u.Set(release.FieldSummary, v)
	return u
}

// UpdateSummary sets the "summary" field to the value that was provided on create.
func (u *ReleaseUpsert) UpdateSummary() *ReleaseUpsert {
	u.SetExcluded(release.FieldSummary)
	return u
}

// ClearSummary clears the value of the "summary" field.
func (u *ReleaseUpsert) ClearSummary() *ReleaseUpsert {
	u.SetNull(release.FieldSummary)
	return u
}

// SetReleaseNotes sets the "release_notes" field.
func (u *ReleaseUpsert) SetReleaseNotes(v string) *ReleaseUpsert {
	u.Set(release.FieldReleaseNotes, v)
	return u
}

// UpdateReleaseNotes sets the "release_notes" field to the value that was provided on create.
func (u *ReleaseUpsert) UpdateReleaseNotes() *ReleaseUpsert {
	u.SetExcluded(release.FieldReleaseNotes)
	return u
}

// ClearReleaseNotes clears the value of the "release_notes" field.
func (u *ReleaseUpsert) ClearReleaseNotes() *ReleaseUpsert {
	u.SetNull(release.FieldReleaseNotes)
	return u
}

// SetFileURL sets the "file_url" field.
func (u *ReleaseUpsert) SetFileURL(v string) *ReleaseUpsert {
	u.Set(release.FieldFileURL, v)
	return u
}

// UpdateFileURL sets the "file_url" field to the value that was provided on create.
func (u *ReleaseUpsert) UpdateFileURL() *ReleaseUpsert {
	u.SetExcluded(release.FieldFileURL)
	return u
}

// ClearFileURL clears the value of the "file_url" field.
func (u *ReleaseUpsert) ClearFileURL() *ReleaseUpsert {
	u.SetNull(release.FieldFileURL)
	return u
}

// SetChecksum sets the "checksum" field.
func (u *ReleaseUpsert) SetChecksum(v string) *ReleaseUpsert {
	u.Set(release.FieldChecksum, v)
	return u
}

// UpdateChecksum sets the "checksum" field to the value that was provided on create.
func (u *ReleaseUpsert) UpdateChecksum() *ReleaseUpsert {
	u.SetExcluded(release.FieldChecksum)
	return u
}

// ClearChecksum clears the value of the "checksum" field.
func (u *ReleaseUpsert) ClearChecksum() *ReleaseUpsert {
	u.SetNull(release.FieldChecksum)
	return u
}

// SetIsCritical sets the "is_critical" field.
func (u *ReleaseUpsert) SetIsCritical(v bool) *ReleaseUpsert {
	u.Set(release.FieldIsCritical, v)
	return u
}

// UpdateIsCritical sets the "is_critical" field to the value that was provided on create.
func (u *ReleaseUpsert) UpdateIsCritical() *ReleaseUpsert {
	u.SetExcluded(release.FieldIsCritical)
	return u
}

// ClearIsCritical clears the value of the "is_critical" field.
func (u *ReleaseUpsert) ClearIsCritical() *ReleaseUpsert {
	u.SetNull(release.FieldIsCritical)
	return u
}

// SetReleaseDate sets the "release_date" field.
func (u *ReleaseUpsert) SetReleaseDate(v time.Time) *ReleaseUpsert {
	u.Set(release.FieldReleaseDate, v)
	return u
}

// UpdateReleaseDate sets the "release_date" field to the value that was provided on create.
func (u *ReleaseUpsert) UpdateReleaseDate() *ReleaseUpsert {
	u.SetExcluded(release.FieldReleaseDate)
	return u
}

// ClearReleaseDate clears the value of the "release_date" field.
func (u *ReleaseUpsert) ClearReleaseDate() *ReleaseUpsert {
	u.SetNull(release.FieldReleaseDate)
	return u
}

// SetOs sets the "os" field.
func (u *ReleaseUpsert) SetOs(v string) *ReleaseUpsert {
	u.Set(release.FieldOs, v)
	return u
}

// UpdateOs sets the "os" field to the value that was provided on create.
func (u *ReleaseUpsert) UpdateOs() *ReleaseUpsert {
	u.SetExcluded(release.FieldOs)
	return u
}

// ClearOs clears the value of the "os" field.
func (u *ReleaseUpsert) ClearOs() *ReleaseUpsert {
	u.SetNull(release.FieldOs)
	return u
}

// SetArch sets the "arch" field.
func (u *ReleaseUpsert) SetArch(v string) *ReleaseUpsert {
	u.Set(release.FieldArch, v)
	return u
}

// UpdateArch sets the "arch" field to the value that was provided on create.
func (u *ReleaseUpsert) UpdateArch() *ReleaseUpsert {
	u.SetExcluded(release.FieldArch)
	return u
}

// ClearArch clears the value of the "arch" field.
func (u *ReleaseUpsert) ClearArch() *ReleaseUpsert {
	u.SetNull(release.FieldArch)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Release.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *ReleaseUpsertOne) UpdateNewValues() *ReleaseUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Release.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ReleaseUpsertOne) Ignore() *ReleaseUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ReleaseUpsertOne) DoNothing() *ReleaseUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ReleaseCreate.OnConflict
// documentation for more info.
func (u *ReleaseUpsertOne) Update(set func(*ReleaseUpsert)) *ReleaseUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ReleaseUpsert{UpdateSet: update})
	}))
	return u
}

// SetReleaseType sets the "release_type" field.
func (u *ReleaseUpsertOne) SetReleaseType(v release.ReleaseType) *ReleaseUpsertOne {
	return u.Update(func(s *ReleaseUpsert) {
		s.SetReleaseType(v)
	})
}

// UpdateReleaseType sets the "release_type" field to the value that was provided on create.
func (u *ReleaseUpsertOne) UpdateReleaseType() *ReleaseUpsertOne {
	return u.Update(func(s *ReleaseUpsert) {
		s.UpdateReleaseType()
	})
}

// ClearReleaseType clears the value of the "release_type" field.
func (u *ReleaseUpsertOne) ClearReleaseType() *ReleaseUpsertOne {
	return u.Update(func(s *ReleaseUpsert) {
		s.ClearReleaseType()
	})
}

// SetVersion sets the "version" field.
func (u *ReleaseUpsertOne) SetVersion(v string) *ReleaseUpsertOne {
	return u.Update(func(s *ReleaseUpsert) {
		s.SetVersion(v)
	})
}

// UpdateVersion sets the "version" field to the value that was provided on create.
func (u *ReleaseUpsertOne) UpdateVersion() *ReleaseUpsertOne {
	return u.Update(func(s *ReleaseUpsert) {
		s.UpdateVersion()
	})
}

// ClearVersion clears the value of the "version" field.
func (u *ReleaseUpsertOne) ClearVersion() *ReleaseUpsertOne {
	return u.Update(func(s *ReleaseUpsert) {
		s.ClearVersion()
	})
}

// SetChannel sets the "channel" field.
func (u *ReleaseUpsertOne) SetChannel(v string) *ReleaseUpsertOne {
	return u.Update(func(s *ReleaseUpsert) {
		s.SetChannel(v)
	})
}

// UpdateChannel sets the "channel" field to the value that was provided on create.
func (u *ReleaseUpsertOne) UpdateChannel() *ReleaseUpsertOne {
	return u.Update(func(s *ReleaseUpsert) {
		s.UpdateChannel()
	})
}

// ClearChannel clears the value of the "channel" field.
func (u *ReleaseUpsertOne) ClearChannel() *ReleaseUpsertOne {
	return u.Update(func(s *ReleaseUpsert) {
		s.ClearChannel()
	})
}

// SetSummary sets the "summary" field.
func (u *ReleaseUpsertOne) SetSummary(v string) *ReleaseUpsertOne {
	return u.Update(func(s *ReleaseUpsert) {
		s.SetSummary(v)
	})
}

// UpdateSummary sets the "summary" field to the value that was provided on create.
func (u *ReleaseUpsertOne) UpdateSummary() *ReleaseUpsertOne {
	return u.Update(func(s *ReleaseUpsert) {
		s.UpdateSummary()
	})
}

// ClearSummary clears the value of the "summary" field.
func (u *ReleaseUpsertOne) ClearSummary() *ReleaseUpsertOne {
	return u.Update(func(s *ReleaseUpsert) {
		s.ClearSummary()
	})
}

// SetReleaseNotes sets the "release_notes" field.
func (u *ReleaseUpsertOne) SetReleaseNotes(v string) *ReleaseUpsertOne {
	return u.Update(func(s *ReleaseUpsert) {
		s.SetReleaseNotes(v)
	})
}

// UpdateReleaseNotes sets the "release_notes" field to the value that was provided on create.
func (u *ReleaseUpsertOne) UpdateReleaseNotes() *ReleaseUpsertOne {
	return u.Update(func(s *ReleaseUpsert) {
		s.UpdateReleaseNotes()
	})
}

// ClearReleaseNotes clears the value of the "release_notes" field.
func (u *ReleaseUpsertOne) ClearReleaseNotes() *ReleaseUpsertOne {
	return u.Update(func(s *ReleaseUpsert) {
		s.ClearReleaseNotes()
	})
}

// SetFileURL sets the "file_url" field.
func (u *ReleaseUpsertOne) SetFileURL(v string) *ReleaseUpsertOne {
	return u.Update(func(s *ReleaseUpsert) {
		s.SetFileURL(v)
	})
}

// UpdateFileURL sets the "file_url" field to the value that was provided on create.
func (u *ReleaseUpsertOne) UpdateFileURL() *ReleaseUpsertOne {
	return u.Update(func(s *ReleaseUpsert) {
		s.UpdateFileURL()
	})
}

// ClearFileURL clears the value of the "file_url" field.
func (u *ReleaseUpsertOne) ClearFileURL() *ReleaseUpsertOne {
	return u.Update(func(s *ReleaseUpsert) {
		s.ClearFileURL()
	})
}

// SetChecksum sets the "checksum" field.
func (u *ReleaseUpsertOne) SetChecksum(v string) *ReleaseUpsertOne {
	return u.Update(func(s *ReleaseUpsert) {
		s.SetChecksum(v)
	})
}

// UpdateChecksum sets the "checksum" field to the value that was provided on create.
func (u *ReleaseUpsertOne) UpdateChecksum() *ReleaseUpsertOne {
	return u.Update(func(s *ReleaseUpsert) {
		s.UpdateChecksum()
	})
}

// ClearChecksum clears the value of the "checksum" field.
func (u *ReleaseUpsertOne) ClearChecksum() *ReleaseUpsertOne {
	return u.Update(func(s *ReleaseUpsert) {
		s.ClearChecksum()
	})
}

// SetIsCritical sets the "is_critical" field.
func (u *ReleaseUpsertOne) SetIsCritical(v bool) *ReleaseUpsertOne {
	return u.Update(func(s *ReleaseUpsert) {
		s.SetIsCritical(v)
	})
}

// UpdateIsCritical sets the "is_critical" field to the value that was provided on create.
func (u *ReleaseUpsertOne) UpdateIsCritical() *ReleaseUpsertOne {
	return u.Update(func(s *ReleaseUpsert) {
		s.UpdateIsCritical()
	})
}

// ClearIsCritical clears the value of the "is_critical" field.
func (u *ReleaseUpsertOne) ClearIsCritical() *ReleaseUpsertOne {
	return u.Update(func(s *ReleaseUpsert) {
		s.ClearIsCritical()
	})
}

// SetReleaseDate sets the "release_date" field.
func (u *ReleaseUpsertOne) SetReleaseDate(v time.Time) *ReleaseUpsertOne {
	return u.Update(func(s *ReleaseUpsert) {
		s.SetReleaseDate(v)
	})
}

// UpdateReleaseDate sets the "release_date" field to the value that was provided on create.
func (u *ReleaseUpsertOne) UpdateReleaseDate() *ReleaseUpsertOne {
	return u.Update(func(s *ReleaseUpsert) {
		s.UpdateReleaseDate()
	})
}

// ClearReleaseDate clears the value of the "release_date" field.
func (u *ReleaseUpsertOne) ClearReleaseDate() *ReleaseUpsertOne {
	return u.Update(func(s *ReleaseUpsert) {
		s.ClearReleaseDate()
	})
}

// SetOs sets the "os" field.
func (u *ReleaseUpsertOne) SetOs(v string) *ReleaseUpsertOne {
	return u.Update(func(s *ReleaseUpsert) {
		s.SetOs(v)
	})
}

// UpdateOs sets the "os" field to the value that was provided on create.
func (u *ReleaseUpsertOne) UpdateOs() *ReleaseUpsertOne {
	return u.Update(func(s *ReleaseUpsert) {
		s.UpdateOs()
	})
}

// ClearOs clears the value of the "os" field.
func (u *ReleaseUpsertOne) ClearOs() *ReleaseUpsertOne {
	return u.Update(func(s *ReleaseUpsert) {
		s.ClearOs()
	})
}

// SetArch sets the "arch" field.
func (u *ReleaseUpsertOne) SetArch(v string) *ReleaseUpsertOne {
	return u.Update(func(s *ReleaseUpsert) {
		s.SetArch(v)
	})
}

// UpdateArch sets the "arch" field to the value that was provided on create.
func (u *ReleaseUpsertOne) UpdateArch() *ReleaseUpsertOne {
	return u.Update(func(s *ReleaseUpsert) {
		s.UpdateArch()
	})
}

// ClearArch clears the value of the "arch" field.
func (u *ReleaseUpsertOne) ClearArch() *ReleaseUpsertOne {
	return u.Update(func(s *ReleaseUpsert) {
		s.ClearArch()
	})
}

// Exec executes the query.
func (u *ReleaseUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("openuem_ent: missing options for ReleaseCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ReleaseUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ReleaseUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ReleaseUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ReleaseCreateBulk is the builder for creating many Release entities in bulk.
type ReleaseCreateBulk struct {
	config
	err      error
	builders []*ReleaseCreate
	conflict []sql.ConflictOption
}

// Save creates the Release entities in the database.
func (rcb *ReleaseCreateBulk) Save(ctx context.Context) ([]*Release, error) {
	if rcb.err != nil {
		return nil, rcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Release, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ReleaseMutation)
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
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = rcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *ReleaseCreateBulk) SaveX(ctx context.Context) []*Release {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *ReleaseCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *ReleaseCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Release.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ReleaseUpsert) {
//			SetReleaseType(v+v).
//		}).
//		Exec(ctx)
func (rcb *ReleaseCreateBulk) OnConflict(opts ...sql.ConflictOption) *ReleaseUpsertBulk {
	rcb.conflict = opts
	return &ReleaseUpsertBulk{
		create: rcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Release.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (rcb *ReleaseCreateBulk) OnConflictColumns(columns ...string) *ReleaseUpsertBulk {
	rcb.conflict = append(rcb.conflict, sql.ConflictColumns(columns...))
	return &ReleaseUpsertBulk{
		create: rcb,
	}
}

// ReleaseUpsertBulk is the builder for "upsert"-ing
// a bulk of Release nodes.
type ReleaseUpsertBulk struct {
	create *ReleaseCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Release.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *ReleaseUpsertBulk) UpdateNewValues() *ReleaseUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Release.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ReleaseUpsertBulk) Ignore() *ReleaseUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ReleaseUpsertBulk) DoNothing() *ReleaseUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ReleaseCreateBulk.OnConflict
// documentation for more info.
func (u *ReleaseUpsertBulk) Update(set func(*ReleaseUpsert)) *ReleaseUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ReleaseUpsert{UpdateSet: update})
	}))
	return u
}

// SetReleaseType sets the "release_type" field.
func (u *ReleaseUpsertBulk) SetReleaseType(v release.ReleaseType) *ReleaseUpsertBulk {
	return u.Update(func(s *ReleaseUpsert) {
		s.SetReleaseType(v)
	})
}

// UpdateReleaseType sets the "release_type" field to the value that was provided on create.
func (u *ReleaseUpsertBulk) UpdateReleaseType() *ReleaseUpsertBulk {
	return u.Update(func(s *ReleaseUpsert) {
		s.UpdateReleaseType()
	})
}

// ClearReleaseType clears the value of the "release_type" field.
func (u *ReleaseUpsertBulk) ClearReleaseType() *ReleaseUpsertBulk {
	return u.Update(func(s *ReleaseUpsert) {
		s.ClearReleaseType()
	})
}

// SetVersion sets the "version" field.
func (u *ReleaseUpsertBulk) SetVersion(v string) *ReleaseUpsertBulk {
	return u.Update(func(s *ReleaseUpsert) {
		s.SetVersion(v)
	})
}

// UpdateVersion sets the "version" field to the value that was provided on create.
func (u *ReleaseUpsertBulk) UpdateVersion() *ReleaseUpsertBulk {
	return u.Update(func(s *ReleaseUpsert) {
		s.UpdateVersion()
	})
}

// ClearVersion clears the value of the "version" field.
func (u *ReleaseUpsertBulk) ClearVersion() *ReleaseUpsertBulk {
	return u.Update(func(s *ReleaseUpsert) {
		s.ClearVersion()
	})
}

// SetChannel sets the "channel" field.
func (u *ReleaseUpsertBulk) SetChannel(v string) *ReleaseUpsertBulk {
	return u.Update(func(s *ReleaseUpsert) {
		s.SetChannel(v)
	})
}

// UpdateChannel sets the "channel" field to the value that was provided on create.
func (u *ReleaseUpsertBulk) UpdateChannel() *ReleaseUpsertBulk {
	return u.Update(func(s *ReleaseUpsert) {
		s.UpdateChannel()
	})
}

// ClearChannel clears the value of the "channel" field.
func (u *ReleaseUpsertBulk) ClearChannel() *ReleaseUpsertBulk {
	return u.Update(func(s *ReleaseUpsert) {
		s.ClearChannel()
	})
}

// SetSummary sets the "summary" field.
func (u *ReleaseUpsertBulk) SetSummary(v string) *ReleaseUpsertBulk {
	return u.Update(func(s *ReleaseUpsert) {
		s.SetSummary(v)
	})
}

// UpdateSummary sets the "summary" field to the value that was provided on create.
func (u *ReleaseUpsertBulk) UpdateSummary() *ReleaseUpsertBulk {
	return u.Update(func(s *ReleaseUpsert) {
		s.UpdateSummary()
	})
}

// ClearSummary clears the value of the "summary" field.
func (u *ReleaseUpsertBulk) ClearSummary() *ReleaseUpsertBulk {
	return u.Update(func(s *ReleaseUpsert) {
		s.ClearSummary()
	})
}

// SetReleaseNotes sets the "release_notes" field.
func (u *ReleaseUpsertBulk) SetReleaseNotes(v string) *ReleaseUpsertBulk {
	return u.Update(func(s *ReleaseUpsert) {
		s.SetReleaseNotes(v)
	})
}

// UpdateReleaseNotes sets the "release_notes" field to the value that was provided on create.
func (u *ReleaseUpsertBulk) UpdateReleaseNotes() *ReleaseUpsertBulk {
	return u.Update(func(s *ReleaseUpsert) {
		s.UpdateReleaseNotes()
	})
}

// ClearReleaseNotes clears the value of the "release_notes" field.
func (u *ReleaseUpsertBulk) ClearReleaseNotes() *ReleaseUpsertBulk {
	return u.Update(func(s *ReleaseUpsert) {
		s.ClearReleaseNotes()
	})
}

// SetFileURL sets the "file_url" field.
func (u *ReleaseUpsertBulk) SetFileURL(v string) *ReleaseUpsertBulk {
	return u.Update(func(s *ReleaseUpsert) {
		s.SetFileURL(v)
	})
}

// UpdateFileURL sets the "file_url" field to the value that was provided on create.
func (u *ReleaseUpsertBulk) UpdateFileURL() *ReleaseUpsertBulk {
	return u.Update(func(s *ReleaseUpsert) {
		s.UpdateFileURL()
	})
}

// ClearFileURL clears the value of the "file_url" field.
func (u *ReleaseUpsertBulk) ClearFileURL() *ReleaseUpsertBulk {
	return u.Update(func(s *ReleaseUpsert) {
		s.ClearFileURL()
	})
}

// SetChecksum sets the "checksum" field.
func (u *ReleaseUpsertBulk) SetChecksum(v string) *ReleaseUpsertBulk {
	return u.Update(func(s *ReleaseUpsert) {
		s.SetChecksum(v)
	})
}

// UpdateChecksum sets the "checksum" field to the value that was provided on create.
func (u *ReleaseUpsertBulk) UpdateChecksum() *ReleaseUpsertBulk {
	return u.Update(func(s *ReleaseUpsert) {
		s.UpdateChecksum()
	})
}

// ClearChecksum clears the value of the "checksum" field.
func (u *ReleaseUpsertBulk) ClearChecksum() *ReleaseUpsertBulk {
	return u.Update(func(s *ReleaseUpsert) {
		s.ClearChecksum()
	})
}

// SetIsCritical sets the "is_critical" field.
func (u *ReleaseUpsertBulk) SetIsCritical(v bool) *ReleaseUpsertBulk {
	return u.Update(func(s *ReleaseUpsert) {
		s.SetIsCritical(v)
	})
}

// UpdateIsCritical sets the "is_critical" field to the value that was provided on create.
func (u *ReleaseUpsertBulk) UpdateIsCritical() *ReleaseUpsertBulk {
	return u.Update(func(s *ReleaseUpsert) {
		s.UpdateIsCritical()
	})
}

// ClearIsCritical clears the value of the "is_critical" field.
func (u *ReleaseUpsertBulk) ClearIsCritical() *ReleaseUpsertBulk {
	return u.Update(func(s *ReleaseUpsert) {
		s.ClearIsCritical()
	})
}

// SetReleaseDate sets the "release_date" field.
func (u *ReleaseUpsertBulk) SetReleaseDate(v time.Time) *ReleaseUpsertBulk {
	return u.Update(func(s *ReleaseUpsert) {
		s.SetReleaseDate(v)
	})
}

// UpdateReleaseDate sets the "release_date" field to the value that was provided on create.
func (u *ReleaseUpsertBulk) UpdateReleaseDate() *ReleaseUpsertBulk {
	return u.Update(func(s *ReleaseUpsert) {
		s.UpdateReleaseDate()
	})
}

// ClearReleaseDate clears the value of the "release_date" field.
func (u *ReleaseUpsertBulk) ClearReleaseDate() *ReleaseUpsertBulk {
	return u.Update(func(s *ReleaseUpsert) {
		s.ClearReleaseDate()
	})
}

// SetOs sets the "os" field.
func (u *ReleaseUpsertBulk) SetOs(v string) *ReleaseUpsertBulk {
	return u.Update(func(s *ReleaseUpsert) {
		s.SetOs(v)
	})
}

// UpdateOs sets the "os" field to the value that was provided on create.
func (u *ReleaseUpsertBulk) UpdateOs() *ReleaseUpsertBulk {
	return u.Update(func(s *ReleaseUpsert) {
		s.UpdateOs()
	})
}

// ClearOs clears the value of the "os" field.
func (u *ReleaseUpsertBulk) ClearOs() *ReleaseUpsertBulk {
	return u.Update(func(s *ReleaseUpsert) {
		s.ClearOs()
	})
}

// SetArch sets the "arch" field.
func (u *ReleaseUpsertBulk) SetArch(v string) *ReleaseUpsertBulk {
	return u.Update(func(s *ReleaseUpsert) {
		s.SetArch(v)
	})
}

// UpdateArch sets the "arch" field to the value that was provided on create.
func (u *ReleaseUpsertBulk) UpdateArch() *ReleaseUpsertBulk {
	return u.Update(func(s *ReleaseUpsert) {
		s.UpdateArch()
	})
}

// ClearArch clears the value of the "arch" field.
func (u *ReleaseUpsertBulk) ClearArch() *ReleaseUpsertBulk {
	return u.Update(func(s *ReleaseUpsert) {
		s.ClearArch()
	})
}

// Exec executes the query.
func (u *ReleaseUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("openuem_ent: OnConflict was set for builder %d. Set it on the ReleaseCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("openuem_ent: missing options for ReleaseCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ReleaseUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
