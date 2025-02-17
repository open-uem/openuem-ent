// Code generated by ent, DO NOT EDIT.

package update

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the update type in the database.
	Label = "update"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldDate holds the string denoting the date field in the database.
	FieldDate = "date"
	// FieldSupportURL holds the string denoting the support_url field in the database.
	FieldSupportURL = "support_url"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// AgentFieldID holds the string denoting the ID field of the Agent.
	AgentFieldID = "oid"
	// Table holds the table name of the update in the database.
	Table = "updates"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "updates"
	// OwnerInverseTable is the table name for the Agent entity.
	// It exists in this package in order to avoid circular dependency with the "agent" package.
	OwnerInverseTable = "agents"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "agent_updates"
)

// Columns holds all SQL columns for update fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldDate,
	FieldSupportURL,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "updates"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"agent_updates",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Update queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByDate orders the results by the date field.
func ByDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDate, opts...).ToFunc()
}

// BySupportURL orders the results by the support_url field.
func BySupportURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSupportURL, opts...).ToFunc()
}

// ByOwnerField orders the results by owner field.
func ByOwnerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOwnerStep(), sql.OrderByField(field, opts...))
	}
}
func newOwnerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OwnerInverseTable, AgentFieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
	)
}
