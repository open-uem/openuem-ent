// Code generated by ent, DO NOT EDIT.

package revocation

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/doncicuto/openuem_ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Revocation {
	return predicate.Revocation(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Revocation {
	return predicate.Revocation(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Revocation {
	return predicate.Revocation(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Revocation {
	return predicate.Revocation(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Revocation {
	return predicate.Revocation(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Revocation {
	return predicate.Revocation(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Revocation {
	return predicate.Revocation(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Revocation {
	return predicate.Revocation(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Revocation {
	return predicate.Revocation(sql.FieldLTE(FieldID, id))
}

// Reason applies equality check predicate on the "reason" field. It's identical to ReasonEQ.
func Reason(v int) predicate.Revocation {
	return predicate.Revocation(sql.FieldEQ(FieldReason, v))
}

// Info applies equality check predicate on the "info" field. It's identical to InfoEQ.
func Info(v string) predicate.Revocation {
	return predicate.Revocation(sql.FieldEQ(FieldInfo, v))
}

// Expiry applies equality check predicate on the "expiry" field. It's identical to ExpiryEQ.
func Expiry(v time.Time) predicate.Revocation {
	return predicate.Revocation(sql.FieldEQ(FieldExpiry, v))
}

// Revoked applies equality check predicate on the "revoked" field. It's identical to RevokedEQ.
func Revoked(v time.Time) predicate.Revocation {
	return predicate.Revocation(sql.FieldEQ(FieldRevoked, v))
}

// ReasonEQ applies the EQ predicate on the "reason" field.
func ReasonEQ(v int) predicate.Revocation {
	return predicate.Revocation(sql.FieldEQ(FieldReason, v))
}

// ReasonNEQ applies the NEQ predicate on the "reason" field.
func ReasonNEQ(v int) predicate.Revocation {
	return predicate.Revocation(sql.FieldNEQ(FieldReason, v))
}

// ReasonIn applies the In predicate on the "reason" field.
func ReasonIn(vs ...int) predicate.Revocation {
	return predicate.Revocation(sql.FieldIn(FieldReason, vs...))
}

// ReasonNotIn applies the NotIn predicate on the "reason" field.
func ReasonNotIn(vs ...int) predicate.Revocation {
	return predicate.Revocation(sql.FieldNotIn(FieldReason, vs...))
}

// ReasonGT applies the GT predicate on the "reason" field.
func ReasonGT(v int) predicate.Revocation {
	return predicate.Revocation(sql.FieldGT(FieldReason, v))
}

// ReasonGTE applies the GTE predicate on the "reason" field.
func ReasonGTE(v int) predicate.Revocation {
	return predicate.Revocation(sql.FieldGTE(FieldReason, v))
}

// ReasonLT applies the LT predicate on the "reason" field.
func ReasonLT(v int) predicate.Revocation {
	return predicate.Revocation(sql.FieldLT(FieldReason, v))
}

// ReasonLTE applies the LTE predicate on the "reason" field.
func ReasonLTE(v int) predicate.Revocation {
	return predicate.Revocation(sql.FieldLTE(FieldReason, v))
}

// ReasonIsNil applies the IsNil predicate on the "reason" field.
func ReasonIsNil() predicate.Revocation {
	return predicate.Revocation(sql.FieldIsNull(FieldReason))
}

// ReasonNotNil applies the NotNil predicate on the "reason" field.
func ReasonNotNil() predicate.Revocation {
	return predicate.Revocation(sql.FieldNotNull(FieldReason))
}

// InfoEQ applies the EQ predicate on the "info" field.
func InfoEQ(v string) predicate.Revocation {
	return predicate.Revocation(sql.FieldEQ(FieldInfo, v))
}

// InfoNEQ applies the NEQ predicate on the "info" field.
func InfoNEQ(v string) predicate.Revocation {
	return predicate.Revocation(sql.FieldNEQ(FieldInfo, v))
}

// InfoIn applies the In predicate on the "info" field.
func InfoIn(vs ...string) predicate.Revocation {
	return predicate.Revocation(sql.FieldIn(FieldInfo, vs...))
}

// InfoNotIn applies the NotIn predicate on the "info" field.
func InfoNotIn(vs ...string) predicate.Revocation {
	return predicate.Revocation(sql.FieldNotIn(FieldInfo, vs...))
}

// InfoGT applies the GT predicate on the "info" field.
func InfoGT(v string) predicate.Revocation {
	return predicate.Revocation(sql.FieldGT(FieldInfo, v))
}

// InfoGTE applies the GTE predicate on the "info" field.
func InfoGTE(v string) predicate.Revocation {
	return predicate.Revocation(sql.FieldGTE(FieldInfo, v))
}

// InfoLT applies the LT predicate on the "info" field.
func InfoLT(v string) predicate.Revocation {
	return predicate.Revocation(sql.FieldLT(FieldInfo, v))
}

// InfoLTE applies the LTE predicate on the "info" field.
func InfoLTE(v string) predicate.Revocation {
	return predicate.Revocation(sql.FieldLTE(FieldInfo, v))
}

// InfoContains applies the Contains predicate on the "info" field.
func InfoContains(v string) predicate.Revocation {
	return predicate.Revocation(sql.FieldContains(FieldInfo, v))
}

// InfoHasPrefix applies the HasPrefix predicate on the "info" field.
func InfoHasPrefix(v string) predicate.Revocation {
	return predicate.Revocation(sql.FieldHasPrefix(FieldInfo, v))
}

// InfoHasSuffix applies the HasSuffix predicate on the "info" field.
func InfoHasSuffix(v string) predicate.Revocation {
	return predicate.Revocation(sql.FieldHasSuffix(FieldInfo, v))
}

// InfoIsNil applies the IsNil predicate on the "info" field.
func InfoIsNil() predicate.Revocation {
	return predicate.Revocation(sql.FieldIsNull(FieldInfo))
}

// InfoNotNil applies the NotNil predicate on the "info" field.
func InfoNotNil() predicate.Revocation {
	return predicate.Revocation(sql.FieldNotNull(FieldInfo))
}

// InfoEqualFold applies the EqualFold predicate on the "info" field.
func InfoEqualFold(v string) predicate.Revocation {
	return predicate.Revocation(sql.FieldEqualFold(FieldInfo, v))
}

// InfoContainsFold applies the ContainsFold predicate on the "info" field.
func InfoContainsFold(v string) predicate.Revocation {
	return predicate.Revocation(sql.FieldContainsFold(FieldInfo, v))
}

// ExpiryEQ applies the EQ predicate on the "expiry" field.
func ExpiryEQ(v time.Time) predicate.Revocation {
	return predicate.Revocation(sql.FieldEQ(FieldExpiry, v))
}

// ExpiryNEQ applies the NEQ predicate on the "expiry" field.
func ExpiryNEQ(v time.Time) predicate.Revocation {
	return predicate.Revocation(sql.FieldNEQ(FieldExpiry, v))
}

// ExpiryIn applies the In predicate on the "expiry" field.
func ExpiryIn(vs ...time.Time) predicate.Revocation {
	return predicate.Revocation(sql.FieldIn(FieldExpiry, vs...))
}

// ExpiryNotIn applies the NotIn predicate on the "expiry" field.
func ExpiryNotIn(vs ...time.Time) predicate.Revocation {
	return predicate.Revocation(sql.FieldNotIn(FieldExpiry, vs...))
}

// ExpiryGT applies the GT predicate on the "expiry" field.
func ExpiryGT(v time.Time) predicate.Revocation {
	return predicate.Revocation(sql.FieldGT(FieldExpiry, v))
}

// ExpiryGTE applies the GTE predicate on the "expiry" field.
func ExpiryGTE(v time.Time) predicate.Revocation {
	return predicate.Revocation(sql.FieldGTE(FieldExpiry, v))
}

// ExpiryLT applies the LT predicate on the "expiry" field.
func ExpiryLT(v time.Time) predicate.Revocation {
	return predicate.Revocation(sql.FieldLT(FieldExpiry, v))
}

// ExpiryLTE applies the LTE predicate on the "expiry" field.
func ExpiryLTE(v time.Time) predicate.Revocation {
	return predicate.Revocation(sql.FieldLTE(FieldExpiry, v))
}

// ExpiryIsNil applies the IsNil predicate on the "expiry" field.
func ExpiryIsNil() predicate.Revocation {
	return predicate.Revocation(sql.FieldIsNull(FieldExpiry))
}

// ExpiryNotNil applies the NotNil predicate on the "expiry" field.
func ExpiryNotNil() predicate.Revocation {
	return predicate.Revocation(sql.FieldNotNull(FieldExpiry))
}

// RevokedEQ applies the EQ predicate on the "revoked" field.
func RevokedEQ(v time.Time) predicate.Revocation {
	return predicate.Revocation(sql.FieldEQ(FieldRevoked, v))
}

// RevokedNEQ applies the NEQ predicate on the "revoked" field.
func RevokedNEQ(v time.Time) predicate.Revocation {
	return predicate.Revocation(sql.FieldNEQ(FieldRevoked, v))
}

// RevokedIn applies the In predicate on the "revoked" field.
func RevokedIn(vs ...time.Time) predicate.Revocation {
	return predicate.Revocation(sql.FieldIn(FieldRevoked, vs...))
}

// RevokedNotIn applies the NotIn predicate on the "revoked" field.
func RevokedNotIn(vs ...time.Time) predicate.Revocation {
	return predicate.Revocation(sql.FieldNotIn(FieldRevoked, vs...))
}

// RevokedGT applies the GT predicate on the "revoked" field.
func RevokedGT(v time.Time) predicate.Revocation {
	return predicate.Revocation(sql.FieldGT(FieldRevoked, v))
}

// RevokedGTE applies the GTE predicate on the "revoked" field.
func RevokedGTE(v time.Time) predicate.Revocation {
	return predicate.Revocation(sql.FieldGTE(FieldRevoked, v))
}

// RevokedLT applies the LT predicate on the "revoked" field.
func RevokedLT(v time.Time) predicate.Revocation {
	return predicate.Revocation(sql.FieldLT(FieldRevoked, v))
}

// RevokedLTE applies the LTE predicate on the "revoked" field.
func RevokedLTE(v time.Time) predicate.Revocation {
	return predicate.Revocation(sql.FieldLTE(FieldRevoked, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Revocation) predicate.Revocation {
	return predicate.Revocation(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Revocation) predicate.Revocation {
	return predicate.Revocation(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Revocation) predicate.Revocation {
	return predicate.Revocation(sql.NotPredicates(p))
}
