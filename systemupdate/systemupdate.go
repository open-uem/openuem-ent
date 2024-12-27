// Code generated by ent, DO NOT EDIT.

package systemupdate

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the systemupdate type in the database.
	Label = "system_update"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSystemUpdateStatus holds the string denoting the system_update_status field in the database.
	FieldSystemUpdateStatus = "system_update_status"
	// FieldLastInstall holds the string denoting the last_install field in the database.
	FieldLastInstall = "last_install"
	// FieldLastSearch holds the string denoting the last_search field in the database.
	FieldLastSearch = "last_search"
	// FieldPendingUpdates holds the string denoting the pending_updates field in the database.
	FieldPendingUpdates = "pending_updates"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// AgentFieldID holds the string denoting the ID field of the Agent.
	AgentFieldID = "oid"
	// Table holds the table name of the systemupdate in the database.
	Table = "system_updates"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "system_updates"
	// OwnerInverseTable is the table name for the Agent entity.
	// It exists in this package in order to avoid circular dependency with the "agent" package.
	OwnerInverseTable = "agents"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "agent_systemupdate"
)

// Columns holds all SQL columns for systemupdate fields.
var Columns = []string{
	FieldID,
	FieldSystemUpdateStatus,
	FieldLastInstall,
	FieldLastSearch,
	FieldPendingUpdates,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "system_updates"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"agent_systemupdate",
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

// OrderOption defines the ordering options for the SystemUpdate queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// BySystemUpdateStatus orders the results by the system_update_status field.
func BySystemUpdateStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSystemUpdateStatus, opts...).ToFunc()
}

// ByLastInstall orders the results by the last_install field.
func ByLastInstall(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastInstall, opts...).ToFunc()
}

// ByLastSearch orders the results by the last_search field.
func ByLastSearch(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastSearch, opts...).ToFunc()
}

// ByPendingUpdates orders the results by the pending_updates field.
func ByPendingUpdates(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPendingUpdates, opts...).ToFunc()
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
		sqlgraph.Edge(sqlgraph.O2O, true, OwnerTable, OwnerColumn),
	)
}
