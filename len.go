package govals

import (
	"fmt"
	"reflect"
	"strconv"
)

// RuneLength is validate text length
func RuneLength(s string, min, max int) (bool, error) {
	return minMax(s, min, max)
}

// Length is validate number length
func Length(s int, min, max int) (bool, error) {
	num := strconv.Itoa(s)
	return minMax(num, min, max)
}

func (d *lengthValsRule) validate() (ok bool, err error) {

	switch d.tagType.Type.Kind() {
	case reflect.String:
		val := d.val.String()
		if d.count > 1 {
			return minMax(val, d.min, d.max)
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		val := d.val.Int()
		if d.count > 1 {
			return minMax(strconv.Itoa(int(val)), d.min, d.max)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		val := d.val.Uint()
		return minMax(strconv.Itoa(int(val)), d.min, d.max)
	}
	return true, nil
}

func minMax(s string, min, max int) (bool, error) {
	if len(s) == 0 {
		return false, errEmptyValue
	}
	if min > max {
		return false, fmt.Errorf("invalid rule, min must less than max")
	}
	if len(s) < min {
		return false, fmt.Errorf("value must greater than %d", min)
	}
	if len(s) > max {
		return false, fmt.Errorf("value must less than %d", max)
	}
	return true, nil
}
