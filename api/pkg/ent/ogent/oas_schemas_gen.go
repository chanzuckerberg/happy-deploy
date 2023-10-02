// Code generated by ogen, DO NOT EDIT.

package ogent

import (
	"time"

	"github.com/go-faster/jx"
)

// Ref: #/components/schemas/AppConfigList
type AppConfigList struct {
	ID          int64       `json:"id"`
	CreatedAt   OptDateTime `json:"created_at"`
	UpdatedAt   OptDateTime `json:"updated_at"`
	DeletedAt   OptDateTime `json:"deleted_at"`
	AppName     OptString   `json:"app_name"`
	Environment OptString   `json:"environment"`
	Stack       OptString   `json:"stack"`
	Key         OptString   `json:"key"`
	Value       OptString   `json:"value"`
}

// GetID returns the value of ID.
func (s *AppConfigList) GetID() int64 {
	return s.ID
}

// GetCreatedAt returns the value of CreatedAt.
func (s *AppConfigList) GetCreatedAt() OptDateTime {
	return s.CreatedAt
}

// GetUpdatedAt returns the value of UpdatedAt.
func (s *AppConfigList) GetUpdatedAt() OptDateTime {
	return s.UpdatedAt
}

// GetDeletedAt returns the value of DeletedAt.
func (s *AppConfigList) GetDeletedAt() OptDateTime {
	return s.DeletedAt
}

// GetAppName returns the value of AppName.
func (s *AppConfigList) GetAppName() OptString {
	return s.AppName
}

// GetEnvironment returns the value of Environment.
func (s *AppConfigList) GetEnvironment() OptString {
	return s.Environment
}

// GetStack returns the value of Stack.
func (s *AppConfigList) GetStack() OptString {
	return s.Stack
}

// GetKey returns the value of Key.
func (s *AppConfigList) GetKey() OptString {
	return s.Key
}

// GetValue returns the value of Value.
func (s *AppConfigList) GetValue() OptString {
	return s.Value
}

// SetID sets the value of ID.
func (s *AppConfigList) SetID(val int64) {
	s.ID = val
}

// SetCreatedAt sets the value of CreatedAt.
func (s *AppConfigList) SetCreatedAt(val OptDateTime) {
	s.CreatedAt = val
}

// SetUpdatedAt sets the value of UpdatedAt.
func (s *AppConfigList) SetUpdatedAt(val OptDateTime) {
	s.UpdatedAt = val
}

// SetDeletedAt sets the value of DeletedAt.
func (s *AppConfigList) SetDeletedAt(val OptDateTime) {
	s.DeletedAt = val
}

// SetAppName sets the value of AppName.
func (s *AppConfigList) SetAppName(val OptString) {
	s.AppName = val
}

// SetEnvironment sets the value of Environment.
func (s *AppConfigList) SetEnvironment(val OptString) {
	s.Environment = val
}

// SetStack sets the value of Stack.
func (s *AppConfigList) SetStack(val OptString) {
	s.Stack = val
}

// SetKey sets the value of Key.
func (s *AppConfigList) SetKey(val OptString) {
	s.Key = val
}

// SetValue sets the value of Value.
func (s *AppConfigList) SetValue(val OptString) {
	s.Value = val
}

// Ref: #/components/schemas/AppConfigRead
type AppConfigRead struct {
	ID          int64       `json:"id"`
	CreatedAt   OptDateTime `json:"created_at"`
	UpdatedAt   OptDateTime `json:"updated_at"`
	DeletedAt   OptDateTime `json:"deleted_at"`
	AppName     OptString   `json:"app_name"`
	Environment OptString   `json:"environment"`
	Stack       OptString   `json:"stack"`
	Key         OptString   `json:"key"`
	Value       OptString   `json:"value"`
}

// GetID returns the value of ID.
func (s *AppConfigRead) GetID() int64 {
	return s.ID
}

// GetCreatedAt returns the value of CreatedAt.
func (s *AppConfigRead) GetCreatedAt() OptDateTime {
	return s.CreatedAt
}

// GetUpdatedAt returns the value of UpdatedAt.
func (s *AppConfigRead) GetUpdatedAt() OptDateTime {
	return s.UpdatedAt
}

// GetDeletedAt returns the value of DeletedAt.
func (s *AppConfigRead) GetDeletedAt() OptDateTime {
	return s.DeletedAt
}

// GetAppName returns the value of AppName.
func (s *AppConfigRead) GetAppName() OptString {
	return s.AppName
}

// GetEnvironment returns the value of Environment.
func (s *AppConfigRead) GetEnvironment() OptString {
	return s.Environment
}

// GetStack returns the value of Stack.
func (s *AppConfigRead) GetStack() OptString {
	return s.Stack
}

// GetKey returns the value of Key.
func (s *AppConfigRead) GetKey() OptString {
	return s.Key
}

// GetValue returns the value of Value.
func (s *AppConfigRead) GetValue() OptString {
	return s.Value
}

// SetID sets the value of ID.
func (s *AppConfigRead) SetID(val int64) {
	s.ID = val
}

// SetCreatedAt sets the value of CreatedAt.
func (s *AppConfigRead) SetCreatedAt(val OptDateTime) {
	s.CreatedAt = val
}

// SetUpdatedAt sets the value of UpdatedAt.
func (s *AppConfigRead) SetUpdatedAt(val OptDateTime) {
	s.UpdatedAt = val
}

// SetDeletedAt sets the value of DeletedAt.
func (s *AppConfigRead) SetDeletedAt(val OptDateTime) {
	s.DeletedAt = val
}

// SetAppName sets the value of AppName.
func (s *AppConfigRead) SetAppName(val OptString) {
	s.AppName = val
}

// SetEnvironment sets the value of Environment.
func (s *AppConfigRead) SetEnvironment(val OptString) {
	s.Environment = val
}

// SetStack sets the value of Stack.
func (s *AppConfigRead) SetStack(val OptString) {
	s.Stack = val
}

// SetKey sets the value of Key.
func (s *AppConfigRead) SetKey(val OptString) {
	s.Key = val
}

// SetValue sets the value of Value.
func (s *AppConfigRead) SetValue(val OptString) {
	s.Value = val
}

func (*AppConfigRead) readAppConfigRes() {}

// Ref: #/components/schemas/AppStackList
type AppStackList struct {
	ID          int64       `json:"id"`
	CreatedAt   OptDateTime `json:"created_at"`
	UpdatedAt   OptDateTime `json:"updated_at"`
	DeletedAt   OptDateTime `json:"deleted_at"`
	AppName     OptString   `json:"app_name"`
	Environment OptString   `json:"environment"`
	Stack       string      `json:"stack"`
}

// GetID returns the value of ID.
func (s *AppStackList) GetID() int64 {
	return s.ID
}

// GetCreatedAt returns the value of CreatedAt.
func (s *AppStackList) GetCreatedAt() OptDateTime {
	return s.CreatedAt
}

// GetUpdatedAt returns the value of UpdatedAt.
func (s *AppStackList) GetUpdatedAt() OptDateTime {
	return s.UpdatedAt
}

// GetDeletedAt returns the value of DeletedAt.
func (s *AppStackList) GetDeletedAt() OptDateTime {
	return s.DeletedAt
}

// GetAppName returns the value of AppName.
func (s *AppStackList) GetAppName() OptString {
	return s.AppName
}

// GetEnvironment returns the value of Environment.
func (s *AppStackList) GetEnvironment() OptString {
	return s.Environment
}

// GetStack returns the value of Stack.
func (s *AppStackList) GetStack() string {
	return s.Stack
}

// SetID sets the value of ID.
func (s *AppStackList) SetID(val int64) {
	s.ID = val
}

// SetCreatedAt sets the value of CreatedAt.
func (s *AppStackList) SetCreatedAt(val OptDateTime) {
	s.CreatedAt = val
}

// SetUpdatedAt sets the value of UpdatedAt.
func (s *AppStackList) SetUpdatedAt(val OptDateTime) {
	s.UpdatedAt = val
}

// SetDeletedAt sets the value of DeletedAt.
func (s *AppStackList) SetDeletedAt(val OptDateTime) {
	s.DeletedAt = val
}

// SetAppName sets the value of AppName.
func (s *AppStackList) SetAppName(val OptString) {
	s.AppName = val
}

// SetEnvironment sets the value of Environment.
func (s *AppStackList) SetEnvironment(val OptString) {
	s.Environment = val
}

// SetStack sets the value of Stack.
func (s *AppStackList) SetStack(val string) {
	s.Stack = val
}

// Ref: #/components/schemas/AppStackRead
type AppStackRead struct {
	ID          int64       `json:"id"`
	CreatedAt   OptDateTime `json:"created_at"`
	UpdatedAt   OptDateTime `json:"updated_at"`
	DeletedAt   OptDateTime `json:"deleted_at"`
	AppName     OptString   `json:"app_name"`
	Environment OptString   `json:"environment"`
	Stack       string      `json:"stack"`
}

// GetID returns the value of ID.
func (s *AppStackRead) GetID() int64 {
	return s.ID
}

// GetCreatedAt returns the value of CreatedAt.
func (s *AppStackRead) GetCreatedAt() OptDateTime {
	return s.CreatedAt
}

// GetUpdatedAt returns the value of UpdatedAt.
func (s *AppStackRead) GetUpdatedAt() OptDateTime {
	return s.UpdatedAt
}

// GetDeletedAt returns the value of DeletedAt.
func (s *AppStackRead) GetDeletedAt() OptDateTime {
	return s.DeletedAt
}

// GetAppName returns the value of AppName.
func (s *AppStackRead) GetAppName() OptString {
	return s.AppName
}

// GetEnvironment returns the value of Environment.
func (s *AppStackRead) GetEnvironment() OptString {
	return s.Environment
}

// GetStack returns the value of Stack.
func (s *AppStackRead) GetStack() string {
	return s.Stack
}

// SetID sets the value of ID.
func (s *AppStackRead) SetID(val int64) {
	s.ID = val
}

// SetCreatedAt sets the value of CreatedAt.
func (s *AppStackRead) SetCreatedAt(val OptDateTime) {
	s.CreatedAt = val
}

// SetUpdatedAt sets the value of UpdatedAt.
func (s *AppStackRead) SetUpdatedAt(val OptDateTime) {
	s.UpdatedAt = val
}

// SetDeletedAt sets the value of DeletedAt.
func (s *AppStackRead) SetDeletedAt(val OptDateTime) {
	s.DeletedAt = val
}

// SetAppName sets the value of AppName.
func (s *AppStackRead) SetAppName(val OptString) {
	s.AppName = val
}

// SetEnvironment sets the value of Environment.
func (s *AppStackRead) SetEnvironment(val OptString) {
	s.Environment = val
}

// SetStack sets the value of Stack.
func (s *AppStackRead) SetStack(val string) {
	s.Stack = val
}

func (*AppStackRead) readAppStackRes() {}

type ListAppConfigOKApplicationJSON []AppConfigList

func (*ListAppConfigOKApplicationJSON) listAppConfigRes() {}

type ListAppStackOKApplicationJSON []AppStackList

func (*ListAppStackOKApplicationJSON) listAppStackRes() {}

// NewOptDateTime returns new OptDateTime with value set to v.
func NewOptDateTime(v time.Time) OptDateTime {
	return OptDateTime{
		Value: v,
		Set:   true,
	}
}

// OptDateTime is optional time.Time.
type OptDateTime struct {
	Value time.Time
	Set   bool
}

// IsSet returns true if OptDateTime was set.
func (o OptDateTime) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptDateTime) Reset() {
	var v time.Time
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptDateTime) SetTo(v time.Time) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptDateTime) Get() (v time.Time, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptDateTime) Or(d time.Time) time.Time {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptInt returns new OptInt with value set to v.
func NewOptInt(v int) OptInt {
	return OptInt{
		Value: v,
		Set:   true,
	}
}

// OptInt is optional int.
type OptInt struct {
	Value int
	Set   bool
}

// IsSet returns true if OptInt was set.
func (o OptInt) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt) Reset() {
	var v int
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt) SetTo(v int) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt) Get() (v int, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt) Or(d int) int {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

type R400 struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Errors jx.Raw `json:"errors"`
}

// GetCode returns the value of Code.
func (s *R400) GetCode() int {
	return s.Code
}

// GetStatus returns the value of Status.
func (s *R400) GetStatus() string {
	return s.Status
}

// GetErrors returns the value of Errors.
func (s *R400) GetErrors() jx.Raw {
	return s.Errors
}

// SetCode sets the value of Code.
func (s *R400) SetCode(val int) {
	s.Code = val
}

// SetStatus sets the value of Status.
func (s *R400) SetStatus(val string) {
	s.Status = val
}

// SetErrors sets the value of Errors.
func (s *R400) SetErrors(val jx.Raw) {
	s.Errors = val
}

func (*R400) listAppConfigRes() {}
func (*R400) listAppStackRes()  {}
func (*R400) readAppConfigRes() {}
func (*R400) readAppStackRes()  {}

type R404 struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Errors jx.Raw `json:"errors"`
}

// GetCode returns the value of Code.
func (s *R404) GetCode() int {
	return s.Code
}

// GetStatus returns the value of Status.
func (s *R404) GetStatus() string {
	return s.Status
}

// GetErrors returns the value of Errors.
func (s *R404) GetErrors() jx.Raw {
	return s.Errors
}

// SetCode sets the value of Code.
func (s *R404) SetCode(val int) {
	s.Code = val
}

// SetStatus sets the value of Status.
func (s *R404) SetStatus(val string) {
	s.Status = val
}

// SetErrors sets the value of Errors.
func (s *R404) SetErrors(val jx.Raw) {
	s.Errors = val
}

func (*R404) listAppConfigRes() {}
func (*R404) listAppStackRes()  {}
func (*R404) readAppConfigRes() {}
func (*R404) readAppStackRes()  {}

type R409 struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Errors jx.Raw `json:"errors"`
}

// GetCode returns the value of Code.
func (s *R409) GetCode() int {
	return s.Code
}

// GetStatus returns the value of Status.
func (s *R409) GetStatus() string {
	return s.Status
}

// GetErrors returns the value of Errors.
func (s *R409) GetErrors() jx.Raw {
	return s.Errors
}

// SetCode sets the value of Code.
func (s *R409) SetCode(val int) {
	s.Code = val
}

// SetStatus sets the value of Status.
func (s *R409) SetStatus(val string) {
	s.Status = val
}

// SetErrors sets the value of Errors.
func (s *R409) SetErrors(val jx.Raw) {
	s.Errors = val
}

func (*R409) listAppConfigRes() {}
func (*R409) listAppStackRes()  {}
func (*R409) readAppConfigRes() {}
func (*R409) readAppStackRes()  {}

type R500 struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Errors jx.Raw `json:"errors"`
}

// GetCode returns the value of Code.
func (s *R500) GetCode() int {
	return s.Code
}

// GetStatus returns the value of Status.
func (s *R500) GetStatus() string {
	return s.Status
}

// GetErrors returns the value of Errors.
func (s *R500) GetErrors() jx.Raw {
	return s.Errors
}

// SetCode sets the value of Code.
func (s *R500) SetCode(val int) {
	s.Code = val
}

// SetStatus sets the value of Status.
func (s *R500) SetStatus(val string) {
	s.Status = val
}

// SetErrors sets the value of Errors.
func (s *R500) SetErrors(val jx.Raw) {
	s.Errors = val
}

func (*R500) listAppConfigRes() {}
func (*R500) listAppStackRes()  {}
func (*R500) readAppConfigRes() {}
func (*R500) readAppStackRes()  {}
