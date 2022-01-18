package errorc

// custom error
type CustomErr struct {
	code int
	text string
}

// new error without code
func New(text string) error {
	return &CustomErr{
		text: text,
	}
}

// new error with code, purposely for http response
func NewWithCode(code int, text string) error {
	return &CustomErr{
		code: code,
		text: text,
	}
}

func (e *CustomErr) Error() string {
	return e.text
}

func (e *CustomErr) Code() int {
	return e.code
}
