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
	"log"

	"go.infratographer.com/server-api/internal/ent/generated/migrate"
	"go.infratographer.com/x/gidx"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"go.infratographer.com/server-api/internal/ent/generated/provider"
	"go.infratographer.com/server-api/internal/ent/generated/server"
	"go.infratographer.com/server-api/internal/ent/generated/servercomponent"
	"go.infratographer.com/server-api/internal/ent/generated/servercomponenttype"
	"go.infratographer.com/server-api/internal/ent/generated/servertype"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Provider is the client for interacting with the Provider builders.
	Provider *ProviderClient
	// Server is the client for interacting with the Server builders.
	Server *ServerClient
	// ServerComponent is the client for interacting with the ServerComponent builders.
	ServerComponent *ServerComponentClient
	// ServerComponentType is the client for interacting with the ServerComponentType builders.
	ServerComponentType *ServerComponentTypeClient
	// ServerType is the client for interacting with the ServerType builders.
	ServerType *ServerTypeClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Provider = NewProviderClient(c.config)
	c.Server = NewServerClient(c.config)
	c.ServerComponent = NewServerComponentClient(c.config)
	c.ServerComponentType = NewServerComponentTypeClient(c.config)
	c.ServerType = NewServerTypeClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("generated: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("generated: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:                 ctx,
		config:              cfg,
		Provider:            NewProviderClient(cfg),
		Server:              NewServerClient(cfg),
		ServerComponent:     NewServerComponentClient(cfg),
		ServerComponentType: NewServerComponentTypeClient(cfg),
		ServerType:          NewServerTypeClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:                 ctx,
		config:              cfg,
		Provider:            NewProviderClient(cfg),
		Server:              NewServerClient(cfg),
		ServerComponent:     NewServerComponentClient(cfg),
		ServerComponentType: NewServerComponentTypeClient(cfg),
		ServerType:          NewServerTypeClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Provider.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Provider.Use(hooks...)
	c.Server.Use(hooks...)
	c.ServerComponent.Use(hooks...)
	c.ServerComponentType.Use(hooks...)
	c.ServerType.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Provider.Intercept(interceptors...)
	c.Server.Intercept(interceptors...)
	c.ServerComponent.Intercept(interceptors...)
	c.ServerComponentType.Intercept(interceptors...)
	c.ServerType.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *ProviderMutation:
		return c.Provider.mutate(ctx, m)
	case *ServerMutation:
		return c.Server.mutate(ctx, m)
	case *ServerComponentMutation:
		return c.ServerComponent.mutate(ctx, m)
	case *ServerComponentTypeMutation:
		return c.ServerComponentType.mutate(ctx, m)
	case *ServerTypeMutation:
		return c.ServerType.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("generated: unknown mutation type %T", m)
	}
}

// ProviderClient is a client for the Provider schema.
type ProviderClient struct {
	config
}

// NewProviderClient returns a client for the Provider from the given config.
func NewProviderClient(c config) *ProviderClient {
	return &ProviderClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `provider.Hooks(f(g(h())))`.
func (c *ProviderClient) Use(hooks ...Hook) {
	c.hooks.Provider = append(c.hooks.Provider, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `provider.Intercept(f(g(h())))`.
func (c *ProviderClient) Intercept(interceptors ...Interceptor) {
	c.inters.Provider = append(c.inters.Provider, interceptors...)
}

// Create returns a builder for creating a Provider entity.
func (c *ProviderClient) Create() *ProviderCreate {
	mutation := newProviderMutation(c.config, OpCreate)
	return &ProviderCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Provider entities.
func (c *ProviderClient) CreateBulk(builders ...*ProviderCreate) *ProviderCreateBulk {
	return &ProviderCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Provider.
func (c *ProviderClient) Update() *ProviderUpdate {
	mutation := newProviderMutation(c.config, OpUpdate)
	return &ProviderUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ProviderClient) UpdateOne(pr *Provider) *ProviderUpdateOne {
	mutation := newProviderMutation(c.config, OpUpdateOne, withProvider(pr))
	return &ProviderUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ProviderClient) UpdateOneID(id gidx.PrefixedID) *ProviderUpdateOne {
	mutation := newProviderMutation(c.config, OpUpdateOne, withProviderID(id))
	return &ProviderUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Provider.
func (c *ProviderClient) Delete() *ProviderDelete {
	mutation := newProviderMutation(c.config, OpDelete)
	return &ProviderDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ProviderClient) DeleteOne(pr *Provider) *ProviderDeleteOne {
	return c.DeleteOneID(pr.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ProviderClient) DeleteOneID(id gidx.PrefixedID) *ProviderDeleteOne {
	builder := c.Delete().Where(provider.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ProviderDeleteOne{builder}
}

// Query returns a query builder for Provider.
func (c *ProviderClient) Query() *ProviderQuery {
	return &ProviderQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeProvider},
		inters: c.Interceptors(),
	}
}

// Get returns a Provider entity by its id.
func (c *ProviderClient) Get(ctx context.Context, id gidx.PrefixedID) (*Provider, error) {
	return c.Query().Where(provider.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ProviderClient) GetX(ctx context.Context, id gidx.PrefixedID) *Provider {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryServers queries the servers edge of a Provider.
func (c *ProviderClient) QueryServers(pr *Provider) *ServerQuery {
	query := (&ServerClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(provider.Table, provider.FieldID, id),
			sqlgraph.To(server.Table, server.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, provider.ServersTable, provider.ServersColumn),
		)
		fromV = sqlgraph.Neighbors(pr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ProviderClient) Hooks() []Hook {
	return c.hooks.Provider
}

// Interceptors returns the client interceptors.
func (c *ProviderClient) Interceptors() []Interceptor {
	return c.inters.Provider
}

func (c *ProviderClient) mutate(ctx context.Context, m *ProviderMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ProviderCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ProviderUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ProviderUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ProviderDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("generated: unknown Provider mutation op: %q", m.Op())
	}
}

// ServerClient is a client for the Server schema.
type ServerClient struct {
	config
}

// NewServerClient returns a client for the Server from the given config.
func NewServerClient(c config) *ServerClient {
	return &ServerClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `server.Hooks(f(g(h())))`.
func (c *ServerClient) Use(hooks ...Hook) {
	c.hooks.Server = append(c.hooks.Server, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `server.Intercept(f(g(h())))`.
func (c *ServerClient) Intercept(interceptors ...Interceptor) {
	c.inters.Server = append(c.inters.Server, interceptors...)
}

// Create returns a builder for creating a Server entity.
func (c *ServerClient) Create() *ServerCreate {
	mutation := newServerMutation(c.config, OpCreate)
	return &ServerCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Server entities.
func (c *ServerClient) CreateBulk(builders ...*ServerCreate) *ServerCreateBulk {
	return &ServerCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Server.
func (c *ServerClient) Update() *ServerUpdate {
	mutation := newServerMutation(c.config, OpUpdate)
	return &ServerUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ServerClient) UpdateOne(s *Server) *ServerUpdateOne {
	mutation := newServerMutation(c.config, OpUpdateOne, withServer(s))
	return &ServerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ServerClient) UpdateOneID(id gidx.PrefixedID) *ServerUpdateOne {
	mutation := newServerMutation(c.config, OpUpdateOne, withServerID(id))
	return &ServerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Server.
func (c *ServerClient) Delete() *ServerDelete {
	mutation := newServerMutation(c.config, OpDelete)
	return &ServerDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ServerClient) DeleteOne(s *Server) *ServerDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ServerClient) DeleteOneID(id gidx.PrefixedID) *ServerDeleteOne {
	builder := c.Delete().Where(server.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ServerDeleteOne{builder}
}

// Query returns a query builder for Server.
func (c *ServerClient) Query() *ServerQuery {
	return &ServerQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeServer},
		inters: c.Interceptors(),
	}
}

// Get returns a Server entity by its id.
func (c *ServerClient) Get(ctx context.Context, id gidx.PrefixedID) (*Server, error) {
	return c.Query().Where(server.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ServerClient) GetX(ctx context.Context, id gidx.PrefixedID) *Server {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryProvider queries the provider edge of a Server.
func (c *ServerClient) QueryProvider(s *Server) *ProviderQuery {
	query := (&ProviderClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(server.Table, server.FieldID, id),
			sqlgraph.To(provider.Table, provider.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, server.ProviderTable, server.ProviderColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryServerType queries the server_type edge of a Server.
func (c *ServerClient) QueryServerType(s *Server) *ServerTypeQuery {
	query := (&ServerTypeClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(server.Table, server.FieldID, id),
			sqlgraph.To(servertype.Table, servertype.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, server.ServerTypeTable, server.ServerTypeColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryComponents queries the components edge of a Server.
func (c *ServerClient) QueryComponents(s *Server) *ServerComponentQuery {
	query := (&ServerComponentClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(server.Table, server.FieldID, id),
			sqlgraph.To(servercomponent.Table, servercomponent.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, server.ComponentsTable, server.ComponentsColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ServerClient) Hooks() []Hook {
	return c.hooks.Server
}

// Interceptors returns the client interceptors.
func (c *ServerClient) Interceptors() []Interceptor {
	return c.inters.Server
}

func (c *ServerClient) mutate(ctx context.Context, m *ServerMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ServerCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ServerUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ServerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ServerDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("generated: unknown Server mutation op: %q", m.Op())
	}
}

// ServerComponentClient is a client for the ServerComponent schema.
type ServerComponentClient struct {
	config
}

// NewServerComponentClient returns a client for the ServerComponent from the given config.
func NewServerComponentClient(c config) *ServerComponentClient {
	return &ServerComponentClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `servercomponent.Hooks(f(g(h())))`.
func (c *ServerComponentClient) Use(hooks ...Hook) {
	c.hooks.ServerComponent = append(c.hooks.ServerComponent, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `servercomponent.Intercept(f(g(h())))`.
func (c *ServerComponentClient) Intercept(interceptors ...Interceptor) {
	c.inters.ServerComponent = append(c.inters.ServerComponent, interceptors...)
}

// Create returns a builder for creating a ServerComponent entity.
func (c *ServerComponentClient) Create() *ServerComponentCreate {
	mutation := newServerComponentMutation(c.config, OpCreate)
	return &ServerComponentCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ServerComponent entities.
func (c *ServerComponentClient) CreateBulk(builders ...*ServerComponentCreate) *ServerComponentCreateBulk {
	return &ServerComponentCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ServerComponent.
func (c *ServerComponentClient) Update() *ServerComponentUpdate {
	mutation := newServerComponentMutation(c.config, OpUpdate)
	return &ServerComponentUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ServerComponentClient) UpdateOne(sc *ServerComponent) *ServerComponentUpdateOne {
	mutation := newServerComponentMutation(c.config, OpUpdateOne, withServerComponent(sc))
	return &ServerComponentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ServerComponentClient) UpdateOneID(id gidx.PrefixedID) *ServerComponentUpdateOne {
	mutation := newServerComponentMutation(c.config, OpUpdateOne, withServerComponentID(id))
	return &ServerComponentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ServerComponent.
func (c *ServerComponentClient) Delete() *ServerComponentDelete {
	mutation := newServerComponentMutation(c.config, OpDelete)
	return &ServerComponentDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ServerComponentClient) DeleteOne(sc *ServerComponent) *ServerComponentDeleteOne {
	return c.DeleteOneID(sc.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ServerComponentClient) DeleteOneID(id gidx.PrefixedID) *ServerComponentDeleteOne {
	builder := c.Delete().Where(servercomponent.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ServerComponentDeleteOne{builder}
}

// Query returns a query builder for ServerComponent.
func (c *ServerComponentClient) Query() *ServerComponentQuery {
	return &ServerComponentQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeServerComponent},
		inters: c.Interceptors(),
	}
}

// Get returns a ServerComponent entity by its id.
func (c *ServerComponentClient) Get(ctx context.Context, id gidx.PrefixedID) (*ServerComponent, error) {
	return c.Query().Where(servercomponent.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ServerComponentClient) GetX(ctx context.Context, id gidx.PrefixedID) *ServerComponent {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryComponentType queries the component_type edge of a ServerComponent.
func (c *ServerComponentClient) QueryComponentType(sc *ServerComponent) *ServerComponentTypeQuery {
	query := (&ServerComponentTypeClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := sc.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(servercomponent.Table, servercomponent.FieldID, id),
			sqlgraph.To(servercomponenttype.Table, servercomponenttype.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, servercomponent.ComponentTypeTable, servercomponent.ComponentTypeColumn),
		)
		fromV = sqlgraph.Neighbors(sc.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryServer queries the server edge of a ServerComponent.
func (c *ServerComponentClient) QueryServer(sc *ServerComponent) *ServerQuery {
	query := (&ServerClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := sc.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(servercomponent.Table, servercomponent.FieldID, id),
			sqlgraph.To(server.Table, server.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, servercomponent.ServerTable, servercomponent.ServerColumn),
		)
		fromV = sqlgraph.Neighbors(sc.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ServerComponentClient) Hooks() []Hook {
	return c.hooks.ServerComponent
}

// Interceptors returns the client interceptors.
func (c *ServerComponentClient) Interceptors() []Interceptor {
	return c.inters.ServerComponent
}

func (c *ServerComponentClient) mutate(ctx context.Context, m *ServerComponentMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ServerComponentCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ServerComponentUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ServerComponentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ServerComponentDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("generated: unknown ServerComponent mutation op: %q", m.Op())
	}
}

// ServerComponentTypeClient is a client for the ServerComponentType schema.
type ServerComponentTypeClient struct {
	config
}

// NewServerComponentTypeClient returns a client for the ServerComponentType from the given config.
func NewServerComponentTypeClient(c config) *ServerComponentTypeClient {
	return &ServerComponentTypeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `servercomponenttype.Hooks(f(g(h())))`.
func (c *ServerComponentTypeClient) Use(hooks ...Hook) {
	c.hooks.ServerComponentType = append(c.hooks.ServerComponentType, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `servercomponenttype.Intercept(f(g(h())))`.
func (c *ServerComponentTypeClient) Intercept(interceptors ...Interceptor) {
	c.inters.ServerComponentType = append(c.inters.ServerComponentType, interceptors...)
}

// Create returns a builder for creating a ServerComponentType entity.
func (c *ServerComponentTypeClient) Create() *ServerComponentTypeCreate {
	mutation := newServerComponentTypeMutation(c.config, OpCreate)
	return &ServerComponentTypeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ServerComponentType entities.
func (c *ServerComponentTypeClient) CreateBulk(builders ...*ServerComponentTypeCreate) *ServerComponentTypeCreateBulk {
	return &ServerComponentTypeCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ServerComponentType.
func (c *ServerComponentTypeClient) Update() *ServerComponentTypeUpdate {
	mutation := newServerComponentTypeMutation(c.config, OpUpdate)
	return &ServerComponentTypeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ServerComponentTypeClient) UpdateOne(sct *ServerComponentType) *ServerComponentTypeUpdateOne {
	mutation := newServerComponentTypeMutation(c.config, OpUpdateOne, withServerComponentType(sct))
	return &ServerComponentTypeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ServerComponentTypeClient) UpdateOneID(id gidx.PrefixedID) *ServerComponentTypeUpdateOne {
	mutation := newServerComponentTypeMutation(c.config, OpUpdateOne, withServerComponentTypeID(id))
	return &ServerComponentTypeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ServerComponentType.
func (c *ServerComponentTypeClient) Delete() *ServerComponentTypeDelete {
	mutation := newServerComponentTypeMutation(c.config, OpDelete)
	return &ServerComponentTypeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ServerComponentTypeClient) DeleteOne(sct *ServerComponentType) *ServerComponentTypeDeleteOne {
	return c.DeleteOneID(sct.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ServerComponentTypeClient) DeleteOneID(id gidx.PrefixedID) *ServerComponentTypeDeleteOne {
	builder := c.Delete().Where(servercomponenttype.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ServerComponentTypeDeleteOne{builder}
}

// Query returns a query builder for ServerComponentType.
func (c *ServerComponentTypeClient) Query() *ServerComponentTypeQuery {
	return &ServerComponentTypeQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeServerComponentType},
		inters: c.Interceptors(),
	}
}

// Get returns a ServerComponentType entity by its id.
func (c *ServerComponentTypeClient) Get(ctx context.Context, id gidx.PrefixedID) (*ServerComponentType, error) {
	return c.Query().Where(servercomponenttype.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ServerComponentTypeClient) GetX(ctx context.Context, id gidx.PrefixedID) *ServerComponentType {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ServerComponentTypeClient) Hooks() []Hook {
	return c.hooks.ServerComponentType
}

// Interceptors returns the client interceptors.
func (c *ServerComponentTypeClient) Interceptors() []Interceptor {
	return c.inters.ServerComponentType
}

func (c *ServerComponentTypeClient) mutate(ctx context.Context, m *ServerComponentTypeMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ServerComponentTypeCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ServerComponentTypeUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ServerComponentTypeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ServerComponentTypeDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("generated: unknown ServerComponentType mutation op: %q", m.Op())
	}
}

// ServerTypeClient is a client for the ServerType schema.
type ServerTypeClient struct {
	config
}

// NewServerTypeClient returns a client for the ServerType from the given config.
func NewServerTypeClient(c config) *ServerTypeClient {
	return &ServerTypeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `servertype.Hooks(f(g(h())))`.
func (c *ServerTypeClient) Use(hooks ...Hook) {
	c.hooks.ServerType = append(c.hooks.ServerType, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `servertype.Intercept(f(g(h())))`.
func (c *ServerTypeClient) Intercept(interceptors ...Interceptor) {
	c.inters.ServerType = append(c.inters.ServerType, interceptors...)
}

// Create returns a builder for creating a ServerType entity.
func (c *ServerTypeClient) Create() *ServerTypeCreate {
	mutation := newServerTypeMutation(c.config, OpCreate)
	return &ServerTypeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ServerType entities.
func (c *ServerTypeClient) CreateBulk(builders ...*ServerTypeCreate) *ServerTypeCreateBulk {
	return &ServerTypeCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ServerType.
func (c *ServerTypeClient) Update() *ServerTypeUpdate {
	mutation := newServerTypeMutation(c.config, OpUpdate)
	return &ServerTypeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ServerTypeClient) UpdateOne(st *ServerType) *ServerTypeUpdateOne {
	mutation := newServerTypeMutation(c.config, OpUpdateOne, withServerType(st))
	return &ServerTypeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ServerTypeClient) UpdateOneID(id gidx.PrefixedID) *ServerTypeUpdateOne {
	mutation := newServerTypeMutation(c.config, OpUpdateOne, withServerTypeID(id))
	return &ServerTypeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ServerType.
func (c *ServerTypeClient) Delete() *ServerTypeDelete {
	mutation := newServerTypeMutation(c.config, OpDelete)
	return &ServerTypeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ServerTypeClient) DeleteOne(st *ServerType) *ServerTypeDeleteOne {
	return c.DeleteOneID(st.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ServerTypeClient) DeleteOneID(id gidx.PrefixedID) *ServerTypeDeleteOne {
	builder := c.Delete().Where(servertype.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ServerTypeDeleteOne{builder}
}

// Query returns a query builder for ServerType.
func (c *ServerTypeClient) Query() *ServerTypeQuery {
	return &ServerTypeQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeServerType},
		inters: c.Interceptors(),
	}
}

// Get returns a ServerType entity by its id.
func (c *ServerTypeClient) Get(ctx context.Context, id gidx.PrefixedID) (*ServerType, error) {
	return c.Query().Where(servertype.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ServerTypeClient) GetX(ctx context.Context, id gidx.PrefixedID) *ServerType {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryServers queries the servers edge of a ServerType.
func (c *ServerTypeClient) QueryServers(st *ServerType) *ServerQuery {
	query := (&ServerClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := st.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(servertype.Table, servertype.FieldID, id),
			sqlgraph.To(server.Table, server.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, servertype.ServersTable, servertype.ServersColumn),
		)
		fromV = sqlgraph.Neighbors(st.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ServerTypeClient) Hooks() []Hook {
	return c.hooks.ServerType
}

// Interceptors returns the client interceptors.
func (c *ServerTypeClient) Interceptors() []Interceptor {
	return c.inters.ServerType
}

func (c *ServerTypeClient) mutate(ctx context.Context, m *ServerTypeMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ServerTypeCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ServerTypeUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ServerTypeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ServerTypeDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("generated: unknown ServerType mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Provider, Server, ServerComponent, ServerComponentType, ServerType []ent.Hook
	}
	inters struct {
		Provider, Server, ServerComponent, ServerComponentType,
		ServerType []ent.Interceptor
	}
)
