package ErrorExample

import (
	"errors"
	"fmt"
)

func Add(a, b interface{}) (int, error) {
	aValue, ok := a.(int)
	if ok != true {
		return -1, errors.New("arg1 is not integer")
	}
	bValue, ok := b.(int)
	if ok != true {
		return -1, errors.New("arg2 is not integer")
	}

	return aValue + bValue, nil

}

type argError struct {
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%s", e.prob)
}

func Divide(arg1, arg2 int) (int, error) {
	if arg2 == 0 {
		return -1, &argError{"Zero Divide Error"}
	}
	return arg1 / arg2, nil
}
