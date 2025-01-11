package enums

import "reflect"

type enum[T ~int64] struct{}

func (this *enum[T]) Contains(value T) bool {
	v := reflect.ValueOf(*this)

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Int() == int64(value) {
			return true
		}
	}
	return false
}
