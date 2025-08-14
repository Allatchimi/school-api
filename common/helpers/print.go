package helpers

import (
	"fmt"
	"reflect"
)

func PrintAny(v any) string {
	if v == nil {
		return ""
	}

	rv := reflect.ValueOf(v)

	// If v is a pointer an non-nil value
	for rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
		if rv.IsNil() {
			return ""
		}
		rv = rv.Elem()
	}

	//  Otherwise, print the value
	return fmt.Sprint(rv.Interface())
}
