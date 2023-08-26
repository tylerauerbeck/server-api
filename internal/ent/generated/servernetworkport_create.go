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
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"go.infratographer.com/server-api/internal/ent/generated/servernetworkcard"
	"go.infratographer.com/server-api/internal/ent/generated/servernetworkport"
	"go.infratographer.com/x/gidx"
)

// ServerNetworkPortCreate is the builder for creating a ServerNetworkPort entity.
type ServerNetworkPortCreate struct {
	config
	mutation *ServerNetworkPortMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (snpc *ServerNetworkPortCreate) SetCreatedAt(t time.Time) *ServerNetworkPortCreate {
	snpc.mutation.SetCreatedAt(t)
	return snpc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (snpc *ServerNetworkPortCreate) SetNillableCreatedAt(t *time.Time) *ServerNetworkPortCreate {
	if t != nil {
		snpc.SetCreatedAt(*t)
	}
	return snpc
}

// SetUpdatedAt sets the "updated_at" field.
func (snpc *ServerNetworkPortCreate) SetUpdatedAt(t time.Time) *ServerNetworkPortCreate {
	snpc.mutation.SetUpdatedAt(t)
	return snpc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (snpc *ServerNetworkPortCreate) SetNillableUpdatedAt(t *time.Time) *ServerNetworkPortCreate {
	if t != nil {
		snpc.SetUpdatedAt(*t)
	}
	return snpc
}

// SetMACAddress sets the "mac_address" field.
func (snpc *ServerNetworkPortCreate) SetMACAddress(s string) *ServerNetworkPortCreate {
	snpc.mutation.SetMACAddress(s)
	return snpc
}

// SetNetworkCardID sets the "network_card_id" field.
func (snpc *ServerNetworkPortCreate) SetNetworkCardID(gi gidx.PrefixedID) *ServerNetworkPortCreate {
	snpc.mutation.SetNetworkCardID(gi)
	return snpc
}

// SetID sets the "id" field.
func (snpc *ServerNetworkPortCreate) SetID(gi gidx.PrefixedID) *ServerNetworkPortCreate {
	snpc.mutation.SetID(gi)
	return snpc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (snpc *ServerNetworkPortCreate) SetNillableID(gi *gidx.PrefixedID) *ServerNetworkPortCreate {
	if gi != nil {
		snpc.SetID(*gi)
	}
	return snpc
}

// SetNetworkCard sets the "network_card" edge to the ServerNetworkCard entity.
func (snpc *ServerNetworkPortCreate) SetNetworkCard(s *ServerNetworkCard) *ServerNetworkPortCreate {
	return snpc.SetNetworkCardID(s.ID)
}

// Mutation returns the ServerNetworkPortMutation object of the builder.
func (snpc *ServerNetworkPortCreate) Mutation() *ServerNetworkPortMutation {
	return snpc.mutation
}

// Save creates the ServerNetworkPort in the database.
func (snpc *ServerNetworkPortCreate) Save(ctx context.Context) (*ServerNetworkPort, error) {
	snpc.defaults()
	return withHooks(ctx, snpc.sqlSave, snpc.mutation, snpc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (snpc *ServerNetworkPortCreate) SaveX(ctx context.Context) *ServerNetworkPort {
	v, err := snpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (snpc *ServerNetworkPortCreate) Exec(ctx context.Context) error {
	_, err := snpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (snpc *ServerNetworkPortCreate) ExecX(ctx context.Context) {
	if err := snpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (snpc *ServerNetworkPortCreate) defaults() {
	if _, ok := snpc.mutation.CreatedAt(); !ok {
		v := servernetworkport.DefaultCreatedAt()
		snpc.mutation.SetCreatedAt(v)
	}
	if _, ok := snpc.mutation.UpdatedAt(); !ok {
		v := servernetworkport.DefaultUpdatedAt()
		snpc.mutation.SetUpdatedAt(v)
	}
	if _, ok := snpc.mutation.ID(); !ok {
		v := servernetworkport.DefaultID()
		snpc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (snpc *ServerNetworkPortCreate) check() error {
	if _, ok := snpc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`generated: missing required field "ServerNetworkPort.created_at"`)}
	}
	if _, ok := snpc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`generated: missing required field "ServerNetworkPort.updated_at"`)}
	}
	if _, ok := snpc.mutation.MACAddress(); !ok {
		return &ValidationError{Name: "mac_address", err: errors.New(`generated: missing required field "ServerNetworkPort.mac_address"`)}
	}
	if v, ok := snpc.mutation.MACAddress(); ok {
		if err := servernetworkport.MACAddressValidator(v); err != nil {
			return &ValidationError{Name: "mac_address", err: fmt.Errorf(`generated: validator failed for field "ServerNetworkPort.mac_address": %w`, err)}
		}
	}
	if _, ok := snpc.mutation.NetworkCardID(); !ok {
		return &ValidationError{Name: "network_card_id", err: errors.New(`generated: missing required field "ServerNetworkPort.network_card_id"`)}
	}
	if _, ok := snpc.mutation.NetworkCardID(); !ok {
		return &ValidationError{Name: "network_card", err: errors.New(`generated: missing required edge "ServerNetworkPort.network_card"`)}
	}
	return nil
}

func (snpc *ServerNetworkPortCreate) sqlSave(ctx context.Context) (*ServerNetworkPort, error) {
	if err := snpc.check(); err != nil {
		return nil, err
	}
	_node, _spec := snpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, snpc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*gidx.PrefixedID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	snpc.mutation.id = &_node.ID
	snpc.mutation.done = true
	return _node, nil
}

func (snpc *ServerNetworkPortCreate) createSpec() (*ServerNetworkPort, *sqlgraph.CreateSpec) {
	var (
		_node = &ServerNetworkPort{config: snpc.config}
		_spec = sqlgraph.NewCreateSpec(servernetworkport.Table, sqlgraph.NewFieldSpec(servernetworkport.FieldID, field.TypeString))
	)
	if id, ok := snpc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := snpc.mutation.CreatedAt(); ok {
		_spec.SetField(servernetworkport.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := snpc.mutation.UpdatedAt(); ok {
		_spec.SetField(servernetworkport.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := snpc.mutation.MACAddress(); ok {
		_spec.SetField(servernetworkport.FieldMACAddress, field.TypeString, value)
		_node.MACAddress = value
	}
	if nodes := snpc.mutation.NetworkCardIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   servernetworkport.NetworkCardTable,
			Columns: []string{servernetworkport.NetworkCardColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(servernetworkcard.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.NetworkCardID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ServerNetworkPortCreateBulk is the builder for creating many ServerNetworkPort entities in bulk.
type ServerNetworkPortCreateBulk struct {
	config
	builders []*ServerNetworkPortCreate
}

// Save creates the ServerNetworkPort entities in the database.
func (snpcb *ServerNetworkPortCreateBulk) Save(ctx context.Context) ([]*ServerNetworkPort, error) {
	specs := make([]*sqlgraph.CreateSpec, len(snpcb.builders))
	nodes := make([]*ServerNetworkPort, len(snpcb.builders))
	mutators := make([]Mutator, len(snpcb.builders))
	for i := range snpcb.builders {
		func(i int, root context.Context) {
			builder := snpcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ServerNetworkPortMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, snpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, snpcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, snpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (snpcb *ServerNetworkPortCreateBulk) SaveX(ctx context.Context) []*ServerNetworkPort {
	v, err := snpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (snpcb *ServerNetworkPortCreateBulk) Exec(ctx context.Context) error {
	_, err := snpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (snpcb *ServerNetworkPortCreateBulk) ExecX(ctx context.Context) {
	if err := snpcb.Exec(ctx); err != nil {
		panic(err)
	}
}
