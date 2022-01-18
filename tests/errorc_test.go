package tests

import (
	"errors"
	"fmt"
	"testing"

	"github.com/mhafizhrh/errorc"
)

func TestErrorc(t *testing.T) {
	err := errorc.NewWithCode(400, "success")
	errWihoutCode := errorc.New("err without code")
	builtinErr := errors.New("builtin err")

	e1, ok1 := err.(*errorc.CustomErr)
	e2, ok2 := errWihoutCode.(*errorc.CustomErr)
	e3, ok3 := builtinErr.(*errorc.CustomErr)

	fmt.Println(e1, ok1)
	fmt.Println(e2, ok2)
	fmt.Println(e3, ok3)

	fmt.Println(err)
	fmt.Println(err.Error())
	fmt.Println(err.(*errorc.CustomErr))
	fmt.Println(err.(*errorc.CustomErr).Error())
	fmt.Println(err.(*errorc.CustomErr).Code())
}
