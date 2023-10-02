// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/chanzuckerberg/happy/api/pkg/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/chanzuckerberg/happy/api/pkg/ent/appconfig"
	"github.com/chanzuckerberg/happy/api/pkg/ent/appstack"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// AppConfig is the client for interacting with the AppConfig builders.
	AppConfig *AppConfigClient
	// AppStack is the client for interacting with the AppStack builders.
	AppStack *AppStackClient
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
	c.AppConfig = NewAppConfigClient(c.config)
	c.AppStack = NewAppStackClient(c.config)
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
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:       ctx,
		config:    cfg,
		AppConfig: NewAppConfigClient(cfg),
		AppStack:  NewAppStackClient(cfg),
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
		ctx:       ctx,
		config:    cfg,
		AppConfig: NewAppConfigClient(cfg),
		AppStack:  NewAppStackClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		AppConfig.
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
	c.AppConfig.Use(hooks...)
	c.AppStack.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.AppConfig.Intercept(interceptors...)
	c.AppStack.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *AppConfigMutation:
		return c.AppConfig.mutate(ctx, m)
	case *AppStackMutation:
		return c.AppStack.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// AppConfigClient is a client for the AppConfig schema.
type AppConfigClient struct {
	config
}

// NewAppConfigClient returns a client for the AppConfig from the given config.
func NewAppConfigClient(c config) *AppConfigClient {
	return &AppConfigClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `appconfig.Hooks(f(g(h())))`.
func (c *AppConfigClient) Use(hooks ...Hook) {
	c.hooks.AppConfig = append(c.hooks.AppConfig, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `appconfig.Intercept(f(g(h())))`.
func (c *AppConfigClient) Intercept(interceptors ...Interceptor) {
	c.inters.AppConfig = append(c.inters.AppConfig, interceptors...)
}

// Create returns a builder for creating a AppConfig entity.
func (c *AppConfigClient) Create() *AppConfigCreate {
	mutation := newAppConfigMutation(c.config, OpCreate)
	return &AppConfigCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of AppConfig entities.
func (c *AppConfigClient) CreateBulk(builders ...*AppConfigCreate) *AppConfigCreateBulk {
	return &AppConfigCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for AppConfig.
func (c *AppConfigClient) Update() *AppConfigUpdate {
	mutation := newAppConfigMutation(c.config, OpUpdate)
	return &AppConfigUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AppConfigClient) UpdateOne(ac *AppConfig) *AppConfigUpdateOne {
	mutation := newAppConfigMutation(c.config, OpUpdateOne, withAppConfig(ac))
	return &AppConfigUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AppConfigClient) UpdateOneID(id uint) *AppConfigUpdateOne {
	mutation := newAppConfigMutation(c.config, OpUpdateOne, withAppConfigID(id))
	return &AppConfigUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for AppConfig.
func (c *AppConfigClient) Delete() *AppConfigDelete {
	mutation := newAppConfigMutation(c.config, OpDelete)
	return &AppConfigDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *AppConfigClient) DeleteOne(ac *AppConfig) *AppConfigDeleteOne {
	return c.DeleteOneID(ac.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *AppConfigClient) DeleteOneID(id uint) *AppConfigDeleteOne {
	builder := c.Delete().Where(appconfig.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AppConfigDeleteOne{builder}
}

// Query returns a query builder for AppConfig.
func (c *AppConfigClient) Query() *AppConfigQuery {
	return &AppConfigQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeAppConfig},
		inters: c.Interceptors(),
	}
}

// Get returns a AppConfig entity by its id.
func (c *AppConfigClient) Get(ctx context.Context, id uint) (*AppConfig, error) {
	return c.Query().Where(appconfig.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AppConfigClient) GetX(ctx context.Context, id uint) *AppConfig {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *AppConfigClient) Hooks() []Hook {
	return c.hooks.AppConfig
}

// Interceptors returns the client interceptors.
func (c *AppConfigClient) Interceptors() []Interceptor {
	return c.inters.AppConfig
}

func (c *AppConfigClient) mutate(ctx context.Context, m *AppConfigMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&AppConfigCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&AppConfigUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&AppConfigUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&AppConfigDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown AppConfig mutation op: %q", m.Op())
	}
}

// AppStackClient is a client for the AppStack schema.
type AppStackClient struct {
	config
}

// NewAppStackClient returns a client for the AppStack from the given config.
func NewAppStackClient(c config) *AppStackClient {
	return &AppStackClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `appstack.Hooks(f(g(h())))`.
func (c *AppStackClient) Use(hooks ...Hook) {
	c.hooks.AppStack = append(c.hooks.AppStack, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `appstack.Intercept(f(g(h())))`.
func (c *AppStackClient) Intercept(interceptors ...Interceptor) {
	c.inters.AppStack = append(c.inters.AppStack, interceptors...)
}

// Create returns a builder for creating a AppStack entity.
func (c *AppStackClient) Create() *AppStackCreate {
	mutation := newAppStackMutation(c.config, OpCreate)
	return &AppStackCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of AppStack entities.
func (c *AppStackClient) CreateBulk(builders ...*AppStackCreate) *AppStackCreateBulk {
	return &AppStackCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for AppStack.
func (c *AppStackClient) Update() *AppStackUpdate {
	mutation := newAppStackMutation(c.config, OpUpdate)
	return &AppStackUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AppStackClient) UpdateOne(as *AppStack) *AppStackUpdateOne {
	mutation := newAppStackMutation(c.config, OpUpdateOne, withAppStack(as))
	return &AppStackUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AppStackClient) UpdateOneID(id uint) *AppStackUpdateOne {
	mutation := newAppStackMutation(c.config, OpUpdateOne, withAppStackID(id))
	return &AppStackUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for AppStack.
func (c *AppStackClient) Delete() *AppStackDelete {
	mutation := newAppStackMutation(c.config, OpDelete)
	return &AppStackDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *AppStackClient) DeleteOne(as *AppStack) *AppStackDeleteOne {
	return c.DeleteOneID(as.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *AppStackClient) DeleteOneID(id uint) *AppStackDeleteOne {
	builder := c.Delete().Where(appstack.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AppStackDeleteOne{builder}
}

// Query returns a query builder for AppStack.
func (c *AppStackClient) Query() *AppStackQuery {
	return &AppStackQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeAppStack},
		inters: c.Interceptors(),
	}
}

// Get returns a AppStack entity by its id.
func (c *AppStackClient) Get(ctx context.Context, id uint) (*AppStack, error) {
	return c.Query().Where(appstack.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AppStackClient) GetX(ctx context.Context, id uint) *AppStack {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *AppStackClient) Hooks() []Hook {
	return c.hooks.AppStack
}

// Interceptors returns the client interceptors.
func (c *AppStackClient) Interceptors() []Interceptor {
	return c.inters.AppStack
}

func (c *AppStackClient) mutate(ctx context.Context, m *AppStackMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&AppStackCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&AppStackUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&AppStackUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&AppStackDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown AppStack mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		AppConfig, AppStack []ent.Hook
	}
	inters struct {
		AppConfig, AppStack []ent.Interceptor
	}
)
