// Copyright 2023 The Infratographer Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"go.infratographer.com/server-api/internal/ent/generated/predicate"
	"go.infratographer.com/server-api/internal/ent/generated/serverchassistype"
)

// ServerChassisTypeUpdate is the builder for updating ServerChassisType entities.
type ServerChassisTypeUpdate struct {
	config
	hooks    []Hook
	mutation *ServerChassisTypeMutation
}

// Where appends a list predicates to the ServerChassisTypeUpdate builder.
func (sctu *ServerChassisTypeUpdate) Where(ps ...predicate.ServerChassisType) *ServerChassisTypeUpdate {
	sctu.mutation.Where(ps...)
	return sctu
}

// SetVendor sets the "vendor" field.
func (sctu *ServerChassisTypeUpdate) SetVendor(s string) *ServerChassisTypeUpdate {
	sctu.mutation.SetVendor(s)
	return sctu
}

// SetModel sets the "model" field.
func (sctu *ServerChassisTypeUpdate) SetModel(s string) *ServerChassisTypeUpdate {
	sctu.mutation.SetModel(s)
	return sctu
}

// SetHeight sets the "height" field.
func (sctu *ServerChassisTypeUpdate) SetHeight(s string) *ServerChassisTypeUpdate {
	sctu.mutation.SetHeight(s)
	return sctu
}

// SetIsFullDepth sets the "is_full_depth" field.
func (sctu *ServerChassisTypeUpdate) SetIsFullDepth(b bool) *ServerChassisTypeUpdate {
	sctu.mutation.SetIsFullDepth(b)
	return sctu
}

// Mutation returns the ServerChassisTypeMutation object of the builder.
func (sctu *ServerChassisTypeUpdate) Mutation() *ServerChassisTypeMutation {
	return sctu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (sctu *ServerChassisTypeUpdate) Save(ctx context.Context) (int, error) {
	sctu.defaults()
	return withHooks(ctx, sctu.sqlSave, sctu.mutation, sctu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (sctu *ServerChassisTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := sctu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (sctu *ServerChassisTypeUpdate) Exec(ctx context.Context) error {
	_, err := sctu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sctu *ServerChassisTypeUpdate) ExecX(ctx context.Context) {
	if err := sctu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sctu *ServerChassisTypeUpdate) defaults() {
	if _, ok := sctu.mutation.UpdatedAt(); !ok {
		v := serverchassistype.UpdateDefaultUpdatedAt()
		sctu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sctu *ServerChassisTypeUpdate) check() error {
	if v, ok := sctu.mutation.Vendor(); ok {
		if err := serverchassistype.VendorValidator(v); err != nil {
			return &ValidationError{Name: "vendor", err: fmt.Errorf(`generated: validator failed for field "ServerChassisType.vendor": %w`, err)}
		}
	}
	if v, ok := sctu.mutation.Model(); ok {
		if err := serverchassistype.ModelValidator(v); err != nil {
			return &ValidationError{Name: "model", err: fmt.Errorf(`generated: validator failed for field "ServerChassisType.model": %w`, err)}
		}
	}
	if v, ok := sctu.mutation.Height(); ok {
		if err := serverchassistype.HeightValidator(v); err != nil {
			return &ValidationError{Name: "height", err: fmt.Errorf(`generated: validator failed for field "ServerChassisType.height": %w`, err)}
		}
	}
	return nil
}

func (sctu *ServerChassisTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := sctu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(serverchassistype.Table, serverchassistype.Columns, sqlgraph.NewFieldSpec(serverchassistype.FieldID, field.TypeString))
	if ps := sctu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sctu.mutation.UpdatedAt(); ok {
		_spec.SetField(serverchassistype.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := sctu.mutation.Vendor(); ok {
		_spec.SetField(serverchassistype.FieldVendor, field.TypeString, value)
	}
	if value, ok := sctu.mutation.Model(); ok {
		_spec.SetField(serverchassistype.FieldModel, field.TypeString, value)
	}
	if value, ok := sctu.mutation.Height(); ok {
		_spec.SetField(serverchassistype.FieldHeight, field.TypeString, value)
	}
	if value, ok := sctu.mutation.IsFullDepth(); ok {
		_spec.SetField(serverchassistype.FieldIsFullDepth, field.TypeBool, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, sctu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{serverchassistype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	sctu.mutation.done = true
	return n, nil
}

// ServerChassisTypeUpdateOne is the builder for updating a single ServerChassisType entity.
type ServerChassisTypeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ServerChassisTypeMutation
}

// SetVendor sets the "vendor" field.
func (sctuo *ServerChassisTypeUpdateOne) SetVendor(s string) *ServerChassisTypeUpdateOne {
	sctuo.mutation.SetVendor(s)
	return sctuo
}

// SetModel sets the "model" field.
func (sctuo *ServerChassisTypeUpdateOne) SetModel(s string) *ServerChassisTypeUpdateOne {
	sctuo.mutation.SetModel(s)
	return sctuo
}

// SetHeight sets the "height" field.
func (sctuo *ServerChassisTypeUpdateOne) SetHeight(s string) *ServerChassisTypeUpdateOne {
	sctuo.mutation.SetHeight(s)
	return sctuo
}

// SetIsFullDepth sets the "is_full_depth" field.
func (sctuo *ServerChassisTypeUpdateOne) SetIsFullDepth(b bool) *ServerChassisTypeUpdateOne {
	sctuo.mutation.SetIsFullDepth(b)
	return sctuo
}

// Mutation returns the ServerChassisTypeMutation object of the builder.
func (sctuo *ServerChassisTypeUpdateOne) Mutation() *ServerChassisTypeMutation {
	return sctuo.mutation
}

// Where appends a list predicates to the ServerChassisTypeUpdate builder.
func (sctuo *ServerChassisTypeUpdateOne) Where(ps ...predicate.ServerChassisType) *ServerChassisTypeUpdateOne {
	sctuo.mutation.Where(ps...)
	return sctuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sctuo *ServerChassisTypeUpdateOne) Select(field string, fields ...string) *ServerChassisTypeUpdateOne {
	sctuo.fields = append([]string{field}, fields...)
	return sctuo
}

// Save executes the query and returns the updated ServerChassisType entity.
func (sctuo *ServerChassisTypeUpdateOne) Save(ctx context.Context) (*ServerChassisType, error) {
	sctuo.defaults()
	return withHooks(ctx, sctuo.sqlSave, sctuo.mutation, sctuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (sctuo *ServerChassisTypeUpdateOne) SaveX(ctx context.Context) *ServerChassisType {
	node, err := sctuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sctuo *ServerChassisTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := sctuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sctuo *ServerChassisTypeUpdateOne) ExecX(ctx context.Context) {
	if err := sctuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sctuo *ServerChassisTypeUpdateOne) defaults() {
	if _, ok := sctuo.mutation.UpdatedAt(); !ok {
		v := serverchassistype.UpdateDefaultUpdatedAt()
		sctuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sctuo *ServerChassisTypeUpdateOne) check() error {
	if v, ok := sctuo.mutation.Vendor(); ok {
		if err := serverchassistype.VendorValidator(v); err != nil {
			return &ValidationError{Name: "vendor", err: fmt.Errorf(`generated: validator failed for field "ServerChassisType.vendor": %w`, err)}
		}
	}
	if v, ok := sctuo.mutation.Model(); ok {
		if err := serverchassistype.ModelValidator(v); err != nil {
			return &ValidationError{Name: "model", err: fmt.Errorf(`generated: validator failed for field "ServerChassisType.model": %w`, err)}
		}
	}
	if v, ok := sctuo.mutation.Height(); ok {
		if err := serverchassistype.HeightValidator(v); err != nil {
			return &ValidationError{Name: "height", err: fmt.Errorf(`generated: validator failed for field "ServerChassisType.height": %w`, err)}
		}
	}
	return nil
}

func (sctuo *ServerChassisTypeUpdateOne) sqlSave(ctx context.Context) (_node *ServerChassisType, err error) {
	if err := sctuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(serverchassistype.Table, serverchassistype.Columns, sqlgraph.NewFieldSpec(serverchassistype.FieldID, field.TypeString))
	id, ok := sctuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "ServerChassisType.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := sctuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, serverchassistype.FieldID)
		for _, f := range fields {
			if !serverchassistype.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != serverchassistype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := sctuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sctuo.mutation.UpdatedAt(); ok {
		_spec.SetField(serverchassistype.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := sctuo.mutation.Vendor(); ok {
		_spec.SetField(serverchassistype.FieldVendor, field.TypeString, value)
	}
	if value, ok := sctuo.mutation.Model(); ok {
		_spec.SetField(serverchassistype.FieldModel, field.TypeString, value)
	}
	if value, ok := sctuo.mutation.Height(); ok {
		_spec.SetField(serverchassistype.FieldHeight, field.TypeString, value)
	}
	if value, ok := sctuo.mutation.IsFullDepth(); ok {
		_spec.SetField(serverchassistype.FieldIsFullDepth, field.TypeBool, value)
	}
	_node = &ServerChassisType{config: sctuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sctuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{serverchassistype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	sctuo.mutation.done = true
	return _node, nil
}
