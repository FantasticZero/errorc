package errorc

import (
	"fmt"
)

type error interface {
	Code() int
	Text() string
	FullError() string
}

// custom error
type errorC struct {
	code int
	text string
}

// new error without code
func New(text string) error {
	return &errorC{
		text: text,
	}
}

// new error with code, purposely for http response
func NewWithCode(code int, text string) error {
	return &errorC{
		code: code,
		text: text,
	}
}

func (e *errorC) Code() int {
	return e.code
}

func (e *errorC) Text() string {
	return e.text
}

func (e *errorC) FullError() string {
	return fmt.Sprintf("status: %d. error: %s", e.code, e.text)
}

func (e *errorC) Error() string {
	return e.text
}
