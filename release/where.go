// Code generated by ent, DO NOT EDIT.

package release

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/doncicuto/openuem_ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Release {
	return predicate.Release(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Release {
	return predicate.Release(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Release {
	return predicate.Release(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Release {
	return predicate.Release(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Release {
	return predicate.Release(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Release {
	return predicate.Release(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Release {
	return predicate.Release(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Release {
	return predicate.Release(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Release {
	return predicate.Release(sql.FieldLTE(FieldID, id))
}

// Version applies equality check predicate on the "version" field. It's identical to VersionEQ.
func Version(v string) predicate.Release {
	return predicate.Release(sql.FieldEQ(FieldVersion, v))
}

// Channel applies equality check predicate on the "channel" field. It's identical to ChannelEQ.
func Channel(v string) predicate.Release {
	return predicate.Release(sql.FieldEQ(FieldChannel, v))
}

// Summary applies equality check predicate on the "summary" field. It's identical to SummaryEQ.
func Summary(v string) predicate.Release {
	return predicate.Release(sql.FieldEQ(FieldSummary, v))
}

// ReleaseNotes applies equality check predicate on the "release_notes" field. It's identical to ReleaseNotesEQ.
func ReleaseNotes(v string) predicate.Release {
	return predicate.Release(sql.FieldEQ(FieldReleaseNotes, v))
}

// FileURL applies equality check predicate on the "file_url" field. It's identical to FileURLEQ.
func FileURL(v string) predicate.Release {
	return predicate.Release(sql.FieldEQ(FieldFileURL, v))
}

// Checksum applies equality check predicate on the "checksum" field. It's identical to ChecksumEQ.
func Checksum(v string) predicate.Release {
	return predicate.Release(sql.FieldEQ(FieldChecksum, v))
}

// IsCritical applies equality check predicate on the "is_critical" field. It's identical to IsCriticalEQ.
func IsCritical(v bool) predicate.Release {
	return predicate.Release(sql.FieldEQ(FieldIsCritical, v))
}

// ReleaseDate applies equality check predicate on the "release_date" field. It's identical to ReleaseDateEQ.
func ReleaseDate(v time.Time) predicate.Release {
	return predicate.Release(sql.FieldEQ(FieldReleaseDate, v))
}

// Os applies equality check predicate on the "os" field. It's identical to OsEQ.
func Os(v string) predicate.Release {
	return predicate.Release(sql.FieldEQ(FieldOs, v))
}

// Arch applies equality check predicate on the "arch" field. It's identical to ArchEQ.
func Arch(v string) predicate.Release {
	return predicate.Release(sql.FieldEQ(FieldArch, v))
}

// ReleaseTypeEQ applies the EQ predicate on the "release_type" field.
func ReleaseTypeEQ(v ReleaseType) predicate.Release {
	return predicate.Release(sql.FieldEQ(FieldReleaseType, v))
}

// ReleaseTypeNEQ applies the NEQ predicate on the "release_type" field.
func ReleaseTypeNEQ(v ReleaseType) predicate.Release {
	return predicate.Release(sql.FieldNEQ(FieldReleaseType, v))
}

// ReleaseTypeIn applies the In predicate on the "release_type" field.
func ReleaseTypeIn(vs ...ReleaseType) predicate.Release {
	return predicate.Release(sql.FieldIn(FieldReleaseType, vs...))
}

// ReleaseTypeNotIn applies the NotIn predicate on the "release_type" field.
func ReleaseTypeNotIn(vs ...ReleaseType) predicate.Release {
	return predicate.Release(sql.FieldNotIn(FieldReleaseType, vs...))
}

// ReleaseTypeIsNil applies the IsNil predicate on the "release_type" field.
func ReleaseTypeIsNil() predicate.Release {
	return predicate.Release(sql.FieldIsNull(FieldReleaseType))
}

// ReleaseTypeNotNil applies the NotNil predicate on the "release_type" field.
func ReleaseTypeNotNil() predicate.Release {
	return predicate.Release(sql.FieldNotNull(FieldReleaseType))
}

// VersionEQ applies the EQ predicate on the "version" field.
func VersionEQ(v string) predicate.Release {
	return predicate.Release(sql.FieldEQ(FieldVersion, v))
}

// VersionNEQ applies the NEQ predicate on the "version" field.
func VersionNEQ(v string) predicate.Release {
	return predicate.Release(sql.FieldNEQ(FieldVersion, v))
}

// VersionIn applies the In predicate on the "version" field.
func VersionIn(vs ...string) predicate.Release {
	return predicate.Release(sql.FieldIn(FieldVersion, vs...))
}

// VersionNotIn applies the NotIn predicate on the "version" field.
func VersionNotIn(vs ...string) predicate.Release {
	return predicate.Release(sql.FieldNotIn(FieldVersion, vs...))
}

// VersionGT applies the GT predicate on the "version" field.
func VersionGT(v string) predicate.Release {
	return predicate.Release(sql.FieldGT(FieldVersion, v))
}

// VersionGTE applies the GTE predicate on the "version" field.
func VersionGTE(v string) predicate.Release {
	return predicate.Release(sql.FieldGTE(FieldVersion, v))
}

// VersionLT applies the LT predicate on the "version" field.
func VersionLT(v string) predicate.Release {
	return predicate.Release(sql.FieldLT(FieldVersion, v))
}

// VersionLTE applies the LTE predicate on the "version" field.
func VersionLTE(v string) predicate.Release {
	return predicate.Release(sql.FieldLTE(FieldVersion, v))
}

// VersionContains applies the Contains predicate on the "version" field.
func VersionContains(v string) predicate.Release {
	return predicate.Release(sql.FieldContains(FieldVersion, v))
}

// VersionHasPrefix applies the HasPrefix predicate on the "version" field.
func VersionHasPrefix(v string) predicate.Release {
	return predicate.Release(sql.FieldHasPrefix(FieldVersion, v))
}

// VersionHasSuffix applies the HasSuffix predicate on the "version" field.
func VersionHasSuffix(v string) predicate.Release {
	return predicate.Release(sql.FieldHasSuffix(FieldVersion, v))
}

// VersionIsNil applies the IsNil predicate on the "version" field.
func VersionIsNil() predicate.Release {
	return predicate.Release(sql.FieldIsNull(FieldVersion))
}

// VersionNotNil applies the NotNil predicate on the "version" field.
func VersionNotNil() predicate.Release {
	return predicate.Release(sql.FieldNotNull(FieldVersion))
}

// VersionEqualFold applies the EqualFold predicate on the "version" field.
func VersionEqualFold(v string) predicate.Release {
	return predicate.Release(sql.FieldEqualFold(FieldVersion, v))
}

// VersionContainsFold applies the ContainsFold predicate on the "version" field.
func VersionContainsFold(v string) predicate.Release {
	return predicate.Release(sql.FieldContainsFold(FieldVersion, v))
}

// ChannelEQ applies the EQ predicate on the "channel" field.
func ChannelEQ(v string) predicate.Release {
	return predicate.Release(sql.FieldEQ(FieldChannel, v))
}

// ChannelNEQ applies the NEQ predicate on the "channel" field.
func ChannelNEQ(v string) predicate.Release {
	return predicate.Release(sql.FieldNEQ(FieldChannel, v))
}

// ChannelIn applies the In predicate on the "channel" field.
func ChannelIn(vs ...string) predicate.Release {
	return predicate.Release(sql.FieldIn(FieldChannel, vs...))
}

// ChannelNotIn applies the NotIn predicate on the "channel" field.
func ChannelNotIn(vs ...string) predicate.Release {
	return predicate.Release(sql.FieldNotIn(FieldChannel, vs...))
}

// ChannelGT applies the GT predicate on the "channel" field.
func ChannelGT(v string) predicate.Release {
	return predicate.Release(sql.FieldGT(FieldChannel, v))
}

// ChannelGTE applies the GTE predicate on the "channel" field.
func ChannelGTE(v string) predicate.Release {
	return predicate.Release(sql.FieldGTE(FieldChannel, v))
}

// ChannelLT applies the LT predicate on the "channel" field.
func ChannelLT(v string) predicate.Release {
	return predicate.Release(sql.FieldLT(FieldChannel, v))
}

// ChannelLTE applies the LTE predicate on the "channel" field.
func ChannelLTE(v string) predicate.Release {
	return predicate.Release(sql.FieldLTE(FieldChannel, v))
}

// ChannelContains applies the Contains predicate on the "channel" field.
func ChannelContains(v string) predicate.Release {
	return predicate.Release(sql.FieldContains(FieldChannel, v))
}

// ChannelHasPrefix applies the HasPrefix predicate on the "channel" field.
func ChannelHasPrefix(v string) predicate.Release {
	return predicate.Release(sql.FieldHasPrefix(FieldChannel, v))
}

// ChannelHasSuffix applies the HasSuffix predicate on the "channel" field.
func ChannelHasSuffix(v string) predicate.Release {
	return predicate.Release(sql.FieldHasSuffix(FieldChannel, v))
}

// ChannelIsNil applies the IsNil predicate on the "channel" field.
func ChannelIsNil() predicate.Release {
	return predicate.Release(sql.FieldIsNull(FieldChannel))
}

// ChannelNotNil applies the NotNil predicate on the "channel" field.
func ChannelNotNil() predicate.Release {
	return predicate.Release(sql.FieldNotNull(FieldChannel))
}

// ChannelEqualFold applies the EqualFold predicate on the "channel" field.
func ChannelEqualFold(v string) predicate.Release {
	return predicate.Release(sql.FieldEqualFold(FieldChannel, v))
}

// ChannelContainsFold applies the ContainsFold predicate on the "channel" field.
func ChannelContainsFold(v string) predicate.Release {
	return predicate.Release(sql.FieldContainsFold(FieldChannel, v))
}

// SummaryEQ applies the EQ predicate on the "summary" field.
func SummaryEQ(v string) predicate.Release {
	return predicate.Release(sql.FieldEQ(FieldSummary, v))
}

// SummaryNEQ applies the NEQ predicate on the "summary" field.
func SummaryNEQ(v string) predicate.Release {
	return predicate.Release(sql.FieldNEQ(FieldSummary, v))
}

// SummaryIn applies the In predicate on the "summary" field.
func SummaryIn(vs ...string) predicate.Release {
	return predicate.Release(sql.FieldIn(FieldSummary, vs...))
}

// SummaryNotIn applies the NotIn predicate on the "summary" field.
func SummaryNotIn(vs ...string) predicate.Release {
	return predicate.Release(sql.FieldNotIn(FieldSummary, vs...))
}

// SummaryGT applies the GT predicate on the "summary" field.
func SummaryGT(v string) predicate.Release {
	return predicate.Release(sql.FieldGT(FieldSummary, v))
}

// SummaryGTE applies the GTE predicate on the "summary" field.
func SummaryGTE(v string) predicate.Release {
	return predicate.Release(sql.FieldGTE(FieldSummary, v))
}

// SummaryLT applies the LT predicate on the "summary" field.
func SummaryLT(v string) predicate.Release {
	return predicate.Release(sql.FieldLT(FieldSummary, v))
}

// SummaryLTE applies the LTE predicate on the "summary" field.
func SummaryLTE(v string) predicate.Release {
	return predicate.Release(sql.FieldLTE(FieldSummary, v))
}

// SummaryContains applies the Contains predicate on the "summary" field.
func SummaryContains(v string) predicate.Release {
	return predicate.Release(sql.FieldContains(FieldSummary, v))
}

// SummaryHasPrefix applies the HasPrefix predicate on the "summary" field.
func SummaryHasPrefix(v string) predicate.Release {
	return predicate.Release(sql.FieldHasPrefix(FieldSummary, v))
}

// SummaryHasSuffix applies the HasSuffix predicate on the "summary" field.
func SummaryHasSuffix(v string) predicate.Release {
	return predicate.Release(sql.FieldHasSuffix(FieldSummary, v))
}

// SummaryIsNil applies the IsNil predicate on the "summary" field.
func SummaryIsNil() predicate.Release {
	return predicate.Release(sql.FieldIsNull(FieldSummary))
}

// SummaryNotNil applies the NotNil predicate on the "summary" field.
func SummaryNotNil() predicate.Release {
	return predicate.Release(sql.FieldNotNull(FieldSummary))
}

// SummaryEqualFold applies the EqualFold predicate on the "summary" field.
func SummaryEqualFold(v string) predicate.Release {
	return predicate.Release(sql.FieldEqualFold(FieldSummary, v))
}

// SummaryContainsFold applies the ContainsFold predicate on the "summary" field.
func SummaryContainsFold(v string) predicate.Release {
	return predicate.Release(sql.FieldContainsFold(FieldSummary, v))
}

// ReleaseNotesEQ applies the EQ predicate on the "release_notes" field.
func ReleaseNotesEQ(v string) predicate.Release {
	return predicate.Release(sql.FieldEQ(FieldReleaseNotes, v))
}

// ReleaseNotesNEQ applies the NEQ predicate on the "release_notes" field.
func ReleaseNotesNEQ(v string) predicate.Release {
	return predicate.Release(sql.FieldNEQ(FieldReleaseNotes, v))
}

// ReleaseNotesIn applies the In predicate on the "release_notes" field.
func ReleaseNotesIn(vs ...string) predicate.Release {
	return predicate.Release(sql.FieldIn(FieldReleaseNotes, vs...))
}

// ReleaseNotesNotIn applies the NotIn predicate on the "release_notes" field.
func ReleaseNotesNotIn(vs ...string) predicate.Release {
	return predicate.Release(sql.FieldNotIn(FieldReleaseNotes, vs...))
}

// ReleaseNotesGT applies the GT predicate on the "release_notes" field.
func ReleaseNotesGT(v string) predicate.Release {
	return predicate.Release(sql.FieldGT(FieldReleaseNotes, v))
}

// ReleaseNotesGTE applies the GTE predicate on the "release_notes" field.
func ReleaseNotesGTE(v string) predicate.Release {
	return predicate.Release(sql.FieldGTE(FieldReleaseNotes, v))
}

// ReleaseNotesLT applies the LT predicate on the "release_notes" field.
func ReleaseNotesLT(v string) predicate.Release {
	return predicate.Release(sql.FieldLT(FieldReleaseNotes, v))
}

// ReleaseNotesLTE applies the LTE predicate on the "release_notes" field.
func ReleaseNotesLTE(v string) predicate.Release {
	return predicate.Release(sql.FieldLTE(FieldReleaseNotes, v))
}

// ReleaseNotesContains applies the Contains predicate on the "release_notes" field.
func ReleaseNotesContains(v string) predicate.Release {
	return predicate.Release(sql.FieldContains(FieldReleaseNotes, v))
}

// ReleaseNotesHasPrefix applies the HasPrefix predicate on the "release_notes" field.
func ReleaseNotesHasPrefix(v string) predicate.Release {
	return predicate.Release(sql.FieldHasPrefix(FieldReleaseNotes, v))
}

// ReleaseNotesHasSuffix applies the HasSuffix predicate on the "release_notes" field.
func ReleaseNotesHasSuffix(v string) predicate.Release {
	return predicate.Release(sql.FieldHasSuffix(FieldReleaseNotes, v))
}

// ReleaseNotesIsNil applies the IsNil predicate on the "release_notes" field.
func ReleaseNotesIsNil() predicate.Release {
	return predicate.Release(sql.FieldIsNull(FieldReleaseNotes))
}

// ReleaseNotesNotNil applies the NotNil predicate on the "release_notes" field.
func ReleaseNotesNotNil() predicate.Release {
	return predicate.Release(sql.FieldNotNull(FieldReleaseNotes))
}

// ReleaseNotesEqualFold applies the EqualFold predicate on the "release_notes" field.
func ReleaseNotesEqualFold(v string) predicate.Release {
	return predicate.Release(sql.FieldEqualFold(FieldReleaseNotes, v))
}

// ReleaseNotesContainsFold applies the ContainsFold predicate on the "release_notes" field.
func ReleaseNotesContainsFold(v string) predicate.Release {
	return predicate.Release(sql.FieldContainsFold(FieldReleaseNotes, v))
}

// FileURLEQ applies the EQ predicate on the "file_url" field.
func FileURLEQ(v string) predicate.Release {
	return predicate.Release(sql.FieldEQ(FieldFileURL, v))
}

// FileURLNEQ applies the NEQ predicate on the "file_url" field.
func FileURLNEQ(v string) predicate.Release {
	return predicate.Release(sql.FieldNEQ(FieldFileURL, v))
}

// FileURLIn applies the In predicate on the "file_url" field.
func FileURLIn(vs ...string) predicate.Release {
	return predicate.Release(sql.FieldIn(FieldFileURL, vs...))
}

// FileURLNotIn applies the NotIn predicate on the "file_url" field.
func FileURLNotIn(vs ...string) predicate.Release {
	return predicate.Release(sql.FieldNotIn(FieldFileURL, vs...))
}

// FileURLGT applies the GT predicate on the "file_url" field.
func FileURLGT(v string) predicate.Release {
	return predicate.Release(sql.FieldGT(FieldFileURL, v))
}

// FileURLGTE applies the GTE predicate on the "file_url" field.
func FileURLGTE(v string) predicate.Release {
	return predicate.Release(sql.FieldGTE(FieldFileURL, v))
}

// FileURLLT applies the LT predicate on the "file_url" field.
func FileURLLT(v string) predicate.Release {
	return predicate.Release(sql.FieldLT(FieldFileURL, v))
}

// FileURLLTE applies the LTE predicate on the "file_url" field.
func FileURLLTE(v string) predicate.Release {
	return predicate.Release(sql.FieldLTE(FieldFileURL, v))
}

// FileURLContains applies the Contains predicate on the "file_url" field.
func FileURLContains(v string) predicate.Release {
	return predicate.Release(sql.FieldContains(FieldFileURL, v))
}

// FileURLHasPrefix applies the HasPrefix predicate on the "file_url" field.
func FileURLHasPrefix(v string) predicate.Release {
	return predicate.Release(sql.FieldHasPrefix(FieldFileURL, v))
}

// FileURLHasSuffix applies the HasSuffix predicate on the "file_url" field.
func FileURLHasSuffix(v string) predicate.Release {
	return predicate.Release(sql.FieldHasSuffix(FieldFileURL, v))
}

// FileURLIsNil applies the IsNil predicate on the "file_url" field.
func FileURLIsNil() predicate.Release {
	return predicate.Release(sql.FieldIsNull(FieldFileURL))
}

// FileURLNotNil applies the NotNil predicate on the "file_url" field.
func FileURLNotNil() predicate.Release {
	return predicate.Release(sql.FieldNotNull(FieldFileURL))
}

// FileURLEqualFold applies the EqualFold predicate on the "file_url" field.
func FileURLEqualFold(v string) predicate.Release {
	return predicate.Release(sql.FieldEqualFold(FieldFileURL, v))
}

// FileURLContainsFold applies the ContainsFold predicate on the "file_url" field.
func FileURLContainsFold(v string) predicate.Release {
	return predicate.Release(sql.FieldContainsFold(FieldFileURL, v))
}

// ChecksumEQ applies the EQ predicate on the "checksum" field.
func ChecksumEQ(v string) predicate.Release {
	return predicate.Release(sql.FieldEQ(FieldChecksum, v))
}

// ChecksumNEQ applies the NEQ predicate on the "checksum" field.
func ChecksumNEQ(v string) predicate.Release {
	return predicate.Release(sql.FieldNEQ(FieldChecksum, v))
}

// ChecksumIn applies the In predicate on the "checksum" field.
func ChecksumIn(vs ...string) predicate.Release {
	return predicate.Release(sql.FieldIn(FieldChecksum, vs...))
}

// ChecksumNotIn applies the NotIn predicate on the "checksum" field.
func ChecksumNotIn(vs ...string) predicate.Release {
	return predicate.Release(sql.FieldNotIn(FieldChecksum, vs...))
}

// ChecksumGT applies the GT predicate on the "checksum" field.
func ChecksumGT(v string) predicate.Release {
	return predicate.Release(sql.FieldGT(FieldChecksum, v))
}

// ChecksumGTE applies the GTE predicate on the "checksum" field.
func ChecksumGTE(v string) predicate.Release {
	return predicate.Release(sql.FieldGTE(FieldChecksum, v))
}

// ChecksumLT applies the LT predicate on the "checksum" field.
func ChecksumLT(v string) predicate.Release {
	return predicate.Release(sql.FieldLT(FieldChecksum, v))
}

// ChecksumLTE applies the LTE predicate on the "checksum" field.
func ChecksumLTE(v string) predicate.Release {
	return predicate.Release(sql.FieldLTE(FieldChecksum, v))
}

// ChecksumContains applies the Contains predicate on the "checksum" field.
func ChecksumContains(v string) predicate.Release {
	return predicate.Release(sql.FieldContains(FieldChecksum, v))
}

// ChecksumHasPrefix applies the HasPrefix predicate on the "checksum" field.
func ChecksumHasPrefix(v string) predicate.Release {
	return predicate.Release(sql.FieldHasPrefix(FieldChecksum, v))
}

// ChecksumHasSuffix applies the HasSuffix predicate on the "checksum" field.
func ChecksumHasSuffix(v string) predicate.Release {
	return predicate.Release(sql.FieldHasSuffix(FieldChecksum, v))
}

// ChecksumIsNil applies the IsNil predicate on the "checksum" field.
func ChecksumIsNil() predicate.Release {
	return predicate.Release(sql.FieldIsNull(FieldChecksum))
}

// ChecksumNotNil applies the NotNil predicate on the "checksum" field.
func ChecksumNotNil() predicate.Release {
	return predicate.Release(sql.FieldNotNull(FieldChecksum))
}

// ChecksumEqualFold applies the EqualFold predicate on the "checksum" field.
func ChecksumEqualFold(v string) predicate.Release {
	return predicate.Release(sql.FieldEqualFold(FieldChecksum, v))
}

// ChecksumContainsFold applies the ContainsFold predicate on the "checksum" field.
func ChecksumContainsFold(v string) predicate.Release {
	return predicate.Release(sql.FieldContainsFold(FieldChecksum, v))
}

// IsCriticalEQ applies the EQ predicate on the "is_critical" field.
func IsCriticalEQ(v bool) predicate.Release {
	return predicate.Release(sql.FieldEQ(FieldIsCritical, v))
}

// IsCriticalNEQ applies the NEQ predicate on the "is_critical" field.
func IsCriticalNEQ(v bool) predicate.Release {
	return predicate.Release(sql.FieldNEQ(FieldIsCritical, v))
}

// IsCriticalIsNil applies the IsNil predicate on the "is_critical" field.
func IsCriticalIsNil() predicate.Release {
	return predicate.Release(sql.FieldIsNull(FieldIsCritical))
}

// IsCriticalNotNil applies the NotNil predicate on the "is_critical" field.
func IsCriticalNotNil() predicate.Release {
	return predicate.Release(sql.FieldNotNull(FieldIsCritical))
}

// ReleaseDateEQ applies the EQ predicate on the "release_date" field.
func ReleaseDateEQ(v time.Time) predicate.Release {
	return predicate.Release(sql.FieldEQ(FieldReleaseDate, v))
}

// ReleaseDateNEQ applies the NEQ predicate on the "release_date" field.
func ReleaseDateNEQ(v time.Time) predicate.Release {
	return predicate.Release(sql.FieldNEQ(FieldReleaseDate, v))
}

// ReleaseDateIn applies the In predicate on the "release_date" field.
func ReleaseDateIn(vs ...time.Time) predicate.Release {
	return predicate.Release(sql.FieldIn(FieldReleaseDate, vs...))
}

// ReleaseDateNotIn applies the NotIn predicate on the "release_date" field.
func ReleaseDateNotIn(vs ...time.Time) predicate.Release {
	return predicate.Release(sql.FieldNotIn(FieldReleaseDate, vs...))
}

// ReleaseDateGT applies the GT predicate on the "release_date" field.
func ReleaseDateGT(v time.Time) predicate.Release {
	return predicate.Release(sql.FieldGT(FieldReleaseDate, v))
}

// ReleaseDateGTE applies the GTE predicate on the "release_date" field.
func ReleaseDateGTE(v time.Time) predicate.Release {
	return predicate.Release(sql.FieldGTE(FieldReleaseDate, v))
}

// ReleaseDateLT applies the LT predicate on the "release_date" field.
func ReleaseDateLT(v time.Time) predicate.Release {
	return predicate.Release(sql.FieldLT(FieldReleaseDate, v))
}

// ReleaseDateLTE applies the LTE predicate on the "release_date" field.
func ReleaseDateLTE(v time.Time) predicate.Release {
	return predicate.Release(sql.FieldLTE(FieldReleaseDate, v))
}

// ReleaseDateIsNil applies the IsNil predicate on the "release_date" field.
func ReleaseDateIsNil() predicate.Release {
	return predicate.Release(sql.FieldIsNull(FieldReleaseDate))
}

// ReleaseDateNotNil applies the NotNil predicate on the "release_date" field.
func ReleaseDateNotNil() predicate.Release {
	return predicate.Release(sql.FieldNotNull(FieldReleaseDate))
}

// OsEQ applies the EQ predicate on the "os" field.
func OsEQ(v string) predicate.Release {
	return predicate.Release(sql.FieldEQ(FieldOs, v))
}

// OsNEQ applies the NEQ predicate on the "os" field.
func OsNEQ(v string) predicate.Release {
	return predicate.Release(sql.FieldNEQ(FieldOs, v))
}

// OsIn applies the In predicate on the "os" field.
func OsIn(vs ...string) predicate.Release {
	return predicate.Release(sql.FieldIn(FieldOs, vs...))
}

// OsNotIn applies the NotIn predicate on the "os" field.
func OsNotIn(vs ...string) predicate.Release {
	return predicate.Release(sql.FieldNotIn(FieldOs, vs...))
}

// OsGT applies the GT predicate on the "os" field.
func OsGT(v string) predicate.Release {
	return predicate.Release(sql.FieldGT(FieldOs, v))
}

// OsGTE applies the GTE predicate on the "os" field.
func OsGTE(v string) predicate.Release {
	return predicate.Release(sql.FieldGTE(FieldOs, v))
}

// OsLT applies the LT predicate on the "os" field.
func OsLT(v string) predicate.Release {
	return predicate.Release(sql.FieldLT(FieldOs, v))
}

// OsLTE applies the LTE predicate on the "os" field.
func OsLTE(v string) predicate.Release {
	return predicate.Release(sql.FieldLTE(FieldOs, v))
}

// OsContains applies the Contains predicate on the "os" field.
func OsContains(v string) predicate.Release {
	return predicate.Release(sql.FieldContains(FieldOs, v))
}

// OsHasPrefix applies the HasPrefix predicate on the "os" field.
func OsHasPrefix(v string) predicate.Release {
	return predicate.Release(sql.FieldHasPrefix(FieldOs, v))
}

// OsHasSuffix applies the HasSuffix predicate on the "os" field.
func OsHasSuffix(v string) predicate.Release {
	return predicate.Release(sql.FieldHasSuffix(FieldOs, v))
}

// OsIsNil applies the IsNil predicate on the "os" field.
func OsIsNil() predicate.Release {
	return predicate.Release(sql.FieldIsNull(FieldOs))
}

// OsNotNil applies the NotNil predicate on the "os" field.
func OsNotNil() predicate.Release {
	return predicate.Release(sql.FieldNotNull(FieldOs))
}

// OsEqualFold applies the EqualFold predicate on the "os" field.
func OsEqualFold(v string) predicate.Release {
	return predicate.Release(sql.FieldEqualFold(FieldOs, v))
}

// OsContainsFold applies the ContainsFold predicate on the "os" field.
func OsContainsFold(v string) predicate.Release {
	return predicate.Release(sql.FieldContainsFold(FieldOs, v))
}

// ArchEQ applies the EQ predicate on the "arch" field.
func ArchEQ(v string) predicate.Release {
	return predicate.Release(sql.FieldEQ(FieldArch, v))
}

// ArchNEQ applies the NEQ predicate on the "arch" field.
func ArchNEQ(v string) predicate.Release {
	return predicate.Release(sql.FieldNEQ(FieldArch, v))
}

// ArchIn applies the In predicate on the "arch" field.
func ArchIn(vs ...string) predicate.Release {
	return predicate.Release(sql.FieldIn(FieldArch, vs...))
}

// ArchNotIn applies the NotIn predicate on the "arch" field.
func ArchNotIn(vs ...string) predicate.Release {
	return predicate.Release(sql.FieldNotIn(FieldArch, vs...))
}

// ArchGT applies the GT predicate on the "arch" field.
func ArchGT(v string) predicate.Release {
	return predicate.Release(sql.FieldGT(FieldArch, v))
}

// ArchGTE applies the GTE predicate on the "arch" field.
func ArchGTE(v string) predicate.Release {
	return predicate.Release(sql.FieldGTE(FieldArch, v))
}

// ArchLT applies the LT predicate on the "arch" field.
func ArchLT(v string) predicate.Release {
	return predicate.Release(sql.FieldLT(FieldArch, v))
}

// ArchLTE applies the LTE predicate on the "arch" field.
func ArchLTE(v string) predicate.Release {
	return predicate.Release(sql.FieldLTE(FieldArch, v))
}

// ArchContains applies the Contains predicate on the "arch" field.
func ArchContains(v string) predicate.Release {
	return predicate.Release(sql.FieldContains(FieldArch, v))
}

// ArchHasPrefix applies the HasPrefix predicate on the "arch" field.
func ArchHasPrefix(v string) predicate.Release {
	return predicate.Release(sql.FieldHasPrefix(FieldArch, v))
}

// ArchHasSuffix applies the HasSuffix predicate on the "arch" field.
func ArchHasSuffix(v string) predicate.Release {
	return predicate.Release(sql.FieldHasSuffix(FieldArch, v))
}

// ArchIsNil applies the IsNil predicate on the "arch" field.
func ArchIsNil() predicate.Release {
	return predicate.Release(sql.FieldIsNull(FieldArch))
}

// ArchNotNil applies the NotNil predicate on the "arch" field.
func ArchNotNil() predicate.Release {
	return predicate.Release(sql.FieldNotNull(FieldArch))
}

// ArchEqualFold applies the EqualFold predicate on the "arch" field.
func ArchEqualFold(v string) predicate.Release {
	return predicate.Release(sql.FieldEqualFold(FieldArch, v))
}

// ArchContainsFold applies the ContainsFold predicate on the "arch" field.
func ArchContainsFold(v string) predicate.Release {
	return predicate.Release(sql.FieldContainsFold(FieldArch, v))
}

// HasAgents applies the HasEdge predicate on the "agents" edge.
func HasAgents() predicate.Release {
	return predicate.Release(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AgentsTable, AgentsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAgentsWith applies the HasEdge predicate on the "agents" edge with a given conditions (other predicates).
func HasAgentsWith(preds ...predicate.Agent) predicate.Release {
	return predicate.Release(func(s *sql.Selector) {
		step := newAgentsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasServers applies the HasEdge predicate on the "servers" edge.
func HasServers() predicate.Release {
	return predicate.Release(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ServersTable, ServersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasServersWith applies the HasEdge predicate on the "servers" edge with a given conditions (other predicates).
func HasServersWith(preds ...predicate.Server) predicate.Release {
	return predicate.Release(func(s *sql.Selector) {
		step := newServersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Release) predicate.Release {
	return predicate.Release(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Release) predicate.Release {
	return predicate.Release(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Release) predicate.Release {
	return predicate.Release(sql.NotPredicates(p))
}
