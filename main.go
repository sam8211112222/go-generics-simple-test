package main

import (
	"fmt"
	"reflect"
)

func main() {
	boolTestCase := []bool{false, true}
	if err := truthyPrint(boolTestCase); err != nil {
		fmt.Println(err)
	}
	intTestCase := []int{0, 1}
	if err := truthyPrint(intTestCase); err != nil {
		fmt.Println(err)
	}
	stringTestCase := []string{"", "  ", "test"}
	if err := truthyPrint(stringTestCase); err != nil {
		fmt.Println(err)
	}
}

type allowedType interface {
	string | int | bool
}

func truthyPrint[T allowedType](values []T) error {
	var defaultValue T
	for _, value := range values {
		if !isAllowedType(reflect.TypeOf(value)) {
			return fmt.Errorf("%v is unsupported type %T", value, value)
		}
		if value != defaultValue {
			fmt.Println(value)
		}
	}
	return nil
}

func isAllowedType(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Bool, reflect.Int, reflect.String:
		return true
	default:
		return false
	}
}
