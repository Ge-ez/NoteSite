package form

import (
	"fmt"
	"net/url"
	"regexp"
	"unicode/utf8"
)

// EmailRX represents email address maching pattern
var EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// Input represents form input values and validations
type Input struct {
	Values  url.Values
	VErrors ValidationErrors
	CSRF    string
}

// MinLength checks if a given minium length is satisfied
func (inVal *Input) MinLength(field string, d int) {
	value := inVal.Values.Get(field)
	if value == "" {
		return
	}
	if utf8.RuneCountInString(value) < d {
		inVal.VErrors.Add(field, fmt.Sprintf("This field is too short (minimum is %d characters)", d))
	}
}

// Required checks if list of provided form input fields have values
func (inVal *Input) Required(fields ...string) {
	for _, f := range fields {
		value := inVal.Values.Get(f)
		if value == "" {
			inVal.VErrors.Add(f, "This field is required field")
		}
	}
}

// MatchesPattern checks if a given input form field matchs a given pattern
func (inVal *Input) MatchesPattern(field string, pattern *regexp.Regexp) {
	value := inVal.Values.Get(field)
	if value == "" {
		return
	}
	if !pattern.MatchString(value) {
		inVal.VErrors.Add(field, "The value entered is invalid")
	}
}
