package govals

import (
	"errors"
)

var (
	errFieldNotResgister = errors.New("field is not register")
	errEmptyValue        = errors.New("value cannot be empty")
	errInvalidFormat     = errors.New("value format is invalid")
	errInvalidRuleFormat = errors.New("Rule format is invalid")
)
