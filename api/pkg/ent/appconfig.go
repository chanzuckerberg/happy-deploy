// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/chanzuckerberg/happy/api/pkg/ent/appconfig"
)

// AppConfig is the model entity for the AppConfig schema.
type AppConfig struct {
	config `json:"-"`
	// ID of the ent.
	ID uint `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// AppName holds the value of the "app_name" field.
	AppName string `json:"app_name,omitempty"`
	// Environment holds the value of the "environment" field.
	Environment string `json:"environment,omitempty"`
	// Stack holds the value of the "stack" field.
	Stack string `json:"stack,omitempty"`
	// Key holds the value of the "key" field.
	Key string `json:"key,omitempty"`
	// Value holds the value of the "value" field.
	Value        string `json:"value,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AppConfig) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case appconfig.FieldID:
			values[i] = new(sql.NullInt64)
		case appconfig.FieldAppName, appconfig.FieldEnvironment, appconfig.FieldStack, appconfig.FieldKey, appconfig.FieldValue:
			values[i] = new(sql.NullString)
		case appconfig.FieldCreatedAt, appconfig.FieldUpdatedAt, appconfig.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AppConfig fields.
func (ac *AppConfig) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case appconfig.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ac.ID = uint(value.Int64)
		case appconfig.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ac.CreatedAt = value.Time
			}
		case appconfig.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ac.UpdatedAt = value.Time
			}
		case appconfig.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				ac.DeletedAt = new(time.Time)
				*ac.DeletedAt = value.Time
			}
		case appconfig.FieldAppName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field app_name", values[i])
			} else if value.Valid {
				ac.AppName = value.String
			}
		case appconfig.FieldEnvironment:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field environment", values[i])
			} else if value.Valid {
				ac.Environment = value.String
			}
		case appconfig.FieldStack:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field stack", values[i])
			} else if value.Valid {
				ac.Stack = value.String
			}
		case appconfig.FieldKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field key", values[i])
			} else if value.Valid {
				ac.Key = value.String
			}
		case appconfig.FieldValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				ac.Value = value.String
			}
		default:
			ac.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// GetValue returns the ent.Value that was dynamically selected and assigned to the AppConfig.
// This includes values selected through modifiers, order, etc.
func (ac *AppConfig) GetValue(name string) (ent.Value, error) {
	return ac.selectValues.Get(name)
}

// Update returns a builder for updating this AppConfig.
// Note that you need to call AppConfig.Unwrap() before calling this method if this AppConfig
// was returned from a transaction, and the transaction was committed or rolled back.
func (ac *AppConfig) Update() *AppConfigUpdateOne {
	return NewAppConfigClient(ac.config).UpdateOne(ac)
}

// Unwrap unwraps the AppConfig entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ac *AppConfig) Unwrap() *AppConfig {
	_tx, ok := ac.config.driver.(*txDriver)
	if !ok {
		panic("ent: AppConfig is not a transactional entity")
	}
	ac.config.driver = _tx.drv
	return ac
}

// String implements the fmt.Stringer.
func (ac *AppConfig) String() string {
	var builder strings.Builder
	builder.WriteString("AppConfig(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ac.ID))
	builder.WriteString("created_at=")
	builder.WriteString(ac.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ac.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := ac.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("app_name=")
	builder.WriteString(ac.AppName)
	builder.WriteString(", ")
	builder.WriteString("environment=")
	builder.WriteString(ac.Environment)
	builder.WriteString(", ")
	builder.WriteString("stack=")
	builder.WriteString(ac.Stack)
	builder.WriteString(", ")
	builder.WriteString("key=")
	builder.WriteString(ac.Key)
	builder.WriteString(", ")
	builder.WriteString("value=")
	builder.WriteString(ac.Value)
	builder.WriteByte(')')
	return builder.String()
}

// AppConfigs is a parsable slice of AppConfig.
type AppConfigs []*AppConfig
