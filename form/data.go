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
