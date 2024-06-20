package defaults

import "errors"

var (
	ErrInvalidValue    = errors.New("value must be a non-nil pointer to a struct")
	ErrUnsupportedType = errors.New("field is an unsupported type")
	ErrUnexportedField = errors.New("field must be exported")
)
