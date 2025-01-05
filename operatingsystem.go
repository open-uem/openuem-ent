// Code generated by ent, DO NOT EDIT.

package openuem_ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/open-uem/openuem_ent/agent"
	"github.com/open-uem/openuem_ent/operatingsystem"
)

// OperatingSystem is the model entity for the OperatingSystem schema.
type OperatingSystem struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// Version holds the value of the "version" field.
	Version string `json:"version,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Edition holds the value of the "edition" field.
	Edition string `json:"edition,omitempty"`
	// InstallDate holds the value of the "install_date" field.
	InstallDate time.Time `json:"install_date,omitempty"`
	// Arch holds the value of the "arch" field.
	Arch string `json:"arch,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// LastBootupTime holds the value of the "last_bootup_time" field.
	LastBootupTime time.Time `json:"last_bootup_time,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OperatingSystemQuery when eager-loading is set.
	Edges                 OperatingSystemEdges `json:"edges"`
	agent_operatingsystem *string
	selectValues          sql.SelectValues
}

// OperatingSystemEdges holds the relations/edges for other nodes in the graph.
type OperatingSystemEdges struct {
	// Owner holds the value of the owner edge.
	Owner *Agent `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OperatingSystemEdges) OwnerOrErr() (*Agent, error) {
	if e.Owner != nil {
		return e.Owner, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: agent.Label}
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OperatingSystem) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case operatingsystem.FieldID:
			values[i] = new(sql.NullInt64)
		case operatingsystem.FieldType, operatingsystem.FieldVersion, operatingsystem.FieldDescription, operatingsystem.FieldEdition, operatingsystem.FieldArch, operatingsystem.FieldUsername:
			values[i] = new(sql.NullString)
		case operatingsystem.FieldInstallDate, operatingsystem.FieldLastBootupTime:
			values[i] = new(sql.NullTime)
		case operatingsystem.ForeignKeys[0]: // agent_operatingsystem
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OperatingSystem fields.
func (os *OperatingSystem) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case operatingsystem.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			os.ID = int(value.Int64)
		case operatingsystem.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				os.Type = value.String
			}
		case operatingsystem.FieldVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				os.Version = value.String
			}
		case operatingsystem.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				os.Description = value.String
			}
		case operatingsystem.FieldEdition:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field edition", values[i])
			} else if value.Valid {
				os.Edition = value.String
			}
		case operatingsystem.FieldInstallDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field install_date", values[i])
			} else if value.Valid {
				os.InstallDate = value.Time
			}
		case operatingsystem.FieldArch:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field arch", values[i])
			} else if value.Valid {
				os.Arch = value.String
			}
		case operatingsystem.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				os.Username = value.String
			}
		case operatingsystem.FieldLastBootupTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_bootup_time", values[i])
			} else if value.Valid {
				os.LastBootupTime = value.Time
			}
		case operatingsystem.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field agent_operatingsystem", values[i])
			} else if value.Valid {
				os.agent_operatingsystem = new(string)
				*os.agent_operatingsystem = value.String
			}
		default:
			os.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the OperatingSystem.
// This includes values selected through modifiers, order, etc.
func (os *OperatingSystem) Value(name string) (ent.Value, error) {
	return os.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the OperatingSystem entity.
func (os *OperatingSystem) QueryOwner() *AgentQuery {
	return NewOperatingSystemClient(os.config).QueryOwner(os)
}

// Update returns a builder for updating this OperatingSystem.
// Note that you need to call OperatingSystem.Unwrap() before calling this method if this OperatingSystem
// was returned from a transaction, and the transaction was committed or rolled back.
func (os *OperatingSystem) Update() *OperatingSystemUpdateOne {
	return NewOperatingSystemClient(os.config).UpdateOne(os)
}

// Unwrap unwraps the OperatingSystem entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (os *OperatingSystem) Unwrap() *OperatingSystem {
	_tx, ok := os.config.driver.(*txDriver)
	if !ok {
		panic("openuem_ent: OperatingSystem is not a transactional entity")
	}
	os.config.driver = _tx.drv
	return os
}

// String implements the fmt.Stringer.
func (os *OperatingSystem) String() string {
	var builder strings.Builder
	builder.WriteString("OperatingSystem(")
	builder.WriteString(fmt.Sprintf("id=%v, ", os.ID))
	builder.WriteString("type=")
	builder.WriteString(os.Type)
	builder.WriteString(", ")
	builder.WriteString("version=")
	builder.WriteString(os.Version)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(os.Description)
	builder.WriteString(", ")
	builder.WriteString("edition=")
	builder.WriteString(os.Edition)
	builder.WriteString(", ")
	builder.WriteString("install_date=")
	builder.WriteString(os.InstallDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("arch=")
	builder.WriteString(os.Arch)
	builder.WriteString(", ")
	builder.WriteString("username=")
	builder.WriteString(os.Username)
	builder.WriteString(", ")
	builder.WriteString("last_bootup_time=")
	builder.WriteString(os.LastBootupTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// OperatingSystems is a parsable slice of OperatingSystem.
type OperatingSystems []*OperatingSystem
