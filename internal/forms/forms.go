package forms

import "net/url"

// Form custom form struct
type Form struct {
	url.Values
	Errors errors
}

// New initializes form struct
func New(data url.Values) *Form {

	return &Form{
		data,
		errors(map[string][]string{}),
	}
}
