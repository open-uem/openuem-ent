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
	"github.com/doncicuto/openuem_ent/predicate"
	"github.com/doncicuto/openuem_ent/settings"
)

// SettingsUpdate is the builder for updating Settings entities.
type SettingsUpdate struct {
	config
	hooks     []Hook
	mutation  *SettingsMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the SettingsUpdate builder.
func (su *SettingsUpdate) Where(ps ...predicate.Settings) *SettingsUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetLanguage sets the "language" field.
func (su *SettingsUpdate) SetLanguage(s string) *SettingsUpdate {
	su.mutation.SetLanguage(s)
	return su
}

// SetNillableLanguage sets the "language" field if the given value is not nil.
func (su *SettingsUpdate) SetNillableLanguage(s *string) *SettingsUpdate {
	if s != nil {
		su.SetLanguage(*s)
	}
	return su
}

// ClearLanguage clears the value of the "language" field.
func (su *SettingsUpdate) ClearLanguage() *SettingsUpdate {
	su.mutation.ClearLanguage()
	return su
}

// SetOrganization sets the "organization" field.
func (su *SettingsUpdate) SetOrganization(s string) *SettingsUpdate {
	su.mutation.SetOrganization(s)
	return su
}

// SetNillableOrganization sets the "organization" field if the given value is not nil.
func (su *SettingsUpdate) SetNillableOrganization(s *string) *SettingsUpdate {
	if s != nil {
		su.SetOrganization(*s)
	}
	return su
}

// ClearOrganization clears the value of the "organization" field.
func (su *SettingsUpdate) ClearOrganization() *SettingsUpdate {
	su.mutation.ClearOrganization()
	return su
}

// SetPostalAddress sets the "postal_address" field.
func (su *SettingsUpdate) SetPostalAddress(s string) *SettingsUpdate {
	su.mutation.SetPostalAddress(s)
	return su
}

// SetNillablePostalAddress sets the "postal_address" field if the given value is not nil.
func (su *SettingsUpdate) SetNillablePostalAddress(s *string) *SettingsUpdate {
	if s != nil {
		su.SetPostalAddress(*s)
	}
	return su
}

// ClearPostalAddress clears the value of the "postal_address" field.
func (su *SettingsUpdate) ClearPostalAddress() *SettingsUpdate {
	su.mutation.ClearPostalAddress()
	return su
}

// SetPostalCode sets the "postal_code" field.
func (su *SettingsUpdate) SetPostalCode(s string) *SettingsUpdate {
	su.mutation.SetPostalCode(s)
	return su
}

// SetNillablePostalCode sets the "postal_code" field if the given value is not nil.
func (su *SettingsUpdate) SetNillablePostalCode(s *string) *SettingsUpdate {
	if s != nil {
		su.SetPostalCode(*s)
	}
	return su
}

// ClearPostalCode clears the value of the "postal_code" field.
func (su *SettingsUpdate) ClearPostalCode() *SettingsUpdate {
	su.mutation.ClearPostalCode()
	return su
}

// SetLocality sets the "locality" field.
func (su *SettingsUpdate) SetLocality(s string) *SettingsUpdate {
	su.mutation.SetLocality(s)
	return su
}

// SetNillableLocality sets the "locality" field if the given value is not nil.
func (su *SettingsUpdate) SetNillableLocality(s *string) *SettingsUpdate {
	if s != nil {
		su.SetLocality(*s)
	}
	return su
}

// ClearLocality clears the value of the "locality" field.
func (su *SettingsUpdate) ClearLocality() *SettingsUpdate {
	su.mutation.ClearLocality()
	return su
}

// SetProvince sets the "province" field.
func (su *SettingsUpdate) SetProvince(s string) *SettingsUpdate {
	su.mutation.SetProvince(s)
	return su
}

// SetNillableProvince sets the "province" field if the given value is not nil.
func (su *SettingsUpdate) SetNillableProvince(s *string) *SettingsUpdate {
	if s != nil {
		su.SetProvince(*s)
	}
	return su
}

// ClearProvince clears the value of the "province" field.
func (su *SettingsUpdate) ClearProvince() *SettingsUpdate {
	su.mutation.ClearProvince()
	return su
}

// SetState sets the "state" field.
func (su *SettingsUpdate) SetState(s string) *SettingsUpdate {
	su.mutation.SetState(s)
	return su
}

// SetNillableState sets the "state" field if the given value is not nil.
func (su *SettingsUpdate) SetNillableState(s *string) *SettingsUpdate {
	if s != nil {
		su.SetState(*s)
	}
	return su
}

// ClearState clears the value of the "state" field.
func (su *SettingsUpdate) ClearState() *SettingsUpdate {
	su.mutation.ClearState()
	return su
}

// SetCountry sets the "country" field.
func (su *SettingsUpdate) SetCountry(s string) *SettingsUpdate {
	su.mutation.SetCountry(s)
	return su
}

// SetNillableCountry sets the "country" field if the given value is not nil.
func (su *SettingsUpdate) SetNillableCountry(s *string) *SettingsUpdate {
	if s != nil {
		su.SetCountry(*s)
	}
	return su
}

// ClearCountry clears the value of the "country" field.
func (su *SettingsUpdate) ClearCountry() *SettingsUpdate {
	su.mutation.ClearCountry()
	return su
}

// SetSMTPServer sets the "smtp_server" field.
func (su *SettingsUpdate) SetSMTPServer(s string) *SettingsUpdate {
	su.mutation.SetSMTPServer(s)
	return su
}

// SetNillableSMTPServer sets the "smtp_server" field if the given value is not nil.
func (su *SettingsUpdate) SetNillableSMTPServer(s *string) *SettingsUpdate {
	if s != nil {
		su.SetSMTPServer(*s)
	}
	return su
}

// ClearSMTPServer clears the value of the "smtp_server" field.
func (su *SettingsUpdate) ClearSMTPServer() *SettingsUpdate {
	su.mutation.ClearSMTPServer()
	return su
}

// SetSMTPPort sets the "smtp_port" field.
func (su *SettingsUpdate) SetSMTPPort(i int) *SettingsUpdate {
	su.mutation.ResetSMTPPort()
	su.mutation.SetSMTPPort(i)
	return su
}

// SetNillableSMTPPort sets the "smtp_port" field if the given value is not nil.
func (su *SettingsUpdate) SetNillableSMTPPort(i *int) *SettingsUpdate {
	if i != nil {
		su.SetSMTPPort(*i)
	}
	return su
}

// AddSMTPPort adds i to the "smtp_port" field.
func (su *SettingsUpdate) AddSMTPPort(i int) *SettingsUpdate {
	su.mutation.AddSMTPPort(i)
	return su
}

// ClearSMTPPort clears the value of the "smtp_port" field.
func (su *SettingsUpdate) ClearSMTPPort() *SettingsUpdate {
	su.mutation.ClearSMTPPort()
	return su
}

// SetSMTPUser sets the "smtp_user" field.
func (su *SettingsUpdate) SetSMTPUser(s string) *SettingsUpdate {
	su.mutation.SetSMTPUser(s)
	return su
}

// SetNillableSMTPUser sets the "smtp_user" field if the given value is not nil.
func (su *SettingsUpdate) SetNillableSMTPUser(s *string) *SettingsUpdate {
	if s != nil {
		su.SetSMTPUser(*s)
	}
	return su
}

// ClearSMTPUser clears the value of the "smtp_user" field.
func (su *SettingsUpdate) ClearSMTPUser() *SettingsUpdate {
	su.mutation.ClearSMTPUser()
	return su
}

// SetSMTPPassword sets the "smtp_password" field.
func (su *SettingsUpdate) SetSMTPPassword(s string) *SettingsUpdate {
	su.mutation.SetSMTPPassword(s)
	return su
}

// SetNillableSMTPPassword sets the "smtp_password" field if the given value is not nil.
func (su *SettingsUpdate) SetNillableSMTPPassword(s *string) *SettingsUpdate {
	if s != nil {
		su.SetSMTPPassword(*s)
	}
	return su
}

// ClearSMTPPassword clears the value of the "smtp_password" field.
func (su *SettingsUpdate) ClearSMTPPassword() *SettingsUpdate {
	su.mutation.ClearSMTPPassword()
	return su
}

// SetSMTPAuth sets the "smtp_auth" field.
func (su *SettingsUpdate) SetSMTPAuth(s string) *SettingsUpdate {
	su.mutation.SetSMTPAuth(s)
	return su
}

// SetNillableSMTPAuth sets the "smtp_auth" field if the given value is not nil.
func (su *SettingsUpdate) SetNillableSMTPAuth(s *string) *SettingsUpdate {
	if s != nil {
		su.SetSMTPAuth(*s)
	}
	return su
}

// ClearSMTPAuth clears the value of the "smtp_auth" field.
func (su *SettingsUpdate) ClearSMTPAuth() *SettingsUpdate {
	su.mutation.ClearSMTPAuth()
	return su
}

// SetSMTPTLS sets the "smtp_tls" field.
func (su *SettingsUpdate) SetSMTPTLS(b bool) *SettingsUpdate {
	su.mutation.SetSMTPTLS(b)
	return su
}

// SetNillableSMTPTLS sets the "smtp_tls" field if the given value is not nil.
func (su *SettingsUpdate) SetNillableSMTPTLS(b *bool) *SettingsUpdate {
	if b != nil {
		su.SetSMTPTLS(*b)
	}
	return su
}

// ClearSMTPTLS clears the value of the "smtp_tls" field.
func (su *SettingsUpdate) ClearSMTPTLS() *SettingsUpdate {
	su.mutation.ClearSMTPTLS()
	return su
}

// SetSMTPStarttls sets the "smtp_starttls" field.
func (su *SettingsUpdate) SetSMTPStarttls(b bool) *SettingsUpdate {
	su.mutation.SetSMTPStarttls(b)
	return su
}

// SetNillableSMTPStarttls sets the "smtp_starttls" field if the given value is not nil.
func (su *SettingsUpdate) SetNillableSMTPStarttls(b *bool) *SettingsUpdate {
	if b != nil {
		su.SetSMTPStarttls(*b)
	}
	return su
}

// ClearSMTPStarttls clears the value of the "smtp_starttls" field.
func (su *SettingsUpdate) ClearSMTPStarttls() *SettingsUpdate {
	su.mutation.ClearSMTPStarttls()
	return su
}

// SetCreated sets the "created" field.
func (su *SettingsUpdate) SetCreated(t time.Time) *SettingsUpdate {
	su.mutation.SetCreated(t)
	return su
}

// SetNillableCreated sets the "created" field if the given value is not nil.
func (su *SettingsUpdate) SetNillableCreated(t *time.Time) *SettingsUpdate {
	if t != nil {
		su.SetCreated(*t)
	}
	return su
}

// ClearCreated clears the value of the "created" field.
func (su *SettingsUpdate) ClearCreated() *SettingsUpdate {
	su.mutation.ClearCreated()
	return su
}

// SetModified sets the "modified" field.
func (su *SettingsUpdate) SetModified(t time.Time) *SettingsUpdate {
	su.mutation.SetModified(t)
	return su
}

// ClearModified clears the value of the "modified" field.
func (su *SettingsUpdate) ClearModified() *SettingsUpdate {
	su.mutation.ClearModified()
	return su
}

// Mutation returns the SettingsMutation object of the builder.
func (su *SettingsUpdate) Mutation() *SettingsMutation {
	return su.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SettingsUpdate) Save(ctx context.Context) (int, error) {
	su.defaults()
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SettingsUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SettingsUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SettingsUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (su *SettingsUpdate) defaults() {
	if _, ok := su.mutation.Modified(); !ok && !su.mutation.ModifiedCleared() {
		v := settings.UpdateDefaultModified()
		su.mutation.SetModified(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (su *SettingsUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SettingsUpdate {
	su.modifiers = append(su.modifiers, modifiers...)
	return su
}

func (su *SettingsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(settings.Table, settings.Columns, sqlgraph.NewFieldSpec(settings.FieldID, field.TypeInt))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Language(); ok {
		_spec.SetField(settings.FieldLanguage, field.TypeString, value)
	}
	if su.mutation.LanguageCleared() {
		_spec.ClearField(settings.FieldLanguage, field.TypeString)
	}
	if value, ok := su.mutation.Organization(); ok {
		_spec.SetField(settings.FieldOrganization, field.TypeString, value)
	}
	if su.mutation.OrganizationCleared() {
		_spec.ClearField(settings.FieldOrganization, field.TypeString)
	}
	if value, ok := su.mutation.PostalAddress(); ok {
		_spec.SetField(settings.FieldPostalAddress, field.TypeString, value)
	}
	if su.mutation.PostalAddressCleared() {
		_spec.ClearField(settings.FieldPostalAddress, field.TypeString)
	}
	if value, ok := su.mutation.PostalCode(); ok {
		_spec.SetField(settings.FieldPostalCode, field.TypeString, value)
	}
	if su.mutation.PostalCodeCleared() {
		_spec.ClearField(settings.FieldPostalCode, field.TypeString)
	}
	if value, ok := su.mutation.Locality(); ok {
		_spec.SetField(settings.FieldLocality, field.TypeString, value)
	}
	if su.mutation.LocalityCleared() {
		_spec.ClearField(settings.FieldLocality, field.TypeString)
	}
	if value, ok := su.mutation.Province(); ok {
		_spec.SetField(settings.FieldProvince, field.TypeString, value)
	}
	if su.mutation.ProvinceCleared() {
		_spec.ClearField(settings.FieldProvince, field.TypeString)
	}
	if value, ok := su.mutation.State(); ok {
		_spec.SetField(settings.FieldState, field.TypeString, value)
	}
	if su.mutation.StateCleared() {
		_spec.ClearField(settings.FieldState, field.TypeString)
	}
	if value, ok := su.mutation.Country(); ok {
		_spec.SetField(settings.FieldCountry, field.TypeString, value)
	}
	if su.mutation.CountryCleared() {
		_spec.ClearField(settings.FieldCountry, field.TypeString)
	}
	if value, ok := su.mutation.SMTPServer(); ok {
		_spec.SetField(settings.FieldSMTPServer, field.TypeString, value)
	}
	if su.mutation.SMTPServerCleared() {
		_spec.ClearField(settings.FieldSMTPServer, field.TypeString)
	}
	if value, ok := su.mutation.SMTPPort(); ok {
		_spec.SetField(settings.FieldSMTPPort, field.TypeInt, value)
	}
	if value, ok := su.mutation.AddedSMTPPort(); ok {
		_spec.AddField(settings.FieldSMTPPort, field.TypeInt, value)
	}
	if su.mutation.SMTPPortCleared() {
		_spec.ClearField(settings.FieldSMTPPort, field.TypeInt)
	}
	if value, ok := su.mutation.SMTPUser(); ok {
		_spec.SetField(settings.FieldSMTPUser, field.TypeString, value)
	}
	if su.mutation.SMTPUserCleared() {
		_spec.ClearField(settings.FieldSMTPUser, field.TypeString)
	}
	if value, ok := su.mutation.SMTPPassword(); ok {
		_spec.SetField(settings.FieldSMTPPassword, field.TypeString, value)
	}
	if su.mutation.SMTPPasswordCleared() {
		_spec.ClearField(settings.FieldSMTPPassword, field.TypeString)
	}
	if value, ok := su.mutation.SMTPAuth(); ok {
		_spec.SetField(settings.FieldSMTPAuth, field.TypeString, value)
	}
	if su.mutation.SMTPAuthCleared() {
		_spec.ClearField(settings.FieldSMTPAuth, field.TypeString)
	}
	if value, ok := su.mutation.SMTPTLS(); ok {
		_spec.SetField(settings.FieldSMTPTLS, field.TypeBool, value)
	}
	if su.mutation.SMTPTLSCleared() {
		_spec.ClearField(settings.FieldSMTPTLS, field.TypeBool)
	}
	if value, ok := su.mutation.SMTPStarttls(); ok {
		_spec.SetField(settings.FieldSMTPStarttls, field.TypeBool, value)
	}
	if su.mutation.SMTPStarttlsCleared() {
		_spec.ClearField(settings.FieldSMTPStarttls, field.TypeBool)
	}
	if value, ok := su.mutation.Created(); ok {
		_spec.SetField(settings.FieldCreated, field.TypeTime, value)
	}
	if su.mutation.CreatedCleared() {
		_spec.ClearField(settings.FieldCreated, field.TypeTime)
	}
	if value, ok := su.mutation.Modified(); ok {
		_spec.SetField(settings.FieldModified, field.TypeTime, value)
	}
	if su.mutation.ModifiedCleared() {
		_spec.ClearField(settings.FieldModified, field.TypeTime)
	}
	_spec.AddModifiers(su.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{settings.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SettingsUpdateOne is the builder for updating a single Settings entity.
type SettingsUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *SettingsMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetLanguage sets the "language" field.
func (suo *SettingsUpdateOne) SetLanguage(s string) *SettingsUpdateOne {
	suo.mutation.SetLanguage(s)
	return suo
}

// SetNillableLanguage sets the "language" field if the given value is not nil.
func (suo *SettingsUpdateOne) SetNillableLanguage(s *string) *SettingsUpdateOne {
	if s != nil {
		suo.SetLanguage(*s)
	}
	return suo
}

// ClearLanguage clears the value of the "language" field.
func (suo *SettingsUpdateOne) ClearLanguage() *SettingsUpdateOne {
	suo.mutation.ClearLanguage()
	return suo
}

// SetOrganization sets the "organization" field.
func (suo *SettingsUpdateOne) SetOrganization(s string) *SettingsUpdateOne {
	suo.mutation.SetOrganization(s)
	return suo
}

// SetNillableOrganization sets the "organization" field if the given value is not nil.
func (suo *SettingsUpdateOne) SetNillableOrganization(s *string) *SettingsUpdateOne {
	if s != nil {
		suo.SetOrganization(*s)
	}
	return suo
}

// ClearOrganization clears the value of the "organization" field.
func (suo *SettingsUpdateOne) ClearOrganization() *SettingsUpdateOne {
	suo.mutation.ClearOrganization()
	return suo
}

// SetPostalAddress sets the "postal_address" field.
func (suo *SettingsUpdateOne) SetPostalAddress(s string) *SettingsUpdateOne {
	suo.mutation.SetPostalAddress(s)
	return suo
}

// SetNillablePostalAddress sets the "postal_address" field if the given value is not nil.
func (suo *SettingsUpdateOne) SetNillablePostalAddress(s *string) *SettingsUpdateOne {
	if s != nil {
		suo.SetPostalAddress(*s)
	}
	return suo
}

// ClearPostalAddress clears the value of the "postal_address" field.
func (suo *SettingsUpdateOne) ClearPostalAddress() *SettingsUpdateOne {
	suo.mutation.ClearPostalAddress()
	return suo
}

// SetPostalCode sets the "postal_code" field.
func (suo *SettingsUpdateOne) SetPostalCode(s string) *SettingsUpdateOne {
	suo.mutation.SetPostalCode(s)
	return suo
}

// SetNillablePostalCode sets the "postal_code" field if the given value is not nil.
func (suo *SettingsUpdateOne) SetNillablePostalCode(s *string) *SettingsUpdateOne {
	if s != nil {
		suo.SetPostalCode(*s)
	}
	return suo
}

// ClearPostalCode clears the value of the "postal_code" field.
func (suo *SettingsUpdateOne) ClearPostalCode() *SettingsUpdateOne {
	suo.mutation.ClearPostalCode()
	return suo
}

// SetLocality sets the "locality" field.
func (suo *SettingsUpdateOne) SetLocality(s string) *SettingsUpdateOne {
	suo.mutation.SetLocality(s)
	return suo
}

// SetNillableLocality sets the "locality" field if the given value is not nil.
func (suo *SettingsUpdateOne) SetNillableLocality(s *string) *SettingsUpdateOne {
	if s != nil {
		suo.SetLocality(*s)
	}
	return suo
}

// ClearLocality clears the value of the "locality" field.
func (suo *SettingsUpdateOne) ClearLocality() *SettingsUpdateOne {
	suo.mutation.ClearLocality()
	return suo
}

// SetProvince sets the "province" field.
func (suo *SettingsUpdateOne) SetProvince(s string) *SettingsUpdateOne {
	suo.mutation.SetProvince(s)
	return suo
}

// SetNillableProvince sets the "province" field if the given value is not nil.
func (suo *SettingsUpdateOne) SetNillableProvince(s *string) *SettingsUpdateOne {
	if s != nil {
		suo.SetProvince(*s)
	}
	return suo
}

// ClearProvince clears the value of the "province" field.
func (suo *SettingsUpdateOne) ClearProvince() *SettingsUpdateOne {
	suo.mutation.ClearProvince()
	return suo
}

// SetState sets the "state" field.
func (suo *SettingsUpdateOne) SetState(s string) *SettingsUpdateOne {
	suo.mutation.SetState(s)
	return suo
}

// SetNillableState sets the "state" field if the given value is not nil.
func (suo *SettingsUpdateOne) SetNillableState(s *string) *SettingsUpdateOne {
	if s != nil {
		suo.SetState(*s)
	}
	return suo
}

// ClearState clears the value of the "state" field.
func (suo *SettingsUpdateOne) ClearState() *SettingsUpdateOne {
	suo.mutation.ClearState()
	return suo
}

// SetCountry sets the "country" field.
func (suo *SettingsUpdateOne) SetCountry(s string) *SettingsUpdateOne {
	suo.mutation.SetCountry(s)
	return suo
}

// SetNillableCountry sets the "country" field if the given value is not nil.
func (suo *SettingsUpdateOne) SetNillableCountry(s *string) *SettingsUpdateOne {
	if s != nil {
		suo.SetCountry(*s)
	}
	return suo
}

// ClearCountry clears the value of the "country" field.
func (suo *SettingsUpdateOne) ClearCountry() *SettingsUpdateOne {
	suo.mutation.ClearCountry()
	return suo
}

// SetSMTPServer sets the "smtp_server" field.
func (suo *SettingsUpdateOne) SetSMTPServer(s string) *SettingsUpdateOne {
	suo.mutation.SetSMTPServer(s)
	return suo
}

// SetNillableSMTPServer sets the "smtp_server" field if the given value is not nil.
func (suo *SettingsUpdateOne) SetNillableSMTPServer(s *string) *SettingsUpdateOne {
	if s != nil {
		suo.SetSMTPServer(*s)
	}
	return suo
}

// ClearSMTPServer clears the value of the "smtp_server" field.
func (suo *SettingsUpdateOne) ClearSMTPServer() *SettingsUpdateOne {
	suo.mutation.ClearSMTPServer()
	return suo
}

// SetSMTPPort sets the "smtp_port" field.
func (suo *SettingsUpdateOne) SetSMTPPort(i int) *SettingsUpdateOne {
	suo.mutation.ResetSMTPPort()
	suo.mutation.SetSMTPPort(i)
	return suo
}

// SetNillableSMTPPort sets the "smtp_port" field if the given value is not nil.
func (suo *SettingsUpdateOne) SetNillableSMTPPort(i *int) *SettingsUpdateOne {
	if i != nil {
		suo.SetSMTPPort(*i)
	}
	return suo
}

// AddSMTPPort adds i to the "smtp_port" field.
func (suo *SettingsUpdateOne) AddSMTPPort(i int) *SettingsUpdateOne {
	suo.mutation.AddSMTPPort(i)
	return suo
}

// ClearSMTPPort clears the value of the "smtp_port" field.
func (suo *SettingsUpdateOne) ClearSMTPPort() *SettingsUpdateOne {
	suo.mutation.ClearSMTPPort()
	return suo
}

// SetSMTPUser sets the "smtp_user" field.
func (suo *SettingsUpdateOne) SetSMTPUser(s string) *SettingsUpdateOne {
	suo.mutation.SetSMTPUser(s)
	return suo
}

// SetNillableSMTPUser sets the "smtp_user" field if the given value is not nil.
func (suo *SettingsUpdateOne) SetNillableSMTPUser(s *string) *SettingsUpdateOne {
	if s != nil {
		suo.SetSMTPUser(*s)
	}
	return suo
}

// ClearSMTPUser clears the value of the "smtp_user" field.
func (suo *SettingsUpdateOne) ClearSMTPUser() *SettingsUpdateOne {
	suo.mutation.ClearSMTPUser()
	return suo
}

// SetSMTPPassword sets the "smtp_password" field.
func (suo *SettingsUpdateOne) SetSMTPPassword(s string) *SettingsUpdateOne {
	suo.mutation.SetSMTPPassword(s)
	return suo
}

// SetNillableSMTPPassword sets the "smtp_password" field if the given value is not nil.
func (suo *SettingsUpdateOne) SetNillableSMTPPassword(s *string) *SettingsUpdateOne {
	if s != nil {
		suo.SetSMTPPassword(*s)
	}
	return suo
}

// ClearSMTPPassword clears the value of the "smtp_password" field.
func (suo *SettingsUpdateOne) ClearSMTPPassword() *SettingsUpdateOne {
	suo.mutation.ClearSMTPPassword()
	return suo
}

// SetSMTPAuth sets the "smtp_auth" field.
func (suo *SettingsUpdateOne) SetSMTPAuth(s string) *SettingsUpdateOne {
	suo.mutation.SetSMTPAuth(s)
	return suo
}

// SetNillableSMTPAuth sets the "smtp_auth" field if the given value is not nil.
func (suo *SettingsUpdateOne) SetNillableSMTPAuth(s *string) *SettingsUpdateOne {
	if s != nil {
		suo.SetSMTPAuth(*s)
	}
	return suo
}

// ClearSMTPAuth clears the value of the "smtp_auth" field.
func (suo *SettingsUpdateOne) ClearSMTPAuth() *SettingsUpdateOne {
	suo.mutation.ClearSMTPAuth()
	return suo
}

// SetSMTPTLS sets the "smtp_tls" field.
func (suo *SettingsUpdateOne) SetSMTPTLS(b bool) *SettingsUpdateOne {
	suo.mutation.SetSMTPTLS(b)
	return suo
}

// SetNillableSMTPTLS sets the "smtp_tls" field if the given value is not nil.
func (suo *SettingsUpdateOne) SetNillableSMTPTLS(b *bool) *SettingsUpdateOne {
	if b != nil {
		suo.SetSMTPTLS(*b)
	}
	return suo
}

// ClearSMTPTLS clears the value of the "smtp_tls" field.
func (suo *SettingsUpdateOne) ClearSMTPTLS() *SettingsUpdateOne {
	suo.mutation.ClearSMTPTLS()
	return suo
}

// SetSMTPStarttls sets the "smtp_starttls" field.
func (suo *SettingsUpdateOne) SetSMTPStarttls(b bool) *SettingsUpdateOne {
	suo.mutation.SetSMTPStarttls(b)
	return suo
}

// SetNillableSMTPStarttls sets the "smtp_starttls" field if the given value is not nil.
func (suo *SettingsUpdateOne) SetNillableSMTPStarttls(b *bool) *SettingsUpdateOne {
	if b != nil {
		suo.SetSMTPStarttls(*b)
	}
	return suo
}

// ClearSMTPStarttls clears the value of the "smtp_starttls" field.
func (suo *SettingsUpdateOne) ClearSMTPStarttls() *SettingsUpdateOne {
	suo.mutation.ClearSMTPStarttls()
	return suo
}

// SetCreated sets the "created" field.
func (suo *SettingsUpdateOne) SetCreated(t time.Time) *SettingsUpdateOne {
	suo.mutation.SetCreated(t)
	return suo
}

// SetNillableCreated sets the "created" field if the given value is not nil.
func (suo *SettingsUpdateOne) SetNillableCreated(t *time.Time) *SettingsUpdateOne {
	if t != nil {
		suo.SetCreated(*t)
	}
	return suo
}

// ClearCreated clears the value of the "created" field.
func (suo *SettingsUpdateOne) ClearCreated() *SettingsUpdateOne {
	suo.mutation.ClearCreated()
	return suo
}

// SetModified sets the "modified" field.
func (suo *SettingsUpdateOne) SetModified(t time.Time) *SettingsUpdateOne {
	suo.mutation.SetModified(t)
	return suo
}

// ClearModified clears the value of the "modified" field.
func (suo *SettingsUpdateOne) ClearModified() *SettingsUpdateOne {
	suo.mutation.ClearModified()
	return suo
}

// Mutation returns the SettingsMutation object of the builder.
func (suo *SettingsUpdateOne) Mutation() *SettingsMutation {
	return suo.mutation
}

// Where appends a list predicates to the SettingsUpdate builder.
func (suo *SettingsUpdateOne) Where(ps ...predicate.Settings) *SettingsUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SettingsUpdateOne) Select(field string, fields ...string) *SettingsUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Settings entity.
func (suo *SettingsUpdateOne) Save(ctx context.Context) (*Settings, error) {
	suo.defaults()
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SettingsUpdateOne) SaveX(ctx context.Context) *Settings {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SettingsUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SettingsUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suo *SettingsUpdateOne) defaults() {
	if _, ok := suo.mutation.Modified(); !ok && !suo.mutation.ModifiedCleared() {
		v := settings.UpdateDefaultModified()
		suo.mutation.SetModified(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (suo *SettingsUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SettingsUpdateOne {
	suo.modifiers = append(suo.modifiers, modifiers...)
	return suo
}

func (suo *SettingsUpdateOne) sqlSave(ctx context.Context) (_node *Settings, err error) {
	_spec := sqlgraph.NewUpdateSpec(settings.Table, settings.Columns, sqlgraph.NewFieldSpec(settings.FieldID, field.TypeInt))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`openuem_ent: missing "Settings.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, settings.FieldID)
		for _, f := range fields {
			if !settings.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("openuem_ent: invalid field %q for query", f)}
			}
			if f != settings.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Language(); ok {
		_spec.SetField(settings.FieldLanguage, field.TypeString, value)
	}
	if suo.mutation.LanguageCleared() {
		_spec.ClearField(settings.FieldLanguage, field.TypeString)
	}
	if value, ok := suo.mutation.Organization(); ok {
		_spec.SetField(settings.FieldOrganization, field.TypeString, value)
	}
	if suo.mutation.OrganizationCleared() {
		_spec.ClearField(settings.FieldOrganization, field.TypeString)
	}
	if value, ok := suo.mutation.PostalAddress(); ok {
		_spec.SetField(settings.FieldPostalAddress, field.TypeString, value)
	}
	if suo.mutation.PostalAddressCleared() {
		_spec.ClearField(settings.FieldPostalAddress, field.TypeString)
	}
	if value, ok := suo.mutation.PostalCode(); ok {
		_spec.SetField(settings.FieldPostalCode, field.TypeString, value)
	}
	if suo.mutation.PostalCodeCleared() {
		_spec.ClearField(settings.FieldPostalCode, field.TypeString)
	}
	if value, ok := suo.mutation.Locality(); ok {
		_spec.SetField(settings.FieldLocality, field.TypeString, value)
	}
	if suo.mutation.LocalityCleared() {
		_spec.ClearField(settings.FieldLocality, field.TypeString)
	}
	if value, ok := suo.mutation.Province(); ok {
		_spec.SetField(settings.FieldProvince, field.TypeString, value)
	}
	if suo.mutation.ProvinceCleared() {
		_spec.ClearField(settings.FieldProvince, field.TypeString)
	}
	if value, ok := suo.mutation.State(); ok {
		_spec.SetField(settings.FieldState, field.TypeString, value)
	}
	if suo.mutation.StateCleared() {
		_spec.ClearField(settings.FieldState, field.TypeString)
	}
	if value, ok := suo.mutation.Country(); ok {
		_spec.SetField(settings.FieldCountry, field.TypeString, value)
	}
	if suo.mutation.CountryCleared() {
		_spec.ClearField(settings.FieldCountry, field.TypeString)
	}
	if value, ok := suo.mutation.SMTPServer(); ok {
		_spec.SetField(settings.FieldSMTPServer, field.TypeString, value)
	}
	if suo.mutation.SMTPServerCleared() {
		_spec.ClearField(settings.FieldSMTPServer, field.TypeString)
	}
	if value, ok := suo.mutation.SMTPPort(); ok {
		_spec.SetField(settings.FieldSMTPPort, field.TypeInt, value)
	}
	if value, ok := suo.mutation.AddedSMTPPort(); ok {
		_spec.AddField(settings.FieldSMTPPort, field.TypeInt, value)
	}
	if suo.mutation.SMTPPortCleared() {
		_spec.ClearField(settings.FieldSMTPPort, field.TypeInt)
	}
	if value, ok := suo.mutation.SMTPUser(); ok {
		_spec.SetField(settings.FieldSMTPUser, field.TypeString, value)
	}
	if suo.mutation.SMTPUserCleared() {
		_spec.ClearField(settings.FieldSMTPUser, field.TypeString)
	}
	if value, ok := suo.mutation.SMTPPassword(); ok {
		_spec.SetField(settings.FieldSMTPPassword, field.TypeString, value)
	}
	if suo.mutation.SMTPPasswordCleared() {
		_spec.ClearField(settings.FieldSMTPPassword, field.TypeString)
	}
	if value, ok := suo.mutation.SMTPAuth(); ok {
		_spec.SetField(settings.FieldSMTPAuth, field.TypeString, value)
	}
	if suo.mutation.SMTPAuthCleared() {
		_spec.ClearField(settings.FieldSMTPAuth, field.TypeString)
	}
	if value, ok := suo.mutation.SMTPTLS(); ok {
		_spec.SetField(settings.FieldSMTPTLS, field.TypeBool, value)
	}
	if suo.mutation.SMTPTLSCleared() {
		_spec.ClearField(settings.FieldSMTPTLS, field.TypeBool)
	}
	if value, ok := suo.mutation.SMTPStarttls(); ok {
		_spec.SetField(settings.FieldSMTPStarttls, field.TypeBool, value)
	}
	if suo.mutation.SMTPStarttlsCleared() {
		_spec.ClearField(settings.FieldSMTPStarttls, field.TypeBool)
	}
	if value, ok := suo.mutation.Created(); ok {
		_spec.SetField(settings.FieldCreated, field.TypeTime, value)
	}
	if suo.mutation.CreatedCleared() {
		_spec.ClearField(settings.FieldCreated, field.TypeTime)
	}
	if value, ok := suo.mutation.Modified(); ok {
		_spec.SetField(settings.FieldModified, field.TypeTime, value)
	}
	if suo.mutation.ModifiedCleared() {
		_spec.ClearField(settings.FieldModified, field.TypeTime)
	}
	_spec.AddModifiers(suo.modifiers...)
	_node = &Settings{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{settings.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
