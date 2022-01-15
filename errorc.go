package errorc

import (
	"errors"
	"fmt"
)

type errorc interface {
	Code() int
	Text() string
	FullError() string
}

// custom error
type errorcS struct {
	code  int
	text  string
	Error error
}

// new error without code
func New(text string) *errorcS {
	return &errorcS{
		text: text,
	}
}

// new error with code, purposely for http response
func NewWithCode(code int, text string) *errorcS {
	return &errorcS{
		code:  code,
		text:  text,
		Error: errors.New(text),
	}
}

func (e *errorcS) Code() int {
	return e.code
}

func (e *errorcS) Text() string {
	return e.text
}

func (e *errorcS) FullError() string {
	return fmt.Sprintf("status: %d. error: %s", e.code, e.text)
}
