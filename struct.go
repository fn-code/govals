package govals

// StructVals is for store validation status
import (
	"reflect"
)

type StructVals struct {
	fields map[string]tagField
}

// Struct set new validator
func Struct(s interface{}) *StructVals {
	gov := &StructVals{fields: make(map[string]tagField)}
	val := reflect.ValueOf(s)
	num := 0
	if val.Kind() == reflect.Ptr {
		num = val.Elem().NumField()
		val = val.Elem()
	} else {
		num = val.NumField()
	}

	for i := 0; i < num; i++ {
		tg := val.Type().Field(i).Tag.Get(tagName)
		if tg == "-" {
			continue
		} else if tg == "" {
			continue
		}

		tag := &tagField{
			val:     val.Field(i),
			tagVal:  tg,
			tagType: val.Type().Field(i),
		}

		valid := tag.process()
		ok, err := valid.validate()
		tag.err = err
		tag.status = ok

		gov.fields[tag.tagType.Name] = *tag
	}
	return gov
}

// GetStatus return validation bool and error
// of validation fields
func (gov *StructVals) GetStatus(str string) (bool, error) {
	val, ok := gov.fields[str]
	if !ok {
		return false, errFieldNotResgister
	}
	return val.status, val.err
}
