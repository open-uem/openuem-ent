// Code generated by ent, DO NOT EDIT.

package certificate

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/open-uem/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Certificate {
	return predicate.Certificate(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Certificate {
	return predicate.Certificate(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Certificate {
	return predicate.Certificate(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Certificate {
	return predicate.Certificate(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Certificate {
	return predicate.Certificate(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Certificate {
	return predicate.Certificate(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Certificate {
	return predicate.Certificate(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Certificate {
	return predicate.Certificate(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Certificate {
	return predicate.Certificate(sql.FieldLTE(FieldID, id))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldEQ(FieldDescription, v))
}

// Expiry applies equality check predicate on the "expiry" field. It's identical to ExpiryEQ.
func Expiry(v time.Time) predicate.Certificate {
	return predicate.Certificate(sql.FieldEQ(FieldExpiry, v))
}

// UID applies equality check predicate on the "uid" field. It's identical to UIDEQ.
func UID(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldEQ(FieldUID, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v Type) predicate.Certificate {
	return predicate.Certificate(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v Type) predicate.Certificate {
	return predicate.Certificate(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...Type) predicate.Certificate {
	return predicate.Certificate(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...Type) predicate.Certificate {
	return predicate.Certificate(sql.FieldNotIn(FieldType, vs...))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Certificate {
	return predicate.Certificate(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Certificate {
	return predicate.Certificate(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Certificate {
	return predicate.Certificate(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Certificate {
	return predicate.Certificate(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldContainsFold(FieldDescription, v))
}

// ExpiryEQ applies the EQ predicate on the "expiry" field.
func ExpiryEQ(v time.Time) predicate.Certificate {
	return predicate.Certificate(sql.FieldEQ(FieldExpiry, v))
}

// ExpiryNEQ applies the NEQ predicate on the "expiry" field.
func ExpiryNEQ(v time.Time) predicate.Certificate {
	return predicate.Certificate(sql.FieldNEQ(FieldExpiry, v))
}

// ExpiryIn applies the In predicate on the "expiry" field.
func ExpiryIn(vs ...time.Time) predicate.Certificate {
	return predicate.Certificate(sql.FieldIn(FieldExpiry, vs...))
}

// ExpiryNotIn applies the NotIn predicate on the "expiry" field.
func ExpiryNotIn(vs ...time.Time) predicate.Certificate {
	return predicate.Certificate(sql.FieldNotIn(FieldExpiry, vs...))
}

// ExpiryGT applies the GT predicate on the "expiry" field.
func ExpiryGT(v time.Time) predicate.Certificate {
	return predicate.Certificate(sql.FieldGT(FieldExpiry, v))
}

// ExpiryGTE applies the GTE predicate on the "expiry" field.
func ExpiryGTE(v time.Time) predicate.Certificate {
	return predicate.Certificate(sql.FieldGTE(FieldExpiry, v))
}

// ExpiryLT applies the LT predicate on the "expiry" field.
func ExpiryLT(v time.Time) predicate.Certificate {
	return predicate.Certificate(sql.FieldLT(FieldExpiry, v))
}

// ExpiryLTE applies the LTE predicate on the "expiry" field.
func ExpiryLTE(v time.Time) predicate.Certificate {
	return predicate.Certificate(sql.FieldLTE(FieldExpiry, v))
}

// ExpiryIsNil applies the IsNil predicate on the "expiry" field.
func ExpiryIsNil() predicate.Certificate {
	return predicate.Certificate(sql.FieldIsNull(FieldExpiry))
}

// ExpiryNotNil applies the NotNil predicate on the "expiry" field.
func ExpiryNotNil() predicate.Certificate {
	return predicate.Certificate(sql.FieldNotNull(FieldExpiry))
}

// UIDEQ applies the EQ predicate on the "uid" field.
func UIDEQ(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldEQ(FieldUID, v))
}

// UIDNEQ applies the NEQ predicate on the "uid" field.
func UIDNEQ(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldNEQ(FieldUID, v))
}

// UIDIn applies the In predicate on the "uid" field.
func UIDIn(vs ...string) predicate.Certificate {
	return predicate.Certificate(sql.FieldIn(FieldUID, vs...))
}

// UIDNotIn applies the NotIn predicate on the "uid" field.
func UIDNotIn(vs ...string) predicate.Certificate {
	return predicate.Certificate(sql.FieldNotIn(FieldUID, vs...))
}

// UIDGT applies the GT predicate on the "uid" field.
func UIDGT(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldGT(FieldUID, v))
}

// UIDGTE applies the GTE predicate on the "uid" field.
func UIDGTE(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldGTE(FieldUID, v))
}

// UIDLT applies the LT predicate on the "uid" field.
func UIDLT(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldLT(FieldUID, v))
}

// UIDLTE applies the LTE predicate on the "uid" field.
func UIDLTE(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldLTE(FieldUID, v))
}

// UIDContains applies the Contains predicate on the "uid" field.
func UIDContains(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldContains(FieldUID, v))
}

// UIDHasPrefix applies the HasPrefix predicate on the "uid" field.
func UIDHasPrefix(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldHasPrefix(FieldUID, v))
}

// UIDHasSuffix applies the HasSuffix predicate on the "uid" field.
func UIDHasSuffix(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldHasSuffix(FieldUID, v))
}

// UIDIsNil applies the IsNil predicate on the "uid" field.
func UIDIsNil() predicate.Certificate {
	return predicate.Certificate(sql.FieldIsNull(FieldUID))
}

// UIDNotNil applies the NotNil predicate on the "uid" field.
func UIDNotNil() predicate.Certificate {
	return predicate.Certificate(sql.FieldNotNull(FieldUID))
}

// UIDEqualFold applies the EqualFold predicate on the "uid" field.
func UIDEqualFold(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldEqualFold(FieldUID, v))
}

// UIDContainsFold applies the ContainsFold predicate on the "uid" field.
func UIDContainsFold(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldContainsFold(FieldUID, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Certificate) predicate.Certificate {
	return predicate.Certificate(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Certificate) predicate.Certificate {
	return predicate.Certificate(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Certificate) predicate.Certificate {
	return predicate.Certificate(sql.NotPredicates(p))
}
