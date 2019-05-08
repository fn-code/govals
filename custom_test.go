package govals

import "testing"

func TestCustomValidate(t *testing.T) {
	tc := []struct {
		name string
		val  string
		rule string
		err  error
	}{
		{name: "valid test", val: "ludinnento", rule: "^[a-zA-Z]*$"},
		{name: "empty value", val: "", rule: "^[a-zA-Z]*$", err: errEmptyValue},
		{name: "invalid value", val: "ludin1233", rule: "^[a-zA-Z]*$", err: errInvalidFormat},
		{name: "invalid rule", val: "ludin", rule: "^[a-zA-Z", err: errInvalidRuleFormat},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			ok, err := CustomValidate(tt.val, tt.rule)
			if tt.err != nil {
				if tt.err != err {
					t.Errorf("error value expects %v got %v\n", tt.err, err)
				}
				return
			}

			if !ok {
				t.Errorf("%s invalid result", tt.val)
			}
		})
	}

}
