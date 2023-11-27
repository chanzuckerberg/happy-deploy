// Code generated by ogen, DO NOT EDIT.

package ogent

import (
	"fmt"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/validate"
)

func (s *AppConfigList) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.Int{
			MinSet:        true,
			Min:           0,
			MaxSet:        true,
			Max:           4294967295,
			MinExclusive:  false,
			MaxExclusive:  false,
			MultipleOfSet: false,
			MultipleOf:    0,
		}).Validate(int64(s.ID)); err != nil {
			return errors.Wrap(err, "int")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "id",
			Error: err,
		})
	}
	if err := func() error {
		if err := s.Source.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "source",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s AppConfigListSource) Validate() error {
	switch s {
	case "stack":
		return nil
	case "environment":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s ListAppConfigOKApplicationJSON) Validate() error {
	alias := ([]AppConfigList)(s)
	if alias == nil {
		return errors.New("nil is invalid value")
	}
	var failures []validate.FieldError
	for i, elem := range alias {
		if err := func() error {
			if err := elem.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			failures = append(failures, validate.FieldError{
				Name:  fmt.Sprintf("[%d]", i),
				Error: err,
			})
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
