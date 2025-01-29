// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"time"

	"github.com/chanzuckerberg/happy/api/pkg/ent/appconfig"
	"github.com/chanzuckerberg/happy/api/pkg/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	appconfigHooks := schema.AppConfig{}.Hooks()
	appconfig.Hooks[0] = appconfigHooks[0]
	appconfigFields := schema.AppConfig{}.Fields()
	_ = appconfigFields
	// appconfigDescCreatedAt is the schema descriptor for created_at field.
	appconfigDescCreatedAt := appconfigFields[1].Descriptor()
	// appconfig.DefaultCreatedAt holds the default value on creation for the created_at field.
	appconfig.DefaultCreatedAt = appconfigDescCreatedAt.Default.(func() time.Time)
	// appconfigDescUpdatedAt is the schema descriptor for updated_at field.
	appconfigDescUpdatedAt := appconfigFields[2].Descriptor()
	// appconfig.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	appconfig.DefaultUpdatedAt = appconfigDescUpdatedAt.Default.(func() time.Time)
	// appconfig.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	appconfig.UpdateDefaultUpdatedAt = appconfigDescUpdatedAt.UpdateDefault.(func() time.Time)
	// appconfigDescStack is the schema descriptor for stack field.
	appconfigDescStack := appconfigFields[6].Descriptor()
	// appconfig.DefaultStack holds the default value on creation for the stack field.
	appconfig.DefaultStack = appconfigDescStack.Default.(string)
}

const (
	Version = "v0.13.2-0.20240717044502-34158f2c129b"           // Version of ent codegen.
	Sum     = "h1:kC+uzL8UFWwtXQ+yY0wUdvVUgPlJPGU3Fx1uttM8PJA=" // Sum of ent codegen.
)
