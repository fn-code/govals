package govals

import (
	"fmt"
	"strings"
)

func (t *tagField) process() validator {
	tag := t.getTag()
	switch tag[0] {
	case "len":
		return t.setLenRule(tag)
	case "email", "time", "date", "phone", "alpha", "numeric", "alphaNumeric":
		return t.setDefaultRule(tag)
	case "custom":
		return t.setCustomRule(tag)
	}
	return &noValRule{}
}

func (t *tagField) getTag() []string {
	args := strings.Split(t.tagVal, ",")
	for i, v := range args {
		args[i] = strings.Trim(v, " ")
	}
	return args
}

func (t *tagField) setLenRule(s []string) *lengthValsRule {
	rl := lengthValsRule{}
	rl.tag = s[0]
	rl.name = t.tagType.Name
	rl.val = t.val
	rl.tagType = t.tagType
	if len(s) > 1 {
		for _, v := range s[1:] {
			ok := strings.HasPrefix(v, "min")
			if ok {
				rl.count++
				fmt.Sscanf(v, "min=%d", &rl.min)
			}
			ok = strings.HasPrefix(v, "max")
			if ok {
				rl.count++
				fmt.Sscanf(v, "max=%d", &rl.max)
			}
		}
	}
	return &rl
}

func (t *tagField) setDefaultRule(s []string) *defaultValsRule {
	rl := defaultValsRule{}
	rl.name = t.tagType.Name
	rl.val = t.val
	rl.tag = s[0]
	return &rl
}

func (t *tagField) setCustomRule(s []string) *customValRule {
	rl := customValRule{}
	rl.tagType = t.tagType
	rl.name = t.tagType.Name
	rl.val = t.val
	rl.tag = s[0]
	if len(s) > 1 {
		ok := strings.HasPrefix(s[1], "regex")
		if ok {
			rl.count++
			fmt.Sscanf(s[1], "regex=%s", &rl.rule)
		}
	}
	return &rl
}
