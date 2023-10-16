// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/chanzuckerberg/happy/api/pkg/ent/appconfig"
)

// AppConfigCreate is the builder for creating a AppConfig entity.
type AppConfigCreate struct {
	config
	mutation *AppConfigMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (acc *AppConfigCreate) SetCreatedAt(t time.Time) *AppConfigCreate {
	acc.mutation.SetCreatedAt(t)
	return acc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (acc *AppConfigCreate) SetNillableCreatedAt(t *time.Time) *AppConfigCreate {
	if t != nil {
		acc.SetCreatedAt(*t)
	}
	return acc
}

// SetUpdatedAt sets the "updated_at" field.
func (acc *AppConfigCreate) SetUpdatedAt(t time.Time) *AppConfigCreate {
	acc.mutation.SetUpdatedAt(t)
	return acc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (acc *AppConfigCreate) SetNillableUpdatedAt(t *time.Time) *AppConfigCreate {
	if t != nil {
		acc.SetUpdatedAt(*t)
	}
	return acc
}

// SetDeletedAt sets the "deleted_at" field.
func (acc *AppConfigCreate) SetDeletedAt(t time.Time) *AppConfigCreate {
	acc.mutation.SetDeletedAt(t)
	return acc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (acc *AppConfigCreate) SetNillableDeletedAt(t *time.Time) *AppConfigCreate {
	if t != nil {
		acc.SetDeletedAt(*t)
	}
	return acc
}

// SetAppName sets the "app_name" field.
func (acc *AppConfigCreate) SetAppName(s string) *AppConfigCreate {
	acc.mutation.SetAppName(s)
	return acc
}

// SetEnvironment sets the "environment" field.
func (acc *AppConfigCreate) SetEnvironment(s string) *AppConfigCreate {
	acc.mutation.SetEnvironment(s)
	return acc
}

// SetStack sets the "stack" field.
func (acc *AppConfigCreate) SetStack(s string) *AppConfigCreate {
	acc.mutation.SetStack(s)
	return acc
}

// SetNillableStack sets the "stack" field if the given value is not nil.
func (acc *AppConfigCreate) SetNillableStack(s *string) *AppConfigCreate {
	if s != nil {
		acc.SetStack(*s)
	}
	return acc
}

// SetKey sets the "key" field.
func (acc *AppConfigCreate) SetKey(s string) *AppConfigCreate {
	acc.mutation.SetKey(s)
	return acc
}

// SetValue sets the "value" field.
func (acc *AppConfigCreate) SetValue(s string) *AppConfigCreate {
	acc.mutation.SetValue(s)
	return acc
}

// SetSource sets the "source" field.
func (acc *AppConfigCreate) SetSource(a appconfig.Source) *AppConfigCreate {
	acc.mutation.SetSource(a)
	return acc
}

// SetNillableSource sets the "source" field if the given value is not nil.
func (acc *AppConfigCreate) SetNillableSource(a *appconfig.Source) *AppConfigCreate {
	if a != nil {
		acc.SetSource(*a)
	}
	return acc
}

// SetID sets the "id" field.
func (acc *AppConfigCreate) SetID(u uint) *AppConfigCreate {
	acc.mutation.SetID(u)
	return acc
}

// Mutation returns the AppConfigMutation object of the builder.
func (acc *AppConfigCreate) Mutation() *AppConfigMutation {
	return acc.mutation
}

// Save creates the AppConfig in the database.
func (acc *AppConfigCreate) Save(ctx context.Context) (*AppConfig, error) {
	if err := acc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, acc.sqlSave, acc.mutation, acc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (acc *AppConfigCreate) SaveX(ctx context.Context) *AppConfig {
	v, err := acc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acc *AppConfigCreate) Exec(ctx context.Context) error {
	_, err := acc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acc *AppConfigCreate) ExecX(ctx context.Context) {
	if err := acc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (acc *AppConfigCreate) defaults() error {
	if _, ok := acc.mutation.CreatedAt(); !ok {
		if appconfig.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized appconfig.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := appconfig.DefaultCreatedAt()
		acc.mutation.SetCreatedAt(v)
	}
	if _, ok := acc.mutation.UpdatedAt(); !ok {
		if appconfig.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized appconfig.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := appconfig.DefaultUpdatedAt()
		acc.mutation.SetUpdatedAt(v)
	}
	if _, ok := acc.mutation.Stack(); !ok {
		v := appconfig.DefaultStack
		acc.mutation.SetStack(v)
	}
	if _, ok := acc.mutation.Source(); !ok {
		v := appconfig.DefaultSource
		acc.mutation.SetSource(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (acc *AppConfigCreate) check() error {
	if _, ok := acc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "AppConfig.created_at"`)}
	}
	if _, ok := acc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "AppConfig.updated_at"`)}
	}
	if _, ok := acc.mutation.AppName(); !ok {
		return &ValidationError{Name: "app_name", err: errors.New(`ent: missing required field "AppConfig.app_name"`)}
	}
	if _, ok := acc.mutation.Environment(); !ok {
		return &ValidationError{Name: "environment", err: errors.New(`ent: missing required field "AppConfig.environment"`)}
	}
	if _, ok := acc.mutation.Stack(); !ok {
		return &ValidationError{Name: "stack", err: errors.New(`ent: missing required field "AppConfig.stack"`)}
	}
	if _, ok := acc.mutation.Key(); !ok {
		return &ValidationError{Name: "key", err: errors.New(`ent: missing required field "AppConfig.key"`)}
	}
	if _, ok := acc.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New(`ent: missing required field "AppConfig.value"`)}
	}
	if _, ok := acc.mutation.Source(); !ok {
		return &ValidationError{Name: "source", err: errors.New(`ent: missing required field "AppConfig.source"`)}
	}
	if v, ok := acc.mutation.Source(); ok {
		if err := appconfig.SourceValidator(v); err != nil {
			return &ValidationError{Name: "source", err: fmt.Errorf(`ent: validator failed for field "AppConfig.source": %w`, err)}
		}
	}
	return nil
}

func (acc *AppConfigCreate) sqlSave(ctx context.Context) (*AppConfig, error) {
	if err := acc.check(); err != nil {
		return nil, err
	}
	_node, _spec := acc.createSpec()
	if err := sqlgraph.CreateNode(ctx, acc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint(id)
	}
	acc.mutation.id = &_node.ID
	acc.mutation.done = true
	return _node, nil
}

func (acc *AppConfigCreate) createSpec() (*AppConfig, *sqlgraph.CreateSpec) {
	var (
		_node = &AppConfig{config: acc.config}
		_spec = sqlgraph.NewCreateSpec(appconfig.Table, sqlgraph.NewFieldSpec(appconfig.FieldID, field.TypeUint))
	)
	_spec.OnConflict = acc.conflict
	if id, ok := acc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := acc.mutation.CreatedAt(); ok {
		_spec.SetField(appconfig.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := acc.mutation.UpdatedAt(); ok {
		_spec.SetField(appconfig.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := acc.mutation.DeletedAt(); ok {
		_spec.SetField(appconfig.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := acc.mutation.AppName(); ok {
		_spec.SetField(appconfig.FieldAppName, field.TypeString, value)
		_node.AppName = value
	}
	if value, ok := acc.mutation.Environment(); ok {
		_spec.SetField(appconfig.FieldEnvironment, field.TypeString, value)
		_node.Environment = value
	}
	if value, ok := acc.mutation.Stack(); ok {
		_spec.SetField(appconfig.FieldStack, field.TypeString, value)
		_node.Stack = value
	}
	if value, ok := acc.mutation.Key(); ok {
		_spec.SetField(appconfig.FieldKey, field.TypeString, value)
		_node.Key = value
	}
	if value, ok := acc.mutation.Value(); ok {
		_spec.SetField(appconfig.FieldValue, field.TypeString, value)
		_node.Value = value
	}
	if value, ok := acc.mutation.Source(); ok {
		_spec.SetField(appconfig.FieldSource, field.TypeEnum, value)
		_node.Source = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppConfig.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppConfigUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (acc *AppConfigCreate) OnConflict(opts ...sql.ConflictOption) *AppConfigUpsertOne {
	acc.conflict = opts
	return &AppConfigUpsertOne{
		create: acc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppConfig.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (acc *AppConfigCreate) OnConflictColumns(columns ...string) *AppConfigUpsertOne {
	acc.conflict = append(acc.conflict, sql.ConflictColumns(columns...))
	return &AppConfigUpsertOne{
		create: acc,
	}
}

type (
	// AppConfigUpsertOne is the builder for "upsert"-ing
	//  one AppConfig node.
	AppConfigUpsertOne struct {
		create *AppConfigCreate
	}

	// AppConfigUpsert is the "OnConflict" setter.
	AppConfigUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *AppConfigUpsert) SetUpdatedAt(v time.Time) *AppConfigUpsert {
	u.Set(appconfig.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppConfigUpsert) UpdateUpdatedAt() *AppConfigUpsert {
	u.SetExcluded(appconfig.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppConfigUpsert) SetDeletedAt(v time.Time) *AppConfigUpsert {
	u.Set(appconfig.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppConfigUpsert) UpdateDeletedAt() *AppConfigUpsert {
	u.SetExcluded(appconfig.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *AppConfigUpsert) ClearDeletedAt() *AppConfigUpsert {
	u.SetNull(appconfig.FieldDeletedAt)
	return u
}

// SetAppName sets the "app_name" field.
func (u *AppConfigUpsert) SetAppName(v string) *AppConfigUpsert {
	u.Set(appconfig.FieldAppName, v)
	return u
}

// UpdateAppName sets the "app_name" field to the value that was provided on create.
func (u *AppConfigUpsert) UpdateAppName() *AppConfigUpsert {
	u.SetExcluded(appconfig.FieldAppName)
	return u
}

// SetEnvironment sets the "environment" field.
func (u *AppConfigUpsert) SetEnvironment(v string) *AppConfigUpsert {
	u.Set(appconfig.FieldEnvironment, v)
	return u
}

// UpdateEnvironment sets the "environment" field to the value that was provided on create.
func (u *AppConfigUpsert) UpdateEnvironment() *AppConfigUpsert {
	u.SetExcluded(appconfig.FieldEnvironment)
	return u
}

// SetStack sets the "stack" field.
func (u *AppConfigUpsert) SetStack(v string) *AppConfigUpsert {
	u.Set(appconfig.FieldStack, v)
	return u
}

// UpdateStack sets the "stack" field to the value that was provided on create.
func (u *AppConfigUpsert) UpdateStack() *AppConfigUpsert {
	u.SetExcluded(appconfig.FieldStack)
	return u
}

// SetKey sets the "key" field.
func (u *AppConfigUpsert) SetKey(v string) *AppConfigUpsert {
	u.Set(appconfig.FieldKey, v)
	return u
}

// UpdateKey sets the "key" field to the value that was provided on create.
func (u *AppConfigUpsert) UpdateKey() *AppConfigUpsert {
	u.SetExcluded(appconfig.FieldKey)
	return u
}

// SetValue sets the "value" field.
func (u *AppConfigUpsert) SetValue(v string) *AppConfigUpsert {
	u.Set(appconfig.FieldValue, v)
	return u
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *AppConfigUpsert) UpdateValue() *AppConfigUpsert {
	u.SetExcluded(appconfig.FieldValue)
	return u
}

// SetSource sets the "source" field.
func (u *AppConfigUpsert) SetSource(v appconfig.Source) *AppConfigUpsert {
	u.Set(appconfig.FieldSource, v)
	return u
}

// UpdateSource sets the "source" field to the value that was provided on create.
func (u *AppConfigUpsert) UpdateSource() *AppConfigUpsert {
	u.SetExcluded(appconfig.FieldSource)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.AppConfig.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(appconfig.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AppConfigUpsertOne) UpdateNewValues() *AppConfigUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(appconfig.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(appconfig.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AppConfig.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *AppConfigUpsertOne) Ignore() *AppConfigUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppConfigUpsertOne) DoNothing() *AppConfigUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppConfigCreate.OnConflict
// documentation for more info.
func (u *AppConfigUpsertOne) Update(set func(*AppConfigUpsert)) *AppConfigUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppConfigUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppConfigUpsertOne) SetUpdatedAt(v time.Time) *AppConfigUpsertOne {
	return u.Update(func(s *AppConfigUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppConfigUpsertOne) UpdateUpdatedAt() *AppConfigUpsertOne {
	return u.Update(func(s *AppConfigUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppConfigUpsertOne) SetDeletedAt(v time.Time) *AppConfigUpsertOne {
	return u.Update(func(s *AppConfigUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppConfigUpsertOne) UpdateDeletedAt() *AppConfigUpsertOne {
	return u.Update(func(s *AppConfigUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *AppConfigUpsertOne) ClearDeletedAt() *AppConfigUpsertOne {
	return u.Update(func(s *AppConfigUpsert) {
		s.ClearDeletedAt()
	})
}

// SetAppName sets the "app_name" field.
func (u *AppConfigUpsertOne) SetAppName(v string) *AppConfigUpsertOne {
	return u.Update(func(s *AppConfigUpsert) {
		s.SetAppName(v)
	})
}

// UpdateAppName sets the "app_name" field to the value that was provided on create.
func (u *AppConfigUpsertOne) UpdateAppName() *AppConfigUpsertOne {
	return u.Update(func(s *AppConfigUpsert) {
		s.UpdateAppName()
	})
}

// SetEnvironment sets the "environment" field.
func (u *AppConfigUpsertOne) SetEnvironment(v string) *AppConfigUpsertOne {
	return u.Update(func(s *AppConfigUpsert) {
		s.SetEnvironment(v)
	})
}

// UpdateEnvironment sets the "environment" field to the value that was provided on create.
func (u *AppConfigUpsertOne) UpdateEnvironment() *AppConfigUpsertOne {
	return u.Update(func(s *AppConfigUpsert) {
		s.UpdateEnvironment()
	})
}

// SetStack sets the "stack" field.
func (u *AppConfigUpsertOne) SetStack(v string) *AppConfigUpsertOne {
	return u.Update(func(s *AppConfigUpsert) {
		s.SetStack(v)
	})
}

// UpdateStack sets the "stack" field to the value that was provided on create.
func (u *AppConfigUpsertOne) UpdateStack() *AppConfigUpsertOne {
	return u.Update(func(s *AppConfigUpsert) {
		s.UpdateStack()
	})
}

// SetKey sets the "key" field.
func (u *AppConfigUpsertOne) SetKey(v string) *AppConfigUpsertOne {
	return u.Update(func(s *AppConfigUpsert) {
		s.SetKey(v)
	})
}

// UpdateKey sets the "key" field to the value that was provided on create.
func (u *AppConfigUpsertOne) UpdateKey() *AppConfigUpsertOne {
	return u.Update(func(s *AppConfigUpsert) {
		s.UpdateKey()
	})
}

// SetValue sets the "value" field.
func (u *AppConfigUpsertOne) SetValue(v string) *AppConfigUpsertOne {
	return u.Update(func(s *AppConfigUpsert) {
		s.SetValue(v)
	})
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *AppConfigUpsertOne) UpdateValue() *AppConfigUpsertOne {
	return u.Update(func(s *AppConfigUpsert) {
		s.UpdateValue()
	})
}

// SetSource sets the "source" field.
func (u *AppConfigUpsertOne) SetSource(v appconfig.Source) *AppConfigUpsertOne {
	return u.Update(func(s *AppConfigUpsert) {
		s.SetSource(v)
	})
}

// UpdateSource sets the "source" field to the value that was provided on create.
func (u *AppConfigUpsertOne) UpdateSource() *AppConfigUpsertOne {
	return u.Update(func(s *AppConfigUpsert) {
		s.UpdateSource()
	})
}

// Exec executes the query.
func (u *AppConfigUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppConfigCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppConfigUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AppConfigUpsertOne) ID(ctx context.Context) (id uint, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AppConfigUpsertOne) IDX(ctx context.Context) uint {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AppConfigCreateBulk is the builder for creating many AppConfig entities in bulk.
type AppConfigCreateBulk struct {
	config
	err      error
	builders []*AppConfigCreate
	conflict []sql.ConflictOption
}

// Save creates the AppConfig entities in the database.
func (accb *AppConfigCreateBulk) Save(ctx context.Context) ([]*AppConfig, error) {
	if accb.err != nil {
		return nil, accb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(accb.builders))
	nodes := make([]*AppConfig, len(accb.builders))
	mutators := make([]Mutator, len(accb.builders))
	for i := range accb.builders {
		func(i int, root context.Context) {
			builder := accb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AppConfigMutation)
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
					_, err = mutators[i+1].Mutate(root, accb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = accb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, accb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint(id)
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
		if _, err := mutators[0].Mutate(ctx, accb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (accb *AppConfigCreateBulk) SaveX(ctx context.Context) []*AppConfig {
	v, err := accb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (accb *AppConfigCreateBulk) Exec(ctx context.Context) error {
	_, err := accb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (accb *AppConfigCreateBulk) ExecX(ctx context.Context) {
	if err := accb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppConfig.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppConfigUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (accb *AppConfigCreateBulk) OnConflict(opts ...sql.ConflictOption) *AppConfigUpsertBulk {
	accb.conflict = opts
	return &AppConfigUpsertBulk{
		create: accb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppConfig.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (accb *AppConfigCreateBulk) OnConflictColumns(columns ...string) *AppConfigUpsertBulk {
	accb.conflict = append(accb.conflict, sql.ConflictColumns(columns...))
	return &AppConfigUpsertBulk{
		create: accb,
	}
}

// AppConfigUpsertBulk is the builder for "upsert"-ing
// a bulk of AppConfig nodes.
type AppConfigUpsertBulk struct {
	create *AppConfigCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.AppConfig.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(appconfig.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AppConfigUpsertBulk) UpdateNewValues() *AppConfigUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(appconfig.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(appconfig.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AppConfig.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *AppConfigUpsertBulk) Ignore() *AppConfigUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppConfigUpsertBulk) DoNothing() *AppConfigUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppConfigCreateBulk.OnConflict
// documentation for more info.
func (u *AppConfigUpsertBulk) Update(set func(*AppConfigUpsert)) *AppConfigUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppConfigUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppConfigUpsertBulk) SetUpdatedAt(v time.Time) *AppConfigUpsertBulk {
	return u.Update(func(s *AppConfigUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppConfigUpsertBulk) UpdateUpdatedAt() *AppConfigUpsertBulk {
	return u.Update(func(s *AppConfigUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppConfigUpsertBulk) SetDeletedAt(v time.Time) *AppConfigUpsertBulk {
	return u.Update(func(s *AppConfigUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppConfigUpsertBulk) UpdateDeletedAt() *AppConfigUpsertBulk {
	return u.Update(func(s *AppConfigUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *AppConfigUpsertBulk) ClearDeletedAt() *AppConfigUpsertBulk {
	return u.Update(func(s *AppConfigUpsert) {
		s.ClearDeletedAt()
	})
}

// SetAppName sets the "app_name" field.
func (u *AppConfigUpsertBulk) SetAppName(v string) *AppConfigUpsertBulk {
	return u.Update(func(s *AppConfigUpsert) {
		s.SetAppName(v)
	})
}

// UpdateAppName sets the "app_name" field to the value that was provided on create.
func (u *AppConfigUpsertBulk) UpdateAppName() *AppConfigUpsertBulk {
	return u.Update(func(s *AppConfigUpsert) {
		s.UpdateAppName()
	})
}

// SetEnvironment sets the "environment" field.
func (u *AppConfigUpsertBulk) SetEnvironment(v string) *AppConfigUpsertBulk {
	return u.Update(func(s *AppConfigUpsert) {
		s.SetEnvironment(v)
	})
}

// UpdateEnvironment sets the "environment" field to the value that was provided on create.
func (u *AppConfigUpsertBulk) UpdateEnvironment() *AppConfigUpsertBulk {
	return u.Update(func(s *AppConfigUpsert) {
		s.UpdateEnvironment()
	})
}

// SetStack sets the "stack" field.
func (u *AppConfigUpsertBulk) SetStack(v string) *AppConfigUpsertBulk {
	return u.Update(func(s *AppConfigUpsert) {
		s.SetStack(v)
	})
}

// UpdateStack sets the "stack" field to the value that was provided on create.
func (u *AppConfigUpsertBulk) UpdateStack() *AppConfigUpsertBulk {
	return u.Update(func(s *AppConfigUpsert) {
		s.UpdateStack()
	})
}

// SetKey sets the "key" field.
func (u *AppConfigUpsertBulk) SetKey(v string) *AppConfigUpsertBulk {
	return u.Update(func(s *AppConfigUpsert) {
		s.SetKey(v)
	})
}

// UpdateKey sets the "key" field to the value that was provided on create.
func (u *AppConfigUpsertBulk) UpdateKey() *AppConfigUpsertBulk {
	return u.Update(func(s *AppConfigUpsert) {
		s.UpdateKey()
	})
}

// SetValue sets the "value" field.
func (u *AppConfigUpsertBulk) SetValue(v string) *AppConfigUpsertBulk {
	return u.Update(func(s *AppConfigUpsert) {
		s.SetValue(v)
	})
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *AppConfigUpsertBulk) UpdateValue() *AppConfigUpsertBulk {
	return u.Update(func(s *AppConfigUpsert) {
		s.UpdateValue()
	})
}

// SetSource sets the "source" field.
func (u *AppConfigUpsertBulk) SetSource(v appconfig.Source) *AppConfigUpsertBulk {
	return u.Update(func(s *AppConfigUpsert) {
		s.SetSource(v)
	})
}

// UpdateSource sets the "source" field to the value that was provided on create.
func (u *AppConfigUpsertBulk) UpdateSource() *AppConfigUpsertBulk {
	return u.Update(func(s *AppConfigUpsert) {
		s.UpdateSource()
	})
}

// Exec executes the query.
func (u *AppConfigUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AppConfigCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppConfigCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppConfigUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
