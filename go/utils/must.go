package utils

import "log"

func Must[T any](value T, err error) T {
	if err != nil {
		log.Fatalln(err.Error())
	}
	return value
}
