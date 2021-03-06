// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team10/app/ent/nurse"
	"github.com/team10/app/ent/user"
)

// Nurse is the model entity for the Nurse schema.
type Nurse struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Nursinglicense holds the value of the "nursinglicense" field.
	Nursinglicense string `json:"nursinglicense,omitempty"`
	// Position holds the value of the "position" field.
	Position string `json:"position,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the NurseQuery when eager-loading is set.
	Edges   NurseEdges `json:"edges"`
	user_id *int
}

// NurseEdges holds the relations/edges for other nodes in the graph.
type NurseEdges struct {
	// Historytaking holds the value of the historytaking edge.
	Historytaking []*Historytaking
	// User holds the value of the user edge.
	User *User
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// HistorytakingOrErr returns the Historytaking value or an error if the edge
// was not loaded in eager-loading.
func (e NurseEdges) HistorytakingOrErr() ([]*Historytaking, error) {
	if e.loadedTypes[0] {
		return e.Historytaking, nil
	}
	return nil, &NotLoadedError{edge: "historytaking"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e NurseEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Nurse) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // name
		&sql.NullString{}, // nursinglicense
		&sql.NullString{}, // position
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Nurse) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // user_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Nurse fields.
func (n *Nurse) assignValues(values ...interface{}) error {
	if m, n := len(values), len(nurse.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	n.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		n.Name = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field nursinglicense", values[1])
	} else if value.Valid {
		n.Nursinglicense = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field position", values[2])
	} else if value.Valid {
		n.Position = value.String
	}
	values = values[3:]
	if len(values) == len(nurse.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field user_id", value)
		} else if value.Valid {
			n.user_id = new(int)
			*n.user_id = int(value.Int64)
		}
	}
	return nil
}

// QueryHistorytaking queries the historytaking edge of the Nurse.
func (n *Nurse) QueryHistorytaking() *HistorytakingQuery {
	return (&NurseClient{config: n.config}).QueryHistorytaking(n)
}

// QueryUser queries the user edge of the Nurse.
func (n *Nurse) QueryUser() *UserQuery {
	return (&NurseClient{config: n.config}).QueryUser(n)
}

// Update returns a builder for updating this Nurse.
// Note that, you need to call Nurse.Unwrap() before calling this method, if this Nurse
// was returned from a transaction, and the transaction was committed or rolled back.
func (n *Nurse) Update() *NurseUpdateOne {
	return (&NurseClient{config: n.config}).UpdateOne(n)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (n *Nurse) Unwrap() *Nurse {
	tx, ok := n.config.driver.(*txDriver)
	if !ok {
		panic("ent: Nurse is not a transactional entity")
	}
	n.config.driver = tx.drv
	return n
}

// String implements the fmt.Stringer.
func (n *Nurse) String() string {
	var builder strings.Builder
	builder.WriteString("Nurse(")
	builder.WriteString(fmt.Sprintf("id=%v", n.ID))
	builder.WriteString(", name=")
	builder.WriteString(n.Name)
	builder.WriteString(", nursinglicense=")
	builder.WriteString(n.Nursinglicense)
	builder.WriteString(", position=")
	builder.WriteString(n.Position)
	builder.WriteByte(')')
	return builder.String()
}

// Nurses is a parsable slice of Nurse.
type Nurses []*Nurse

func (n Nurses) config(cfg config) {
	for _i := range n {
		n[_i].config = cfg
	}
}
