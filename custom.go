package govals

import (
	"fmt"
	"reflect"
)

// CustomValidate is validate data value
// with custom rule, your can write your regexp custom rule
func CustomValidate(val string, rule string) (bool, error) {
	return checkRule(val, rule)
}

func (d *customValRule) validate() (bool, error) {
	if d.count > 0 {
		if d.tagType.Type.Kind() != reflect.String {
			return false, fmt.Errorf("invalid value type, value type must be string")
		}
		val := d.val.String()
		return checkRule(val, d.rule)
	}
	return false, nil
}
