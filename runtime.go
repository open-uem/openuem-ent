// Code generated by ent, DO NOT EDIT.

package openuem_ent

import (
	"time"

	"github.com/doncicuto/openuem_ent/agent"
	"github.com/doncicuto/openuem_ent/logicaldisk"
	"github.com/doncicuto/openuem_ent/revocation"
	"github.com/doncicuto/openuem_ent/schema"
	"github.com/doncicuto/openuem_ent/sessions"
	"github.com/doncicuto/openuem_ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	agentFields := schema.Agent{}.Fields()
	_ = agentFields
	// agentDescOs is the schema descriptor for os field.
	agentDescOs := agentFields[1].Descriptor()
	// agent.OsValidator is a validator for the "os" field. It is called by the builders before save.
	agent.OsValidator = agentDescOs.Validators[0].(func(string) error)
	// agentDescHostname is the schema descriptor for hostname field.
	agentDescHostname := agentFields[2].Descriptor()
	// agent.HostnameValidator is a validator for the "hostname" field. It is called by the builders before save.
	agent.HostnameValidator = agentDescHostname.Validators[0].(func(string) error)
	// agentDescVersion is the schema descriptor for version field.
	agentDescVersion := agentFields[3].Descriptor()
	// agent.VersionValidator is a validator for the "version" field. It is called by the builders before save.
	agent.VersionValidator = agentDescVersion.Validators[0].(func(string) error)
	// agentDescIP is the schema descriptor for ip field.
	agentDescIP := agentFields[4].Descriptor()
	// agent.DefaultIP holds the default value on creation for the ip field.
	agent.DefaultIP = agentDescIP.Default.(string)
	// agentDescEnabled is the schema descriptor for enabled field.
	agentDescEnabled := agentFields[7].Descriptor()
	// agent.DefaultEnabled holds the default value on creation for the enabled field.
	agent.DefaultEnabled = agentDescEnabled.Default.(bool)
	// agentDescID is the schema descriptor for id field.
	agentDescID := agentFields[0].Descriptor()
	// agent.IDValidator is a validator for the "id" field. It is called by the builders before save.
	agent.IDValidator = agentDescID.Validators[0].(func(string) error)
	logicaldiskFields := schema.LogicalDisk{}.Fields()
	_ = logicaldiskFields
	// logicaldiskDescUsage is the schema descriptor for usage field.
	logicaldiskDescUsage := logicaldiskFields[2].Descriptor()
	// logicaldisk.DefaultUsage holds the default value on creation for the usage field.
	logicaldisk.DefaultUsage = logicaldiskDescUsage.Default.(int8)
	revocationFields := schema.Revocation{}.Fields()
	_ = revocationFields
	// revocationDescRevoked is the schema descriptor for revoked field.
	revocationDescRevoked := revocationFields[2].Descriptor()
	// revocation.DefaultRevoked holds the default value on creation for the revoked field.
	revocation.DefaultRevoked = revocationDescRevoked.Default.(func() time.Time)
	sessionsFields := schema.Sessions{}.Fields()
	_ = sessionsFields
	// sessionsDescData is the schema descriptor for data field.
	sessionsDescData := sessionsFields[1].Descriptor()
	// sessions.DataValidator is a validator for the "data" field. It is called by the builders before save.
	sessions.DataValidator = sessionsDescData.Validators[0].(func([]byte) error)
	// sessionsDescID is the schema descriptor for id field.
	sessionsDescID := sessionsFields[0].Descriptor()
	// sessions.IDValidator is a validator for the "id" field. It is called by the builders before save.
	sessions.IDValidator = sessionsDescID.Validators[0].(func(string) error)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreated is the schema descriptor for created field.
	userDescCreated := userFields[4].Descriptor()
	// user.DefaultCreated holds the default value on creation for the created field.
	user.DefaultCreated = userDescCreated.Default.(func() time.Time)
	// userDescModified is the schema descriptor for modified field.
	userDescModified := userFields[5].Descriptor()
	// user.DefaultModified holds the default value on creation for the modified field.
	user.DefaultModified = userDescModified.Default.(func() time.Time)
	// user.UpdateDefaultModified holds the default value on update for the modified field.
	user.UpdateDefaultModified = userDescModified.UpdateDefault.(func() time.Time)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.IDValidator is a validator for the "id" field. It is called by the builders before save.
	user.IDValidator = userDescID.Validators[0].(func(string) error)
}
