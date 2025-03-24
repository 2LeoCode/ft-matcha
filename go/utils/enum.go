package utils

import "reflect"

type Enum[T ~int64] struct{}

func (this *Enum[T]) Contains(value T) bool {
	v := reflect.ValueOf(*this)

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Int() == int64(value) {
			return true
		}
	}
	return false
}
