package errors

import "strings"

type ErrorsEntity struct {
	Error  bool              `json:"error"`
	Errors map[string]string `json:"errors"`
}

func (e *ErrorsEntity) AddError(typeError, mot, libel string) {
	if e.Errors == nil {
		e.Errors = make(map[string]string)
	}

	switch strings.ToLower(typeError) {
	case "e":
		e.Errors[mot] = libel
		e.Error = true
	}
}
