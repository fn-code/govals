package govals

import (
	"fmt"
	"testing"
)

func TestRuneLength(t *testing.T) {
	tc := []struct {
		name  string
		value string
		min   int
		max   int
		err   error
	}{
		{name: "valid rune length", value: "ludin", min: 2, max: 10},
		{name: "invalid rune length min", value: "ludin", min: 8, max: 10, err: fmt.Errorf("error value")},
		{name: "invalid rune length max", value: "ludidsfsdfdfsdf234343", min: 8, max: 10, err: fmt.Errorf("error value")},
		{name: "invalid len parameter", value: "ludin", min: 11, max: 10, err: fmt.Errorf("error value")},
		{name: "empty rune", value: "", min: 11, max: 10, err: fmt.Errorf("error value")},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			ok, err := RuneLength(tt.value, tt.min, tt.max)
			if tt.err != nil {
				if err == nil && ok {
					t.Errorf("error expect %v, got %v\n", tt.err, err)
				}
			}
		})
	}
}

func TestLength(t *testing.T) {
	tc := []struct {
		name  string
		value int
		min   int
		max   int
		err   error
	}{
		{name: "valid length", value: 123, min: 2, max: 10},
		{name: "invalid length min", value: 123, min: 8, max: 10, err: fmt.Errorf("error value")},
		{name: "invalid length max", value: 123234342342343, min: 8, max: 10, err: fmt.Errorf("error value")},
		{name: "invalid len parameter", value: 1234, min: 11, max: 10, err: fmt.Errorf("error value")},
		{name: "empty length", min: 3, max: 10, err: fmt.Errorf("error value")},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			ok, err := Length(tt.value, tt.min, tt.max)
			if tt.err != nil {
				if err == nil && ok {
					t.Errorf("error expect %v, got %v\n", tt.err, err)
				}
			}
		})
	}
}
