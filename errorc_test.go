package errorc

import (
	"fmt"
	"testing"
)

func Test__(t *testing.T) {
	err := NewWithCode(200, "ok")
	fmt.Println(err.Error)
	fmt.Println(err.Code())
	fmt.Println(err.Text())
	fmt.Println(err.FullError())
	fmt.Println()

	err = NewWithCode(400, "bad request")
	fmt.Println(err.Error)
	fmt.Println(err.Code())
	fmt.Println(err.Text())
	fmt.Println(err.FullError())
	fmt.Println()

	err = NewWithCode(500, "internal server error")
	fmt.Println(err.Error)
	fmt.Println(err.Code())
	fmt.Println(err.Text())
	fmt.Println(err.FullError())
	fmt.Println()

	err = New("parameter nil")
	fmt.Println(err.Error)
	fmt.Println(err.Code())
	fmt.Println(err.Text())
	fmt.Println(err.FullError())
	fmt.Println()
}
