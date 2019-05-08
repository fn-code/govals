package govals

import (
	"regexp"
)

// Email validate string if is valid email format,
// this function return bool and error
// if value does not valid
func Email(s string) (bool, error) {
	return checkRule(s, emailRules)
}

// Numeric validate string contains digits only [0-9]
// this function return bool and error
// if value does not valid
func Numeric(s string) (bool, error) {
	return checkRule(s, numericRule)
}

// Alpha validate string contains letters only [a-zA-Z]
// this function return bool and error
// if value does not valid
func Alpha(s string) (bool, error) {
	return checkRule(s, alphaRule)
}

// AlphaNumeric validate string contains letters
// and number only [a-zA-Z0-9]
// this function return bool and error
// if value does not valid
func AlphaNumeric(s string) (bool, error) {
	return checkRule(s, alphaNumericRule)
}

// Date checking string format as date
// like mm-dd-yyyy or mm/dd/yyyy
// this function return bool and error
// if value does not valid
func Date(s string) (bool, error) {
	return checkRule(s, dateRule)
}

// Times checking string format time
// like 12:00:01 or 12:00
func Times(s string) (bool, error) {
	return checkRule(s, timeRule)
}

// PhoneNumber checking string match phone number
// like +62828882888 or 082323333333
// this function return bool and error
// if value does not valid
func PhoneNumber(s string) (bool, error) {
	return checkRule(s, phoneNumberRule)
}

func (d *defaultValsRule) validate() (bool, error) {
	val := d.val.String()
	switch d.tag {
	case "email":
		return checkRule(val, emailRules)
	case "time":
		return checkRule(val, timeRule)
	case "date":
		return checkRule(val, dateRule)
	case "alpha":
		return checkRule(val, alphaRule)
	case "numeric":
		return checkRule(val, numericRule)
	case "alphaNumeric":
		return checkRule(val, alphaNumericRule)
	case "phone":
		return checkRule(val, phoneNumberRule)
	}
	return true, nil

}

func checkRule(s, rule string) (bool, error) {
	if len(s) == 0 {
		return false, errEmptyValue
	}
	re, err := regexp.Compile(rule)
	if err != nil {
		return false, errInvalidRuleFormat
	}
	ok := re.MatchString(s)
	if !ok {
		return false, errInvalidFormat
	}
	return ok, nil
}
