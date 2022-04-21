// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/ogen-go/ent2ogen/example/ent/city"
	"github.com/ogen-go/ent2ogen/example/ent/predicate"
	"github.com/ogen-go/ent2ogen/example/ent/user"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeCity = "City"
	TypeUser = "User"
)

// CityMutation represents an operation that mutates the City nodes in the graph.
type CityMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	created_at    *time.Time
	updated_at    *time.Time
	name          *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*City, error)
	predicates    []predicate.City
}

var _ ent.Mutation = (*CityMutation)(nil)

// cityOption allows management of the mutation configuration using functional options.
type cityOption func(*CityMutation)

// newCityMutation creates new mutation for the City entity.
func newCityMutation(c config, op Op, opts ...cityOption) *CityMutation {
	m := &CityMutation{
		config:        c,
		op:            op,
		typ:           TypeCity,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withCityID sets the ID field of the mutation.
func withCityID(id uuid.UUID) cityOption {
	return func(m *CityMutation) {
		var (
			err   error
			once  sync.Once
			value *City
		)
		m.oldValue = func(ctx context.Context) (*City, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().City.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withCity sets the old City of the mutation.
func withCity(node *City) cityOption {
	return func(m *CityMutation) {
		m.oldValue = func(context.Context) (*City, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m CityMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m CityMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of City entities.
func (m *CityMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *CityMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *CityMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().City.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreatedAt sets the "created_at" field.
func (m *CityMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *CityMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the City entity.
// If the City object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CityMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *CityMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *CityMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *CityMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the City entity.
// If the City object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CityMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *CityMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetName sets the "name" field.
func (m *CityMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *CityMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the City entity.
// If the City object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CityMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *CityMutation) ResetName() {
	m.name = nil
}

// Where appends a list predicates to the CityMutation builder.
func (m *CityMutation) Where(ps ...predicate.City) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *CityMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (City).
func (m *CityMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *CityMutation) Fields() []string {
	fields := make([]string, 0, 3)
	if m.created_at != nil {
		fields = append(fields, city.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, city.FieldUpdatedAt)
	}
	if m.name != nil {
		fields = append(fields, city.FieldName)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *CityMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case city.FieldCreatedAt:
		return m.CreatedAt()
	case city.FieldUpdatedAt:
		return m.UpdatedAt()
	case city.FieldName:
		return m.Name()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *CityMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case city.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case city.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case city.FieldName:
		return m.OldName(ctx)
	}
	return nil, fmt.Errorf("unknown City field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CityMutation) SetField(name string, value ent.Value) error {
	switch name {
	case city.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case city.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case city.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown City field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *CityMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *CityMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CityMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown City numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *CityMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *CityMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *CityMutation) ClearField(name string) error {
	return fmt.Errorf("unknown City nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *CityMutation) ResetField(name string) error {
	switch name {
	case city.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case city.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case city.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown City field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *CityMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *CityMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *CityMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *CityMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *CityMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *CityMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *CityMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown City unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *CityMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown City edge %s", name)
}

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op                     Op
	typ                    string
	id                     *uuid.UUID
	created_at             *time.Time
	updated_at             *time.Time
	first_name             *string
	last_name              *string
	username               *string
	optional_nullable_bool *bool
	clearedFields          map[string]struct{}
	city                   *uuid.UUID
	clearedcity            bool
	friends                map[uuid.UUID]struct{}
	removedfriends         map[uuid.UUID]struct{}
	clearedfriends         bool
	done                   bool
	oldValue               func(context.Context) (*User, error)
	predicates             []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id uuid.UUID) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of User entities.
func (m *UserMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().User.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreatedAt sets the "created_at" field.
func (m *UserMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *UserMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *UserMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *UserMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *UserMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *UserMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetFirstName sets the "first_name" field.
func (m *UserMutation) SetFirstName(s string) {
	m.first_name = &s
}

// FirstName returns the value of the "first_name" field in the mutation.
func (m *UserMutation) FirstName() (r string, exists bool) {
	v := m.first_name
	if v == nil {
		return
	}
	return *v, true
}

// OldFirstName returns the old "first_name" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldFirstName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldFirstName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldFirstName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldFirstName: %w", err)
	}
	return oldValue.FirstName, nil
}

// ResetFirstName resets all changes to the "first_name" field.
func (m *UserMutation) ResetFirstName() {
	m.first_name = nil
}

// SetLastName sets the "last_name" field.
func (m *UserMutation) SetLastName(s string) {
	m.last_name = &s
}

// LastName returns the value of the "last_name" field in the mutation.
func (m *UserMutation) LastName() (r string, exists bool) {
	v := m.last_name
	if v == nil {
		return
	}
	return *v, true
}

// OldLastName returns the old "last_name" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldLastName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldLastName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldLastName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLastName: %w", err)
	}
	return oldValue.LastName, nil
}

// ResetLastName resets all changes to the "last_name" field.
func (m *UserMutation) ResetLastName() {
	m.last_name = nil
}

// SetUsername sets the "username" field.
func (m *UserMutation) SetUsername(s string) {
	m.username = &s
}

// Username returns the value of the "username" field in the mutation.
func (m *UserMutation) Username() (r string, exists bool) {
	v := m.username
	if v == nil {
		return
	}
	return *v, true
}

// OldUsername returns the old "username" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldUsername(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUsername is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUsername requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUsername: %w", err)
	}
	return oldValue.Username, nil
}

// ResetUsername resets all changes to the "username" field.
func (m *UserMutation) ResetUsername() {
	m.username = nil
}

// SetOptionalNullableBool sets the "optional_nullable_bool" field.
func (m *UserMutation) SetOptionalNullableBool(b bool) {
	m.optional_nullable_bool = &b
}

// OptionalNullableBool returns the value of the "optional_nullable_bool" field in the mutation.
func (m *UserMutation) OptionalNullableBool() (r bool, exists bool) {
	v := m.optional_nullable_bool
	if v == nil {
		return
	}
	return *v, true
}

// OldOptionalNullableBool returns the old "optional_nullable_bool" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldOptionalNullableBool(ctx context.Context) (v *bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldOptionalNullableBool is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldOptionalNullableBool requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldOptionalNullableBool: %w", err)
	}
	return oldValue.OptionalNullableBool, nil
}

// ClearOptionalNullableBool clears the value of the "optional_nullable_bool" field.
func (m *UserMutation) ClearOptionalNullableBool() {
	m.optional_nullable_bool = nil
	m.clearedFields[user.FieldOptionalNullableBool] = struct{}{}
}

// OptionalNullableBoolCleared returns if the "optional_nullable_bool" field was cleared in this mutation.
func (m *UserMutation) OptionalNullableBoolCleared() bool {
	_, ok := m.clearedFields[user.FieldOptionalNullableBool]
	return ok
}

// ResetOptionalNullableBool resets all changes to the "optional_nullable_bool" field.
func (m *UserMutation) ResetOptionalNullableBool() {
	m.optional_nullable_bool = nil
	delete(m.clearedFields, user.FieldOptionalNullableBool)
}

// SetCityID sets the "city" edge to the City entity by id.
func (m *UserMutation) SetCityID(id uuid.UUID) {
	m.city = &id
}

// ClearCity clears the "city" edge to the City entity.
func (m *UserMutation) ClearCity() {
	m.clearedcity = true
}

// CityCleared reports if the "city" edge to the City entity was cleared.
func (m *UserMutation) CityCleared() bool {
	return m.clearedcity
}

// CityID returns the "city" edge ID in the mutation.
func (m *UserMutation) CityID() (id uuid.UUID, exists bool) {
	if m.city != nil {
		return *m.city, true
	}
	return
}

// CityIDs returns the "city" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// CityID instead. It exists only for internal usage by the builders.
func (m *UserMutation) CityIDs() (ids []uuid.UUID) {
	if id := m.city; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetCity resets all changes to the "city" edge.
func (m *UserMutation) ResetCity() {
	m.city = nil
	m.clearedcity = false
}

// AddFriendIDs adds the "friends" edge to the User entity by ids.
func (m *UserMutation) AddFriendIDs(ids ...uuid.UUID) {
	if m.friends == nil {
		m.friends = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		m.friends[ids[i]] = struct{}{}
	}
}

// ClearFriends clears the "friends" edge to the User entity.
func (m *UserMutation) ClearFriends() {
	m.clearedfriends = true
}

// FriendsCleared reports if the "friends" edge to the User entity was cleared.
func (m *UserMutation) FriendsCleared() bool {
	return m.clearedfriends
}

// RemoveFriendIDs removes the "friends" edge to the User entity by IDs.
func (m *UserMutation) RemoveFriendIDs(ids ...uuid.UUID) {
	if m.removedfriends == nil {
		m.removedfriends = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		delete(m.friends, ids[i])
		m.removedfriends[ids[i]] = struct{}{}
	}
}

// RemovedFriends returns the removed IDs of the "friends" edge to the User entity.
func (m *UserMutation) RemovedFriendsIDs() (ids []uuid.UUID) {
	for id := range m.removedfriends {
		ids = append(ids, id)
	}
	return
}

// FriendsIDs returns the "friends" edge IDs in the mutation.
func (m *UserMutation) FriendsIDs() (ids []uuid.UUID) {
	for id := range m.friends {
		ids = append(ids, id)
	}
	return
}

// ResetFriends resets all changes to the "friends" edge.
func (m *UserMutation) ResetFriends() {
	m.friends = nil
	m.clearedfriends = false
	m.removedfriends = nil
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 6)
	if m.created_at != nil {
		fields = append(fields, user.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, user.FieldUpdatedAt)
	}
	if m.first_name != nil {
		fields = append(fields, user.FieldFirstName)
	}
	if m.last_name != nil {
		fields = append(fields, user.FieldLastName)
	}
	if m.username != nil {
		fields = append(fields, user.FieldUsername)
	}
	if m.optional_nullable_bool != nil {
		fields = append(fields, user.FieldOptionalNullableBool)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldCreatedAt:
		return m.CreatedAt()
	case user.FieldUpdatedAt:
		return m.UpdatedAt()
	case user.FieldFirstName:
		return m.FirstName()
	case user.FieldLastName:
		return m.LastName()
	case user.FieldUsername:
		return m.Username()
	case user.FieldOptionalNullableBool:
		return m.OptionalNullableBool()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case user.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case user.FieldFirstName:
		return m.OldFirstName(ctx)
	case user.FieldLastName:
		return m.OldLastName(ctx)
	case user.FieldUsername:
		return m.OldUsername(ctx)
	case user.FieldOptionalNullableBool:
		return m.OldOptionalNullableBool(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case user.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case user.FieldFirstName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetFirstName(v)
		return nil
	case user.FieldLastName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLastName(v)
		return nil
	case user.FieldUsername:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUsername(v)
		return nil
	case user.FieldOptionalNullableBool:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetOptionalNullableBool(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(user.FieldOptionalNullableBool) {
		fields = append(fields, user.FieldOptionalNullableBool)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	switch name {
	case user.FieldOptionalNullableBool:
		m.ClearOptionalNullableBool()
		return nil
	}
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case user.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case user.FieldFirstName:
		m.ResetFirstName()
		return nil
	case user.FieldLastName:
		m.ResetLastName()
		return nil
	case user.FieldUsername:
		m.ResetUsername()
		return nil
	case user.FieldOptionalNullableBool:
		m.ResetOptionalNullableBool()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 2)
	if m.city != nil {
		edges = append(edges, user.EdgeCity)
	}
	if m.friends != nil {
		edges = append(edges, user.EdgeFriends)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeCity:
		if id := m.city; id != nil {
			return []ent.Value{*id}
		}
	case user.EdgeFriends:
		ids := make([]ent.Value, 0, len(m.friends))
		for id := range m.friends {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 2)
	if m.removedfriends != nil {
		edges = append(edges, user.EdgeFriends)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeFriends:
		ids := make([]ent.Value, 0, len(m.removedfriends))
		for id := range m.removedfriends {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 2)
	if m.clearedcity {
		edges = append(edges, user.EdgeCity)
	}
	if m.clearedfriends {
		edges = append(edges, user.EdgeFriends)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	switch name {
	case user.EdgeCity:
		return m.clearedcity
	case user.EdgeFriends:
		return m.clearedfriends
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	switch name {
	case user.EdgeCity:
		m.ClearCity()
		return nil
	}
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	switch name {
	case user.EdgeCity:
		m.ResetCity()
		return nil
	case user.EdgeFriends:
		m.ResetFriends()
		return nil
	}
	return fmt.Errorf("unknown User edge %s", name)
}
