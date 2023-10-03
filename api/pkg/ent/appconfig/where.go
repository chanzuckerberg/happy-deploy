// Code generated by ent, DO NOT EDIT.

package appconfig

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/chanzuckerberg/happy/api/pkg/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEQ(FieldDeletedAt, v))
}

// AppName applies equality check predicate on the "app_name" field. It's identical to AppNameEQ.
func AppName(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEQ(FieldAppName, v))
}

// Environment applies equality check predicate on the "environment" field. It's identical to EnvironmentEQ.
func Environment(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEQ(FieldEnvironment, v))
}

// Stack applies equality check predicate on the "stack" field. It's identical to StackEQ.
func Stack(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEQ(FieldStack, v))
}

// Key applies equality check predicate on the "key" field. It's identical to KeyEQ.
func Key(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEQ(FieldKey, v))
}

// Value applies equality check predicate on the "value" field. It's identical to ValueEQ.
func Value(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEQ(FieldValue, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.AppConfig {
	return predicate.AppConfig(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNotNull(FieldDeletedAt))
}

// AppNameEQ applies the EQ predicate on the "app_name" field.
func AppNameEQ(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEQ(FieldAppName, v))
}

// AppNameNEQ applies the NEQ predicate on the "app_name" field.
func AppNameNEQ(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNEQ(FieldAppName, v))
}

// AppNameIn applies the In predicate on the "app_name" field.
func AppNameIn(vs ...string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldIn(FieldAppName, vs...))
}

// AppNameNotIn applies the NotIn predicate on the "app_name" field.
func AppNameNotIn(vs ...string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNotIn(FieldAppName, vs...))
}

// AppNameGT applies the GT predicate on the "app_name" field.
func AppNameGT(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldGT(FieldAppName, v))
}

// AppNameGTE applies the GTE predicate on the "app_name" field.
func AppNameGTE(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldGTE(FieldAppName, v))
}

// AppNameLT applies the LT predicate on the "app_name" field.
func AppNameLT(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldLT(FieldAppName, v))
}

// AppNameLTE applies the LTE predicate on the "app_name" field.
func AppNameLTE(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldLTE(FieldAppName, v))
}

// AppNameContains applies the Contains predicate on the "app_name" field.
func AppNameContains(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldContains(FieldAppName, v))
}

// AppNameHasPrefix applies the HasPrefix predicate on the "app_name" field.
func AppNameHasPrefix(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldHasPrefix(FieldAppName, v))
}

// AppNameHasSuffix applies the HasSuffix predicate on the "app_name" field.
func AppNameHasSuffix(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldHasSuffix(FieldAppName, v))
}

// AppNameEqualFold applies the EqualFold predicate on the "app_name" field.
func AppNameEqualFold(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEqualFold(FieldAppName, v))
}

// AppNameContainsFold applies the ContainsFold predicate on the "app_name" field.
func AppNameContainsFold(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldContainsFold(FieldAppName, v))
}

// EnvironmentEQ applies the EQ predicate on the "environment" field.
func EnvironmentEQ(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEQ(FieldEnvironment, v))
}

// EnvironmentNEQ applies the NEQ predicate on the "environment" field.
func EnvironmentNEQ(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNEQ(FieldEnvironment, v))
}

// EnvironmentIn applies the In predicate on the "environment" field.
func EnvironmentIn(vs ...string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldIn(FieldEnvironment, vs...))
}

// EnvironmentNotIn applies the NotIn predicate on the "environment" field.
func EnvironmentNotIn(vs ...string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNotIn(FieldEnvironment, vs...))
}

// EnvironmentGT applies the GT predicate on the "environment" field.
func EnvironmentGT(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldGT(FieldEnvironment, v))
}

// EnvironmentGTE applies the GTE predicate on the "environment" field.
func EnvironmentGTE(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldGTE(FieldEnvironment, v))
}

// EnvironmentLT applies the LT predicate on the "environment" field.
func EnvironmentLT(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldLT(FieldEnvironment, v))
}

// EnvironmentLTE applies the LTE predicate on the "environment" field.
func EnvironmentLTE(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldLTE(FieldEnvironment, v))
}

// EnvironmentContains applies the Contains predicate on the "environment" field.
func EnvironmentContains(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldContains(FieldEnvironment, v))
}

// EnvironmentHasPrefix applies the HasPrefix predicate on the "environment" field.
func EnvironmentHasPrefix(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldHasPrefix(FieldEnvironment, v))
}

// EnvironmentHasSuffix applies the HasSuffix predicate on the "environment" field.
func EnvironmentHasSuffix(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldHasSuffix(FieldEnvironment, v))
}

// EnvironmentEqualFold applies the EqualFold predicate on the "environment" field.
func EnvironmentEqualFold(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEqualFold(FieldEnvironment, v))
}

// EnvironmentContainsFold applies the ContainsFold predicate on the "environment" field.
func EnvironmentContainsFold(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldContainsFold(FieldEnvironment, v))
}

// StackEQ applies the EQ predicate on the "stack" field.
func StackEQ(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEQ(FieldStack, v))
}

// StackNEQ applies the NEQ predicate on the "stack" field.
func StackNEQ(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNEQ(FieldStack, v))
}

// StackIn applies the In predicate on the "stack" field.
func StackIn(vs ...string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldIn(FieldStack, vs...))
}

// StackNotIn applies the NotIn predicate on the "stack" field.
func StackNotIn(vs ...string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNotIn(FieldStack, vs...))
}

// StackGT applies the GT predicate on the "stack" field.
func StackGT(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldGT(FieldStack, v))
}

// StackGTE applies the GTE predicate on the "stack" field.
func StackGTE(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldGTE(FieldStack, v))
}

// StackLT applies the LT predicate on the "stack" field.
func StackLT(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldLT(FieldStack, v))
}

// StackLTE applies the LTE predicate on the "stack" field.
func StackLTE(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldLTE(FieldStack, v))
}

// StackContains applies the Contains predicate on the "stack" field.
func StackContains(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldContains(FieldStack, v))
}

// StackHasPrefix applies the HasPrefix predicate on the "stack" field.
func StackHasPrefix(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldHasPrefix(FieldStack, v))
}

// StackHasSuffix applies the HasSuffix predicate on the "stack" field.
func StackHasSuffix(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldHasSuffix(FieldStack, v))
}

// StackIsNil applies the IsNil predicate on the "stack" field.
func StackIsNil() predicate.AppConfig {
	return predicate.AppConfig(sql.FieldIsNull(FieldStack))
}

// StackNotNil applies the NotNil predicate on the "stack" field.
func StackNotNil() predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNotNull(FieldStack))
}

// StackEqualFold applies the EqualFold predicate on the "stack" field.
func StackEqualFold(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEqualFold(FieldStack, v))
}

// StackContainsFold applies the ContainsFold predicate on the "stack" field.
func StackContainsFold(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldContainsFold(FieldStack, v))
}

// KeyEQ applies the EQ predicate on the "key" field.
func KeyEQ(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEQ(FieldKey, v))
}

// KeyNEQ applies the NEQ predicate on the "key" field.
func KeyNEQ(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNEQ(FieldKey, v))
}

// KeyIn applies the In predicate on the "key" field.
func KeyIn(vs ...string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldIn(FieldKey, vs...))
}

// KeyNotIn applies the NotIn predicate on the "key" field.
func KeyNotIn(vs ...string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNotIn(FieldKey, vs...))
}

// KeyGT applies the GT predicate on the "key" field.
func KeyGT(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldGT(FieldKey, v))
}

// KeyGTE applies the GTE predicate on the "key" field.
func KeyGTE(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldGTE(FieldKey, v))
}

// KeyLT applies the LT predicate on the "key" field.
func KeyLT(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldLT(FieldKey, v))
}

// KeyLTE applies the LTE predicate on the "key" field.
func KeyLTE(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldLTE(FieldKey, v))
}

// KeyContains applies the Contains predicate on the "key" field.
func KeyContains(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldContains(FieldKey, v))
}

// KeyHasPrefix applies the HasPrefix predicate on the "key" field.
func KeyHasPrefix(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldHasPrefix(FieldKey, v))
}

// KeyHasSuffix applies the HasSuffix predicate on the "key" field.
func KeyHasSuffix(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldHasSuffix(FieldKey, v))
}

// KeyEqualFold applies the EqualFold predicate on the "key" field.
func KeyEqualFold(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEqualFold(FieldKey, v))
}

// KeyContainsFold applies the ContainsFold predicate on the "key" field.
func KeyContainsFold(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldContainsFold(FieldKey, v))
}

// ValueEQ applies the EQ predicate on the "value" field.
func ValueEQ(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEQ(FieldValue, v))
}

// ValueNEQ applies the NEQ predicate on the "value" field.
func ValueNEQ(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNEQ(FieldValue, v))
}

// ValueIn applies the In predicate on the "value" field.
func ValueIn(vs ...string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldIn(FieldValue, vs...))
}

// ValueNotIn applies the NotIn predicate on the "value" field.
func ValueNotIn(vs ...string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNotIn(FieldValue, vs...))
}

// ValueGT applies the GT predicate on the "value" field.
func ValueGT(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldGT(FieldValue, v))
}

// ValueGTE applies the GTE predicate on the "value" field.
func ValueGTE(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldGTE(FieldValue, v))
}

// ValueLT applies the LT predicate on the "value" field.
func ValueLT(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldLT(FieldValue, v))
}

// ValueLTE applies the LTE predicate on the "value" field.
func ValueLTE(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldLTE(FieldValue, v))
}

// ValueContains applies the Contains predicate on the "value" field.
func ValueContains(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldContains(FieldValue, v))
}

// ValueHasPrefix applies the HasPrefix predicate on the "value" field.
func ValueHasPrefix(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldHasPrefix(FieldValue, v))
}

// ValueHasSuffix applies the HasSuffix predicate on the "value" field.
func ValueHasSuffix(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldHasSuffix(FieldValue, v))
}

// ValueEqualFold applies the EqualFold predicate on the "value" field.
func ValueEqualFold(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEqualFold(FieldValue, v))
}

// ValueContainsFold applies the ContainsFold predicate on the "value" field.
func ValueContainsFold(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldContainsFold(FieldValue, v))
}

// SourceEQ applies the EQ predicate on the "source" field.
func SourceEQ(v Source) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEQ(FieldSource, v))
}

// SourceNEQ applies the NEQ predicate on the "source" field.
func SourceNEQ(v Source) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNEQ(FieldSource, v))
}

// SourceIn applies the In predicate on the "source" field.
func SourceIn(vs ...Source) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldIn(FieldSource, vs...))
}

// SourceNotIn applies the NotIn predicate on the "source" field.
func SourceNotIn(vs ...Source) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNotIn(FieldSource, vs...))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AppConfig) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AppConfig) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.AppConfig) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		p(s.Not())
	})
}
