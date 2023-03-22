// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ogen-go/ent2ogen/internal/test/ent/schemaa"
	"github.com/ogen-go/ent2ogen/internal/test/ent/schemab"
)

// SchemaACreate is the builder for creating a SchemaA entity.
type SchemaACreate struct {
	config
	mutation *SchemaAMutation
	hooks    []Hook
}

// SetInt64 sets the "int64" field.
func (sa *SchemaACreate) SetInt64(i int64) *SchemaACreate {
	sa.mutation.SetInt64(i)
	return sa
}

// SetStringBindtoFoobar sets the "string_bindto_foobar" field.
func (sa *SchemaACreate) SetStringBindtoFoobar(s string) *SchemaACreate {
	sa.mutation.SetStringBindtoFoobar(s)
	return sa
}

// SetStringOptionalNullable sets the "string_optional_nullable" field.
func (sa *SchemaACreate) SetStringOptionalNullable(s string) *SchemaACreate {
	sa.mutation.SetStringOptionalNullable(s)
	return sa
}

// SetNillableStringOptionalNullable sets the "string_optional_nullable" field if the given value is not nil.
func (sa *SchemaACreate) SetNillableStringOptionalNullable(s *string) *SchemaACreate {
	if s != nil {
		sa.SetStringOptionalNullable(*s)
	}
	return sa
}

// SetOptionalNullableBool sets the "optional_nullable_bool" field.
func (sa *SchemaACreate) SetOptionalNullableBool(b bool) *SchemaACreate {
	sa.mutation.SetOptionalNullableBool(b)
	return sa
}

// SetNillableOptionalNullableBool sets the "optional_nullable_bool" field if the given value is not nil.
func (sa *SchemaACreate) SetNillableOptionalNullableBool(b *bool) *SchemaACreate {
	if b != nil {
		sa.SetOptionalNullableBool(*b)
	}
	return sa
}

// SetJsontypeStrings sets the "jsontype_strings" field.
func (sa *SchemaACreate) SetJsontypeStrings(s []string) *SchemaACreate {
	sa.mutation.SetJsontypeStrings(s)
	return sa
}

// SetJsontypeStringsOptional sets the "jsontype_strings_optional" field.
func (sa *SchemaACreate) SetJsontypeStringsOptional(s []string) *SchemaACreate {
	sa.mutation.SetJsontypeStringsOptional(s)
	return sa
}

// SetJsontypeInts sets the "jsontype_ints" field.
func (sa *SchemaACreate) SetJsontypeInts(i []int) *SchemaACreate {
	sa.mutation.SetJsontypeInts(i)
	return sa
}

// SetJsontypeIntsOptional sets the "jsontype_ints_optional" field.
func (sa *SchemaACreate) SetJsontypeIntsOptional(i []int) *SchemaACreate {
	sa.mutation.SetJsontypeIntsOptional(i)
	return sa
}

// SetRequiredEnum sets the "required_enum" field.
func (sa *SchemaACreate) SetRequiredEnum(se schemaa.RequiredEnum) *SchemaACreate {
	sa.mutation.SetRequiredEnum(se)
	return sa
}

// SetOptionalNullableEnum sets the "optional_nullable_enum" field.
func (sa *SchemaACreate) SetOptionalNullableEnum(sne schemaa.OptionalNullableEnum) *SchemaACreate {
	sa.mutation.SetOptionalNullableEnum(sne)
	return sa
}

// SetNillableOptionalNullableEnum sets the "optional_nullable_enum" field if the given value is not nil.
func (sa *SchemaACreate) SetNillableOptionalNullableEnum(sne *schemaa.OptionalNullableEnum) *SchemaACreate {
	if sne != nil {
		sa.SetOptionalNullableEnum(*sne)
	}
	return sa
}

// SetBytes sets the "bytes" field.
func (sa *SchemaACreate) SetBytes(b []byte) *SchemaACreate {
	sa.mutation.SetBytes(b)
	return sa
}

// SetEdgeSchemabUniqueRequiredID sets the "edge_schemab_unique_required" edge to the SchemaB entity by ID.
func (sa *SchemaACreate) SetEdgeSchemabUniqueRequiredID(id int64) *SchemaACreate {
	sa.mutation.SetEdgeSchemabUniqueRequiredID(id)
	return sa
}

// SetEdgeSchemabUniqueRequired sets the "edge_schemab_unique_required" edge to the SchemaB entity.
func (sa *SchemaACreate) SetEdgeSchemabUniqueRequired(s *SchemaB) *SchemaACreate {
	return sa.SetEdgeSchemabUniqueRequiredID(s.ID)
}

// SetEdgeSchemabUniqueRequiredBindtoBsID sets the "edge_schemab_unique_required_bindto_bs" edge to the SchemaB entity by ID.
func (sa *SchemaACreate) SetEdgeSchemabUniqueRequiredBindtoBsID(id int64) *SchemaACreate {
	sa.mutation.SetEdgeSchemabUniqueRequiredBindtoBsID(id)
	return sa
}

// SetEdgeSchemabUniqueRequiredBindtoBs sets the "edge_schemab_unique_required_bindto_bs" edge to the SchemaB entity.
func (sa *SchemaACreate) SetEdgeSchemabUniqueRequiredBindtoBs(s *SchemaB) *SchemaACreate {
	return sa.SetEdgeSchemabUniqueRequiredBindtoBsID(s.ID)
}

// SetEdgeSchemabUniqueOptionalID sets the "edge_schemab_unique_optional" edge to the SchemaB entity by ID.
func (sa *SchemaACreate) SetEdgeSchemabUniqueOptionalID(id int64) *SchemaACreate {
	sa.mutation.SetEdgeSchemabUniqueOptionalID(id)
	return sa
}

// SetNillableEdgeSchemabUniqueOptionalID sets the "edge_schemab_unique_optional" edge to the SchemaB entity by ID if the given value is not nil.
func (sa *SchemaACreate) SetNillableEdgeSchemabUniqueOptionalID(id *int64) *SchemaACreate {
	if id != nil {
		sa = sa.SetEdgeSchemabUniqueOptionalID(*id)
	}
	return sa
}

// SetEdgeSchemabUniqueOptional sets the "edge_schemab_unique_optional" edge to the SchemaB entity.
func (sa *SchemaACreate) SetEdgeSchemabUniqueOptional(s *SchemaB) *SchemaACreate {
	return sa.SetEdgeSchemabUniqueOptionalID(s.ID)
}

// AddEdgeSchemabIDs adds the "edge_schemab" edge to the SchemaB entity by IDs.
func (sa *SchemaACreate) AddEdgeSchemabIDs(ids ...int64) *SchemaACreate {
	sa.mutation.AddEdgeSchemabIDs(ids...)
	return sa
}

// AddEdgeSchemab adds the "edge_schemab" edges to the SchemaB entity.
func (sa *SchemaACreate) AddEdgeSchemab(s ...*SchemaB) *SchemaACreate {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sa.AddEdgeSchemabIDs(ids...)
}

// AddEdgeSchemaaRecursiveIDs adds the "edge_schemaa_recursive" edge to the SchemaA entity by IDs.
func (sa *SchemaACreate) AddEdgeSchemaaRecursiveIDs(ids ...int) *SchemaACreate {
	sa.mutation.AddEdgeSchemaaRecursiveIDs(ids...)
	return sa
}

// AddEdgeSchemaaRecursive adds the "edge_schemaa_recursive" edges to the SchemaA entity.
func (sa *SchemaACreate) AddEdgeSchemaaRecursive(s ...*SchemaA) *SchemaACreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sa.AddEdgeSchemaaRecursiveIDs(ids...)
}

// Mutation returns the SchemaAMutation object of the builder.
func (sa *SchemaACreate) Mutation() *SchemaAMutation {
	return sa.mutation
}

// Save creates the SchemaA in the database.
func (sa *SchemaACreate) Save(ctx context.Context) (*SchemaA, error) {
	return withHooks[*SchemaA, SchemaAMutation](ctx, sa.sqlSave, sa.mutation, sa.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sa *SchemaACreate) SaveX(ctx context.Context) *SchemaA {
	v, err := sa.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sa *SchemaACreate) Exec(ctx context.Context) error {
	_, err := sa.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sa *SchemaACreate) ExecX(ctx context.Context) {
	if err := sa.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sa *SchemaACreate) check() error {
	if _, ok := sa.mutation.Int64(); !ok {
		return &ValidationError{Name: "int64", err: errors.New(`ent: missing required field "SchemaA.int64"`)}
	}
	if _, ok := sa.mutation.StringBindtoFoobar(); !ok {
		return &ValidationError{Name: "string_bindto_foobar", err: errors.New(`ent: missing required field "SchemaA.string_bindto_foobar"`)}
	}
	if _, ok := sa.mutation.JsontypeStrings(); !ok {
		return &ValidationError{Name: "jsontype_strings", err: errors.New(`ent: missing required field "SchemaA.jsontype_strings"`)}
	}
	if _, ok := sa.mutation.JsontypeInts(); !ok {
		return &ValidationError{Name: "jsontype_ints", err: errors.New(`ent: missing required field "SchemaA.jsontype_ints"`)}
	}
	if _, ok := sa.mutation.RequiredEnum(); !ok {
		return &ValidationError{Name: "required_enum", err: errors.New(`ent: missing required field "SchemaA.required_enum"`)}
	}
	if v, ok := sa.mutation.RequiredEnum(); ok {
		if err := schemaa.RequiredEnumValidator(v); err != nil {
			return &ValidationError{Name: "required_enum", err: fmt.Errorf(`ent: validator failed for field "SchemaA.required_enum": %w`, err)}
		}
	}
	if v, ok := sa.mutation.OptionalNullableEnum(); ok {
		if err := schemaa.OptionalNullableEnumValidator(v); err != nil {
			return &ValidationError{Name: "optional_nullable_enum", err: fmt.Errorf(`ent: validator failed for field "SchemaA.optional_nullable_enum": %w`, err)}
		}
	}
	if _, ok := sa.mutation.Bytes(); !ok {
		return &ValidationError{Name: "bytes", err: errors.New(`ent: missing required field "SchemaA.bytes"`)}
	}
	if _, ok := sa.mutation.EdgeSchemabUniqueRequiredID(); !ok {
		return &ValidationError{Name: "edge_schemab_unique_required", err: errors.New(`ent: missing required edge "SchemaA.edge_schemab_unique_required"`)}
	}
	if _, ok := sa.mutation.EdgeSchemabUniqueRequiredBindtoBsID(); !ok {
		return &ValidationError{Name: "edge_schemab_unique_required_bindto_bs", err: errors.New(`ent: missing required edge "SchemaA.edge_schemab_unique_required_bindto_bs"`)}
	}
	return nil
}

func (sa *SchemaACreate) sqlSave(ctx context.Context) (*SchemaA, error) {
	if err := sa.check(); err != nil {
		return nil, err
	}
	_node, _spec := sa.createSpec()
	if err := sqlgraph.CreateNode(ctx, sa.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	sa.mutation.id = &_node.ID
	sa.mutation.done = true
	return _node, nil
}

func (sa *SchemaACreate) createSpec() (*SchemaA, *sqlgraph.CreateSpec) {
	var (
		_node = &SchemaA{config: sa.config}
		_spec = sqlgraph.NewCreateSpec(schemaa.Table, sqlgraph.NewFieldSpec(schemaa.FieldID, field.TypeInt))
	)
	if value, ok := sa.mutation.Int64(); ok {
		_spec.SetField(schemaa.FieldInt64, field.TypeInt64, value)
		_node.Int64 = value
	}
	if value, ok := sa.mutation.StringBindtoFoobar(); ok {
		_spec.SetField(schemaa.FieldStringBindtoFoobar, field.TypeString, value)
		_node.StringBindtoFoobar = value
	}
	if value, ok := sa.mutation.StringOptionalNullable(); ok {
		_spec.SetField(schemaa.FieldStringOptionalNullable, field.TypeString, value)
		_node.StringOptionalNullable = &value
	}
	if value, ok := sa.mutation.OptionalNullableBool(); ok {
		_spec.SetField(schemaa.FieldOptionalNullableBool, field.TypeBool, value)
		_node.OptionalNullableBool = &value
	}
	if value, ok := sa.mutation.JsontypeStrings(); ok {
		_spec.SetField(schemaa.FieldJsontypeStrings, field.TypeJSON, value)
		_node.JsontypeStrings = value
	}
	if value, ok := sa.mutation.JsontypeStringsOptional(); ok {
		_spec.SetField(schemaa.FieldJsontypeStringsOptional, field.TypeJSON, value)
		_node.JsontypeStringsOptional = value
	}
	if value, ok := sa.mutation.JsontypeInts(); ok {
		_spec.SetField(schemaa.FieldJsontypeInts, field.TypeJSON, value)
		_node.JsontypeInts = value
	}
	if value, ok := sa.mutation.JsontypeIntsOptional(); ok {
		_spec.SetField(schemaa.FieldJsontypeIntsOptional, field.TypeJSON, value)
		_node.JsontypeIntsOptional = value
	}
	if value, ok := sa.mutation.RequiredEnum(); ok {
		_spec.SetField(schemaa.FieldRequiredEnum, field.TypeEnum, value)
		_node.RequiredEnum = value
	}
	if value, ok := sa.mutation.OptionalNullableEnum(); ok {
		_spec.SetField(schemaa.FieldOptionalNullableEnum, field.TypeEnum, value)
		_node.OptionalNullableEnum = &value
	}
	if value, ok := sa.mutation.Bytes(); ok {
		_spec.SetField(schemaa.FieldBytes, field.TypeBytes, value)
		_node.Bytes = value
	}
	if nodes := sa.mutation.EdgeSchemabUniqueRequiredIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   schemaa.EdgeSchemabUniqueRequiredTable,
			Columns: []string{schemaa.EdgeSchemabUniqueRequiredColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(schemab.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.schemaa_edge_schemab_unique_required = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sa.mutation.EdgeSchemabUniqueRequiredBindtoBsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   schemaa.EdgeSchemabUniqueRequiredBindtoBsTable,
			Columns: []string{schemaa.EdgeSchemabUniqueRequiredBindtoBsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(schemab.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.schemaa_edge_schemab_unique_required_bindto_bs = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sa.mutation.EdgeSchemabUniqueOptionalIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   schemaa.EdgeSchemabUniqueOptionalTable,
			Columns: []string{schemaa.EdgeSchemabUniqueOptionalColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(schemab.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.schemaa_edge_schemab_unique_optional = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sa.mutation.EdgeSchemabIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   schemaa.EdgeSchemabTable,
			Columns: []string{schemaa.EdgeSchemabColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(schemab.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sa.mutation.EdgeSchemaaRecursiveIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   schemaa.EdgeSchemaaRecursiveTable,
			Columns: schemaa.EdgeSchemaaRecursivePrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(schemaa.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SchemaACreateBulk is the builder for creating many SchemaA entities in bulk.
type SchemaACreateBulk struct {
	config
	builders []*SchemaACreate
}

// Save creates the SchemaA entities in the database.
func (sab *SchemaACreateBulk) Save(ctx context.Context) ([]*SchemaA, error) {
	specs := make([]*sqlgraph.CreateSpec, len(sab.builders))
	nodes := make([]*SchemaA, len(sab.builders))
	mutators := make([]Mutator, len(sab.builders))
	for i := range sab.builders {
		func(i int, root context.Context) {
			builder := sab.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SchemaAMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, sab.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sab.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, sab.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sab *SchemaACreateBulk) SaveX(ctx context.Context) []*SchemaA {
	v, err := sab.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sab *SchemaACreateBulk) Exec(ctx context.Context) error {
	_, err := sab.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sab *SchemaACreateBulk) ExecX(ctx context.Context) {
	if err := sab.Exec(ctx); err != nil {
		panic(err)
	}
}
