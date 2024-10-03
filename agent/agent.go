// Code generated by ent, DO NOT EDIT.

package agent

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the agent type in the database.
	Label = "agent"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "oid"
	// FieldOs holds the string denoting the os field in the database.
	FieldOs = "os"
	// FieldHostname holds the string denoting the hostname field in the database.
	FieldHostname = "hostname"
	// FieldVersion holds the string denoting the version field in the database.
	FieldVersion = "version"
	// FieldIP holds the string denoting the ip field in the database.
	FieldIP = "ip"
	// FieldFirstContact holds the string denoting the first_contact field in the database.
	FieldFirstContact = "first_contact"
	// FieldLastContact holds the string denoting the last_contact field in the database.
	FieldLastContact = "last_contact"
	// FieldEnabled holds the string denoting the enabled field in the database.
	FieldEnabled = "enabled"
	// FieldVnc holds the string denoting the vnc field in the database.
	FieldVnc = "vnc"
	// EdgeComputer holds the string denoting the computer edge name in mutations.
	EdgeComputer = "computer"
	// EdgeOperatingsystem holds the string denoting the operatingsystem edge name in mutations.
	EdgeOperatingsystem = "operatingsystem"
	// EdgeSystemupdate holds the string denoting the systemupdate edge name in mutations.
	EdgeSystemupdate = "systemupdate"
	// EdgeAntivirus holds the string denoting the antivirus edge name in mutations.
	EdgeAntivirus = "antivirus"
	// EdgeLogicaldisks holds the string denoting the logicaldisks edge name in mutations.
	EdgeLogicaldisks = "logicaldisks"
	// EdgeApps holds the string denoting the apps edge name in mutations.
	EdgeApps = "apps"
	// EdgeMonitors holds the string denoting the monitors edge name in mutations.
	EdgeMonitors = "monitors"
	// EdgeShares holds the string denoting the shares edge name in mutations.
	EdgeShares = "shares"
	// EdgePrinters holds the string denoting the printers edge name in mutations.
	EdgePrinters = "printers"
	// EdgeNetworkadapters holds the string denoting the networkadapters edge name in mutations.
	EdgeNetworkadapters = "networkadapters"
	// ComputerFieldID holds the string denoting the ID field of the Computer.
	ComputerFieldID = "id"
	// OperatingSystemFieldID holds the string denoting the ID field of the OperatingSystem.
	OperatingSystemFieldID = "id"
	// SystemUpdateFieldID holds the string denoting the ID field of the SystemUpdate.
	SystemUpdateFieldID = "id"
	// AntivirusFieldID holds the string denoting the ID field of the Antivirus.
	AntivirusFieldID = "id"
	// LogicalDiskFieldID holds the string denoting the ID field of the LogicalDisk.
	LogicalDiskFieldID = "id"
	// AppFieldID holds the string denoting the ID field of the App.
	AppFieldID = "id"
	// MonitorFieldID holds the string denoting the ID field of the Monitor.
	MonitorFieldID = "id"
	// ShareFieldID holds the string denoting the ID field of the Share.
	ShareFieldID = "id"
	// PrinterFieldID holds the string denoting the ID field of the Printer.
	PrinterFieldID = "id"
	// NetworkAdapterFieldID holds the string denoting the ID field of the NetworkAdapter.
	NetworkAdapterFieldID = "id"
	// Table holds the table name of the agent in the database.
	Table = "agents"
	// ComputerTable is the table that holds the computer relation/edge.
	ComputerTable = "computers"
	// ComputerInverseTable is the table name for the Computer entity.
	// It exists in this package in order to avoid circular dependency with the "computer" package.
	ComputerInverseTable = "computers"
	// ComputerColumn is the table column denoting the computer relation/edge.
	ComputerColumn = "agent_computer"
	// OperatingsystemTable is the table that holds the operatingsystem relation/edge.
	OperatingsystemTable = "operating_systems"
	// OperatingsystemInverseTable is the table name for the OperatingSystem entity.
	// It exists in this package in order to avoid circular dependency with the "operatingsystem" package.
	OperatingsystemInverseTable = "operating_systems"
	// OperatingsystemColumn is the table column denoting the operatingsystem relation/edge.
	OperatingsystemColumn = "agent_operatingsystem"
	// SystemupdateTable is the table that holds the systemupdate relation/edge.
	SystemupdateTable = "system_updates"
	// SystemupdateInverseTable is the table name for the SystemUpdate entity.
	// It exists in this package in order to avoid circular dependency with the "systemupdate" package.
	SystemupdateInverseTable = "system_updates"
	// SystemupdateColumn is the table column denoting the systemupdate relation/edge.
	SystemupdateColumn = "agent_systemupdate"
	// AntivirusTable is the table that holds the antivirus relation/edge.
	AntivirusTable = "antiviri"
	// AntivirusInverseTable is the table name for the Antivirus entity.
	// It exists in this package in order to avoid circular dependency with the "antivirus" package.
	AntivirusInverseTable = "antiviri"
	// AntivirusColumn is the table column denoting the antivirus relation/edge.
	AntivirusColumn = "agent_antivirus"
	// LogicaldisksTable is the table that holds the logicaldisks relation/edge.
	LogicaldisksTable = "logical_disks"
	// LogicaldisksInverseTable is the table name for the LogicalDisk entity.
	// It exists in this package in order to avoid circular dependency with the "logicaldisk" package.
	LogicaldisksInverseTable = "logical_disks"
	// LogicaldisksColumn is the table column denoting the logicaldisks relation/edge.
	LogicaldisksColumn = "agent_logicaldisks"
	// AppsTable is the table that holds the apps relation/edge.
	AppsTable = "apps"
	// AppsInverseTable is the table name for the App entity.
	// It exists in this package in order to avoid circular dependency with the "app" package.
	AppsInverseTable = "apps"
	// AppsColumn is the table column denoting the apps relation/edge.
	AppsColumn = "agent_apps"
	// MonitorsTable is the table that holds the monitors relation/edge.
	MonitorsTable = "monitors"
	// MonitorsInverseTable is the table name for the Monitor entity.
	// It exists in this package in order to avoid circular dependency with the "monitor" package.
	MonitorsInverseTable = "monitors"
	// MonitorsColumn is the table column denoting the monitors relation/edge.
	MonitorsColumn = "agent_monitors"
	// SharesTable is the table that holds the shares relation/edge.
	SharesTable = "shares"
	// SharesInverseTable is the table name for the Share entity.
	// It exists in this package in order to avoid circular dependency with the "share" package.
	SharesInverseTable = "shares"
	// SharesColumn is the table column denoting the shares relation/edge.
	SharesColumn = "agent_shares"
	// PrintersTable is the table that holds the printers relation/edge.
	PrintersTable = "printers"
	// PrintersInverseTable is the table name for the Printer entity.
	// It exists in this package in order to avoid circular dependency with the "printer" package.
	PrintersInverseTable = "printers"
	// PrintersColumn is the table column denoting the printers relation/edge.
	PrintersColumn = "agent_printers"
	// NetworkadaptersTable is the table that holds the networkadapters relation/edge.
	NetworkadaptersTable = "network_adapters"
	// NetworkadaptersInverseTable is the table name for the NetworkAdapter entity.
	// It exists in this package in order to avoid circular dependency with the "networkadapter" package.
	NetworkadaptersInverseTable = "network_adapters"
	// NetworkadaptersColumn is the table column denoting the networkadapters relation/edge.
	NetworkadaptersColumn = "agent_networkadapters"
)

// Columns holds all SQL columns for agent fields.
var Columns = []string{
	FieldID,
	FieldOs,
	FieldHostname,
	FieldVersion,
	FieldIP,
	FieldFirstContact,
	FieldLastContact,
	FieldEnabled,
	FieldVnc,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// OsValidator is a validator for the "os" field. It is called by the builders before save.
	OsValidator func(string) error
	// HostnameValidator is a validator for the "hostname" field. It is called by the builders before save.
	HostnameValidator func(string) error
	// VersionValidator is a validator for the "version" field. It is called by the builders before save.
	VersionValidator func(string) error
	// DefaultIP holds the default value on creation for the "ip" field.
	DefaultIP string
	// DefaultEnabled holds the default value on creation for the "enabled" field.
	DefaultEnabled bool
	// DefaultVnc holds the default value on creation for the "vnc" field.
	DefaultVnc string
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// OrderOption defines the ordering options for the Agent queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByOs orders the results by the os field.
func ByOs(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOs, opts...).ToFunc()
}

// ByHostname orders the results by the hostname field.
func ByHostname(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHostname, opts...).ToFunc()
}

// ByVersion orders the results by the version field.
func ByVersion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVersion, opts...).ToFunc()
}

// ByIP orders the results by the ip field.
func ByIP(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIP, opts...).ToFunc()
}

// ByFirstContact orders the results by the first_contact field.
func ByFirstContact(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFirstContact, opts...).ToFunc()
}

// ByLastContact orders the results by the last_contact field.
func ByLastContact(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastContact, opts...).ToFunc()
}

// ByEnabled orders the results by the enabled field.
func ByEnabled(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEnabled, opts...).ToFunc()
}

// ByVnc orders the results by the vnc field.
func ByVnc(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVnc, opts...).ToFunc()
}

// ByComputerField orders the results by computer field.
func ByComputerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newComputerStep(), sql.OrderByField(field, opts...))
	}
}

// ByOperatingsystemField orders the results by operatingsystem field.
func ByOperatingsystemField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOperatingsystemStep(), sql.OrderByField(field, opts...))
	}
}

// BySystemupdateField orders the results by systemupdate field.
func BySystemupdateField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSystemupdateStep(), sql.OrderByField(field, opts...))
	}
}

// ByAntivirusField orders the results by antivirus field.
func ByAntivirusField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAntivirusStep(), sql.OrderByField(field, opts...))
	}
}

// ByLogicaldisksCount orders the results by logicaldisks count.
func ByLogicaldisksCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newLogicaldisksStep(), opts...)
	}
}

// ByLogicaldisks orders the results by logicaldisks terms.
func ByLogicaldisks(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newLogicaldisksStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByAppsCount orders the results by apps count.
func ByAppsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAppsStep(), opts...)
	}
}

// ByApps orders the results by apps terms.
func ByApps(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAppsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByMonitorsCount orders the results by monitors count.
func ByMonitorsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMonitorsStep(), opts...)
	}
}

// ByMonitors orders the results by monitors terms.
func ByMonitors(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMonitorsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// BySharesCount orders the results by shares count.
func BySharesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSharesStep(), opts...)
	}
}

// ByShares orders the results by shares terms.
func ByShares(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSharesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPrintersCount orders the results by printers count.
func ByPrintersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPrintersStep(), opts...)
	}
}

// ByPrinters orders the results by printers terms.
func ByPrinters(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPrintersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByNetworkadaptersCount orders the results by networkadapters count.
func ByNetworkadaptersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newNetworkadaptersStep(), opts...)
	}
}

// ByNetworkadapters orders the results by networkadapters terms.
func ByNetworkadapters(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newNetworkadaptersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newComputerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ComputerInverseTable, ComputerFieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, ComputerTable, ComputerColumn),
	)
}
func newOperatingsystemStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OperatingsystemInverseTable, OperatingSystemFieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, OperatingsystemTable, OperatingsystemColumn),
	)
}
func newSystemupdateStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SystemupdateInverseTable, SystemUpdateFieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, SystemupdateTable, SystemupdateColumn),
	)
}
func newAntivirusStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AntivirusInverseTable, AntivirusFieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, AntivirusTable, AntivirusColumn),
	)
}
func newLogicaldisksStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(LogicaldisksInverseTable, LogicalDiskFieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, LogicaldisksTable, LogicaldisksColumn),
	)
}
func newAppsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AppsInverseTable, AppFieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, AppsTable, AppsColumn),
	)
}
func newMonitorsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MonitorsInverseTable, MonitorFieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, MonitorsTable, MonitorsColumn),
	)
}
func newSharesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SharesInverseTable, ShareFieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, SharesTable, SharesColumn),
	)
}
func newPrintersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PrintersInverseTable, PrinterFieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PrintersTable, PrintersColumn),
	)
}
func newNetworkadaptersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(NetworkadaptersInverseTable, NetworkAdapterFieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, NetworkadaptersTable, NetworkadaptersColumn),
	)
}
