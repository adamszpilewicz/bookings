package forms

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"net/http"
	"net/url"
	"strings"
)

type Form struct {
	url.Values
	Errors errors
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "this field cannot be blank")
		}
	}
}

func (f *Form) MinLength(field string, minLen int) bool {
	x := f.Get(field)
	if len(x) < minLen {
		f.Errors.Add(field, fmt.Sprintf("this field has to be at least %d characters length", minLen))
		return false
	}
	return true
}

func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get("field")
	if x == "" {
		//f.Errors.Add(field,"this field cannot be blank")
		return false
	}
	return true
}

func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get("email")) {
		f.Errors.Add(field, "incorrect email")
	}
}
