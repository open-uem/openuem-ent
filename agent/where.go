// Code generated by ent, DO NOT EDIT.

package agent

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/doncicuto/openuem_ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Agent {
	return predicate.Agent(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Agent {
	return predicate.Agent(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Agent {
	return predicate.Agent(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Agent {
	return predicate.Agent(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Agent {
	return predicate.Agent(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Agent {
	return predicate.Agent(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Agent {
	return predicate.Agent(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Agent {
	return predicate.Agent(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Agent {
	return predicate.Agent(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Agent {
	return predicate.Agent(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Agent {
	return predicate.Agent(sql.FieldContainsFold(FieldID, id))
}

// Os applies equality check predicate on the "os" field. It's identical to OsEQ.
func Os(v string) predicate.Agent {
	return predicate.Agent(sql.FieldEQ(FieldOs, v))
}

// Hostname applies equality check predicate on the "hostname" field. It's identical to HostnameEQ.
func Hostname(v string) predicate.Agent {
	return predicate.Agent(sql.FieldEQ(FieldHostname, v))
}

// Version applies equality check predicate on the "version" field. It's identical to VersionEQ.
func Version(v string) predicate.Agent {
	return predicate.Agent(sql.FieldEQ(FieldVersion, v))
}

// IP applies equality check predicate on the "ip" field. It's identical to IPEQ.
func IP(v string) predicate.Agent {
	return predicate.Agent(sql.FieldEQ(FieldIP, v))
}

// MAC applies equality check predicate on the "mac" field. It's identical to MACEQ.
func MAC(v string) predicate.Agent {
	return predicate.Agent(sql.FieldEQ(FieldMAC, v))
}

// FirstContact applies equality check predicate on the "first_contact" field. It's identical to FirstContactEQ.
func FirstContact(v time.Time) predicate.Agent {
	return predicate.Agent(sql.FieldEQ(FieldFirstContact, v))
}

// LastContact applies equality check predicate on the "last_contact" field. It's identical to LastContactEQ.
func LastContact(v time.Time) predicate.Agent {
	return predicate.Agent(sql.FieldEQ(FieldLastContact, v))
}

// Enabled applies equality check predicate on the "enabled" field. It's identical to EnabledEQ.
func Enabled(v bool) predicate.Agent {
	return predicate.Agent(sql.FieldEQ(FieldEnabled, v))
}

// Vnc applies equality check predicate on the "vnc" field. It's identical to VncEQ.
func Vnc(v string) predicate.Agent {
	return predicate.Agent(sql.FieldEQ(FieldVnc, v))
}

// Notes applies equality check predicate on the "notes" field. It's identical to NotesEQ.
func Notes(v string) predicate.Agent {
	return predicate.Agent(sql.FieldEQ(FieldNotes, v))
}

// OsEQ applies the EQ predicate on the "os" field.
func OsEQ(v string) predicate.Agent {
	return predicate.Agent(sql.FieldEQ(FieldOs, v))
}

// OsNEQ applies the NEQ predicate on the "os" field.
func OsNEQ(v string) predicate.Agent {
	return predicate.Agent(sql.FieldNEQ(FieldOs, v))
}

// OsIn applies the In predicate on the "os" field.
func OsIn(vs ...string) predicate.Agent {
	return predicate.Agent(sql.FieldIn(FieldOs, vs...))
}

// OsNotIn applies the NotIn predicate on the "os" field.
func OsNotIn(vs ...string) predicate.Agent {
	return predicate.Agent(sql.FieldNotIn(FieldOs, vs...))
}

// OsGT applies the GT predicate on the "os" field.
func OsGT(v string) predicate.Agent {
	return predicate.Agent(sql.FieldGT(FieldOs, v))
}

// OsGTE applies the GTE predicate on the "os" field.
func OsGTE(v string) predicate.Agent {
	return predicate.Agent(sql.FieldGTE(FieldOs, v))
}

// OsLT applies the LT predicate on the "os" field.
func OsLT(v string) predicate.Agent {
	return predicate.Agent(sql.FieldLT(FieldOs, v))
}

// OsLTE applies the LTE predicate on the "os" field.
func OsLTE(v string) predicate.Agent {
	return predicate.Agent(sql.FieldLTE(FieldOs, v))
}

// OsContains applies the Contains predicate on the "os" field.
func OsContains(v string) predicate.Agent {
	return predicate.Agent(sql.FieldContains(FieldOs, v))
}

// OsHasPrefix applies the HasPrefix predicate on the "os" field.
func OsHasPrefix(v string) predicate.Agent {
	return predicate.Agent(sql.FieldHasPrefix(FieldOs, v))
}

// OsHasSuffix applies the HasSuffix predicate on the "os" field.
func OsHasSuffix(v string) predicate.Agent {
	return predicate.Agent(sql.FieldHasSuffix(FieldOs, v))
}

// OsEqualFold applies the EqualFold predicate on the "os" field.
func OsEqualFold(v string) predicate.Agent {
	return predicate.Agent(sql.FieldEqualFold(FieldOs, v))
}

// OsContainsFold applies the ContainsFold predicate on the "os" field.
func OsContainsFold(v string) predicate.Agent {
	return predicate.Agent(sql.FieldContainsFold(FieldOs, v))
}

// HostnameEQ applies the EQ predicate on the "hostname" field.
func HostnameEQ(v string) predicate.Agent {
	return predicate.Agent(sql.FieldEQ(FieldHostname, v))
}

// HostnameNEQ applies the NEQ predicate on the "hostname" field.
func HostnameNEQ(v string) predicate.Agent {
	return predicate.Agent(sql.FieldNEQ(FieldHostname, v))
}

// HostnameIn applies the In predicate on the "hostname" field.
func HostnameIn(vs ...string) predicate.Agent {
	return predicate.Agent(sql.FieldIn(FieldHostname, vs...))
}

// HostnameNotIn applies the NotIn predicate on the "hostname" field.
func HostnameNotIn(vs ...string) predicate.Agent {
	return predicate.Agent(sql.FieldNotIn(FieldHostname, vs...))
}

// HostnameGT applies the GT predicate on the "hostname" field.
func HostnameGT(v string) predicate.Agent {
	return predicate.Agent(sql.FieldGT(FieldHostname, v))
}

// HostnameGTE applies the GTE predicate on the "hostname" field.
func HostnameGTE(v string) predicate.Agent {
	return predicate.Agent(sql.FieldGTE(FieldHostname, v))
}

// HostnameLT applies the LT predicate on the "hostname" field.
func HostnameLT(v string) predicate.Agent {
	return predicate.Agent(sql.FieldLT(FieldHostname, v))
}

// HostnameLTE applies the LTE predicate on the "hostname" field.
func HostnameLTE(v string) predicate.Agent {
	return predicate.Agent(sql.FieldLTE(FieldHostname, v))
}

// HostnameContains applies the Contains predicate on the "hostname" field.
func HostnameContains(v string) predicate.Agent {
	return predicate.Agent(sql.FieldContains(FieldHostname, v))
}

// HostnameHasPrefix applies the HasPrefix predicate on the "hostname" field.
func HostnameHasPrefix(v string) predicate.Agent {
	return predicate.Agent(sql.FieldHasPrefix(FieldHostname, v))
}

// HostnameHasSuffix applies the HasSuffix predicate on the "hostname" field.
func HostnameHasSuffix(v string) predicate.Agent {
	return predicate.Agent(sql.FieldHasSuffix(FieldHostname, v))
}

// HostnameEqualFold applies the EqualFold predicate on the "hostname" field.
func HostnameEqualFold(v string) predicate.Agent {
	return predicate.Agent(sql.FieldEqualFold(FieldHostname, v))
}

// HostnameContainsFold applies the ContainsFold predicate on the "hostname" field.
func HostnameContainsFold(v string) predicate.Agent {
	return predicate.Agent(sql.FieldContainsFold(FieldHostname, v))
}

// VersionEQ applies the EQ predicate on the "version" field.
func VersionEQ(v string) predicate.Agent {
	return predicate.Agent(sql.FieldEQ(FieldVersion, v))
}

// VersionNEQ applies the NEQ predicate on the "version" field.
func VersionNEQ(v string) predicate.Agent {
	return predicate.Agent(sql.FieldNEQ(FieldVersion, v))
}

// VersionIn applies the In predicate on the "version" field.
func VersionIn(vs ...string) predicate.Agent {
	return predicate.Agent(sql.FieldIn(FieldVersion, vs...))
}

// VersionNotIn applies the NotIn predicate on the "version" field.
func VersionNotIn(vs ...string) predicate.Agent {
	return predicate.Agent(sql.FieldNotIn(FieldVersion, vs...))
}

// VersionGT applies the GT predicate on the "version" field.
func VersionGT(v string) predicate.Agent {
	return predicate.Agent(sql.FieldGT(FieldVersion, v))
}

// VersionGTE applies the GTE predicate on the "version" field.
func VersionGTE(v string) predicate.Agent {
	return predicate.Agent(sql.FieldGTE(FieldVersion, v))
}

// VersionLT applies the LT predicate on the "version" field.
func VersionLT(v string) predicate.Agent {
	return predicate.Agent(sql.FieldLT(FieldVersion, v))
}

// VersionLTE applies the LTE predicate on the "version" field.
func VersionLTE(v string) predicate.Agent {
	return predicate.Agent(sql.FieldLTE(FieldVersion, v))
}

// VersionContains applies the Contains predicate on the "version" field.
func VersionContains(v string) predicate.Agent {
	return predicate.Agent(sql.FieldContains(FieldVersion, v))
}

// VersionHasPrefix applies the HasPrefix predicate on the "version" field.
func VersionHasPrefix(v string) predicate.Agent {
	return predicate.Agent(sql.FieldHasPrefix(FieldVersion, v))
}

// VersionHasSuffix applies the HasSuffix predicate on the "version" field.
func VersionHasSuffix(v string) predicate.Agent {
	return predicate.Agent(sql.FieldHasSuffix(FieldVersion, v))
}

// VersionEqualFold applies the EqualFold predicate on the "version" field.
func VersionEqualFold(v string) predicate.Agent {
	return predicate.Agent(sql.FieldEqualFold(FieldVersion, v))
}

// VersionContainsFold applies the ContainsFold predicate on the "version" field.
func VersionContainsFold(v string) predicate.Agent {
	return predicate.Agent(sql.FieldContainsFold(FieldVersion, v))
}

// IPEQ applies the EQ predicate on the "ip" field.
func IPEQ(v string) predicate.Agent {
	return predicate.Agent(sql.FieldEQ(FieldIP, v))
}

// IPNEQ applies the NEQ predicate on the "ip" field.
func IPNEQ(v string) predicate.Agent {
	return predicate.Agent(sql.FieldNEQ(FieldIP, v))
}

// IPIn applies the In predicate on the "ip" field.
func IPIn(vs ...string) predicate.Agent {
	return predicate.Agent(sql.FieldIn(FieldIP, vs...))
}

// IPNotIn applies the NotIn predicate on the "ip" field.
func IPNotIn(vs ...string) predicate.Agent {
	return predicate.Agent(sql.FieldNotIn(FieldIP, vs...))
}

// IPGT applies the GT predicate on the "ip" field.
func IPGT(v string) predicate.Agent {
	return predicate.Agent(sql.FieldGT(FieldIP, v))
}

// IPGTE applies the GTE predicate on the "ip" field.
func IPGTE(v string) predicate.Agent {
	return predicate.Agent(sql.FieldGTE(FieldIP, v))
}

// IPLT applies the LT predicate on the "ip" field.
func IPLT(v string) predicate.Agent {
	return predicate.Agent(sql.FieldLT(FieldIP, v))
}

// IPLTE applies the LTE predicate on the "ip" field.
func IPLTE(v string) predicate.Agent {
	return predicate.Agent(sql.FieldLTE(FieldIP, v))
}

// IPContains applies the Contains predicate on the "ip" field.
func IPContains(v string) predicate.Agent {
	return predicate.Agent(sql.FieldContains(FieldIP, v))
}

// IPHasPrefix applies the HasPrefix predicate on the "ip" field.
func IPHasPrefix(v string) predicate.Agent {
	return predicate.Agent(sql.FieldHasPrefix(FieldIP, v))
}

// IPHasSuffix applies the HasSuffix predicate on the "ip" field.
func IPHasSuffix(v string) predicate.Agent {
	return predicate.Agent(sql.FieldHasSuffix(FieldIP, v))
}

// IPEqualFold applies the EqualFold predicate on the "ip" field.
func IPEqualFold(v string) predicate.Agent {
	return predicate.Agent(sql.FieldEqualFold(FieldIP, v))
}

// IPContainsFold applies the ContainsFold predicate on the "ip" field.
func IPContainsFold(v string) predicate.Agent {
	return predicate.Agent(sql.FieldContainsFold(FieldIP, v))
}

// MACEQ applies the EQ predicate on the "mac" field.
func MACEQ(v string) predicate.Agent {
	return predicate.Agent(sql.FieldEQ(FieldMAC, v))
}

// MACNEQ applies the NEQ predicate on the "mac" field.
func MACNEQ(v string) predicate.Agent {
	return predicate.Agent(sql.FieldNEQ(FieldMAC, v))
}

// MACIn applies the In predicate on the "mac" field.
func MACIn(vs ...string) predicate.Agent {
	return predicate.Agent(sql.FieldIn(FieldMAC, vs...))
}

// MACNotIn applies the NotIn predicate on the "mac" field.
func MACNotIn(vs ...string) predicate.Agent {
	return predicate.Agent(sql.FieldNotIn(FieldMAC, vs...))
}

// MACGT applies the GT predicate on the "mac" field.
func MACGT(v string) predicate.Agent {
	return predicate.Agent(sql.FieldGT(FieldMAC, v))
}

// MACGTE applies the GTE predicate on the "mac" field.
func MACGTE(v string) predicate.Agent {
	return predicate.Agent(sql.FieldGTE(FieldMAC, v))
}

// MACLT applies the LT predicate on the "mac" field.
func MACLT(v string) predicate.Agent {
	return predicate.Agent(sql.FieldLT(FieldMAC, v))
}

// MACLTE applies the LTE predicate on the "mac" field.
func MACLTE(v string) predicate.Agent {
	return predicate.Agent(sql.FieldLTE(FieldMAC, v))
}

// MACContains applies the Contains predicate on the "mac" field.
func MACContains(v string) predicate.Agent {
	return predicate.Agent(sql.FieldContains(FieldMAC, v))
}

// MACHasPrefix applies the HasPrefix predicate on the "mac" field.
func MACHasPrefix(v string) predicate.Agent {
	return predicate.Agent(sql.FieldHasPrefix(FieldMAC, v))
}

// MACHasSuffix applies the HasSuffix predicate on the "mac" field.
func MACHasSuffix(v string) predicate.Agent {
	return predicate.Agent(sql.FieldHasSuffix(FieldMAC, v))
}

// MACEqualFold applies the EqualFold predicate on the "mac" field.
func MACEqualFold(v string) predicate.Agent {
	return predicate.Agent(sql.FieldEqualFold(FieldMAC, v))
}

// MACContainsFold applies the ContainsFold predicate on the "mac" field.
func MACContainsFold(v string) predicate.Agent {
	return predicate.Agent(sql.FieldContainsFold(FieldMAC, v))
}

// FirstContactEQ applies the EQ predicate on the "first_contact" field.
func FirstContactEQ(v time.Time) predicate.Agent {
	return predicate.Agent(sql.FieldEQ(FieldFirstContact, v))
}

// FirstContactNEQ applies the NEQ predicate on the "first_contact" field.
func FirstContactNEQ(v time.Time) predicate.Agent {
	return predicate.Agent(sql.FieldNEQ(FieldFirstContact, v))
}

// FirstContactIn applies the In predicate on the "first_contact" field.
func FirstContactIn(vs ...time.Time) predicate.Agent {
	return predicate.Agent(sql.FieldIn(FieldFirstContact, vs...))
}

// FirstContactNotIn applies the NotIn predicate on the "first_contact" field.
func FirstContactNotIn(vs ...time.Time) predicate.Agent {
	return predicate.Agent(sql.FieldNotIn(FieldFirstContact, vs...))
}

// FirstContactGT applies the GT predicate on the "first_contact" field.
func FirstContactGT(v time.Time) predicate.Agent {
	return predicate.Agent(sql.FieldGT(FieldFirstContact, v))
}

// FirstContactGTE applies the GTE predicate on the "first_contact" field.
func FirstContactGTE(v time.Time) predicate.Agent {
	return predicate.Agent(sql.FieldGTE(FieldFirstContact, v))
}

// FirstContactLT applies the LT predicate on the "first_contact" field.
func FirstContactLT(v time.Time) predicate.Agent {
	return predicate.Agent(sql.FieldLT(FieldFirstContact, v))
}

// FirstContactLTE applies the LTE predicate on the "first_contact" field.
func FirstContactLTE(v time.Time) predicate.Agent {
	return predicate.Agent(sql.FieldLTE(FieldFirstContact, v))
}

// FirstContactIsNil applies the IsNil predicate on the "first_contact" field.
func FirstContactIsNil() predicate.Agent {
	return predicate.Agent(sql.FieldIsNull(FieldFirstContact))
}

// FirstContactNotNil applies the NotNil predicate on the "first_contact" field.
func FirstContactNotNil() predicate.Agent {
	return predicate.Agent(sql.FieldNotNull(FieldFirstContact))
}

// LastContactEQ applies the EQ predicate on the "last_contact" field.
func LastContactEQ(v time.Time) predicate.Agent {
	return predicate.Agent(sql.FieldEQ(FieldLastContact, v))
}

// LastContactNEQ applies the NEQ predicate on the "last_contact" field.
func LastContactNEQ(v time.Time) predicate.Agent {
	return predicate.Agent(sql.FieldNEQ(FieldLastContact, v))
}

// LastContactIn applies the In predicate on the "last_contact" field.
func LastContactIn(vs ...time.Time) predicate.Agent {
	return predicate.Agent(sql.FieldIn(FieldLastContact, vs...))
}

// LastContactNotIn applies the NotIn predicate on the "last_contact" field.
func LastContactNotIn(vs ...time.Time) predicate.Agent {
	return predicate.Agent(sql.FieldNotIn(FieldLastContact, vs...))
}

// LastContactGT applies the GT predicate on the "last_contact" field.
func LastContactGT(v time.Time) predicate.Agent {
	return predicate.Agent(sql.FieldGT(FieldLastContact, v))
}

// LastContactGTE applies the GTE predicate on the "last_contact" field.
func LastContactGTE(v time.Time) predicate.Agent {
	return predicate.Agent(sql.FieldGTE(FieldLastContact, v))
}

// LastContactLT applies the LT predicate on the "last_contact" field.
func LastContactLT(v time.Time) predicate.Agent {
	return predicate.Agent(sql.FieldLT(FieldLastContact, v))
}

// LastContactLTE applies the LTE predicate on the "last_contact" field.
func LastContactLTE(v time.Time) predicate.Agent {
	return predicate.Agent(sql.FieldLTE(FieldLastContact, v))
}

// LastContactIsNil applies the IsNil predicate on the "last_contact" field.
func LastContactIsNil() predicate.Agent {
	return predicate.Agent(sql.FieldIsNull(FieldLastContact))
}

// LastContactNotNil applies the NotNil predicate on the "last_contact" field.
func LastContactNotNil() predicate.Agent {
	return predicate.Agent(sql.FieldNotNull(FieldLastContact))
}

// EnabledEQ applies the EQ predicate on the "enabled" field.
func EnabledEQ(v bool) predicate.Agent {
	return predicate.Agent(sql.FieldEQ(FieldEnabled, v))
}

// EnabledNEQ applies the NEQ predicate on the "enabled" field.
func EnabledNEQ(v bool) predicate.Agent {
	return predicate.Agent(sql.FieldNEQ(FieldEnabled, v))
}

// VncEQ applies the EQ predicate on the "vnc" field.
func VncEQ(v string) predicate.Agent {
	return predicate.Agent(sql.FieldEQ(FieldVnc, v))
}

// VncNEQ applies the NEQ predicate on the "vnc" field.
func VncNEQ(v string) predicate.Agent {
	return predicate.Agent(sql.FieldNEQ(FieldVnc, v))
}

// VncIn applies the In predicate on the "vnc" field.
func VncIn(vs ...string) predicate.Agent {
	return predicate.Agent(sql.FieldIn(FieldVnc, vs...))
}

// VncNotIn applies the NotIn predicate on the "vnc" field.
func VncNotIn(vs ...string) predicate.Agent {
	return predicate.Agent(sql.FieldNotIn(FieldVnc, vs...))
}

// VncGT applies the GT predicate on the "vnc" field.
func VncGT(v string) predicate.Agent {
	return predicate.Agent(sql.FieldGT(FieldVnc, v))
}

// VncGTE applies the GTE predicate on the "vnc" field.
func VncGTE(v string) predicate.Agent {
	return predicate.Agent(sql.FieldGTE(FieldVnc, v))
}

// VncLT applies the LT predicate on the "vnc" field.
func VncLT(v string) predicate.Agent {
	return predicate.Agent(sql.FieldLT(FieldVnc, v))
}

// VncLTE applies the LTE predicate on the "vnc" field.
func VncLTE(v string) predicate.Agent {
	return predicate.Agent(sql.FieldLTE(FieldVnc, v))
}

// VncContains applies the Contains predicate on the "vnc" field.
func VncContains(v string) predicate.Agent {
	return predicate.Agent(sql.FieldContains(FieldVnc, v))
}

// VncHasPrefix applies the HasPrefix predicate on the "vnc" field.
func VncHasPrefix(v string) predicate.Agent {
	return predicate.Agent(sql.FieldHasPrefix(FieldVnc, v))
}

// VncHasSuffix applies the HasSuffix predicate on the "vnc" field.
func VncHasSuffix(v string) predicate.Agent {
	return predicate.Agent(sql.FieldHasSuffix(FieldVnc, v))
}

// VncIsNil applies the IsNil predicate on the "vnc" field.
func VncIsNil() predicate.Agent {
	return predicate.Agent(sql.FieldIsNull(FieldVnc))
}

// VncNotNil applies the NotNil predicate on the "vnc" field.
func VncNotNil() predicate.Agent {
	return predicate.Agent(sql.FieldNotNull(FieldVnc))
}

// VncEqualFold applies the EqualFold predicate on the "vnc" field.
func VncEqualFold(v string) predicate.Agent {
	return predicate.Agent(sql.FieldEqualFold(FieldVnc, v))
}

// VncContainsFold applies the ContainsFold predicate on the "vnc" field.
func VncContainsFold(v string) predicate.Agent {
	return predicate.Agent(sql.FieldContainsFold(FieldVnc, v))
}

// NotesEQ applies the EQ predicate on the "notes" field.
func NotesEQ(v string) predicate.Agent {
	return predicate.Agent(sql.FieldEQ(FieldNotes, v))
}

// NotesNEQ applies the NEQ predicate on the "notes" field.
func NotesNEQ(v string) predicate.Agent {
	return predicate.Agent(sql.FieldNEQ(FieldNotes, v))
}

// NotesIn applies the In predicate on the "notes" field.
func NotesIn(vs ...string) predicate.Agent {
	return predicate.Agent(sql.FieldIn(FieldNotes, vs...))
}

// NotesNotIn applies the NotIn predicate on the "notes" field.
func NotesNotIn(vs ...string) predicate.Agent {
	return predicate.Agent(sql.FieldNotIn(FieldNotes, vs...))
}

// NotesGT applies the GT predicate on the "notes" field.
func NotesGT(v string) predicate.Agent {
	return predicate.Agent(sql.FieldGT(FieldNotes, v))
}

// NotesGTE applies the GTE predicate on the "notes" field.
func NotesGTE(v string) predicate.Agent {
	return predicate.Agent(sql.FieldGTE(FieldNotes, v))
}

// NotesLT applies the LT predicate on the "notes" field.
func NotesLT(v string) predicate.Agent {
	return predicate.Agent(sql.FieldLT(FieldNotes, v))
}

// NotesLTE applies the LTE predicate on the "notes" field.
func NotesLTE(v string) predicate.Agent {
	return predicate.Agent(sql.FieldLTE(FieldNotes, v))
}

// NotesContains applies the Contains predicate on the "notes" field.
func NotesContains(v string) predicate.Agent {
	return predicate.Agent(sql.FieldContains(FieldNotes, v))
}

// NotesHasPrefix applies the HasPrefix predicate on the "notes" field.
func NotesHasPrefix(v string) predicate.Agent {
	return predicate.Agent(sql.FieldHasPrefix(FieldNotes, v))
}

// NotesHasSuffix applies the HasSuffix predicate on the "notes" field.
func NotesHasSuffix(v string) predicate.Agent {
	return predicate.Agent(sql.FieldHasSuffix(FieldNotes, v))
}

// NotesIsNil applies the IsNil predicate on the "notes" field.
func NotesIsNil() predicate.Agent {
	return predicate.Agent(sql.FieldIsNull(FieldNotes))
}

// NotesNotNil applies the NotNil predicate on the "notes" field.
func NotesNotNil() predicate.Agent {
	return predicate.Agent(sql.FieldNotNull(FieldNotes))
}

// NotesEqualFold applies the EqualFold predicate on the "notes" field.
func NotesEqualFold(v string) predicate.Agent {
	return predicate.Agent(sql.FieldEqualFold(FieldNotes, v))
}

// NotesContainsFold applies the ContainsFold predicate on the "notes" field.
func NotesContainsFold(v string) predicate.Agent {
	return predicate.Agent(sql.FieldContainsFold(FieldNotes, v))
}

// HasComputer applies the HasEdge predicate on the "computer" edge.
func HasComputer() predicate.Agent {
	return predicate.Agent(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, ComputerTable, ComputerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasComputerWith applies the HasEdge predicate on the "computer" edge with a given conditions (other predicates).
func HasComputerWith(preds ...predicate.Computer) predicate.Agent {
	return predicate.Agent(func(s *sql.Selector) {
		step := newComputerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOperatingsystem applies the HasEdge predicate on the "operatingsystem" edge.
func HasOperatingsystem() predicate.Agent {
	return predicate.Agent(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, OperatingsystemTable, OperatingsystemColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOperatingsystemWith applies the HasEdge predicate on the "operatingsystem" edge with a given conditions (other predicates).
func HasOperatingsystemWith(preds ...predicate.OperatingSystem) predicate.Agent {
	return predicate.Agent(func(s *sql.Selector) {
		step := newOperatingsystemStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSystemupdate applies the HasEdge predicate on the "systemupdate" edge.
func HasSystemupdate() predicate.Agent {
	return predicate.Agent(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, SystemupdateTable, SystemupdateColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSystemupdateWith applies the HasEdge predicate on the "systemupdate" edge with a given conditions (other predicates).
func HasSystemupdateWith(preds ...predicate.SystemUpdate) predicate.Agent {
	return predicate.Agent(func(s *sql.Selector) {
		step := newSystemupdateStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAntivirus applies the HasEdge predicate on the "antivirus" edge.
func HasAntivirus() predicate.Agent {
	return predicate.Agent(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, AntivirusTable, AntivirusColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAntivirusWith applies the HasEdge predicate on the "antivirus" edge with a given conditions (other predicates).
func HasAntivirusWith(preds ...predicate.Antivirus) predicate.Agent {
	return predicate.Agent(func(s *sql.Selector) {
		step := newAntivirusStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasLogicaldisks applies the HasEdge predicate on the "logicaldisks" edge.
func HasLogicaldisks() predicate.Agent {
	return predicate.Agent(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, LogicaldisksTable, LogicaldisksColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLogicaldisksWith applies the HasEdge predicate on the "logicaldisks" edge with a given conditions (other predicates).
func HasLogicaldisksWith(preds ...predicate.LogicalDisk) predicate.Agent {
	return predicate.Agent(func(s *sql.Selector) {
		step := newLogicaldisksStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasApps applies the HasEdge predicate on the "apps" edge.
func HasApps() predicate.Agent {
	return predicate.Agent(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AppsTable, AppsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAppsWith applies the HasEdge predicate on the "apps" edge with a given conditions (other predicates).
func HasAppsWith(preds ...predicate.App) predicate.Agent {
	return predicate.Agent(func(s *sql.Selector) {
		step := newAppsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMonitors applies the HasEdge predicate on the "monitors" edge.
func HasMonitors() predicate.Agent {
	return predicate.Agent(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MonitorsTable, MonitorsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMonitorsWith applies the HasEdge predicate on the "monitors" edge with a given conditions (other predicates).
func HasMonitorsWith(preds ...predicate.Monitor) predicate.Agent {
	return predicate.Agent(func(s *sql.Selector) {
		step := newMonitorsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasShares applies the HasEdge predicate on the "shares" edge.
func HasShares() predicate.Agent {
	return predicate.Agent(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SharesTable, SharesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSharesWith applies the HasEdge predicate on the "shares" edge with a given conditions (other predicates).
func HasSharesWith(preds ...predicate.Share) predicate.Agent {
	return predicate.Agent(func(s *sql.Selector) {
		step := newSharesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPrinters applies the HasEdge predicate on the "printers" edge.
func HasPrinters() predicate.Agent {
	return predicate.Agent(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PrintersTable, PrintersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPrintersWith applies the HasEdge predicate on the "printers" edge with a given conditions (other predicates).
func HasPrintersWith(preds ...predicate.Printer) predicate.Agent {
	return predicate.Agent(func(s *sql.Selector) {
		step := newPrintersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasNetworkadapters applies the HasEdge predicate on the "networkadapters" edge.
func HasNetworkadapters() predicate.Agent {
	return predicate.Agent(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, NetworkadaptersTable, NetworkadaptersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNetworkadaptersWith applies the HasEdge predicate on the "networkadapters" edge with a given conditions (other predicates).
func HasNetworkadaptersWith(preds ...predicate.NetworkAdapter) predicate.Agent {
	return predicate.Agent(func(s *sql.Selector) {
		step := newNetworkadaptersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDeployments applies the HasEdge predicate on the "deployments" edge.
func HasDeployments() predicate.Agent {
	return predicate.Agent(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DeploymentsTable, DeploymentsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDeploymentsWith applies the HasEdge predicate on the "deployments" edge with a given conditions (other predicates).
func HasDeploymentsWith(preds ...predicate.Deployment) predicate.Agent {
	return predicate.Agent(func(s *sql.Selector) {
		step := newDeploymentsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUpdates applies the HasEdge predicate on the "updates" edge.
func HasUpdates() predicate.Agent {
	return predicate.Agent(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UpdatesTable, UpdatesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUpdatesWith applies the HasEdge predicate on the "updates" edge with a given conditions (other predicates).
func HasUpdatesWith(preds ...predicate.Update) predicate.Agent {
	return predicate.Agent(func(s *sql.Selector) {
		step := newUpdatesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTags applies the HasEdge predicate on the "tags" edge.
func HasTags() predicate.Agent {
	return predicate.Agent(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, TagsTable, TagsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTagsWith applies the HasEdge predicate on the "tags" edge with a given conditions (other predicates).
func HasTagsWith(preds ...predicate.Tag) predicate.Agent {
	return predicate.Agent(func(s *sql.Selector) {
		step := newTagsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMetadata applies the HasEdge predicate on the "metadata" edge.
func HasMetadata() predicate.Agent {
	return predicate.Agent(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, MetadataTable, MetadataPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMetadataWith applies the HasEdge predicate on the "metadata" edge with a given conditions (other predicates).
func HasMetadataWith(preds ...predicate.Metadata) predicate.Agent {
	return predicate.Agent(func(s *sql.Selector) {
		step := newMetadataStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Agent) predicate.Agent {
	return predicate.Agent(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Agent) predicate.Agent {
	return predicate.Agent(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Agent) predicate.Agent {
	return predicate.Agent(sql.NotPredicates(p))
}
