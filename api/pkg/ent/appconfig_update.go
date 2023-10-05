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
	"github.com/chanzuckerberg/happy/api/pkg/ent/predicate"
)

// AppConfigUpdate is the builder for updating AppConfig entities.
type AppConfigUpdate struct {
	config
	hooks    []Hook
	mutation *AppConfigMutation
}

// Where appends a list predicates to the AppConfigUpdate builder.
func (acu *AppConfigUpdate) Where(ps ...predicate.AppConfig) *AppConfigUpdate {
	acu.mutation.Where(ps...)
	return acu
}

// SetUpdatedAt sets the "updated_at" field.
func (acu *AppConfigUpdate) SetUpdatedAt(t time.Time) *AppConfigUpdate {
	acu.mutation.SetUpdatedAt(t)
	return acu
}

// SetDeletedAt sets the "deleted_at" field.
func (acu *AppConfigUpdate) SetDeletedAt(t time.Time) *AppConfigUpdate {
	acu.mutation.SetDeletedAt(t)
	return acu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (acu *AppConfigUpdate) SetNillableDeletedAt(t *time.Time) *AppConfigUpdate {
	if t != nil {
		acu.SetDeletedAt(*t)
	}
	return acu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (acu *AppConfigUpdate) ClearDeletedAt() *AppConfigUpdate {
	acu.mutation.ClearDeletedAt()
	return acu
}

// SetAppName sets the "app_name" field.
func (acu *AppConfigUpdate) SetAppName(s string) *AppConfigUpdate {
	acu.mutation.SetAppName(s)
	return acu
}

// SetEnvironment sets the "environment" field.
func (acu *AppConfigUpdate) SetEnvironment(s string) *AppConfigUpdate {
	acu.mutation.SetEnvironment(s)
	return acu
}

// SetStack sets the "stack" field.
func (acu *AppConfigUpdate) SetStack(s string) *AppConfigUpdate {
	acu.mutation.SetStack(s)
	return acu
}

// SetKey sets the "key" field.
func (acu *AppConfigUpdate) SetKey(s string) *AppConfigUpdate {
	acu.mutation.SetKey(s)
	return acu
}

// SetValue sets the "value" field.
func (acu *AppConfigUpdate) SetValue(s string) *AppConfigUpdate {
	acu.mutation.SetValue(s)
	return acu
}

// Mutation returns the AppConfigMutation object of the builder.
func (acu *AppConfigUpdate) Mutation() *AppConfigMutation {
	return acu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (acu *AppConfigUpdate) Save(ctx context.Context) (int, error) {
	acu.defaults()
	return withHooks(ctx, acu.sqlSave, acu.mutation, acu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (acu *AppConfigUpdate) SaveX(ctx context.Context) int {
	affected, err := acu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (acu *AppConfigUpdate) Exec(ctx context.Context) error {
	_, err := acu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acu *AppConfigUpdate) ExecX(ctx context.Context) {
	if err := acu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (acu *AppConfigUpdate) defaults() {
	if _, ok := acu.mutation.UpdatedAt(); !ok {
		v := appconfig.UpdateDefaultUpdatedAt()
		acu.mutation.SetUpdatedAt(v)
	}
}

func (acu *AppConfigUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(appconfig.Table, appconfig.Columns, sqlgraph.NewFieldSpec(appconfig.FieldID, field.TypeUint))
	if ps := acu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := acu.mutation.UpdatedAt(); ok {
		_spec.SetField(appconfig.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := acu.mutation.DeletedAt(); ok {
		_spec.SetField(appconfig.FieldDeletedAt, field.TypeTime, value)
	}
	if acu.mutation.DeletedAtCleared() {
		_spec.ClearField(appconfig.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := acu.mutation.AppName(); ok {
		_spec.SetField(appconfig.FieldAppName, field.TypeString, value)
	}
	if value, ok := acu.mutation.Environment(); ok {
		_spec.SetField(appconfig.FieldEnvironment, field.TypeString, value)
	}
	if value, ok := acu.mutation.Stack(); ok {
		_spec.SetField(appconfig.FieldStack, field.TypeString, value)
	}
	if value, ok := acu.mutation.Key(); ok {
		_spec.SetField(appconfig.FieldKey, field.TypeString, value)
	}
	if value, ok := acu.mutation.Value(); ok {
		_spec.SetField(appconfig.FieldValue, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, acu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appconfig.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	acu.mutation.done = true
	return n, nil
}

// AppConfigUpdateOne is the builder for updating a single AppConfig entity.
type AppConfigUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AppConfigMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (acuo *AppConfigUpdateOne) SetUpdatedAt(t time.Time) *AppConfigUpdateOne {
	acuo.mutation.SetUpdatedAt(t)
	return acuo
}

// SetDeletedAt sets the "deleted_at" field.
func (acuo *AppConfigUpdateOne) SetDeletedAt(t time.Time) *AppConfigUpdateOne {
	acuo.mutation.SetDeletedAt(t)
	return acuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (acuo *AppConfigUpdateOne) SetNillableDeletedAt(t *time.Time) *AppConfigUpdateOne {
	if t != nil {
		acuo.SetDeletedAt(*t)
	}
	return acuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (acuo *AppConfigUpdateOne) ClearDeletedAt() *AppConfigUpdateOne {
	acuo.mutation.ClearDeletedAt()
	return acuo
}

// SetAppName sets the "app_name" field.
func (acuo *AppConfigUpdateOne) SetAppName(s string) *AppConfigUpdateOne {
	acuo.mutation.SetAppName(s)
	return acuo
}

// SetEnvironment sets the "environment" field.
func (acuo *AppConfigUpdateOne) SetEnvironment(s string) *AppConfigUpdateOne {
	acuo.mutation.SetEnvironment(s)
	return acuo
}

// SetStack sets the "stack" field.
func (acuo *AppConfigUpdateOne) SetStack(s string) *AppConfigUpdateOne {
	acuo.mutation.SetStack(s)
	return acuo
}

// SetKey sets the "key" field.
func (acuo *AppConfigUpdateOne) SetKey(s string) *AppConfigUpdateOne {
	acuo.mutation.SetKey(s)
	return acuo
}

// SetValue sets the "value" field.
func (acuo *AppConfigUpdateOne) SetValue(s string) *AppConfigUpdateOne {
	acuo.mutation.SetValue(s)
	return acuo
}

// Mutation returns the AppConfigMutation object of the builder.
func (acuo *AppConfigUpdateOne) Mutation() *AppConfigMutation {
	return acuo.mutation
}

// Where appends a list predicates to the AppConfigUpdate builder.
func (acuo *AppConfigUpdateOne) Where(ps ...predicate.AppConfig) *AppConfigUpdateOne {
	acuo.mutation.Where(ps...)
	return acuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (acuo *AppConfigUpdateOne) Select(field string, fields ...string) *AppConfigUpdateOne {
	acuo.fields = append([]string{field}, fields...)
	return acuo
}

// Save executes the query and returns the updated AppConfig entity.
func (acuo *AppConfigUpdateOne) Save(ctx context.Context) (*AppConfig, error) {
	acuo.defaults()
	return withHooks(ctx, acuo.sqlSave, acuo.mutation, acuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (acuo *AppConfigUpdateOne) SaveX(ctx context.Context) *AppConfig {
	node, err := acuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (acuo *AppConfigUpdateOne) Exec(ctx context.Context) error {
	_, err := acuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acuo *AppConfigUpdateOne) ExecX(ctx context.Context) {
	if err := acuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (acuo *AppConfigUpdateOne) defaults() {
	if _, ok := acuo.mutation.UpdatedAt(); !ok {
		v := appconfig.UpdateDefaultUpdatedAt()
		acuo.mutation.SetUpdatedAt(v)
	}
}

func (acuo *AppConfigUpdateOne) sqlSave(ctx context.Context) (_node *AppConfig, err error) {
	_spec := sqlgraph.NewUpdateSpec(appconfig.Table, appconfig.Columns, sqlgraph.NewFieldSpec(appconfig.FieldID, field.TypeUint))
	id, ok := acuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AppConfig.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := acuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appconfig.FieldID)
		for _, f := range fields {
			if !appconfig.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != appconfig.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := acuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := acuo.mutation.UpdatedAt(); ok {
		_spec.SetField(appconfig.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := acuo.mutation.DeletedAt(); ok {
		_spec.SetField(appconfig.FieldDeletedAt, field.TypeTime, value)
	}
	if acuo.mutation.DeletedAtCleared() {
		_spec.ClearField(appconfig.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := acuo.mutation.AppName(); ok {
		_spec.SetField(appconfig.FieldAppName, field.TypeString, value)
	}
	if value, ok := acuo.mutation.Environment(); ok {
		_spec.SetField(appconfig.FieldEnvironment, field.TypeString, value)
	}
	if value, ok := acuo.mutation.Stack(); ok {
		_spec.SetField(appconfig.FieldStack, field.TypeString, value)
	}
	if value, ok := acuo.mutation.Key(); ok {
		_spec.SetField(appconfig.FieldKey, field.TypeString, value)
	}
	if value, ok := acuo.mutation.Value(); ok {
		_spec.SetField(appconfig.FieldValue, field.TypeString, value)
	}
	_node = &AppConfig{config: acuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, acuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appconfig.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	acuo.mutation.done = true
	return _node, nil
}
