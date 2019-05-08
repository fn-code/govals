package govals

import (
	"reflect"
)

var tagName = "govals"

// Validator data validator.
type validator interface {
	validate() (bool, error)
}

type tagField struct {
	val     reflect.Value
	tagVal  string
	status  bool
	err     error
	tagType reflect.StructField
}

type lengthValsRule struct {
	name    string
	val     reflect.Value
	tag     string
	min     int
	max     int
	tagType reflect.StructField
	count   int
}

type customValRule struct {
	name    string
	val     reflect.Value
	tag     string
	rule    string
	tagType reflect.StructField
	count   int
}

type defaultValsRule struct {
	name string
	val  reflect.Value
	tag  string
}
