// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/jamesw201/go-starter/example/ent/keyboard"
	"github.com/jamesw201/go-starter/example/ent/keycapmodel"
	"github.com/jamesw201/go-starter/example/ent/switchmodel"
)

// Keyboard is the model entity for the Keyboard schema.
type Keyboard struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Price holds the value of the "price" field.
	Price int64 `json:"price,omitempty"`
	// Discount holds the value of the "discount" field.
	Discount *int64 `json:"discount,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the KeyboardQuery when eager-loading is set.
	Edges             KeyboardEdges `json:"edges"`
	keyboard_switches *int64
	keyboard_keycaps  *int64
	selectValues      sql.SelectValues
}

// KeyboardEdges holds the relations/edges for other nodes in the graph.
type KeyboardEdges struct {
	// Switches holds the value of the switches edge.
	Switches *SwitchModel `json:"switches,omitempty"`
	// Keycaps holds the value of the keycaps edge.
	Keycaps *KeycapModel `json:"keycaps,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// SwitchesOrErr returns the Switches value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e KeyboardEdges) SwitchesOrErr() (*SwitchModel, error) {
	if e.loadedTypes[0] {
		if e.Switches == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: switchmodel.Label}
		}
		return e.Switches, nil
	}
	return nil, &NotLoadedError{edge: "switches"}
}

// KeycapsOrErr returns the Keycaps value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e KeyboardEdges) KeycapsOrErr() (*KeycapModel, error) {
	if e.loadedTypes[1] {
		if e.Keycaps == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: keycapmodel.Label}
		}
		return e.Keycaps, nil
	}
	return nil, &NotLoadedError{edge: "keycaps"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Keyboard) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case keyboard.FieldID, keyboard.FieldPrice, keyboard.FieldDiscount:
			values[i] = new(sql.NullInt64)
		case keyboard.FieldName:
			values[i] = new(sql.NullString)
		case keyboard.ForeignKeys[0]: // keyboard_switches
			values[i] = new(sql.NullInt64)
		case keyboard.ForeignKeys[1]: // keyboard_keycaps
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Keyboard fields.
func (k *Keyboard) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case keyboard.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			k.ID = int64(value.Int64)
		case keyboard.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				k.Name = value.String
			}
		case keyboard.FieldPrice:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field price", values[i])
			} else if value.Valid {
				k.Price = value.Int64
			}
		case keyboard.FieldDiscount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field discount", values[i])
			} else if value.Valid {
				k.Discount = new(int64)
				*k.Discount = value.Int64
			}
		case keyboard.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field keyboard_switches", value)
			} else if value.Valid {
				k.keyboard_switches = new(int64)
				*k.keyboard_switches = int64(value.Int64)
			}
		case keyboard.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field keyboard_keycaps", value)
			} else if value.Valid {
				k.keyboard_keycaps = new(int64)
				*k.keyboard_keycaps = int64(value.Int64)
			}
		default:
			k.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Keyboard.
// This includes values selected through modifiers, order, etc.
func (k *Keyboard) Value(name string) (ent.Value, error) {
	return k.selectValues.Get(name)
}

// QuerySwitches queries the "switches" edge of the Keyboard entity.
func (k *Keyboard) QuerySwitches() *SwitchModelQuery {
	return NewKeyboardClient(k.config).QuerySwitches(k)
}

// QueryKeycaps queries the "keycaps" edge of the Keyboard entity.
func (k *Keyboard) QueryKeycaps() *KeycapModelQuery {
	return NewKeyboardClient(k.config).QueryKeycaps(k)
}

// Update returns a builder for updating this Keyboard.
// Note that you need to call Keyboard.Unwrap() before calling this method if this Keyboard
// was returned from a transaction, and the transaction was committed or rolled back.
func (k *Keyboard) Update() *KeyboardUpdateOne {
	return NewKeyboardClient(k.config).UpdateOne(k)
}

// Unwrap unwraps the Keyboard entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (k *Keyboard) Unwrap() *Keyboard {
	_tx, ok := k.config.driver.(*txDriver)
	if !ok {
		panic("ent: Keyboard is not a transactional entity")
	}
	k.config.driver = _tx.drv
	return k
}

// String implements the fmt.Stringer.
func (k *Keyboard) String() string {
	var builder strings.Builder
	builder.WriteString("Keyboard(")
	builder.WriteString(fmt.Sprintf("id=%v, ", k.ID))
	builder.WriteString("name=")
	builder.WriteString(k.Name)
	builder.WriteString(", ")
	builder.WriteString("price=")
	builder.WriteString(fmt.Sprintf("%v", k.Price))
	builder.WriteString(", ")
	if v := k.Discount; v != nil {
		builder.WriteString("discount=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Keyboards is a parsable slice of Keyboard.
type Keyboards []*Keyboard
