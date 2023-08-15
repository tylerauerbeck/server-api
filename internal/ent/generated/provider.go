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
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"go.infratographer.com/server-api/internal/ent/generated/provider"
	"go.infratographer.com/x/gidx"
)

// Representation of a server provider. Server providers are responsible for provisioning and managing servers
type Provider struct {
	config `json:"-"`
	// ID of the ent.
	// The ID of the server provider.
	ID gidx.PrefixedID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// The name of the server provider.
	Name string `json:"name,omitempty"`
	// The ID for the owner of this server.
	OwnerID gidx.PrefixedID `json:"owner_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProviderQuery when eager-loading is set.
	Edges        ProviderEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ProviderEdges holds the relations/edges for other nodes in the graph.
type ProviderEdges struct {
	// Servers holds the value of the servers edge.
	Servers []*Server `json:"servers,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int

	namedServers map[string][]*Server
}

// ServersOrErr returns the Servers value or an error if the edge
// was not loaded in eager-loading.
func (e ProviderEdges) ServersOrErr() ([]*Server, error) {
	if e.loadedTypes[0] {
		return e.Servers, nil
	}
	return nil, &NotLoadedError{edge: "servers"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Provider) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case provider.FieldID, provider.FieldOwnerID:
			values[i] = new(gidx.PrefixedID)
		case provider.FieldName:
			values[i] = new(sql.NullString)
		case provider.FieldCreatedAt, provider.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Provider fields.
func (pr *Provider) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case provider.FieldID:
			if value, ok := values[i].(*gidx.PrefixedID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				pr.ID = *value
			}
		case provider.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pr.CreatedAt = value.Time
			}
		case provider.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pr.UpdatedAt = value.Time
			}
		case provider.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pr.Name = value.String
			}
		case provider.FieldOwnerID:
			if value, ok := values[i].(*gidx.PrefixedID); !ok {
				return fmt.Errorf("unexpected type %T for field owner_id", values[i])
			} else if value != nil {
				pr.OwnerID = *value
			}
		default:
			pr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Provider.
// This includes values selected through modifiers, order, etc.
func (pr *Provider) Value(name string) (ent.Value, error) {
	return pr.selectValues.Get(name)
}

// QueryServers queries the "servers" edge of the Provider entity.
func (pr *Provider) QueryServers() *ServerQuery {
	return NewProviderClient(pr.config).QueryServers(pr)
}

// Update returns a builder for updating this Provider.
// Note that you need to call Provider.Unwrap() before calling this method if this Provider
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Provider) Update() *ProviderUpdateOne {
	return NewProviderClient(pr.config).UpdateOne(pr)
}

// Unwrap unwraps the Provider entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Provider) Unwrap() *Provider {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("generated: Provider is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Provider) String() string {
	var builder strings.Builder
	builder.WriteString("Provider(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("created_at=")
	builder.WriteString(pr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pr.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(pr.Name)
	builder.WriteString(", ")
	builder.WriteString("owner_id=")
	builder.WriteString(fmt.Sprintf("%v", pr.OwnerID))
	builder.WriteByte(')')
	return builder.String()
}

// IsEntity implement fedruntime.Entity
func (pr Provider) IsEntity() {}

// NamedServers returns the Servers named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pr *Provider) NamedServers(name string) ([]*Server, error) {
	if pr.Edges.namedServers == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pr.Edges.namedServers[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pr *Provider) appendNamedServers(name string, edges ...*Server) {
	if pr.Edges.namedServers == nil {
		pr.Edges.namedServers = make(map[string][]*Server)
	}
	if len(edges) == 0 {
		pr.Edges.namedServers[name] = []*Server{}
	} else {
		pr.Edges.namedServers[name] = append(pr.Edges.namedServers[name], edges...)
	}
}

// Providers is a parsable slice of Provider.
type Providers []*Provider
