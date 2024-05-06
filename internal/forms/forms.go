package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

type Form struct {
	url.Values
	Errors errors
}

// New initialises a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Check if form field in post is not empty
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	return x != ""
}

// Returns true if there are no errors
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

func (f *Form) MinLength(field string, length int, r *http.Request) bool {
	x := r.Form.Get(field)

	if len(x) < length {
		f.Errors.Add(field, fmt.Sprintf("%s must be at least %d characters long ", field, length))
		return false
	}

	return true
}

func (f *Form) IsEmail(field string) bool {

	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Email should be valid")
		return false
	}

	return true
}
